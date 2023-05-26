package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	var message string = fmt.Sprintf(randomFormat(), randomPeople())
	return message, nil
}

func Hellos() (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range randomPeople() {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

func randomPeople() []string {
	var peopleStore []string = []string{
		"Thai",
		"Khanh",
		"Trung",
		"Thinh",
		"Thang",
	}

	return peopleStore
}
