package generator

import (
	"encoding/json"
	"io"
	"net/http"
)

const mojangURL = "https://raw.githubusercontent.com/Mojang/bedrock-samples/main/metadata/vanilladata_modules/mojang-items.json"

type mojangItems struct {
	DataItems []struct {
		Name  string `json:"name"`
		RawID int64  `json:"raw_id"`
	} `json:"data_items"`
}

func Generate(customItemsRP string) (map[string]int64, error) {
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

	out := make(map[string]int64, len(data.DataItems))

	// Vanilla AUX
	for _, it := range data.DataItems {
		out[it.Name] = it.RawID << 16
	}

	// Custom AUX
	if customItemsRP != "" {
		custom, err := LoadCustomItemsFromRP(customItemsRP, out)
		if err != nil {
			return nil, err
		}
		for k, v := range custom {
			out[k] = v
		}
	}

	return out, nil
}