package structures

import (
	"fmt"
)

type FlatCount int
type SharpCount int

var NoSignatureMajor = C
var NoSignatureMinor = A

var FlatMajorOrder = [7]PitchName{
	F,
	Bflat,
	Eflat,
	Aflat,
	Dflat,
	Gflat,
	Cflat,
}

var FlatMinorOrder = [7]PitchName{
	D,
	G,
	C,
	F,
	Bflat,
	Eflat,
	Aflat,
}

var SharpMajorOrder = [7]PitchName{
	G,
	D,
	A,
	E,
	B,
	Fsharp,
	Csharp,
}

var SharpMinorOrder = [7]PitchName{
	E,
	B,
	Fsharp,
	Csharp,
	Gsharp,
	Dsharp,
	Asharp,
}

func KeySignature(pitchName PitchName, majMin MajorMinor) (flatCount FlatCount, sharpCount SharpCount) {
	targetFlatOrder, targetSharpOrder := targetOrder(majMin)

	flatCount = FlatCount(countKeySignature(pitchName, targetFlatOrder))
	if flatCount == 0 {
		sharpCount = SharpCount(countKeySignature(pitchName, targetSharpOrder))
	}
	if flatCount == 0 && sharpCount == 0 {
		if pitchName != NoSignatureMajor && pitchName != NoSignatureMinor {
			//errors.New("not existed key signature")
			fmt.Println("not existed ...   will delete")
		}
	}
	return
}

func countKeySignature(pitchName PitchName, order [7]PitchName) (count int) {
	for i, pName := range order {
		if pName == pitchName {
			count = i + 1
			break
		}
	}
	return
}

func targetOrder(majMin MajorMinor) (flatOrder [7]PitchName, sharpOrder [7]PitchName) {
	if majMin == Major {
		flatOrder = FlatMajorOrder
		sharpOrder = SharpMajorOrder
	} else {
		flatOrder = FlatMinorOrder
		sharpOrder = SharpMinorOrder
	}
	return
}
