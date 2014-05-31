package main

import (
	"html/template"
	"regexp"
)

/*
Derived configuration globals.

  * Built from a Settings instance.
  * Instantiated as a single global
*/
type Configuration struct {
	Templates        *template.Template
	ValidPath        *regexp.Regexp
}

func (c *Configuration) init(s Settings) {
	c.Templates = template.Must(template.ParseGlob(s.TemplatesPath + "/*"))
	c.ValidPath = regexp.MustCompile("^/((edit|save|view)/([_a-zA-Z0-9]+))?$")
}
