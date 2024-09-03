package main

import "fmt"

func main(){
	age := 23 // regular 
	agePointer := &age //pointer 

	fmt.Println("pointer",*agePointer) //value behind the ptr
	fmt.Println("Age",age) 

	adultYears := getAdultYears(agePointer) //delete the copy after executing this
	fmt.Println("Adult years ",adultYears)
}

func getAdultYears(age *int) int { //  copy value 
 return *age - 18
}