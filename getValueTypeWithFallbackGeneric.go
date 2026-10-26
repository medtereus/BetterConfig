package BetterConfig

func getValueWithFallback[T any](key string, fallback T) T {
	v, ok := valuesVar[key]
	if !ok {
		return fallback
	}

	vGeneric, ok := v.(T)
	if !ok {
		return fallback
	}
	return vGeneric
}
