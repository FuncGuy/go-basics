package main

import "fmt"

func main() {
	p := NewPerson("Linus", "Trovalds")
	err := p.setTwitterHandler("@linus")
	handleError(err)
	println(p.TwitterHandler())
	println(p.fullName())
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}
}
