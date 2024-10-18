package main

import (
	"bufio"
    "fmt"
    "os"
    "strings"
)
// taking input from user
func getinput(input string){
	
	fmt.Println("Enter the input string:");
	fmt.Scan(&input);
	// it reads the character upto first white space
	fmt.Println("The input string is:",input);
}
func getinput2()(string){
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the input string: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input) // Trim the newline character
    fmt.Println("The input string is:", input)
    return input;
	
}
func main() {
	// difference between printf and println
	// printf is used for printing formatted strings
	// while println ius used for nomal strings
	// it inserts a endl and tabspace
	var height int = 10
	var age int = 15
	var name string = "Abhi"
	fmt.Println("My name is",name, "Height is",height, "Age is", age)
	// My name is Abhi Height is 10 Age is 15
	//whitespace after is
	fmt.Println("Now println part is done");

	//printf
	fmt.Printf("My name is %s\t age is %d and height is %d\n",name,age,height);
  
    // getinput(input);
	input:=getinput2();
	fmt.Println("The input string is:",input);
	


}
