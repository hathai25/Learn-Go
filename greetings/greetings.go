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

func randomPeople() string {
	var peopleStore []string = []string{
		"Thai",
		"Khanh",
		"Trung",
		"Thinh",
		"Thang",
	}

	return peopleStore[rand.Intn(len(peopleStore))]
}
