package main

import "fmt"

func main() {

	var a[100][100] int
	var m int
	var n int
	var uprow=0
	var downrow=0
	var rightcol=0
	var colstrt=0 
	var row=0 
	var col=0 
	fmt.Println("Enter the no of rows: ")
	fmt.Scan(&m);
	fmt.Println("Enter the no of colomns: ")
	fmt.Scan(&n)
	value:= 1
	
	for {
		 a[row][col]=value
		 if(uprow==0 || rightcol==n-1){
			if(downrow<(m-1)){
				downrow=downrow+1
				colstrt=0
				row=downrow
				col=colstrt
				uprow=downrow
				rightcol=colstrt

			} else{
				if(colstrt<(n-1)) {
				colstrt=colstrt+1
				rightcol=colstrt
				col=colstrt
				uprow=downrow
				row=uprow
				}
			}

		} else{
			if(rightcol<(n-1)) {
				rightcol=rightcol+1
				uprow=uprow-1
				row=uprow
				col=rightcol
			}
		}
		/*printf("row and colomn value is %d %d\n",row,col);*/
		 a[row][col]=value+1
		 value++
		
		if row>=(m-1)&&col>=(n-1) {
			break
		}


	}
	for row:=0;row<m;row++ {
   		 for col:=0;col<n;col++ {
			fmt.Print(" \t ",a[row][col])
     		
     		}
     		fmt.Print("\n")
     		}

}
