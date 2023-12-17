package helper

import "strings"

func GetBearerToken(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
