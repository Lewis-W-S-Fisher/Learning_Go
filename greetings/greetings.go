package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func randomFormat() string{
	formats := []string{
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
    // Return a greeting that embeds the name in a message.
    // message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

