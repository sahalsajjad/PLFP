package main 
import "fmt" 
func upperPow(n int) int{
	var i int
	for i:=0; n > 0; i++ {
		n=n/2
	}
	return i
}
func binary(n int) int{
	var bin[100] int
	var size int
	for i:= 0 ; n > 0; i++ {
		bin[i] =  n % 2 
		n = n / 2
		size = i
	}
	
	for n:= size ; n >=0; n-- {
		fmt.Printf("%d",bin[n])
	}
	fmt.Println("")
	return 0
}
func main() {
	var n int
	fmt.Print("Enter the value of n:")
	fmt.Scanf("%d\n",&n)
	binary(n)
}