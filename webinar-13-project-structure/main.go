package main

import (
	"fmt"
	"os"
	"strings"
)

// Створити тип студента з можливістю декорування як відмінник, спортсмен і подібне. Студент має оцінку та імʼя, які змінюються в залежності від декоратора

func main() {
	byteContent, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	content := string(byteContent)

	var printer ContentPrinter = basicContentPrinter{}

	var answer string

	fmt.Println("Add lowercasing?")
	fmt.Scanln(&answer)

	if answer == "y" {
		printer = lowerCaseContentPrinter{
			parent: printer,
		}
	}

	fmt.Println("Add new liner?")
	fmt.Scanln(&answer)

	if answer == "y" {
		printer = newLineContentPrinter{
			parent: printer,
		}
	}

	fmt.Println("Add space removal?")
	fmt.Scanln(&answer)

	if answer == "y" {
		printer = noSpacesContentPrinter{
			parent: printer,
		}
	}

	printer.PrintContent(content)
}

type newLineContentPrinter struct {
	parent ContentPrinter
}

func (p newLineContentPrinter) PrintContent(c string) {
	p.parent.PrintContent(c + "\n")
}

type noSpacesContentPrinter struct {
	parent ContentPrinter
}

func (p noSpacesContentPrinter) PrintContent(c string) {
	p.parent.PrintContent(strings.TrimSpace(c))
}

type lowerCaseContentPrinter struct {
	parent ContentPrinter
}

func (p lowerCaseContentPrinter) PrintContent(c string) {
	p.parent.PrintContent(strings.ToLower(c))
}

type ContentPrinter interface {
	PrintContent(string)
}

type basicContentPrinter struct{}

func (p basicContentPrinter) PrintContent(c string) {
	fmt.Print(c)
}
