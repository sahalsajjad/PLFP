package main
import "fmt"
func printspaces(m int){
	for k:=0 ; k < m; k++ {
		fmt.Print(" ")
	}
}
func printstars(m int){ 
	for k:=0 ; k < m; k++ {
		fmt.Print("* ")
	}
}
func invpyramid(n int){
	for i:=n ; i >= 1 ; i-- {
		printspaces(n-i)
		printstars(i)
		printspaces(n-i) // actually not required, but for perfection sake, incase we replace space with something else.
		fmt.Println(" ")
	}
}
func main(){
	var n int
	fmt.Print("Enter the value of n:")
	fmt.Scanf("%d", &n)
	fmt.Println(n)
	// chosen over other alternatives of Scanning so as to stop scanning at a newline / EOF.
	invpyramid(n)
			
}