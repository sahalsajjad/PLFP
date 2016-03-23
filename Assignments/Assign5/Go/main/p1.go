package main
import "fmt"
func pathfind(srcx int, srcy int, dstx int, dsty int) int{
	/* End upon reaching the destination */	
	if(srcx == dstx && srcy == dstx ){
		return 1
	}else if(srcx == dstx){
		/* Traverse only right in case rows are exhaustes */
		return pathfind(srcx, srcy+1, dstx, dsty)	
	}else if(srcy == dsty ){
		/* Traverse only down in case columns are exhausted */
		return pathfind(srcx+1, srcy, dstx, dsty)	
	}else if(srcx < dstx && srcy < dsty ){ 
		return pathfind(srcx+1, srcy, dstx, dsty) + pathfind(srcx, srcy+1, dstx, dsty)
	}else{
		return 0
	}
		
}
func main(){
	var x1,x2,y1,y2 int
	fmt.Print("Enter the values of x and y for source:")
	fmt.Scanf("%d %d", &x1, &y1)
	fmt.Print("Enter the values of x and y for destination:")
	fmt.Scanf("%d %d", &x2, &y2)
	fmt.Println(pathfind(x1,y1,x2,y2))
}
