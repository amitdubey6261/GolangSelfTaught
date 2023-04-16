package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin);
	fmt.Println(("Enter Number :"))

	a , _ := reader.ReadString('\n');
	fmt.Println("The Input is a :" , a)
	b , _ := reader.ReadString('\n');
	fmt.Println("The Input is b :" , b)
	c , _ := reader.ReadString('\n');
	fmt.Println("The Input is c :" , c)

	fmt.Println("The solution is :" , a + b + c )

	//Parsing In integer values

	vala , _ := strconv.ParseFloat(strings.TrimSpace(a) , 32);
	valb , _:= strconv.ParseFloat(strings.TrimSpace(b) , 32);

	fmt.Println("The sum a and b in number is :" , vala + valb )

}