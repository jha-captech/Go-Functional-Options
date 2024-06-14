# Go-Functional-Options

[Source](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

## Introduction

## The Problem

```go
type Service struct {
    OptionA string
    OptionB string
    OptionC string
}

func Run() {
    service := Service{
        OptionA: "new value for A",
        OptionB: "default_b",
        OptionC: "new value for c",
    }
    
    util.PrintAsJSON(service)
    // Output:
    // {
    // 	"OptionA": "new value for A",
    // 	"OptionB": "default_b",
    // 	"OptionC": "new value for c"
    // }
}

```

## Other Solutions

### Constructor
```go
type Service struct {
    OptionA string
    OptionB string
    OptionC string
}

func NewService(optionA, optionB, optionC string) Service {
    return Service{
        OptionA: optionA,
        OptionB: optionB,
        OptionC: optionC,
    }
}

func Run() {
    service := NewService(
        "new value for A",
        "new value for B",
        "new value for C",
    )
    
    util.PrintAsJSON(service)
    // Output:
    // {
    // 	"OptionA": "new value for A",
    // 	"OptionB": "new value for B",
    // 	"OptionC": "new value for C"
    // }
}
```

### Constructor with config struct
```go
type Service struct {
	OptionA string
	OptionB string
	OptionC string
}

type Options struct {
	OptA string
	OptB string
	OptC string
}

func NewService(options Options) Service {
	return Service{
		OptionA: options.OptA,
		OptionB: options.OptB,
		OptionC: options.OptC,
	}
}

func Run() {
	config := NewService(
		Options{
			OptA: "new value for A",
			OptB: "new value for B",
			OptC: "new value for C",
		},
	)

	util.PrintAsJSON(config)
	// Output:
	// {
	// 	"OptionA": "new value for A",
	// 	"OptionB": "new value for B",
	// 	"OptionC": "new value for C"
	// }
}

```

## Functional Options solution

```go
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

```

## Bonus: Functional Options with Required Parameters and Error Handling

```go
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

```