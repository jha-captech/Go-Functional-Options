package util

import (
	"encoding/json"
	"fmt"
)

func PrintAsJSON[T any](data T) {
	JSONData, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(JSONData))
}
