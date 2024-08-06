package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Numeric interface {
	~int | ~float32 | ~float64
}

func add[T Numeric](x T, y T) T {
	return x + y
}

func subtract[T Numeric](x T, y T) T {
	return x - y
}

func multiply[T Numeric](x T, y T) T {
	return x * y
}

func divide[T Numeric](x T, y T) T {
	return x / y
}

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
				return nil, nil, err
			}

			nums = append(nums, num)
			currentNum = ""
		} else {
			currentNum += string(chr)
		}
	}

	// process last number too
	num, _ := strconv.ParseFloat(currentNum, 64)
	nums = append(nums, num)

	// check if there is only one number
	if len(nums) == 1 && len(operations) == 0 {
		return nums, operations, nil
	}
	// check if there aren't enough operations
	if len(nums)-1 != len(operations) {
		return nil, nil, fmt.Errorf("Invalid expression")
	}

	return nums, operations, nil
}

// iterates over the operations and numbers and executes the operations
func calculate(nums []float64, operations []string) float64 {
	if len(nums) == 1 && len(operations) == 0 {
		return nums[0]
	}

	result := nums[0]
	for i := 0; i < len(operations); i++ {
		if operations[i] == "+" {
			result = add(result, nums[i+1])
		} else if operations[i] == "-" {
			result = subtract(result, nums[i+1])
		} else if operations[i] == "*" {
			result = multiply(result, nums[i+1])
		} else if operations[i] == "/" {
			result = divide(result, nums[i+1])
		}
	}

	return result
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! Welcome to the GoLang calculator API!")
	})

	// TODO: Add error handling for each endpoint
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		x, _ := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, _ := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
		fmt.Fprintf(w, strconv.FormatFloat(add(x, y), 'f', -1, 64))
	})

	http.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
		x, _ := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, _ := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
		fmt.Fprintf(w, strconv.FormatFloat(subtract(x, y), 'f', -1, 64))
	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		x, _ := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, _ := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
		fmt.Fprintf(w, strconv.FormatFloat(multiply(x, y), 'f', -1, 64))
	})

	http.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {
		x, _ := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, _ := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
		fmt.Fprintf(w, strconv.FormatFloat(divide(x, y), 'f', -1, 64))
	})

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		expression := r.URL.Query().Get("expression")
		nums, operations, err := format(expression)

		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, strconv.FormatFloat(calculate(nums, operations), 'f', -1, 64))
	})

	http.ListenAndServe(":8080", nil)
}
