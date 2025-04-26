package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello: saludo a persona
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("nombre no puede estar vacio")
	}
	msg := fmt.Sprintf(randFormat(), name)
	return msg, nil
}

// hello para varias personas
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = msg
	}

	return messages, nil
}

// saludos aleatorios
func randFormat() string {
	formats := []string{
		"hola, %v! Bienvenido!",
		"Que copado, %v",
		"Saludo, gil.. %v",
	}

	return formats[rand.Intn((len(formats)))]
}
