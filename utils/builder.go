package utils

func SetAndReturn[R any, T any](typ R, item *T, new T) R {
	*item = new
	return typ
}
