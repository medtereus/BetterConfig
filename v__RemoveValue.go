package BetterConfig

import "log/slog"

func RemoveValue(key string) {
	if !isInit {
		slog.Warn("BetterConfig is not initialized")
		return
	}
	delete(valuesVar, key)
	valuesVar.Save()
}
