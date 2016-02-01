package main

import "fmt"

// ArgumentError is thrown if number of arguments doesn't make sense
type ArgumentError int8

func (e ArgumentError) Error() string {
	return fmt.Sprintf("Argument number error: %v", int8(e))
}

// MethodError is thrown when unknown protocol is specified
type MethodError string

func (e MethodError) Error() string {
	return fmt.Sprintf("unknown method: %v != [https, ssh]", string(e))
}
