package greetings

import "fmt"

// hello returns a greeting for the named person.
func hello(name string) string {
	var msg = fmt.Sprintf("Hello %v!", name)
	return msg
}
