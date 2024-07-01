package main

import (
	"context"
	"fmt"
	"log"
	pb "sudoku/proto/sudoku"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func buildRequest() []int32 {
	grid := []int32{
		5, 3, 0, 0, 7, 0, 0, 0, 0,
		6, 0, 0, 1, 9, 5, 0, 0, 0,
		0, 9, 8, 0, 0, 0, 0, 6, 0,
		8, 0, 0, 0, 6, 0, 0, 0, 3,
		4, 0, 0, 8, 0, 3, 0, 0, 1,
		7, 0, 0, 0, 2, 0, 0, 0, 6,
		0, 6, 0, 0, 0, 0, 2, 8, 0,
		0, 0, 0, 4, 1, 9, 0, 0, 5,
		0, 0, 0, 0, 8, 0, 0, 7, 9,
	}
	return grid
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSudokuSolverClient(conn)

	request := &pb.SudokuRequest{Grid: buildRequest()}
	var totalRTT float64
	n := 10000 

	for i := 0; i < n; i++ {
		start := time.Now()
		response, err := c.SolveSudoku(context.Background(), request)
		if err != nil {
			log.Fatalf("could not solve sudoku: %v", err)
		}
		elapsed := time.Since(start).Seconds() * 1000 
		totalRTT += elapsed
		if i == 0 {
			if response.Success {
				fmt.Println("Solved Sudoku:")
				printGrid(response.SolvedGrid)
			} else {
				fmt.Println("No solution exists")
			}
		}
	}

	avgRTT := totalRTT / float64(n)
	fmt.Printf("Average RTT using gRPC: %f ms\n", avgRTT)
}

func printGrid(grid []int32) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", grid[i*9+j])
		}
		fmt.Println()
	}
}
