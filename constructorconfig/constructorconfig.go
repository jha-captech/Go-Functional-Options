package constructorconfig

import "github.com/jha-captech/go-functional-options/util"

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
