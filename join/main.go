package main

import (
	"fmt"

	"github.com/53jk1/automated-tests/prose"
)

func main() {
	phrases := []string{"moi rodzice", "klaun z rodeo"}
	fmt.Println("Zdjęcie, na którym są", prose.JoinWithCommas(phrases))
	phrases = []string{"moi rodzice", "klaun z rodeo", "nagrodzony byk"}
	fmt.Println("Zdjęcie, na którym są", prose.JoinWithCommas(phrases))
}
