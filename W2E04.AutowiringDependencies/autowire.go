package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type MessageContainer interface {
	Init()
	Message() string
	SetMessage(string)
}

//--------

// Something to implement the MessageContainer interface
type InMemoryMessageContainer struct {
	message string
}

func NewInMemoryMessageContainer() MessageContainer {
	return new(InMemoryMessageContainer)
}

func (m *InMemoryMessageContainer) Init() {
	m.message = "Hello there"
}

func (m *InMemoryMessageContainer) Message() string {
	return m.message
}

func (m *InMemoryMessageContainer) SetMessage(message string) {
	m.message = message
}

//--------

// The thing where to inject the dependency
type MessagingApplication struct {
	Messager MessageContainer
}

func NewMessagingApplication() *MessagingApplication {
	m := new(MessagingApplication)

	if len(possibleMessagers) > 1 {
		contents, err := ioutil.ReadFile("config")
		if err != nil {
			return nil
		}
		test, ok := possibleMessagers[strings.TrimSpace(string(contents))]
		if !ok {
			return nil
		}

		m.Messager = test()
	} else if len(possibleMessagers) == 1 {
		for _, f := range possibleMessagers {
			m.Messager = f()
		}
	}
	m.Messager.Init()

	return m
}

func (m *MessagingApplication) printMessage() {
	fmt.Println(m.Messager.Message())
}

//--------

var possibleMessagers = map[string]func() MessageContainer{
	"inmemory": NewInMemoryMessageContainer,
}

// Käytännössä voidaan lisätä vaikka ulkoisista paketeista MessageContainer implementoivia asioita sekaan ja käyttää niitä tilalla.
func AddMessager(name string, constructor func() MessageContainer) {
	possibleMessagers[name] = constructor
}

//--------

func main() {
	hello := NewMessagingApplication()
	if hello != nil {
		hello.printMessage()
	}
}
