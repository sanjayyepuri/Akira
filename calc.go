package number

import (
	"errors"
	"strconv"
	"strings"
)

func CalculateCommand(c string) (float64, error) {
	command, strErr := getCommand(c)
	if strErr != nil {
		return 0.0, strErr
	}

	op := command[1]
	x, y, err := parseArgs(command)
	if err != nil {
		return 0.0, err
	}

	switch op[0] {
	case '+':
		return plus(x, y), nil
	case '-':
		return minus(x, y), nil
	case '*':
		return times(x, y), nil
	case '/':
		return divides(x, y), nil
	default:
		return 0.0, errors.New("ERROR: Invalid operation")
	}
}

//parse array of 3 args
func parseArgs(c [3]string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

// get 3 args from string
func getCommand(message string) ([3]string, error) {
	var command [3]string
	c := strings.Split(message, " ")
	if len(c) < 3 {
		return command, errors.New("ERROR: some arguments not supplied")
	}

	command[0] = c[0]
	command[1] = c[1]
	command[2] = c[2]
	return command, nil
}

// float calculator

func plus(x, y float64) float64 {
	return x + y
}

func minus(x, y float64) float64 {
	return plus(x, -y)
}

func times(x, y float64) float64 {
	return x * y
}

func divides(x, y float64) float64 {
	return x / y
}
