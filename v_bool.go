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
	return getValueWithFallback[bool](key, fallback)
}

func SetBool(key string, value bool) {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return
	}
	valuesVar[key] = value
	valuesVar.Save()
}
