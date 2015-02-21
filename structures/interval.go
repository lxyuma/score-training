package structures

var IntervalOrder = [12]PitchName{
	C,
	Csharp,
	D,
	Eflat,
	E,
	F,
	Fsharp,
	G,
	Gsharp,
	A,
	Bflat,
	B,
}

var Interval = map[PitchName]int{
	Cflat:  11,
	C:      0,
	Csharp: 1,

	Dflat:  1,
	D:      2,
	Dsharp: 3,

	Eflat:  3,
	E:      4,
	Esharp: 5,

	Fflat:  4,
	F:      5,
	Fsharp: 6,

	Gflat:  6,
	G:      7,
	Gsharp: 8,

	Aflat:  8,
	A:      9,
	Asharp: 10,

	Bflat:  10,
	B:      11,
	Bsharp: 12,
}

const (
	P1 = "Perfect union"
	m2 = "Minor second"
	M2 = "Major second"
	m3 = "Minor third"
	M3 = "Major third"
	d4 = "Diminished fourth"
	P4 = "Perfect fourth"
	d5 = "Diminished fifth"
	P5 = "Perfect fifth"
	m6 = "Minor sixth"
	M6 = "Major sixth"
	m7 = "Minor seventh"
	M7 = "Major seventh"
	P8 = "Perfect Octave"
)

var IntervalNameOrder = [13]string{P1, m2, M2, m3, M3, P4, d5, P5, m6, M6, m7, M7, P8}

func IntervalNum(from PitchName, to PitchName) (num int) {
	num = Interval[to] - Interval[from]
	if num < 0 {
		num -= 12
	}
	return
}

func IntervalName(intervalNum int) string {
	return IntervalNameOrder[intervalNum]
}
