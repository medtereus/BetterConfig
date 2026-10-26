package BetterConfig

import "encoding/json"

func GetGeneric[T any](key string) T {
	var result T
	v, ok := valuesVar[key]
	if !ok {
		return result
	}

	data, err := json.Marshal(v)
	if err != nil {
		return result
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result
	}
	return result
}

func SetGeneric[T any](key string, value T) {
	valuesVar[key] = value
	valuesVar.Save()
}
