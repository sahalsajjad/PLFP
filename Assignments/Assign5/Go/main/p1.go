package main
import(
	"fmt"
)
func max( a int, b int) int{
	if( a >= b){
		return a
	}else{
		return b
	}
}
func main(){	
	var x1,x2,y1,y2 int
	
	fmt.Print("Enter the values of x and y for source:")
	fmt.Scanf("%d %d", &x1, &y1)
	fmt.Print("Enter the values of x and y for destination:")
	fmt.Scanf("%d %d", &x2, &y2)
	
	if( x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0 ){
		fmt.Println("Please provide Positive values")
	}else if(x1 > x2 || y1 > y2){
		fmt.Println("Path not possible")
	}else{
		
		var TAB [1000][1000]int
		
		for i:=1; i<(x2-x1+1); i++ {
			TAB[i][0] = 1
		}
		for j:=1; j<(y2-y1+1) ; j++ {
				TAB[0][j] = 1		
		}
		for i:=1; i<(x2-x1+1); i++ {
			for j:=1; j<(y2-y1+1) ; j++ {
					TAB[i][j] = TAB[i-1][j] + TAB[i][j-1]		
			}
		}
		fmt.Printf("%d", TAB[x2-x1][y2-y1])
	}
	
	
}
