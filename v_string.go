package BetterConfig

func GetString(key string) string {
	return GetStringWithFallback(key, "")
}

func GetStringWithFallback(key, fallback string) string {
	return getValueWithFallback[string](key, fallback)
}

func SetString(key, value string) {
	valuesVar[key] = value
	valuesVar.Save()
}
