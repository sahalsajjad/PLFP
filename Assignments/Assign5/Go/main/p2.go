package main
import ("fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	
)
func eval(s string) string{
	strings.Replace(s, " ","", -1)
	var val string;
	a , _:= strconv.Atoi(s[0:1])
	b, _:= strconv.Atoi(s[2:3])
	op := s[1:2]
	if( op == "+"){
		val = strconv.Itoa(a + b)
	}else if( op == "-"){
		val= strconv.Itoa(a - b)
	}else if( op == "*"){
		val = strconv.Itoa(a * b)
	}else if( op == "/"){
		if(b == 0){
			val = "ZeroDivisionError: Division is not possible"
		}else{
			val = strconv.Itoa(a/b)
		}
	}else if( op == "%"){
		if(b == 0){
			val = "ZeroDivisionError: Division is not possible"
		}else{
			val = strconv.Itoa(a % b)
		}
	}
	return val
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter the expression:")
	
	exp, _ := reader.ReadString('\n')
	
	fmt.Println(eval(exp))
}
