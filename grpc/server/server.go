package main

import (
	"context"
	"log"
	"net"
	pb "sudoku/proto/sudoku"
	"sudoku/utils"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSudokuSolverServer
}

func (s *server) SolveSudoku(ctx context.Context, req *pb.SudokuRequest) (*pb.SudokuResponse, error) {
	var grid [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid[i][j] = int(req.Grid[i*9+j])
		}
	}

	success := utils.SolveSudoku(&grid)
	response := &pb.SudokuResponse{Success: success}
	if success {
		response.SolvedGrid = make([]int32, 81)
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				response.SolvedGrid[i*9+j] = int32(grid[i][j])
			}
		}
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSudokuSolverServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
