package generator

import (
	"encoding/json"
	"os"
)

func WriteJSON(path string, data map[string]int64) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}