package main
import "fmt"
func pow(x int, y int) int{
	if y == 1 {
		return x
	} else {
		return x * pow(x,y-1)
	}
}


func main() {
	x := 0
	y := 0
	z := 0
	fmt.Print("Enter the value of x:")
	fmt.Scanf("%d", &x)
	fmt.Print("Enter the value of y:") 
	fmt.Scanf("%d", &y)
	fmt.Print("Enter the value of z:")
	fmt.Scanf("%d", &z)
	fmt.Println(pow(x,y) % z)
}