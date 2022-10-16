package utils

func CheckIfSliceIsSquare(grid []string) bool {

	if len(grid) == len(grid[0]) {
		return true
	}
	return false
}
