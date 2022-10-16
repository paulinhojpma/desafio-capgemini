package core

type ErrDetail struct {
	Resource    string `json:"resource"`
	Code        string `json:"code"`
	Message     string `json:"message"`
	IDOperation string `json:"idOperation"`
}

// Constantes que representam nomes que mapeiam um tipo de erro
const (
	ErrorReadAllBuffer = "ErrorReadAllBuffer"
	ErrorJSONUnmarshal = "ErrorJSONUnmarshal"
)
