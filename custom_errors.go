package main

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Name    string
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func YouShallNotPass() error {
	c := CustomError{}
	c.Name = "You Shall Not pass"
	c.Code = http.StatusUnauthorized
	c.Message = "YOU SHALL NOT PAASSSS"
	return c
}
