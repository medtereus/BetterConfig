package BetterConfig

func RemoveValue(key string) {
	delete(valuesVar, key)
	valuesVar.Save()
}
