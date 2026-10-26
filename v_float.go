package BetterConfig

import (
	"encoding/json"
)

func GetFloat(key string) float64 {
	return GetFloatWithFallback(key, 0)
}

func GetFloatWithFallback(key string, fallback float64) float64 {
	v, ok := valuesVar[key]
	if !ok {
		return fallback
	}

	switch n := v.(type) {
	case float64:
		return n
	case json.Number:
		i, err := n.Float64()
		if err != nil {
			return fallback
		}
		return i
	default:
		return fallback
	}
}

func SetFloat(key string, value float64) {
	valuesVar[key] = value
	valuesVar.Save()
}
