package astros

import (
	"io/ioutil"
	"net/http"
)

type LiveGetWebRequest struct {
}

func (LiveGetWebRequest) FetchBytes(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("User-Agent", "spacecount")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
