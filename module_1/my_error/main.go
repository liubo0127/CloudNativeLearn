package main

import (
	"errors"
	"fmt"
)

func notFound1() error {
	return errors.New("errNotFound")
}

func notFound2() error {
	err := fmt.Errorf("errNotFound")
	return err
}

func main() {
	if err := notFound1(); err != nil {
		fmt.Printf("%T, %s\n", err, err)
	}
}
