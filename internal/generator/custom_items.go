package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/SkaticNET/bedrock-jsonfix/bedrockjsonfix"
)

type itemTextureFile struct {
	TextureData map[string]any `json:"texture_data"`
}

// LoadCustomItemsFromRP loads custom item identifiers from
// <RP>/textures/item_texture.json and assigns AUX IDs
func LoadCustomItemsFromRP(rpPath string, startID int64) (map[string]int64, error) {
	path := filepath.Join(rpPath, "textures", "item_texture.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read item_texture.json: %w", err)
	}

	opt := bedrockjsonfix.DefaultOptions()
	res, err := bedrockjsonfix.FixBytes(data, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to fix item_texture.json syntax: %w", err)
	}

	var file itemTextureFile
	if err := json.Unmarshal(res.Output, &file); err != nil {
		return nil, fmt.Errorf("invalid item_texture.json format: %w", err)
	}

	if len(file.TextureData) == 0 {
		return map[string]int64{}, nil
	}

	ids := make([]string, 0, len(file.TextureData))
	for id := range file.TextureData {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	out := make(map[string]int64, len(ids))
	for i, id := range ids {
		itemID := startID + int64(i)
		aux := itemID << 16
		out[id] = aux
	}

	return out, nil
}
