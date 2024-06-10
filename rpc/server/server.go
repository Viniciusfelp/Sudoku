package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sudoku/common"
	"sudoku/utils"
)

type SudokuService struct{}

func (s *SudokuService) SolveSudoku(req common.SudokuRequest, res *common.SudokuResponse) error {
	grid := req.Grid
	if utils.SolveSudoku(&grid) {
		res.SolvedGrid = grid
		res.Success = true
	} else {
		res.Success = false
	}
	return nil
}

func main() {
	sudokuService := new(SudokuService)
	rpc.Register(sudokuService)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Serving RPC server on port 1234")
	rpc.Accept(listener)
}
