package main

import (
	"fmt"
	"os"
	"os/exec"
)

// color codes ref: https://gist.github.com/ik5/d8ecde700972d4378d87
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)
var boardData = map[int8]int8{
	14:28,
	13:8,
}

func main(){
	fmt.Printf(DebugColor,"Snakes and Ladders\n")
	fmt.Println("Value from dice: ",doDice())
	fmt.Println("Number board:")
	party := map[string]int8{"W":0,"X":12,"Y":4,"Z":3}
	fmt.Println(boardPrint(party))
	res:=""
	fmt.Scanln(&res)
	if(res!="exit"){
		// clear screen: https://stackoverflow.com/a/19290028
		cmd := exec.Command("cmd", "/c", "cls")
    	cmd.Stdout = os.Stdout
    	cmd.Run()
		// will run code again:
		main()
		// upgrade in login : indeed
	}
}

func doDice() int8{
	return 6 // I will return random number from here.
}
func boardPrint(party map[string]int8) string{
	s, b:="",""
	var i int8
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
				if  v, found :=boardData[i]; found {
					if v<i{
						fmt.Printf(ErrorColor,fmt.Sprintf("%5v",boardData[i]))
					}else{
						fmt.Printf(NoticeColor,fmt.Sprintf("%5v",boardData[i]))
					}
				}else{
					fmt.Printf("%5v",i)
				}
			}
			b=""
			if i % 10 == 0{
				fmt.Print("\n")
			}
	}
	return s
}