package structures

import (
	"math/rand"
	"time"
)

type PitchName int

const (
	C PitchName = iota + 1
	D
	E
	F
	G
	A
	B
	Csharp
	Dsharp
	Esharp
	Fsharp
	Gsharp
	Asharp
	Bsharp
	Cflat
	Dflat
	Eflat
	Fflat
	Gflat
	Aflat
	Bflat
)

func (p PitchName) IsSharp() bool {
	switch p {
	case Csharp:
		fallthrough
	case Dsharp:
		fallthrough
	case Esharp:
		fallthrough
	case Fsharp:
		fallthrough
	case Gsharp:
		fallthrough
	case Asharp:
		fallthrough
	case Bsharp:
		return true
	default:
		return false
	}
}

//func IsFlat(pitchName PitchName) bool {
//
//}

func RandPitch() PitchName {
	rand.Seed(time.Now().Unix())
	return IntervalOrder[rand.Intn(len(IntervalOrder))]
}
