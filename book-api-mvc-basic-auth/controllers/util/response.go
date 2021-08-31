package util

func Response(m string, r interface{}) map[string]interface{} {
	if r == nil {
		return map[string]interface{}{"message": m}
	}
	return map[string]interface{}{"status": m, "user": r}
}
