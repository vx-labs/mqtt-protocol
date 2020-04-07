package decoder

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vx-labs/mqtt-protocol/packet"
)

func TestAsyncDecoder_Packets(t *testing.T) {
	buff := []byte{0x32, 55,
		0x0, 0x1, 'a',
		0, 50}
	for i := 0; i < 50; i++ {
		buff = append(buff, 'a')
	}
	reader := bytes.NewReader(buff)
	decoder := Async(reader)
	<-decoder.Done()
	require.Error(t, io.EOF, decoder.Err())
	p := <-decoder.Packet()
	require.IsType(t, &packet.Publish{}, p)
}
func BenchmarkAsyncDecoder_decodeEncodedPacket(b *testing.B) {
	buff := []byte{0x32, 55,
		0x0, 0x1, 'a',
		0, 50}
	for i := 0; i < 50; i++ {
		buff = append(buff, 'a')
	}
	h := make([]byte, 4)
	reader := bytes.NewReader(buff)
	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		decodeEncodedPacket(h, reader)
	}
}
