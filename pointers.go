package main

import "fmt"

func main(){
	age := 23 // regular 
	agePointer := &age //pointer 

	fmt.Println("pointer",*agePointer) //value behind the ptr
	fmt.Println("Age",age) 

	 getAdultYears(agePointer) //delete the copy after executing this
	fmt.Println("Adult years ",age) //overrides the value and return as 5 
}

func getAdultYears(age *int)  { //  copy value 
//  return *age - 18
*age = *age-18
}