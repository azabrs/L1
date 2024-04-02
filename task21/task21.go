package main

import "fmt"
// interface of adapter classes
type Target interface {
    Operation()
}

// adaptable class
type Adaptee struct {
}

// A method of an adaptable class that needs to be called somewhere
func (adaptee *Adaptee) AdaptedOperation() {
    fmt.Println("I am AdaptedOperation()")
}

// class of a specific adapter
type ConcreteAdapter struct{
    *Adaptee
}

// implementation of an interface method that implements a call to an adaptable class
func (adapter *ConcreteAdapter) Operation() {
    adapter.AdaptedOperation()
}

// adapter constructor
func NewAdapter(adaptee *Adaptee) Target {
    return &ConcreteAdapter{adaptee}
}

func main() {
    fmt.Println("Adapter demo:")
    adapter := NewAdapter(&Adaptee{})
    adapter.Operation()
}