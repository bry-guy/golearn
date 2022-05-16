package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name - who's there?")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hey, %v! Welcome!",
		"Sup %v, beers in the fridge.",
		"%v? THE %v?? You're shittin' me.",
	}

	return formats[rand.Intn(len(formats))]
}
