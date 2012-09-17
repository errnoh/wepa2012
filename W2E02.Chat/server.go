package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

func getTemplate(file string) *template.Template {
	page, err := template.ParseFiles(file)
	if err != nil {
		log.Fatalln(err)
	}
	return page
}

func main() {
	controller := addControllers()

	log.Fatal(http.ListenAndServe(":8080", controller))
}

func addControllers() *FrontController {
	controller := NewFrontController()
	messageservice := NewMessageService()

	controller.AddController("/list", &ListMessagesController{messageservice})
	controller.AddController("/", &IndexController{})
	controller.AddController("/add-message", &AddMessageController{messageservice})

	return controller
}

//-------

type MessageService struct {
	messages     []string
	messagescopy []string
	mutex        *sync.Mutex
}

func NewMessageService() *MessageService {
	return &MessageService{
		messages:     make([]string, 0, 10),
		messagescopy: make([]string, 0),
		mutex:        new(sync.Mutex),
	}
}

func (m *MessageService) Add(message string) {
	m.mutex.Lock()
	m.messages = append(m.messages, message)
	m.mutex.Unlock()
}

// Returns a copy of the messages for reading
func (m *MessageService) Messages() []string {
	m.mutex.Lock()

	// Tarkastetaan löyhästi onko viestien määrä muuttunut
	if len(m.messages) == len(m.messagescopy) {
		m.mutex.Unlock()
		return m.messagescopy
	}

	// Jos on niin päivitetään kopio
	m.messagescopy = make([]string, len(m.messages))
	copy(m.messagescopy, m.messages)
	m.mutex.Unlock()
	return m.messagescopy
}
