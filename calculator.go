package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Numeric interface {
	~int | ~float64
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

func divide[T Numeric](x T, y T) (T, error) {
	if y == 0 {
		return -1, fmt.Errorf("Cannot divide by zero")
	}
	return x / y, nil
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
func calculate(nums []float64, operations []string) (float64, error) {
	if len(nums) == 1 && len(operations) == 0 {
		return nums[0], nil
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
			res, err := divide(result, nums[i+1])
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
			http.Error(w, "Error: Invalid input", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, strconv.FormatFloat(add(x, y), 'f', -1, 64))
	})

	http.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
		x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

		if err1 != nil || err2 != nil {
			http.Error(w, "Error: Invalid input", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, strconv.FormatFloat(subtract(x, y), 'f', -1, 64))
	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

		if err1 != nil || err2 != nil {
			http.Error(w, "Error: Invalid input", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, strconv.FormatFloat(multiply(x, y), 'f', -1, 64))
	})

	http.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {
		x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

		if err1 != nil || err2 != nil {
			http.Error(w, "Error: Invalid input", http.StatusBadRequest)
			return
		}

		result, err := divide(x, y)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, strconv.FormatFloat(result, 'f', -1, 64))
	})

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		expression := r.URL.Query().Get("expression")
		nums, operations, err := format(expression)

		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
			return
		}

		result, err := calculate(nums, operations)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, strconv.FormatFloat(result, 'f', -1, 64))
	})

	http.ListenAndServe(":8080", nil)
}
