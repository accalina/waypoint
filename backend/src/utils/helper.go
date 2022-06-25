package utils

func SetDefaultStr(value1 string, value2 string) string {
	if value1 != "" {
		return value1
	}
	return value2
}
