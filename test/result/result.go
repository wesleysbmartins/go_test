package result_test

import "fmt"

const (
	red   = "\033[31m"
	green = "\033[32m"
	reset = "\033[0m"
)

func Error(msg string) {
	fmt.Println(red, msg, reset)
}

func Success(msg string) {
	fmt.Println(green, msg, reset)
}
