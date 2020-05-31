package main

import (
	"fmt"
	"math/rand"
	"time"
)

func (b *Board) generateBoard(difficulty string) {
	board := [9][9]int{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			board[i][j] = 0
		}
	}

	difficultyMap := make(map[string]int)
	difficultyMap["easy"] = 50
	difficultyMap["medium"] = 60
	difficultyMap["hard"] = 65
	difficultyMap["expert"] = 75

	b.cells = board

	for k := 0; k < 9; k += 3 {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
				shuffleArray := shuffleArray(nums)
				safe := false
				var num int
				x := i + k
				y := j + k
				for l := 0; l < len(shuffleArray); l++ {
					num = shuffleArray[l]
					safe, _, _, _ = b.isSafe(x, y, num)
					if safe {
						break
					} else {
						continue
					}
				}
				b.cells[x][y] = num
				continue
			}
		}
	}

	b.solveSudoku()

	var difValue int
	if val, ok := difficultyMap[difficulty]; ok {
		difValue = val
	} else {
		fmt.Printf("\n\nPLEASE ENTER THE CORRECT DIFFICULTY. (easy, medium, hard, or expert)")
		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			randNum := randInt(1, 100)
			if randNum <= difValue {
				b.cells[i][j] = 0
				continue
			}
		}
	}

	generatedBoard := b.cells

	if boardGood, _, _, _ := b.checkStartingBoard(); !boardGood {
		b.generateBoard(difficulty)
	} else {
		b.cells = generatedBoard
	}
}

func randInt(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Microsecond * 1)
	return r.Intn(max-min) + min
}

func shuffleArray(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Microsecond * 1)
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func inArray(nums []int, num1 int) bool {
	for _, num := range nums {
		if num == num1 {
			return true
		}
	}

	return false
}
