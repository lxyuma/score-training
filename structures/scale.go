package structures

type MajorMinor bool

const (
	Major MajorMinor = true
	Minor MajorMinor = false
)

var MajorScale = [7]int{2, 2, 1, 2, 2, 2, 1}
var NaturalMinorScale = [7]int{2, 1, 2, 2, 1, 2, 2}
var HarmonicMinorScale = [7]int{2, 1, 2, 2, 1, 3, 1}
var MelodicMinorUp = [7]int{2, 1, 2, 2, 2, 2, 1}
var MelodicMinorDown = [7]int{2, 2, 1, 2, 2, 1, 2}
