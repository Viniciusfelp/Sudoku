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

func protoToGrid(req *pb.SudokuRequest) [utils.N][utils.N]int {
	var grid [utils.N][utils.N]int
	for i, row := range req.Grid {
		for j, cell := range row.Cells {
			grid[i][j] = int(cell)
		}
	}
	return grid
}

func gridToProto(grid [utils.N][utils.N]int) *pb.SudokuResponse {
	response := &pb.SudokuResponse{Success: true}
	for i := 0; i < utils.N; i++ {
		row := &pb.Row{}
		for j := 0; j < utils.N; j++ {
			row.Cells = append(row.Cells, int32(grid[i][j]))
		}
		response.SolvedGrid = append(response.SolvedGrid, row)
	}
	return response
}

func (s *server) SolveSudoku(ctx context.Context, req *pb.SudokuRequest) (*pb.SudokuResponse, error) {
	grid := protoToGrid(req)

	success := utils.SolveSudoku(&grid)
	response := &pb.SudokuResponse{Success: success}
	if success {
		response = gridToProto(grid)
	}
	return response, nil
}

func (s *server) SolveSudokuConcurrently(ctx context.Context, req *pb.SudokuRequest) (*pb.SudokuResponse, error) {
	grid := protoToGrid(req)

	success := utils.SolveSudokuConcurrently(&grid)
	response := &pb.SudokuResponse{Success: success}
	if success {
		response = gridToProto(grid)
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
