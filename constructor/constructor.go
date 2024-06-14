package constructor

import "github.com/jha-captech/go-functional-options/util"

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
