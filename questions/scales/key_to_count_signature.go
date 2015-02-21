package scales

import (
	"fmt"
	"github.com/lxyuma/score-training/structures"
	"github.com/lxyuma/score-training/translations"
	"math/rand"
	"time"
)

type Problem struct {
	Question string
	Answer   *AnswerCountKeySignature
}

type AnswerCountKeySignature struct {
	Flat  structures.FlatCount
	Sharp structures.SharpCount
}

func (a *AnswerCountKeySignature) String() string {
	return fmt.Sprintf("Flat = %v, Sharp = %v", int(a.Flat), int(a.Sharp))
}

func NewProblem() (problem *Problem) {
	var pitchName structures.PitchName = structures.RandPitch()
	majMin := randMajorOrMinor()
	flatCount, sharpCount := structures.KeySignature(pitchName, majMin)
	var q string = "How many key signature you write? In " +
		translations.JpPitch[pitchName] +
		translations.JpMajorMinor[majMin]
	var a *AnswerCountKeySignature = &AnswerCountKeySignature{flatCount, sharpCount}

	problem = &Problem{Question: q, Answer: a}
	return
}

//func (p *Problem) CheckAnswer(string ans) (isRight bool, message string) {
//}

func randMajorOrMinor() (majMin structures.MajorMinor) {
	rand.Seed(time.Now().Unix())
	if rand.Intn(2) == 0 {
		majMin = structures.Major
	} else {
		majMin = structures.Minor
	}
	return
}
