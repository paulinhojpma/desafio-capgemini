package web

type Sequence struct {
	ID      int
	IsValid bool
	Letters string
}

type InfoSequence struct {
	Valids   int     `json:"count_valid"`
	Invalids int     `json:"count_invalid"`
	Ratio    float32 `json:"ratio"`
}

type Letters struct {
	Letters []string `json:"letters"`
}

type SequenceValidatorResponse struct {
	IsValid bool `json:"is_valid"`
}
