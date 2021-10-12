package main

import (
	"fmt"
)

import (
	"bufio"
	"os"
	"strconv"
)

const (
	damierSize = 9
	playerX    = "X"
	playerO = "O"
)

var (
	tttArray = [damierSize] string{
		"1", "2", "3",
		"4", "5", "6",
		"7", "8", "9"}
	player1 = true
)

func main() {
	play()
}

func play() {
	var indexCase int
	for true {
		draw()
		indexCase = userEntry()
		fillCase(indexCase)
		if win() {
			draw()
			fmt.Println(playerName(), "vous avez gagné !")
			os.Exit(0)
		} else if nullGame() {
			draw()
			fmt.Println("Partie nulle !")
			os.Exit(0)
		}
		player1 = !player1
	}
}

func draw() {
	for i := 0; i < len(tttArray); i++ {
		fmt.Print(" ", tttArray[i], " ")
		if (i+1) %3 == 0 {
			fmt.Println()
		}
	}
}

func playerName() string {
	if player1 {
		return "Joueur 1 "
	} else {
		return "Joueur2 "
	}
}

func userEntry() int {

	var (
		goodEntry  = false
		caseNumber = 0
		err        error
		scanner     = bufio.NewScanner(os.Stdin)
	)

	for goodEntry == false {
		fmt.Print(playerName(), "entrez un nombre compris entre 1 à ", damierSize, " : ")
		scanner.Scan()
		caseNumber, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Entrez un nombre et non autre chose !")
		} else if caseNumber < 1 || caseNumber > damierSize {
			fmt.Println("Votre nombre doit être compris entre 0 à", damierSize, "!")
		} else if tttArray[caseNumber-1] == playerX || tttArray[caseNumber-1] == playerO {
			fmt.Println("Cette case est déjà prise !")
		} else {
			goodEntry = true
		}
	}
	return caseNumber - 1
}

func fillCase(numeroCase int) {
	if player1 {
		tttArray[numeroCase] = playerX
	} else {
		tttArray[numeroCase] = playerO
	}
}

func win() bool {
	arrayEntry := [][damierSize]bool{
		{
			true, true, true,
			false, false, false,
			false, false, false},

		{
			false, false, true,
			false, false, true,
			false, false, true},
		{
			false, false, false,
			false, false, false,
			true, true, true},
		{
			true, false, false,
			true, false, false,
			true, false, false},
		{
			true, false, false,
			false, true, false,
			false, false, true},
		{
			false, false, true,
			false, true, false,
			true, false, false},
		{
			false, true, false,
			false, true, false,
			false, true, false}}

	var tttArrayBool [damierSize]bool

	for index, valeur := range tttArray {
		if player1 && valeur == playerX {
			tttArrayBool[index] = true
		} else if !player1 && valeur == playerO {
			tttArrayBool[index] = true
		}
	}

	similarity := 0
	for _, tableauGain := range arrayEntry {
		for i := 0; i < len(tttArrayBool); i++ {
			if tttArrayBool[i] == true && tttArrayBool[i] == tableauGain[i] {
				similarity++
				if similarity == 3 {
					return true
				}
			}
		}
		similarity = 0
	}
	return false
}

func nullGame() bool {
	occurence := 0

	for _, valeur := range tttArray {
		if valeur == playerX || valeur == playerO {
			occurence++
		}
	}
	return (occurence == len(tttArray))
}
