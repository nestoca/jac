package live

import (
	"encoding/json"
	"fmt"
	"strings"
)

func loadValues(rawJson []byte) (map[string]interface{}, error) {
	v := make(map[string]interface{})
	if rawJson != nil {
		err := json.Unmarshal(rawJson, &v)
		if err != nil {
			fmt.Errorf("unmarshalling JSON values: %w", err)
		}
	}
	return v, nil
}

func getValue(values map[string]interface{}, keyPath string) (string, bool) {
	keys := strings.Split(keyPath, ".")
	for i, key := range keys {
		child, ok := values[key]
		if !ok {
			return "", false
		}
		values, ok = child.(map[string]interface{})
		if !ok {
			// Only consider leaf values if we are at end of key path
			if i != len(keys)-1 {
				return "", false
			}
			value, ok := child.(string)
			if !ok {
				return "", false
			}
			return value, true
		}
	}
	return "", false
}
