package main

import (
	"fmt"
)

// Environment represents the environment for variable bindings and function definitions.
type Environment struct {
	variables map[string]interface{}
}

// NewEnvironment creates a new environment.
func NewEnvironment() *Environment {
	return &Environment{
		variables: make(map[string]interface{}),
	}
}

// RunAdd evaluates addition operation.
func RunAdd(env *Environment, args []interface{}) interface{} {
	fmt.Printf("Args: %v\n", args) // Print the arguments for debugging
	left, leftOk := args[0].(int)
	right, rightOk := args[1].(int)
	if !leftOk || !rightOk {
		panic("Invalid arguments for addition")
	}
	return left + right
}

// EnvGet gets the value of an environment variable.
func EnvGet(env *Environment, name string) interface{} {
	value, ok := env.variables[name]
	if !ok {
		panic(fmt.Sprintf("Unknown variable: %s", name))
	}
	return value
}

// RunGet returns the stored value of a variable in the global environment.
func RunGet(env *Environment, args []interface{}) interface{} {
	name := args[0].(string)
	return EnvGet(env, name)
}

// EnvSet sets the value of an environment variable.
func EnvSet(env *Environment, name string, value interface{}) {
	env.variables[name] = value
}

// RunMul evaluates multiplication operation.
func RunMul(env *Environment, args []interface{}) interface{} {
	left := args[0].(int)
	right := args[1].(int)
	return left * right
}

// RunDecrement evaluates subtraction operation.
func RunDecrement(env *Environment, args []interface{}) interface{} {
	left := args[0].(int)
	right := args[1].(int)
	return left - right
}

// RunSeq evaluates a sequence of expressions.
func RunSeq(env *Environment, args []interface{}) interface{} {
	var result interface{}
	for _, expression := range args {
		result = Run(env, expression)
	}
	return result
}

// RunPrint prints the arguments.
func RunPrint(env *Environment, args []interface{}) interface{} {
	for _, arg := range args {
		fmt.Print(arg)
	}
	fmt.Println()
	return nil
}

// RunSet sets the value of a variable in the environment.
func RunSet(env *Environment, args []interface{}) interface{} {
	name := args[0].(string)
	value := Run(env, args[1]).(int)
	EnvSet(env, name, value)
	return value
}

// RunIf evaluates an if-else condition.
func RunIf(env *Environment, args []interface{}) interface{} {
	cond := args[0].(bool)
	ifTrue := args[1]
	ifFalse := args[2]
	if cond {
		return Run(env, ifTrue)
	}
	return Run(env, ifFalse)
}

// RunArray creates an array of specified length.
func RunArray(env *Environment, args []interface{}) interface{} {
	length := args[0].(int)
	return make([]interface{}, length)
}

// RunSetxArray sets the value at the specified index in the array.
func RunSetxArray(env *Environment, args []interface{}) interface{} {
	array := env.variables[args[0].(string)].([]interface{})
	index := args[1].(int)
	value := args[2].(int)
	array[index] = value
	return value
}

// RunGetxArray gets the value at the specified index from the array.
func RunGetxArray(env *Environment, args []interface{}) interface{} {
	array := env.variables[args[0].(string)].([]interface{})
	index := args[1].(int)
	return array[index]
}

// RunWhile evaluates a while loop.
func RunWhile(env *Environment, args []interface{}) interface{} {
	cond := args[0].(bool)
	expr := args[1]
	for cond {
		Run(env, expr)
	}
	return nil
}

// RunDef defines a new function.
func RunDef(env *Environment, args []interface{}) interface{} {
	name := args[0].(string)
	params := args[1].([]interface{})
	body := args[2]
	EnvSet(env, name, []interface{}{"func", params, body})
	return nil
}

// RunCall calls a function.
func RunCall(env *Environment, args []interface{}) interface{} {
	name := args[0].(string)
	funcDef := env.variables[name].([]interface{})
	params := funcDef[1].([]interface{})
	body := funcDef[2]
	funcEnv := NewEnvironment()
	for i, param := range params {
		funcEnv.variables[param.(string)] = args[i+1]
	}
	return Run(funcEnv, body)
}

func Run(env *Environment, expr interface{}) interface{} {
	switch v := expr.(type) {
	case int:
		return v
	case string:
		return EnvGet(env, v)
	case []interface{}:
		op := v[0].(string)
		args := v[1:]
		switch op {
		case "add":
			left := Run(env, args[0]).(int)
			right := Run(env, args[1]).(int)
			return left + right
		case "mul":
			left := Run(env, args[0]).(int)
			right := Run(env, args[1]).(int)
			return left * right
		case "get":
			return RunGet(env, args)
		case "seq":
			return RunSeq(env, args)
		case "print":
			return RunPrint(env, args)
		case "set":
			return RunSet(env, args)
		case "if":
			return RunIf(env, args)
		// Add cases for other Lisp operations
		default:
			panic(fmt.Sprintf("Unknown operation: %s", op))
		}
	default:
		panic(fmt.Sprintf("Invalid Lisp expression: %v", expr))
	}
}

func main() {
	// stuff := map[string]interface{}{
	// 	"reiko": 1,
	// 	"alex":  2,
	// }

	// Example Lisp programs
	program := []interface{}{
		"seq",
		[]interface{}{"set", "reiko", 1},
		[]interface{}{"set", "alex", 2},
		[]interface{}{"add", []interface{}{"get", "reiko"}, []interface{}{"mul", []interface{}{"get", "alex"}, 3}},
	}

	// Evaluate the program
	result := Run(NewEnvironment(), program)
	fmt.Println("Result:", result)
}
