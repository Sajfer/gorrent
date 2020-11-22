package peers_test

import (
	"net"
	"testing"

	"github.com/sajfer/gorrent/peers"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshall(t *testing.T) {
	tests := map[string]struct {
		input  string
		output []peers.Peer
		fails  bool
	}{
		"correctly parses peers": {
			input: string([]byte{127, 0, 0, 1, 0x00, 0x50, 1, 1, 1, 1, 0x01, 0xbb}),
			output: []peers.Peer{
				{IP: net.IP{127, 0, 0, 1}, Port: 80},
				{IP: net.IP{1, 1, 1, 1}, Port: 443},
			},
		},
		"not enough bytes in peers": {
			input:  string([]byte{127, 0, 0, 1, 0x00}),
			output: nil,
			fails:  true,
		},
	}

	for _, test := range tests {
		peers, err := peers.Unmarshal([]byte(test.input))
		if test.fails {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, test.output, peers)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		input  peers.Peer
		output string
	}{
		{
			input:  peers.Peer{IP: net.IP{127, 0, 0, 1}, Port: 8080},
			output: "127.0.0.1:8080",
		},
	}
	for _, test := range tests {
		s := test.input.String()
		assert.Equal(t, test.output, s)
	}

}
