package decoder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vx-labs/mqtt-protocol/packet"
)

func TestAsyncDecoder_Packets(t *testing.T) {
	buff := []byte{0x32, 0x81, 0x1,
		0x0, 0x1, 'a',
		0x0, 0x83}
	for i := 0; i < 0x83; i++ {
		buff = append(buff, 'a')
	}
	reader := bytes.NewReader(buff)
	decoder := Async(reader)
	<-decoder.Done()
	require.NoError(t, decoder.Err())
	p := <-decoder.Packet()
	require.IsType(t, &packet.Publish{}, p)
}
