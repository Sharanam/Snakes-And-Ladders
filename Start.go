package main

import (
	"fmt"
)

// color codes ref: https://gist.github.com/ik5/d8ecde700972d4378d87
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)
var boardData = map[int8]int8{}

func main(){
	fmt.Printf(DebugColor,"Snakes and Ladders\n")
	boardInit()
	fmt.Println("Value from dice: ",doDice())
	fmt.Println("Number board:")
	party := map[string]int8{"W":0,"X":12,"Y":4,"Z":3}
	fmt.Println(boardPrint(party))
}

func doDice() int8{
	return 6 // I will return random number from here.
}
func boardPrint(party map[string]int8) string{
	s, b:="",""
	var i int8
	var counter int8 =1
	for i:=0; i<10; i++{
		for j:=0;j<10;j++{
			if party["W"] == counter{
		for i=1; i<=100;i++{
			if party["W"] == i{
				b+="W" 
			} 
			if party["X"] == i{
				b+="X"
			}
			if party["Y"] == i{
				b+="Y"
			} 
			if party["Z"] == i{
				b+="Z"
			} 
			if b!="" {
				fmt.Printf(WarningColor,fmt.Sprintf("%5v",b))
			} else{
				if boardData[i]==i{
					fmt.Printf("%5v",i)
				}else if boardData[i]<i{
					fmt.Printf(ErrorColor,fmt.Sprintf("%5v",boardData[i]))
				}else{
					fmt.Printf(NoticeColor,fmt.Sprintf("%5v",boardData[i]))
				}
			}
			b=""
			if i % 10 == 0{
				fmt.Print("\n")
			}
	}
	return s
}
func boardInit(){
	var sl=map[int8]int8{
		14:28,
		13:8,
	}
	var i int8
	for i=1; i<=100; i++{
		if  v, found :=sl[i]; found {
			boardData[i]=v
		}else{
			boardData[i]=i
		}
	}
}