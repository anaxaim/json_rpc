package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errIncorrectName = errors.New("incorrect name")
)

type GreetingService struct{}

type GreetingArgs struct {
	Name string `json:"name"`
}

func (s *GreetingService) Greeting(r *http.Request, args *GreetingArgs, reply *string) error {
	if args.Name == "" {
		return errIncorrectName
	}
	*reply = fmt.Sprintf("Hello, %s", args.Name)

	return nil
}
