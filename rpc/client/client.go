package main

import (
	"fmt"
	"net/rpc"
	"sudoku/common"
	"sudoku/utils"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Printf("Error dialing: %v\n", err)
		return
	}
	defer client.Close()

	request := utils.BuildSudokuRequest()
	var response common.SudokuResponse

	n := 10000 // Número de requisições
	var totalDuration time.Duration

	for i := 0; i < n; i++ {
		start := time.Now()
		err := client.Call("SudokuService.SolveSudoku", request, &response)
		if err != nil {
			fmt.Printf("Error calling remote procedure: %v\n", err)
			return
		}

		duration := time.Since(start)
		totalDuration += duration
	}
	if response.Success {
		fmt.Printf("Solved Sudoku\n")
		utils.PrintGrid(response.SolvedGrid)
	} else {
		fmt.Printf("No solution exists (Call %d)\n")
	}
	avgDuration := totalDuration / time.Duration(n)
	fmt.Printf("Average RTT using RPC: %v\n", avgDuration)
}
