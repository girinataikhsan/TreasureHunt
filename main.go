package main

import "fmt"

func main() {
	startX := 4
	startY := 1
	arrayGrid := [6][8]string{}
	for o,data := range arrayGrid{
		for i := range data{
			if o ==0 || o== len(arrayGrid)-1{
				arrayGrid[o][i] ="#"
				fmt.Print("#")
			}else{
				if i==0 || i== len(data)-1{
					arrayGrid[o][i] ="#"
					fmt.Print("#")
				}else if o==2 && i > 1 && i<5{
					arrayGrid[o][i] ="#"
					fmt.Print("#")
				}else if o==3 && (i == 4 || i==6){
					arrayGrid[o][i] ="#"
					fmt.Print("#")
				}else if o==4 && i == 2{
					arrayGrid[o][i] ="#"
					fmt.Print("#")
				}else if o==startX && i==startY{
					arrayGrid[o][i] ="X"
					fmt.Print("X")
				}else{
					arrayGrid[o][i] ="."
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}

	var answer [][]int
	for checkUp:=startX-1; checkUp>0;checkUp-=1{
		if arrayGrid[checkUp][startY]=="."{
			for check:= startY+1; check<7; check+=1{
				if arrayGrid[checkUp][check]=="."{
					for checkAns:=checkUp+1; checkAns<5; checkAns+=1{
						if arrayGrid[checkAns][check]=="."{
							answer = append(answer, []int{checkAns, check})
							arrayGrid[checkAns][check] ="$"
						}else{
							break
						}
					}
				}else{
					break
				}
			}
		}
	}

	fmt.Println("Result :")
	fmt.Println(fmt.Sprintf("Probability coordinate treasure : %v", answer))
	fmt.Println("Grid with all probability :")
	for _,data := range arrayGrid{
		fmt.Println(data)
	}
}
