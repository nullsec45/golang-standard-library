package main

import (
	"fmt"
	"time"
)

func main(){
	var now time.Time=time.Now()
	fmt.Println(now)

	var utc time.Time=time.Date(2025, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2025-08-24 10:10:10"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error", err.Error())
	}else{
		fmt.Println(valueTime)
	}

	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 10
	var duration3 time.Duration = duration1 - duration2

	fmt.Printf("duration1 : %d\n", duration3)
}