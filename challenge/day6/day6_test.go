package day6

import (
	"fmt"
	"testing"
)

var testPacketDataBuffer = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
var testMessageDataBuffer = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

func TestNewPacketStartMarker(t *testing.T) {
	mk := NewPacketStartMarker(testPacketDataBuffer)
	fmt.Println(mk.Marker.MarkerStartPos)
}

func TestNewMessageStartMarker(t *testing.T) {
	mk := NewMessageStartMarker(testMessageDataBuffer)
	fmt.Println(mk.Marker.MarkerStartPos)
}
