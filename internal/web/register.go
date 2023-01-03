package web

func checkAuthToken(key string) bool {
	if key == "authorized" {
		return true
	}
	return false
}
