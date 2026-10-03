package main

import "fmt"


func main(){
	i:=1
	for i <=5 {
		fmt.Println("shutup ahmed khaled")
		i += 1
	}
	
	for j:=0; j<=5; j++ {
		fmt.Println("shutup 3abkr")
	}

	for _ = range 5 { // Blank Identifier
		fmt.Println("shutup ahmed khaled")
	}
	
	fuck:=0
	for {
       	  	if fuck == 5 {	break	}
		fmt.Println("shutup 3abkr")
		fuck +=1 
	}
}