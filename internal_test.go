package goidoit

// export getID
func TgetID() int {
	return getID()
}

// reset id
func TResetID() {
	id = 0
}

// export debug func
func TdebugPrint(format string, a ...interface{}) (n int, err error) {
	return debugPrint(format, a...)
}

func TgetDebug() bool {
	return debug
}
