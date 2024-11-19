package util

import (
	"bytes"
	"io"
	"net/http"
)

func ReadAndAssignRequestBody(req *http.Request) ([]byte, error) {
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	req.Body = io.NopCloser(bytes.NewReader(buf))
	return buf, nil
}
