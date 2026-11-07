package BetterConfig

import "log/slog"

func GetString(key string) string {
	return GetStringWithFallback(key, "")
}

func GetStringWithFallback(key, fallback string) string {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return fallback
	}
	reutrnV, isFallback := getValueWithFallback[string](key, fallback)
	if isFallback {
		SetString(key, fallback)
	}
	return reutrnV

}

func SetString(key, value string) {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return
	}
	valuesVar[key] = value
	valuesVar.Save()
}
