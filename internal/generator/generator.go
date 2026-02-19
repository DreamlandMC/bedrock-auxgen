package generator

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const mojangURL = "https://raw.githubusercontent.com/Mojang/bedrock-samples/main/metadata/vanilladata_modules/mojang-items.json"

type mojangItems struct {
	DataItems []item `json:"data_items"`
}

type item struct {
	Name  string `json:"name"`
	RawID int64  `json:"raw_id"`
}

func Generate() (map[string]int64, error) {
	resp, err := http.Get(mojangURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data mojangItems
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	aux := make(map[string]int64, len(data.DataItems))

	for _, it := range data.DataItems {
		aux[it.Name] = it.RawID * 65536
	}

	return aux, nil
}

func WriteJSON(path string, m map[string]int64) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(m)
}
