package BetterConfig

// (value, isFallback)
func getValueWithFallback[T any](key string, fallback T) (T, bool) {
	v, ok := valuesVar[key]
	if !ok {
		return fallback, true
	}

	vGeneric, ok := v.(T)
	if !ok {
		return fallback, true
	}
	return vGeneric, false
}
