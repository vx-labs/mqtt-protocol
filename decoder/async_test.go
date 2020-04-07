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
	buf := []byte{}
	buf = append(buf, 0x32, 55,
		0x0, 0x1, 'a',
		0, 50)
	for i := 0; i < 50; i++ {
		buf = append(buf, 'a')
	}
	h := make([]byte, 4)
	reader := bytes.NewReader(buf)
	for i := 0; i < b.N; i++ {
		reader.Seek(0, io.SeekStart)
		decodeEncodedPacket(h, reader)
	}
}
func BenchmarkAsyncDecoder_unmarshalPacket(b *testing.B) {
	b.Run("pingreq", func(b *testing.B) {
		header := &packet.Header{}
		p := &packet.PingReq{Header: &packet.Header{}}
		buff := make([]byte, p.Length())
		_, err := p.Encode(buff)
		require.NoError(b, err)
		for i := 0; i < b.N; i++ {
			unmarshalPacket(packet.PINGREQ, header, buff)
		}
	})
	b.Run("publish", func(b *testing.B) {
		header := &packet.Header{}
		p := &packet.Publish{Header: &packet.Header{}, MessageId: 1, Payload: []byte("test"), Topic: []byte("test")}
		buff := make([]byte, p.Length())
		_, err := p.Encode(buff)
		require.NoError(b, err)
		for i := 0; i < b.N; i++ {
			unmarshalPacket(packet.PUBLISH, header, buff)
		}
	})
	b.Run("subscribe", func(b *testing.B) {
		header := &packet.Header{}
		p := &packet.Subscribe{
			Header:    &packet.Header{},
			MessageId: 1,
			Qos:       []int32{0},
			Topic:     [][]byte{[]byte("test")},
		}
		buff := make([]byte, p.Length())
		_, err := p.Encode(buff)
		require.NoError(b, err)
		for i := 0; i < b.N; i++ {
			unmarshalPacket(packet.SUBSCRIBE, header, buff)
		}
	})
}
