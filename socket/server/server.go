package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"sudoku/common"
	"sudoku/utils"
)
func handleConnection(conn net.Conn) {
	defer conn.Close()
	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	var request common.SudokuRequest
	var response common.SudokuResponse

	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println("Error decoding request:", err)
		return
	}

	grid := request.Grid
	if utils.SolveSudoku(&grid) {
		response.SolvedGrid = grid
		response.Success = true
	} else {
		response.Success = false
	}

	err = encoder.Encode(&response)
	if err != nil {
		fmt.Println("Error encoding response:", err)
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
