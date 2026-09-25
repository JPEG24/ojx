package new

import "net/http"

func NewRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	client := &http.Client{}

	return client.Do(req)
}

func NewAtCoderRequest(url, session string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.AddCookie(&http.Cookie{
		Name:  "REVEL_SESSION",
		Value: session,
	})

	client := &http.Client{}

	return client.Do(req)
}
