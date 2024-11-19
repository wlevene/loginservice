package util

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func Json2map(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}

	if jsonStr == "" {
		return result, errors.New("content is null")
	}

	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		// err = errors.WithMessage(err, jsonStr)
		return result, err
	}
	return result, nil
}

func Map2json(mapi map[string]interface{}) (string, error) {
	var result []byte
	result, err := json.Marshal(mapi)

	return string(result), err
}
