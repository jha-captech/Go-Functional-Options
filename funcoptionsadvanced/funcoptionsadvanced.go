package funcoptionsadvanced

import (
	"fmt"
	"log"

	"github.com/jha-captech/go-functional-options/util"
)

type Service struct {
	Required string
	OptionA  string
	OptionB  string
	OptionC  string
}

type Options func(*Service) error

func NewService(required string, options ...Options) (Service, error) {
	config := Service{
		Required: required,
		OptionA:  "default_a",
		OptionB:  "default_b",
		OptionC:  "default_c",
	}
	for _, option := range options {
		if err := option(&config); err != nil {
			return Service{}, fmt.Errorf("NewService: %w", err)
		}
	}
	return config, nil
}

func WithOptionA(a string) Options {
	return func(config *Service) error {
		config.OptionA = a
		return nil
	}
}

func WithOptionB(b string) Options {
	return func(config *Service) error {
		config.OptionB = b
		return nil
	}
}

func WithOptionC(c string) Options {
	return func(config *Service) error {
		config.OptionC = c
		return nil
	}
}

func Run() {
	service, err := NewService(
		"this is required",
		WithOptionA("new value for A"),
		WithOptionC("new value for c"),
	)
	if err != nil {
		log.Fatal(err)
	}

	util.PrintAsJSON(service)
	// Output:
	// {
	// 	"Required": "this is required",
	// 	"OptionA": "new value for A",
	// 	"OptionB": "default_b",
	// 	"OptionC": "new value for c"
	// }
}
