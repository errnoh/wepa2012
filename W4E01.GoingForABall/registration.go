// Nimeäminen rupeaa olemaan typerää koska tämä jo vähitellen kuuluisi olla omassa paketissaan :)

package main

import (
	"errors"
	"strings"
)

var (
	nameError         = errors.New("Length of name should be between 4 and 30.")
	addressError      = errors.New("Length of address should be between 4 and 50.")
	emailError        = errors.New("Email should contain a @-character.")
	alreadyThereError = errors.New("Name already registered.")
)

var registrations = make(map[string]*Registration)

type Registration struct {
	Name    string
	Address string
	Email   string
}

func (r *Registration) Equals(r2 *Registration) bool {
	if r.Name == r2.Name && r.Address == r2.Address && r.Email == r2.Email {
		return true
	}
	return false
}

func AddRegistration(r *Registration) {
	registrations[r.Name] = r
}

func RemoveRegistration(r *Registration) bool {
	if !ContainsRegistration(r) {
		return false
	}
	delete(registrations, r.Name)
	return true
}

// Mmh.. pitäisi varmaan käyttää jotain tunnistettavaa hashia :| O(n) hashmap on huono idea :)
func ContainsRegistration(r *Registration) bool {
	_, ok := registrations[r.Name]
	return ok
}

func (r *Registration) NameIsOk() bool {
	if len(r.Name) < 4 || len(r.Name) > 30 {
		return false
	}
	return true
}

func (r *Registration) AddressIsOk() bool {
	if len(r.Address) < 4 || len(r.Address) > 50 {
		return false
	}
	return true
}

func (r *Registration) EmailIsOk() bool {
	return strings.Contains(r.Email, "@")
}

// errlist == nil if no errors
func (r *Registration) Check() (errlist []error) {
	if !r.NameIsOk() {
		errlist = append(errlist, nameError)
	}
	if !r.AddressIsOk() {
		errlist = append(errlist, addressError)
	}
	if !r.EmailIsOk() {
		errlist = append(errlist, emailError)
	}
	return errlist
}
