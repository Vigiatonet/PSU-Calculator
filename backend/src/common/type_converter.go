package common

import "encoding/json"

func TypeConverter[T any](model interface{}) (*T, error) {
	var dest = new(T)
	bsResult, err := json.Marshal(&model)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bsResult, &dest)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
