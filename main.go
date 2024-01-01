/*  GO QUIZ GAME - Quiz built using GO
    Copyright (C) 2024 TAMIM KHAN

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.*/

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Welcome, input your name to continue\n")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v, enter your age to continue.\n", name)
	var age uint
	fmt.Scan(&age)
	if age >= 13 {
		fmt.Printf("You are old enough to play, %v. Best of luck.\n", name)
	} else {
		fmt.Println("You cannot play, you are too young. ")
		return
	}
	fmt.Printf("Who makes the RX6700?: \n1. Nvidia \n2. AMD\n")
	var answer uint
	fmt.Scan(&answer)
	var score int8 = 0

	if answer == 2 {
		fmt.Printf("Correct!\n")
		score = score + 1
	} else {
		fmt.Printf("wrong, good luck next time!\n")
		score = score - 1
	}

	fmt.Printf("How much RAM is recommended for gaming?: \n1. 8 \n2. 16 \n3. 12\n4. 24\n")
	fmt.Scan(&answer)

	if answer == 16 {
		fmt.Printf("Correct!\n")
		score = score + 1
	} else if answer == 24 {
		fmt.Printf("Correct!\n")
		score = score + 1
	} else {
		fmt.Printf("wrong, good luck next time!\n")
		score = score - 1
	}

	fmt.Printf("Which one is worse: \n1. GT710 \n2. GTX1080 \n3. GTX1660 \n4. GTX 4060\n")
	fmt.Scan(&answer)

	if answer == 1 {
		fmt.Printf("Correct!\n")
		score = score + 1
	} else {
		fmt.Printf("wrong, good luck next time!\n")
		score = score - 1
	}

	fmt.Printf("What does GPU stand for?: \n1. Game Processing Unit \n2. Graphics Printing Unit \n3. Graphics Procession Unit \n4. Graphics Processing Unit\n")
	fmt.Scan(&answer)

	if answer == 4 {
		fmt.Printf("Correct!\n")
		score = score + 1
	} else {
		fmt.Printf("wrong, good luck next time!\n")
		score = score - 1
	}

	fmt.Printf("Which one is better?: \n1. Console \n2. PC \n")
	fmt.Scan(&answer)

	if answer == 1 {
		score = score + 1
	} else if answer == 2 {
		score = score + 1
	}
	if score >= 3 {
		fmt.Println("Thanks for playing!")
	} else if score < 3 {
		fmt.Println("You have failed, better luck next time!")
	}

}
