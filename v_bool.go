package BetterConfig

import "log/slog"

func GetBool(key string) bool {
	return GetBoolWithFallback(key, false)
}

func GetBoolWithFallback(key string, fallback bool) bool {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return fallback
	}

	returnV, isFallback := getValueWithFallback[bool](key, fallback)
	if isFallback {
		SetBool(key, fallback)
	}
	return returnV
}

func SetBool(key string, value bool) {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return
	}
	valuesVar[key] = value
	valuesVar.Save()
}
