package BetterConfig

import (
	"encoding/json"
	"log/slog"
)

func GetFloat(key string) float64 {
	return GetFloatWithFallback(key, 0)
}

func GetFloatWithFallback(key string, fallback float64) float64 {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return fallback
	}
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
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return
	}
	valuesVar[key] = value
	valuesVar.Save()
}
