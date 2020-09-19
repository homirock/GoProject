package main
import "fmt"


func main() {
	str1 := "My name"
	str2 := "is"
	str3 := "Partha"
	aNumber := 40
	isTrue := true
	string_length,err := fmt.Println(str1, str2, str3)
	if err==nil{
	fmt.Println("string Length:", string_length)
	fmt.Printf("Value of aNumber:%v\n", aNumber)
	fmt.Printf("Value of isTrue:%v\n", isTrue)
	fmt.Printf("Value of aNumber as float :%0.3f\n", float64(aNumber))
	myString := fmt.Sprintf("Data type: %T,%T,%T,%T,and %T",str1,str2,str3,aNumber,isTrue)
	fmt.Printf(myString)
}
}

