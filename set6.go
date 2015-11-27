package dense

import (
	"bytes"
	"fmt"
)

// A Set6 is a bit vector of 2^6 bits.
// It mainly exists to test the Set63.
type Set6 uint64

// Construct a new Set6 out of the elements provided.
// All elements should be between 0 and 63 inclusive, a violation will cause a panic.
func NewSet6(elem ...int64) Set6 {
	var s Set6
	for _, ee := range elem {
		if ee < 0 || ee > 63 {
			panic("Set6 can only contain elements [0..63]")
		}
		s |= 1 << uint8(ee)
	}
	return s
}

func (s Set6) Union(t Set6) Set6        { return s | t }
func (s Set6) Intersection(t Set6) Set6 { return s & t }
func (s Set6) Intersects(t Set6) bool   { return s&t != 0 }
func (s Set6) Complement() (r Set6)     { return ^s }
func (s Set6) Contains(elem int64) bool { return s&(1<<uint64(elem)) != 0 }
func (s Set6) IsEmpty() bool            { return s == 0 }

// Interval6 returns a Set63 that covers the contiguous closed interval [min, max].
func Interval6(min, max int64) Set6 {
	if min > max {
		return 0
	}
	if min < 0 || max > 63 {
		panic("Set6 can only contain elements [0..63]")
	}
	return ((1 << uint8(max)) | (1<<uint8(max) - 1)) & ^((1 << uint8(min)) - 1)
}
func (s Set6) Equal(t Set6) bool { return s == t }

//count

// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
// test: http://play.golang.org/p/zggoyPDdlp
func (v Set6) Count() int {
	v -= (v >> 1) & 0x5555555555555555
	v = v&0x3333333333333333 + (v>>2)&0x3333333333333333
	v += v >> 4
	v &= 0xf0f0f0f0f0f0f0f
	v *= 0x101010101010101
	return int(v >> 56)
}

func (s Set6) Ordinal(elem int64) (n uint64, ok bool) {
	ok = (s&(1<<uint64(elem)) != 0)
	s &= (1 << uint64(elem)) - 1
	n = uint64(s.Count())
	return
}

func (s Set6) String() string {
	if s == 0 {
		return "∅"
	}
	var b bytes.Buffer
	b.WriteByte('{')
	for i := 0; s != 0; i, s = i+1, s>>1 {
		if s&1 != 0 {
			fmt.Fprintf(&b, "%d, ", i)
		}
	}
	b.Truncate(b.Len() - 2)
	b.WriteByte('}')
	return b.String()
}
