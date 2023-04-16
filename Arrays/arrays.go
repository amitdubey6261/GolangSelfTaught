package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1,2,3,4,5,6,7,8,9}
	array3 := [...]string{"amit" , " never " , "ends"}
	fmt.Println(array1 , array2 , array3)

	//Accessing multiple values
	fmt.Println(array1[1] , array2[2] , array3[0])

	//Intialising specific values
	arr1 := [5]int{1:10,2:40}

	fmt.Println(arr1)

	//for loop on array
	for i := 0 ; i<len(array1) ; i++ { //len function
		fmt.Println(array1[i])
	}
}