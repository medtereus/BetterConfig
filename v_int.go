package BetterConfig

import (
	"encoding/json"
	"log/slog"
)

func GetInt(key string) int {
	return GetIntWithFallback(key, 0)
}

func GetIntWithFallback(key string, fallback int) int {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return fallback
	}
	v, ok := valuesVar[key]
	if !ok {
		SetInt(key, fallback)
		return fallback
	}

	switch n := v.(type) {
	case int:
		return n
	case json.Number:
		i, err := n.Int64()
		if err != nil {
			SetInt(key, fallback)
			return fallback
		}
		return int(i)
	default:
		SetInt(key, fallback)
		return fallback
	}
}

func SetInt(key string, value int) {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return
	}
	valuesVar[key] = value
	valuesVar.Save()
}
