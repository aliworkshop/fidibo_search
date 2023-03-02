package fidibo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *client) Search(keyword string) (*BookResponse, error) {
	url := c.baseURL

	q := url.Query()
	q.Set("q", keyword)
	url.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.cl.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("error on get data from fidibo with status code %d", resp.StatusCode))
	}

	res := new(BookResponse)
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
