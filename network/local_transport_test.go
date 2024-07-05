package network

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// assert.Equal(t, 1, 1) eşitlik kontrolü

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Merhaba Dunya")
	assert.Nil(t, tra.SendMessage(trb.addr, msg))

	rpc := <-trb.Consume()
	b, err := ioutil.ReadAll(rpc.Payload)

	assert.Nil(t, err)
	assert.Equal(t, b, msg)
	assert.Equal(t, rpc.From, tra.addr)
}


func TestBroadcast(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")
	trc := NewLocalTransport("C")

	tra.Connect(trb)
	tra.Connect(trc)

	msg := []byte("testt")
	assert.Nil(t, tra.Broadcast(msg))

	rpcb := <-trb.Consume()
	b, err := ioutil.ReadAll(rpcb.Payload)
	assert.Nil(t, err)
	assert.Equal(t, b, msg)

	rpcC := <-trc.Consume()
	b, err = ioutil.ReadAll(rpcC.Payload)
	assert.Nil(t, err)
	assert.Equal(t, b, msg)
}