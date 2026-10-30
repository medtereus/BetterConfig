package BetterConfig

import (
	"encoding/json"
	"log/slog"
	"strconv"
)

func GetUInt(key string) uint {
	return GetUIntWithFallback(key, 0)
}

func GetUIntWithFallback(key string, fallback uint) uint {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return fallback
	}
	v, ok := valuesVar[key]
	if !ok {
		return fallback
	}

	switch n := v.(type) {
	case uint:
		return n
	case json.Number:
		i, err := strconv.ParseUint(n.String(), 10, strconv.IntSize)
		if err != nil {
			return fallback
		}
		return uint(i)
	default:
		return fallback
	}
}

func SetUInt(key string, value uint) {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return
	}
	valuesVar[key] = value
	valuesVar.Save()
}
