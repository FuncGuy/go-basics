package main

import "fmt"

func main() {

	 p:= NewPerson("Linus", "Trovalds")
	 err := p.setTwitterHandler("@linus")
	 if err != nil {
	 	fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	 }

	 println(p.fullName())
}
