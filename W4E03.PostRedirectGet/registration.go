package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	nameError    = errors.New("Length of name should be between 4 and 30.")
	addressError = errors.New("Length of address should be between 4 and 50.")
	orderError   = errors.New("Order was empty.")
)

var orders = make(map[string]*Order)

var id int // Käytetään nyt vain kasvavaa inttiä tällä kertaa, tietoturvallista jne. Eikä edes isketä mutexia. :)

type Order struct {
	Id      string `schema:"-"` // Estää täyttämisen lähettämällä custom formidataa
	Name    string
	Address string
	Items   []string
}

func NewOrder() *Order {
	id++
	return &Order{Id: fmt.Sprint(id)}
}

func (r *Order) Equals(r2 *Order) bool {
	if r.Name == r2.Name && r.Address == r2.Address {
		return true
	}
	return false
}

func (o *Order) ParseItems(r *http.Request) {
	r.ParseForm()
	for _, v := range r.Form {
		if strings.HasPrefix(v[0], "item:") {
			temp := strings.SplitAfterN(v[0], ":", 2)
			o.Items = append(o.Items, temp[1])
		}
	}
}

func AddOrder(r *Order) {
	orders[r.Id] = r
}

func RemoveOrder(r *Order) bool {
	if !ContainsOrder(r) {
		return false
	}
	delete(orders, r.Name)
	return true
}

func ContainsOrder(r *Order) bool {
	_, ok := orders[r.Id]
	return ok
}

func (r *Order) NameIsOk() bool {
	if len(r.Name) < 4 || len(r.Name) > 30 {
		return false
	}
	return true
}

func (r *Order) AddressIsOk() bool {
	if len(r.Address) < 4 || len(r.Address) > 50 {
		return false
	}
	return true
}

func (r *Order) OrderIsOk() bool {
	if len(r.Items) == 0 {
		return false
	}
	return true
}

// errlist == nil if no errors
func (r *Order) Check() (errlist []error) {
	if !r.NameIsOk() {
		errlist = append(errlist, nameError)
	}
	if !r.AddressIsOk() {
		errlist = append(errlist, addressError)
	}
	if !r.OrderIsOk() {
		errlist = append(errlist, orderError)
	}
	return errlist
}
