package interfaces

type KeppMultiplyingParams struct {
	Nums     []int `json:"nums"`
	Original int   `json:"original"`
	Expected int `json:"expected"`
}
