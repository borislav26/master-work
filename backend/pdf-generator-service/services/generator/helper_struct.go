package generator

type DataWithFileName struct {
	FileName string `json:"fileName"`
	Data     []byte `json:"data"`
}
