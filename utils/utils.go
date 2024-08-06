package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Numeric interface {
	~int | ~float64
}

func Add[T Numeric](x T, y T) T {
	return x + y
}

func Subtract[T Numeric](x T, y T) T {
	return x - y
}

func Multiply[T Numeric](x T, y T) T {
	return x * y
}

func Divide[T Numeric](x T, y T) (T, error) {
	if y == 0 {
		return -1, fmt.Errorf("Cannot divide by zero")
	}
	return x / y, nil
}

type Response struct {
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

var invalidExpressionError = fmt.Errorf("Invalid expression")

// returns the numbers and operations in the expression
func Format(str string) ([]float64, []string, error) {
	if len(str) == 0 {
		return nil, nil, fmt.Errorf("Empty expression")
	}

	str = strings.ReplaceAll(str, " ", "")

	operations := []string{}
	nums := []float64{}

	currentNum := ""
	for _, chr := range str {
		if chr == '+' || chr == '-' || chr == '*' || chr == '/' {
			operations = append(operations, string(chr))

			num, err := strconv.ParseFloat(currentNum, 64)
			if err != nil {
				return nil, nil, invalidExpressionError
			}
			nums = append(nums, num)

			currentNum = ""
		} else {
			currentNum += string(chr)
		}
	}

	// process last number too
	num, err := strconv.ParseFloat(currentNum, 64)
	if err != nil {
		return nil, nil, invalidExpressionError
	}
	nums = append(nums, num)

	// check if there aren't enough operations
	if len(nums)-1 != len(operations) {
		return nil, nil, invalidExpressionError
	}

	return nums, operations, nil
}
