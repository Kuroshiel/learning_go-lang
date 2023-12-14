package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "eko" {
		return &notFoundError{"Data Not Found"}
	}

	return nil // ok
}

func main() {
	err := SaveData("budi", nil)
	if err != nil {
		// terjadi error
		/* if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation eror :", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error :", notFoundErr.Error())
		} else {
			fmt.Println("unknown error", err.Error())
		} */

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation eror :", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error :", finalError.Error())
		default:
			fmt.Println("unknown error", finalError.Error())
		}

	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
