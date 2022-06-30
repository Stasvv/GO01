package main

import (
	"fmt"
	"strconv"
)

func main() {

	var stavka string
	var razmerStavki int
	var ruletka int

	//var color string

	//stavka = "5"
	//razmerStavki = 25
	ruletka = 5

	fmt.Scan(&stavka)

	fmt.Println(stavka)
	fmt.Println(razmerStavki)
	fmt.Println(ruletka)

	ruletkaString := strconv.Itoa(ruletka)
	//ruletkaString := "black"

	if ruletkaString != stavka {
		fmt.Println("game over")

	} else {

		//fmt.Println("else")
		if ruletkaString == "black" || ruletkaString == "red" {
			fmt.Println("black")

		} else {
			fmt.Println("number")
		}

	}

}
