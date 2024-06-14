package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Config struct {
	Required string
	OptionA  string
	OptionB  string
	OptionC  string
}

type Options func(*Config) error

func newConfig(required string, options ...Options) (Config, error) {
	config := Config{
		Required: required,
		OptionA:  "default_a",
		OptionB:  "default_b",
		OptionC:  "default_c",
	}
	for _, option := range options {
		if err := option(&config); err != nil {
			return Config{}, fmt.Errorf("newConfig: %w", err)
		}
	}
	return config, nil
}

func withOptionA(value string) Options {
	return func(config *Config) error {
		config.OptionA = value
		return nil
	}
}

func withOptionB(value string) Options {
	return func(config *Config) error {
		config.OptionB = value
		return nil
	}
}

func withOptionC(value string) Options {
	return func(config *Config) error {
		config.OptionC = value
		return nil
	}
}

func printAsJSON[T any](data T) {
	JSONData, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(JSONData))
}

func main() {
	config1, err := newConfig(
		"this is required",
		withOptionA("new value for A"),
		withOptionC("new value for c"),
	)
	if err != nil {
		log.Fatal(err)
	}

	printAsJSON(config1)
}
