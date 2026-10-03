package main

import "fmt"

func main(){
	var a = "initial"
	var b,c int = 1,-3
	var d bool = true
	var f string="s"
	var e int 
	var x string
	n:= 0.65
	
	fmt.Println(a,b,c,d,f,x,e,n)

	// Formatted multi-line print using Printf
	fmt.Println("--- Variable Overview ---")
	fmt.Printf("a (string):  %q\n", a)
	fmt.Printf("b (int):     %d\n", b)
	fmt.Printf("c (int):     %d\n", c)
	fmt.Printf("d (bool):    %t\n", d)
	fmt.Printf("f (string):  %q\n", f)
	fmt.Printf("x (string):  %q (zero-value: empty string)\n", x)
	fmt.Printf("e (int):     %d  (zero-value)\n", e)
	fmt.Printf("n (float64): %.2f\n", n)

}
