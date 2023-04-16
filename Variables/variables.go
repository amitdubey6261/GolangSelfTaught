package main

import "fmt"

func main() {
	str := "hello world i am amit learning go to win hackathon 2k23"

	// with type specify decided on the basis on I/P
	a := 10
	b := 20
	c := 30
	isloggedIn := false
	var number uint32 = 881992771
	var website string = "amit dubey is a very handsome boy"
	const IamConst string = "constant variable i am hahahahahaha"
	var num int // no garbage value init
	fmt.Println(a+b+c, str, isloggedIn, number, website, IamConst, num)
}
