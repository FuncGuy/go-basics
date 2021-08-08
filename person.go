package main

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler string
}

func NewPerson(firstName, lastName string) Person {
	return Person {
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p Person) fullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p Person) setTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}