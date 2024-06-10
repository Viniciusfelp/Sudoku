package common

const N = 9

type SudokuRequest struct {
	Grid [N][N]int
}

type SudokuResponse struct {
	SolvedGrid [N][N]int
	Success    bool
}
