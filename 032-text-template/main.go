package main

import {
	"text/template"
	"log"
	"os"
}

func main() {
	tpl, err := template.ParseFiles("one.txt", "two.txt")
}