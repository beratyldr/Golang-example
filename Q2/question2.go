package main 

import "fmt"
	
var (
	inpt int
	count int
	sum int
	a int
)

func  main()  {
	var div int=2
	var carpim int=1
	fmt.Scanln(&inpt)
	Recursivefunc(inpt,div,carpim)
}

func Recursivefunc(value int,div int,carpim int){
		
	if div<inpt{
		if value%div==0{
			 a=value/div
			 sum=sum+a
			 count++
			 carpim=div*carpim
			 Recursivefunc(a,div,carpim)	
		}
		if value%div!=0{
			div++
			Recursivefunc(value,div,carpim)
		}
		if div>=inpt {
			fmt.Println(count,sum,carpim)	
		} 
	
	}
	
	
}

