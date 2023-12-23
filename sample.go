package pulseaudio

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lpulse -lpulse-simple
// #include <pulse/sample.h>
import "C"

type PASampleFormat C.pa_sample_format_t

const (
	PA_SAMPLE_U8        PASampleFormat = C.PA_SAMPLE_U8
	PA_SAMPLE_ALAW      PASampleFormat = C.PA_SAMPLE_ALAW
	PA_SAMPLE_ULAW      PASampleFormat = C.PA_SAMPLE_ULAW
	PA_SAMPLE_S16LE     PASampleFormat = C.PA_SAMPLE_S16LE
	PA_SAMPLE_S16BE     PASampleFormat = C.PA_SAMPLE_S16BE
	PA_SAMPLE_FLOAT32LE PASampleFormat = C.PA_SAMPLE_FLOAT32LE
	PA_SAMPLE_FLOAT32BE PASampleFormat = C.PA_SAMPLE_FLOAT32BE
	PA_SAMPLE_S32LE     PASampleFormat = C.PA_SAMPLE_S32LE
	PA_SAMPLE_S32BE     PASampleFormat = C.PA_SAMPLE_S32BE
	PA_SAMPLE_S24LE     PASampleFormat = C.PA_SAMPLE_S24LE
	PA_SAMPLE_S24BE     PASampleFormat = C.PA_SAMPLE_S24BE
	PA_SAMPLE_S24_32LE  PASampleFormat = C.PA_SAMPLE_S24_32LE
	PA_SAMPLE_S24_32BE  PASampleFormat = C.PA_SAMPLE_S24_32BE
)

type PASampleSpec struct {
	Format   PASampleFormat
	Rate     uint32
	Channels uint8
}

func (sampeSpec *PASampleSpec) toCStruct() *C.pa_sample_spec {
	return &C.pa_sample_spec{
		format:   C.pa_sample_format_t(sampeSpec.Format),
		rate:     C.uint32_t(sampeSpec.Rate),
		channels: C.uint8_t(sampeSpec.Channels),
	}
}
