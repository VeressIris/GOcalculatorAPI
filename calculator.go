package main

import (
	"calculator/utils"
	"fmt"
	"net/http"
	"strconv"
)

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

func additionHandler(w http.ResponseWriter, r *http.Request) {
	x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

	if err1 != nil || err2 != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := utils.Response{Operation: fmt.Sprintf("%f+%f", x, y), Result: utils.Add(x, y)}
	utils.WriteJSONResponse(w, result)
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

	if err1 != nil || err2 != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := utils.Response{Operation: fmt.Sprintf("%f+%f", x, y), Result: utils.Subtract(x, y)}
	utils.WriteJSONResponse(w, result)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	x, err1 := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	y, err2 := strconv.ParseFloat(r.URL.Query().Get("y"), 64)

	if err1 != nil || err2 != nil {
		utils.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := utils.Response{Operation: fmt.Sprintf("%f+%f", x, y), Result: utils.Multiply(x, y)}
	utils.WriteJSONResponse(w, result)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
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
	utils.WriteJSONResponse(w, result)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Query().Get("expression")
	nums, operations, err := utils.Format(expression)

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
	utils.WriteJSONResponse(w, result)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! Welcome to the GoLang calculator API!")
	})

	http.HandleFunc("/add", additionHandler)
	http.HandleFunc("/subtract", subtractHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/divide", divideHandler)
	http.HandleFunc("/calculate", calculateHandler)

	http.ListenAndServe(":8080", nil)
}
