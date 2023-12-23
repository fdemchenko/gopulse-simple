# PulseAudio Simple API Go wrapper

## Examples

### Simple audio recorder
```go
package main

import (
	"fmt"
	"io"
	"os"

	pulse "github.com/fdemchenko/gopulse-simple"
)

func main() {
	ss := pulse.PASampleSpec{
		Format:   pulse.PA_SAMPLE_S16LE,
		Rate:     44100,
		Channels: 1,
	}
	sr, err := pulse.PaSimpleNew(ss, pulse.PA_STREAM_RECORD, "Recorder", "My own recorder")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer sr.Free()

	_, err = io.Copy(os.Stdout, sr)
	if err != nil {
		fmt.Println(err.Error())
	}
}
```

### Simple audio player

```go
package main

import (
	"fmt"
	"io"
	"os"

	pulse "github.com/fdemchenko/gopulse-simple"
)

func main() {

	ss := pulse.PASampleSpec{
		Format:   pulse.PA_SAMPLE_S16LE,
		Rate:     44100,
		Channels: 1,
	}
	sp, err := pulse.PaSimpleNew(ss, pulse.PA_STREAM_PLAYBACK, "Player", "My own player")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer sp.Free()

	_, err = io.Copy(sp, os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
	}
}

```
