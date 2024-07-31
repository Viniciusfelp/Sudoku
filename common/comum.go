package common

const N = 9

type SudokuRequest struct {
	Grid [N][N]int
}

type SudokuJsonRequest struct {
	Grid [N][N]int `json:"grid"`
}

type SudokuResponse struct {
	SolvedGrid [N][N]int
	Success    bool
}
