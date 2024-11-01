package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, world")
	// create a new file and that content in it
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// what is the data type of the file
	fmt.Printf("%T\n", file)
	// write data in the file
	content := "hello add contents in the file"
	bytes, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Number of bytes written", bytes)
	// good practice if you are not opening just add defer before file.close()
	file.Close()
	file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// read the file
	// 1st method

	buffer := make([]byte, 1024)
	data := make([]string, 6)
	for {
		n, err := file.Read(buffer)
		// yahan n ka mtlb kitna bytes file se buffer me ja rha h
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		data = append(data, string(buffer[:n]))

	}
	fmt.Println(data)
    // 2nd method
    
	info, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(info))
}

