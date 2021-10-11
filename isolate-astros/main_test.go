package main

import "testing"

type testWebRequest struct {
}

func (testWebRequest) FetchBytes(url string) ([]byte, error) {
	return []byte(`{"number": 3}`), nil
}
func TestGetAstronauts(t *testing.T) {
	want := 3
	got, err := getAstronauts(testWebRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("People in space, want: %d, got: %d.", want, got)
	}
}
