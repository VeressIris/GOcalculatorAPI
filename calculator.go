package main

import (
	"calculator/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var invalidExpressionError = fmt.Errorf("Invalid expression")

// returns the numbers and operations in the expression
func format(str string) ([]float64, []string, error) {
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

// iterates over the operations and numbers and executes the operations
func calculate(nums []float64, operations []string) (float64, error) {
	if len(nums) == 1 && len(operations) == 0 {
		return nums[0], nil
	}

	result := nums[0]
	for i := 0; i < len(operations); i++ {
		if operations[i] == "+" {
			result = utils.Add(result, nums[i+1])
		} else if operations[i] == "-" {
			result = utils.Subtract(result, nums[i+1])
		} else if operations[i] == "*" {
			result = utils.Multiply(result, nums[i+1])
		} else if operations[i] == "/" {
			res, err := utils.Divide(result, nums[i+1])
			if err != nil {
				return -1, err
			}
			result = res
		}
	}

	return result, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! Welcome to the GoLang calculator API!")
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

		if err1 != nil || err2 != nil {
			utils.WriteError(w, "Invalid input", http.StatusBadRequest)
			return
		}

		result := utils.Response{Operation: fmt.Sprintf("%f+%f", x, y), Result: utils.Add(x, y)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
		x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

		if err1 != nil || err2 != nil {
			utils.WriteError(w, "Invalid input", http.StatusBadRequest)
			return
		}

		result := utils.Response{Operation: fmt.Sprintf("%f+%f", x, y), Result: utils.Subtract(x, y)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

		if err1 != nil || err2 != nil {
			utils.WriteError(w, "Invalid input", http.StatusBadRequest)
			return
		}

		result := utils.Response{Operation: fmt.Sprintf("%f+%f", x, y), Result: utils.Multiply(x, y)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {
		x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

		if err1 != nil || err2 != nil {
			utils.WriteError(w, "Invalid input", http.StatusBadRequest)
			return
		}

		res, err := utils.Divide(x, y)
		if err != nil {
			utils.WriteError(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		result := utils.Response{Operation: fmt.Sprintf("%f+%f", x, y), Result: res}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		expression := r.URL.Query().Get("expression")
		nums, operations, err := format(expression)

		if err != nil {
			utils.WriteError(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		res, err := calculate(nums, operations)
		if err != nil {
			utils.WriteError(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		result := utils.Response{Operation: expression, Result: res}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}
