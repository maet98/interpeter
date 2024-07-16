package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Operation struct {
    first int
    second int
    operation string
    start int
    end int
}

func getFirstOperation(value string, operation string) (Operation, error) {
    operators := []string {
        "*",
        "/",
        "+",
        "-",
    }
    tokens := strings.Fields(value)
    for _, operator := range operators {
        currentPosition := 0
        for i := 0;i < len(tokens); i++ {
            if operator == tokens[i] {
                first, _ := strconv.Atoi(tokens[i-1])
                second, _ := strconv.Atoi(tokens[i+1])
                start := currentPosition - len(tokens[i-1])
                end := currentPosition + len(tokens[i-1]) + len(tokens[i]) + len(tokens[i+1]) + 2
                
                operation := Operation{
                    first: first,
                    second: second,
                    operation: tokens[i],
                    start: start,
                    end: end,
                }
                return operation, nil
            }
            currentPosition += len(tokens[i])
        }
    }
    return Operation{}, errors.New("Couldn't find any valid operation")
}


func handle(value string) int {
    tokens := strings.Fields(value)
    if len(tokens) == 1 {
        response, _ := strconv.Atoi(tokens[0])
        return response
    }
    first, _ := strconv.Atoi(tokens[0]) 
    operation := tokens[1]
    second, _ := strconv.Atoi(tokens[2])
    op := Operation{
        first: first,
        second: second,
        operation: operation,
        start: 0,
        end: 10,
    }
    response := op.handle()

    limit := len(tokens[0]) + len(tokens[1]) + len(tokens[2]) + 3
    valueWithoutOperations := ""
    if limit < len(value) {
        valueWithoutOperations = value[limit:]
    }
    newValue := fmt.Sprintf("%d %s", response, valueWithoutOperations)

    return handle(newValue)
}

func (op Operation) handle() int {
    switch op.operation {
    case "+":
        return op.first + op.second
    case "-":
        return op.first - op.second
    case "*":
        return op.first * op.second
    case "/":
        return op.first / op.second
    }
    return 0
}

func main () {
    fmt.Println(handle("4 * 2 - 1"))
}
