package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
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

// variables are being used throughout game.
var players = 0
var party = map[string]int8{"W":0,"X":0,"Y":0,"Z":0}
var temp = [4]string{"W", "X", "Y", "Z"}
var response=""

func clearScreen(){
		// clear screen: https://stackoverflow.com/a/19290028
		cmd := exec.Command("cmd", "/c", "cls")
    	cmd.Stdout = os.Stdout
    	cmd.Run()
	// will print title everytime it clears console.
	fmt.Printf(DebugColor,"Snakes and Ladders\n")
}

func main(){
	clearScreen()
	fmt.Printf("Enter number of players (max 4):")
	getTotalPlayers() // it will set total number of players.
	round() // init game
}

func getTotalPlayers(){
	fmt.Scan(&players)
	if(players>4 || players<1 ){
		fmt.Println("Minimum 1 & Maximum 4 player(s)... enter number again: ")
		getTotalPlayers()
	}
	clearScreen()
	fmt.Println("Total number of players are:", players)
	fmt.Printf(DebugColor,fmt.Sprintln("You can leave game using text command."))
	fmt.Scanln() // IDK why this is being used here, but it solves a bug.
}

func round(){
	var playerAvailable bool = false
	var dice int8 = 0
	for i:=0;i<players;i++{
		if party[temp[i]]==100{
			fmt.Printf(WarningColor,fmt.Sprintln(temp[i], "cleared the game."))
			continue;
		}
		fmt.Println("Now, Dice is in hands of :", fmt.Sprintf(ErrorColor,fmt.Sprintf( temp[i]) ))
		fmt.Println("Press enter to roll the dice, Dear", temp[i])
		fmt.Scanln(&response)
		clearScreen()
		switch strings.Trim(strings.ToLower(response)," ") {
			case "exit", "close","leave", "c", "x", "e", "zz", "f4":
				exitGame()
		}
		dice=doDice()
		fmt.Println(temp[i], "achieved",dice, "by rolling out dice.")
		newValue:=party[temp[i]]+dice
		if newValue<100{
			if  _, found :=boardData[newValue]; found {
				party[temp[i]]=boardData[newValue]
				fmt.Printf(InfoColor,fmt.Sprintln(temp[i], "was landed on",newValue, "Thus, moved to", party[temp[i]]))
			}else{
				party[temp[i]]=newValue
			}
		}else if newValue==100{
			fmt.Printf(WarningColor,fmt.Sprintln(temp[i],"cleared the game."))
			party[temp[i]]=100
		}else{
			fmt.Println("Try again in next round, dear", temp[i])
		}
		fmt.Println(boardPrint())
	}
	for i:=0;i<players;i++{
		if party[temp[i]]!=100{
			playerAvailable=true
			break
		}
	}
	if playerAvailable{
		round()
	}else{
		exitGame()
	}
	}

func exitGame(){
	fmt.Println("Thank you! Visit again :)")
	os.Exit(0)
}

func doDice() int8{
	return int8(rand.Intn(6)+1)
}
func boardPrint() string{
	s, b:="",""
	var i int8
		for i=1; i<=100;i++{
		if  v, found :=boardData[i]; found {
			if v<i{
				fmt.Printf(ErrorColor,fmt.Sprintf("%5v",boardData[i]))
			}else{
				fmt.Printf(NoticeColor,fmt.Sprintf("%5v",boardData[i]))
			}
		}else{
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
					fmt.Printf("%5v",i)
				}
			b=""
		}
			if i % 10 == 0{
				fmt.Print("\n")
			}
	}
	return s
}