package quiver_se

import "fmt"

func Greet(name string) (greeting string) {
	greeting = fmt.Sprintf("Hello, %v!", name)
	return
}

