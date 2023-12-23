package pulseaudio

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lpulse -lpulse-simple
// #include <pulse/simple.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

type PAStreamDirection C.pa_stream_direction_t

const (
	PA_STREAM_NODIRECTION PAStreamDirection = C.PA_STREAM_NODIRECTION
	PA_STREAM_PLAYBACK    PAStreamDirection = C.PA_STREAM_PLAYBACK
	PA_STREAM_RECORD      PAStreamDirection = C.PA_STREAM_RECORD
	PA_STREAM_UPLOAD      PAStreamDirection = C.PA_STREAM_UPLOAD
)

type PASimple C.struct_pa_simple

func PaSimpleNew(spec PASampleSpec, direction PAStreamDirection, programName, description string) (*PASimple, error) {
	var programNameCString *C.char
	if programName != "" {
		programNameCString = C.CString(programName)
		defer C.free(unsafe.Pointer(programNameCString))
	}
	var descriptionCString *C.char
	if description != "" {
		descriptionCString = C.CString(description)
		defer C.free(unsafe.Pointer(descriptionCString))
	}

	ss := spec.toCStruct()
	var errorCode C.int

	s := C.pa_simple_new(nil, programNameCString, C.pa_stream_direction_t(direction), nil, descriptionCString, ss, nil, nil, &errorCode)
	if errorCode != C.PA_OK {
		return nil, getErrorFromCode(errorCode)
	}
	return (*PASimple)(s), nil
}

func (s *PASimple) Free() {
	C.pa_simple_free((*C.struct_pa_simple)(s))
}

func (s *PASimple) Read(buffer []byte) (int, error) {
	var errorCode C.int
	C.pa_simple_read((*C.pa_simple)(s), unsafe.Pointer(&buffer[0]), C.size_t(len(buffer)), &errorCode)
	if errorCode != C.PA_OK {
		return 0, getErrorFromCode(errorCode)
	}
	return len(buffer), nil
}

func (s *PASimple) Write(buffer []byte) (int, error) {
	var errorCode C.int
	C.pa_simple_write((*C.pa_simple)(s), unsafe.Pointer(&buffer[0]), C.size_t(len(buffer)), &errorCode)
	if errorCode != C.PA_OK {
		return 0, getErrorFromCode(errorCode)
	}
	return len(buffer), nil
}

func (s *PASimple) Flush() error {
	var errorCode C.int
	C.pa_simple_flush((*C.pa_simple)(s), &errorCode)
	if errorCode != C.PA_OK {
		return getErrorFromCode(errorCode)
	}
	return nil
}

func (s *PASimple) Drain() error {
	var errorCode C.int
	C.pa_simple_drain((*C.pa_simple)(s), &errorCode)
	if errorCode != C.PA_OK {
		return getErrorFromCode(errorCode)
	}
	return nil
}
