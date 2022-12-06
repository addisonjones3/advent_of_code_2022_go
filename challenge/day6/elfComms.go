package day6

type Marker struct {
	MarkerString   string
	MarkerMap      map[string]int
	MarkerStartPos int
	DataBuffer     string
}

type PacketStartMarker struct {
	Marker Marker
}

type MessageStartMarker struct {
	Marker Marker
}

func NewMarkerFromString(s string, packetLength int) *Marker {
	mk := &Marker{DataBuffer: s}
	for i := 0; i < len(s)-packetLength; i++ {
		uniqueChars := 0
		currChars := s[i : i+packetLength]
		currMarkerMap := make(map[string]int)
		for _, c := range currChars {
			cStr := string(c)
			if _, ok := currMarkerMap[cStr]; ok {
				break
			} else {
				currMarkerMap[cStr]++
				uniqueChars++
			}
		}
		if uniqueChars == packetLength {
			mk.MarkerString = string(currChars)
			mk.MarkerMap = currMarkerMap
			mk.MarkerStartPos = i + packetLength
			return mk
		}
	}
	return nil
}

func NewMessageStartMarker(s string) *MessageStartMarker {
	mk := NewMarkerFromString(s, 14)
	return &MessageStartMarker{Marker: *mk}
}

func NewPacketStartMarker(s string) *PacketStartMarker {
	mk := NewMarkerFromString(s, 4)
	return &PacketStartMarker{Marker: *mk}
}
