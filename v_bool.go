package BetterConfig

func GetBool(key string) bool {
	return GetBoolWithFallback(key, false)
}

func GetBoolWithFallback(key string, fallback bool) bool {
	return getValueWithFallback[bool](key, fallback)
}

func SetBool(key string, value bool) {
	valuesVar[key] = value
	valuesVar.Save()
}
