// W2E03
//
// Beans.
// Lyhyt versio: Beaneissa ei suuremmin ole järkeä Go:n puolella, Go:n interfacet hoitavat kaiken olennaisen.
// Ja voi tosiaan tyydyttää niin monta interfacea kuin haluaa, kunhan pitää sisällään sen interfacen vaatimat metodit.
package main

import (
	"fmt"
)

type Messager interface {
	Init()
	Message() string
	SetMessage(string)
}

type HelloMessage struct {
	message string
}

func (m *HelloMessage) Init() {
	m.message = "My Cool Property"
}

func (m *HelloMessage) Message() string {
	return m.message
}

func (m *HelloMessage) SetMessage(message string) {
	m.message = message
}

func (m *HelloMessage) MethodThatsNotInTheInferface() {
	m.message = "Oh no you didn't"
}

func ThingThatOnlyAcceptsMessagers(m Messager) {
	fmt.Println(m.Message())

	// m.MethodThatsNotInTheInferface() // wouldn't work.
}

func main() {
	hello := new(HelloMessage)
	hello.Init()

	ThingThatOnlyAcceptsMessagers(hello)
}
