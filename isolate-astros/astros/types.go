package astros

type People struct {
	Number int `json:"number"`
}
type GetWebRequest interface {
	FetchBytes(url string) ([]byte, error)
}
