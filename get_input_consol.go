package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

func main(){
	/*var s string
	fmt.Scanln(&s)
	fmt.Print(s)*/
	reader:=bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	str,_:=reader.ReadString('\n')
	fmt.Println(str)
	fmt.Print("Enter number:")
	str,_=reader.ReadString('\n')
	f,err:=strconv.ParseFloat(strings.TrimSpace(str),64)
	if err !=nil{
		fmt.Println(err)

	}else{
		fmt.Println("value of no:",f)
	}

}