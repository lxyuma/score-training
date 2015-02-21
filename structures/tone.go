package structures

type Tone struct {
	Tonic         PitchName
	ScoreFlatNum  int
	ScoreSharpNum int
}

func NewTone(tonic PitchName) *Tone {
	tone := &Tone{Tonic: tonic}
	return tone
}

//func Tonic() String {
//
//}

//func main() {
//	_ = NewTone(C)
//}
