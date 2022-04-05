package cmap

func Remove(data map[string]interface{}, fields []string) map[string]interface{} {
	for _, item := range fields {
		delete(data, item)
	}
	return data
}
