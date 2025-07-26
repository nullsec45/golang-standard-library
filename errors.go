package main

import "errors"
import "fmt"

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "fajar" {
		return NotFoundError
	}

	return nil
}

func main(){
	err := GetById("")

	if err != nil {
		if errors.Is(err, ValidationError){
			fmt.Println("validation error")
		}else if errors.Is(err, NotFoundError){
			fmt.Println("not found error")
		}else{
			fmt.Println("unknown error")
		}
	}
}