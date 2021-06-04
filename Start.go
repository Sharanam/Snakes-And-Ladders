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
func main(){
	fmt.Printf(DebugColor,"Snakes and Ladders\n")
	fmt.Println("Value from dice: ",doDice())
	fmt.Println("Number board:")
	party := map[string]int8{"W":0,"X":12,"Y":4,"Z":3}
	fmt.Println(boardFormat(party))
}

func doDice() int8{
	return 6 // I will return random number from here.
}
func boardFormat(party map[string]int8) string{
	s, b:="",""
	var counter int8 =1
	for i:=0; i<10; i++{
		for j:=0;j<10;j++{
			if party["W"] == counter{
				b+="W" 
			} 
			if party["X"] == counter{
				b+="X"
			}
			if party["Y"] == counter{
				b+="Y"
			} 
			if party["Z"] == counter{
				b+="Z"
			} 
			if b!="" {
				s +=fmt.Sprintf("%5v",b)
			} else{
				s +=fmt.Sprintf("%5v",counter)
			}
			counter++
			b=""
		}
		s +=fmt.Sprint("\n")
	}
	return s
}
