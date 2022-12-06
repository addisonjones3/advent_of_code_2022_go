package day6

import (
	"fmt"
	"testing"
)

var testDataBuffer = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

func TestPacketStartMarkerFromString(t *testing.T) {
	mk := NewPacketStartMarker(testDataBuffer)
	fmt.Println(mk.Marker.MarkerStartPos)
}
