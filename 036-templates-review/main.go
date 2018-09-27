package main

import (
	"html/template"
)
var tpl *template.Template

func init() {
tpl = template.must(template.ParseGlob)

}

func main() {

}

