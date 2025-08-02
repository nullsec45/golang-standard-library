package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Rama Fajar","Fajar"))
	fmt.Println(strings.Split("Rama Fajar"," "))
	fmt.Println(strings.ToLower("Rama Fajar"))
	fmt.Println(strings.ToUpper("Rama Fajar"))
	fmt.Println(strings.Trim("   Rama Fajar   "," "))
	fmt.Println(strings.ReplaceAll("Rama Fajar Ramadhan Tiba","Rama","Entong"))
}