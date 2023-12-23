package pulseaudio

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lpulse -lpulse-simple
// #include <pulse/error.h>
import "C"

import "errors"

func getErrorFromCode(errorCode C.int) error {
	cStrError := C.pa_strerror(C.int(errorCode))
	return errors.New(C.GoString(cStrError))
}
