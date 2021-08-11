package main

import "fmt"

func main(){
	b:=[...]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	fmt.Println(b)
	arrayLen:=len(b)
	for i:=0;i<=arrayLen-2;i++{
		for j:=0;j<=arrayLen-2;j++ {
			//fmt.Println(lengthCheck(b[i],b[j]))
			if lengthCheck(b[j],b[j+1])==true{
				temp:=b[j+1]
				b[j+1]=b[j]
				b[j]=temp
				
			}
			
		}
	}
	fmt.Println(b)
}

func lengthCheck(value string,value1 string) bool{
	var count int=0
	var count1 int=0
	length := len(value)
	length1 := len(value1)
		
			for i := 0; i < length; i++ {
				if string(value[i])=="a" {
					count++
				}
			}
		
			for i := 0; i < length1; i++ {
				if string(value1[i])=="a"{
					count1++
				}
			}
			if count>count1{
				return false
			} else if count<count1{
				return true
			}

			if count==count1{
				if length>=length1{
					return false
				}else{
					return true
				}
				
			}
			return true
}