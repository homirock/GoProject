package main

import (
	"GoProject/package1"
	child "GoProject/package1/package1"
	"GoProject/package1/nested_package1"
	"fmt"
	_"github.com/jinzhu/gorm"

)

var(
	integers [10] int
	i int = j
    j int = f()
//var k int=5
     k int = 5
)
func f() int{
	return k+1
}
func init(){
	fmt.Println("version==>1")
}
func init(){
	fmt.Println("version==>2")
}
func init() {
	for i := 0; i < 10; i++ {
		integers[i] = i
	}
}


func main(){
	//version:="1.0.1"#this will modify the variable version
	var a int=2
	var b int=a
	var c int=b
	fmt.Println("value of a b c are",a,b,c)
	fmt.Println("value of i j k are",i,j,k)
	fmt.Println(package1.Morning)
	fmt.Println(nested_package1.Evening)
	fmt.Println("version==>",version)
	fmt.Println(integers)
	fmt.Println(child.Night)
}
