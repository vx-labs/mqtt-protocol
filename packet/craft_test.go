package packet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := New()
	assert.Equal(t, &craftRequest{}, p)
	p = p.WithQoS(1)
	assert.Equal(t, byte(1), p.h.qos)
	p = p.WithRETAIN()
	assert.Equal(t, byte(1), p.h.retain)
	p = p.WithoutRETAIN()
	assert.Equal(t, byte(0), p.h.retain)
	p = p.WithDUP()
	assert.Equal(t, byte(1), p.h.dup)
	p = p.WithoutDUP()
	assert.Equal(t, byte(0), p.h.dup)
}
func TestCraft_Connect(t *testing.T) {
	p := New()
	c := p.Connect()
	assert.Equal(t, CONNECT, c.Type())
	assert.Equal(t, p.h, c.header)
}
func TestCraft_Suback(t *testing.T) {
	p := New()
	c := p.Suback()
	assert.Equal(t, SUBACK, c.Type())
	assert.Equal(t, p.h, c.header)
}

func TestCraft_Subscribe(t *testing.T) {
	p := New()
	c := p.Subscribe()
	assert.Equal(t, SUBSCRIBE, c.Type())
	assert.Equal(t, p.h, c.header)
}
func TestCraft_Publish(t *testing.T) {
	p := New()
	c := p.Publish()
	assert.Equal(t, PUBLISH, c.Type())
	assert.Equal(t, p.h, c.header)
}
func TestCraft_Puback(t *testing.T) {
	p := New()
	c := p.PubAck()
	assert.Equal(t, PUBACK, c.Type())
	assert.Equal(t, p.h, c.header)
}