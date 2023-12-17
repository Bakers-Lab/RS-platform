package util

import "errors"

func PanicIf(condition bool, errMessage string) {
	if condition {
		panic(errors.New(errMessage))
	}
}
