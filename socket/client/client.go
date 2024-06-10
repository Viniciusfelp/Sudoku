package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"sudoku/common"
	"sudoku/utils"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	request := common.SudokuRequest{
		Grid: [common.N][common.N]int{
			{5, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		},
	}

	var response common.SudokuResponse
	n := 100 // Número de requisições
	totalRTT := 0.0

	for i := 0; i < n; i++ {
		start := time.Now()
		err = encoder.Encode(request)
		if err != nil {
			fmt.Println("Error encoding request:", err)
			return
		}

		err = decoder.Decode(&response)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}
		elapsed := time.Since(start).Seconds()
		totalRTT += elapsed
	}

	avgRTT := totalRTT / float64(n)
	fmt.Printf("Average RTT using Sockets: %f seconds\n", avgRTT)

	if response.Success {
		fmt.Println("Solved Sudoku:")
		utils.PrintGrid(response.SolvedGrid)
	} else {
		fmt.Println("No solution exists")
	}
}
