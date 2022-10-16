package validate

const (
	ValidatorAssync = "Assync"
)

type SequecenValidator interface {
	ValidateSequence(grid []string) bool
}

type AssyncValidator struct {
}

func (a *AssyncValidator) ValidateSequence(grid []string) bool {

	return assincValidateSequence(grid)
}

func CreateValidator(kind string) SequecenValidator {
	var validator SequecenValidator
	switch kind {
	case ValidatorAssync:
		validator = &AssyncValidator{}

	}
	return validator
}
func assincValidateSequence(grid []string) bool {
	hasSequenceChan := make(chan bool)
	done := make(chan struct{})
	sizeValidation := len(grid[0])*2 + 2
	validateSequence := []func(grid []string, hasSequence chan bool, done chan struct{}){checkAllHorizontal, checkAllVertical, checkAllDiagonal}
	for _, seqV := range validateSequence {
		seqV(grid, hasSequenceChan, done)

	}
	i := 1

	for hasSequence := range hasSequenceChan {

		if hasSequence {
			close(done)
			return hasSequence
		}

		if i >= sizeValidation {
			close(hasSequenceChan)
			close(done)
			return hasSequence
		}
		i++
	}
	return false
}

func checkAllHorizontal(grid []string, hasSequence chan bool, done chan struct{}) {
	for _, str := range grid {

		go func(hasSequence chan bool, str string) {
			select {
			case <-done:
				return
			default:
				hasSequence <- checkHorizontal(str)
			}

		}(hasSequence, str)
	}
}

func checkAllVertical(grid []string, hasSequence chan bool, done chan struct{}) {
	for i := range grid {

		go func(hasSequence chan bool, i int) {
			select {
			case <-done:
				return
			default:
				hasSequence <- checkVertical(grid, i)
			}

		}(hasSequence, i)
	}
}

func checkAllDiagonal(grid []string, hasSequence chan bool, done chan struct{}) {

	go func(hasSequence chan bool) {
		select {
		case <-done:
			return
		default:
			hasSequence <- checkDiagonalLeftToRight(grid)
		}

	}(hasSequence)

	go func(hasSequence chan bool) {
		select {
		case <-done:
			return
		default:
			hasSequence <- checkDiagonalRightToLeft(grid)
		}

	}(hasSequence)

}

func checkHorizontal(row string) bool {
	tam := len(row)
	seqSize := 1
	seqChar := rune(row[0])
	for index, ch := range row {
		if index != 0 {
			// fmt.Println("seqSize - ", seqSize)
			// fmt.Println("seqChar - ", string(ch))
			if ch == seqChar {

				seqSize++
				if seqSize >= 4 {
					return true
				}
			} else {
				seqChar = ch
				seqSize = 1
			}

			if tam-index+1+seqSize < 4 {
				return false
			}

		}

	}
	return false
}

func checkVertical(grid []string, col int) bool {
	tam := len(grid[0])
	seqSize := 1
	seqChar := rune(grid[0][col])
	for index := 1; index < tam; index++ {

		if rune(grid[index][col]) == seqChar {

			seqSize++
			if seqSize >= 4 {
				return true
			}

		} else {
			seqChar = rune(grid[index][col])
			seqSize = 1
		}
		if tam-index+1+seqSize < 4 {
			return false
		}
	}
	return false
}

func checkDiagonalLeftToRight(grid []string) bool {

	tam := len(grid[0])

	columnMaster := 3
	rowMaster := 0
	adjustDiagonal := 0
	for rowMaster < tam-3 {

		col := columnMaster
		row := rowMaster
		tamDiagonal := columnMaster + 1 - adjustDiagonal
		seqSize := 1
		seqChar := rune(grid[row][col])

		for index := 1; index < tamDiagonal; index++ {

			if rune(grid[row+index][col-index]) == seqChar {
				seqSize++
				if seqSize >= 4 {
					return true
				}
			} else {
				seqChar = rune(grid[row+index][col-index])
				seqSize = 1
			}
			if tamDiagonal-index+1+seqSize < 4 {
				break
			}
		}
		if tam-1 > columnMaster {

			columnMaster++
		} else {
			adjustDiagonal++
			rowMaster++
		}

	}
	return false
}

func checkDiagonalRightToLeft(grid []string) bool {

	tam := len(grid[0])

	columnMaster := tam - 4
	rowMaster := 0
	adjustDiagonal := 0
	for rowMaster < tam-3 {

		col := columnMaster
		row := rowMaster

		tamDiagonal := tam - columnMaster - adjustDiagonal
		seqSize := 1

		seqChar := rune(grid[row][col])

		for index := 1; index < tamDiagonal; index++ {

			if rune(grid[row+index][col+index]) == seqChar {
				seqSize++
				if seqSize >= 4 {
					return true
				}
			} else {
				seqChar = rune(grid[row+index][col+index])
				seqSize = 1
			}
			if tamDiagonal-index+1+seqSize < 4 {
				break
			}
		}
		if columnMaster >= 1 {

			columnMaster--
		} else {
			adjustDiagonal++
			rowMaster++
		}

	}
	return false
}
