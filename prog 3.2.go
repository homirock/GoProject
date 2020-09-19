package main
import(
	"fmt"
	"strings"

	//"strings"
)
func main(){
	str1:="An implicitly typed string"
	fmt.Printf("String %v is %T type\n",str1,str1)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

}
