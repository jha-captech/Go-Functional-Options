package funcoptions

import (
	"github.com/jha-captech/go-functional-options/util"
)

type Service struct {
	OptionA string
	OptionB string
	OptionC string
}

type Options func(*Service)

func NewService(options ...Options) Service {
	config := Service{
		OptionA: "default_a",
		OptionB: "default_b",
		OptionC: "default_c",
	}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithOptionA(a string) Options {
	return func(config *Service) {
		config.OptionA = a
	}
}

func WithOptionB(b string) Options {
	return func(config *Service) {
		config.OptionB = b
	}
}

func WithOptionC(c string) Options {
	return func(config *Service) {
		config.OptionC = c
	}
}

func Run() {
	service := NewService(
		WithOptionA("new value for A"),
		WithOptionC("new value for c"),
	)

	util.PrintAsJSON(service)
	// Output:
	// {
	// 	"OptionA": "new value for A",
	// 	"OptionB": "default_b",
	// 	"OptionC": "new value for c"
	// }
}
