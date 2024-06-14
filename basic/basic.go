package basic

import (
	"github.com/jha-captech/go-functional-options/util"
)

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
