package BetterConfig

import (
	"encoding/json"
)

func GetInt(key string) int {
	return GetIntWithFallback(key, 0)
}

func GetIntWithFallback(key string, fallback int) int {
	v, ok := valuesVar[key]
	if !ok {
		return fallback
	}

	switch n := v.(type) {
	case int:
		return n
	case json.Number:
		i, err := n.Int64()
		if err != nil {
			return fallback
		}
		return int(i)
	default:
		return fallback
	}
}

func SetInt(key string, value int) {
	valuesVar[key] = value
	valuesVar.Save()
}
