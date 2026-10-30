package BetterConfig

import (
	"encoding/json"
	"log/slog"
)

func GetGeneric[T any](key string) T {
	var result T

	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return result
	}

	v, ok := valuesVar[key]
	if !ok {
		slog.Error("GetGeneric: Failed to get values from RAM. !ok")
		return result
	}

	data, err := json.Marshal(v)
	if err != nil {
		slog.Error("GetGeneric: Failed to json.Marshal generic",
			"err", err)
		return result
	}

	if err := json.Unmarshal(data, &result); err != nil {
		slog.Error("GetGeneric: Failed to json.Marshal generic",
			"err", err)
		return result
	}
	return result
}

func SetGeneric[T any](key string, value T) {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return
	}
	valuesVar[key] = value
	valuesVar.Save()
}
