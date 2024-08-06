package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
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