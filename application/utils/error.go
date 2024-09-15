package utils

// ErrorPanic panics if the provided error is not nil.
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
