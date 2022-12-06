package day6

import (
	"bufio"
	"os"
)

func day6PreBuild(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	bufferInput := ""
	for s.Scan() {
		bufferInput = s.Text()
	}
	return bufferInput
}

func partA(bufferInput string, packetLength int) int {
	mk := NewPacketStartMarker(bufferInput)
	return mk.Marker.MarkerStartPos
}

func partB(bufferInput string) int {
	mk := NewMessageStartMarker(bufferInput)
	return mk.Marker.MarkerStartPos
}
