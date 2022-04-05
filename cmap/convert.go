package cmap

import "encoding/json"

func ConvertStruct2MapByJSON(obj interface{}) (map[string]interface{}, error) {
	var data map[string]interface{}
	content, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ConvertMap2StructByJSON(from map[string]interface{}, to interface{}) (interface{}, error) {
	content, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &to)
	if err != nil {
		return nil, err
	}
	return to, nil
}
