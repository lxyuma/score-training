package translations

import (
	"github.com/lxyuma/score-training/structures"
)

var JpPitch = map[structures.PitchName]string{
	structures.C:      "ハ",
	structures.D:      "ニ",
	structures.E:      "ホ",
	structures.F:      "ヘ",
	structures.G:      "ト",
	structures.A:      "イ",
	structures.B:      "ロ",
	structures.Csharp: "嬰ハ",
	structures.Dsharp: "嬰イ",
	structures.Esharp: "嬰ホ",
	structures.Fsharp: "嬰ヘ",
	structures.Gsharp: "嬰ホ",
	structures.Asharp: "嬰イ",
	structures.Bsharp: "嬰ロ",
	structures.Cflat:  "変ハ",
	structures.Dflat:  "変イ",
	structures.Eflat:  "変ホ",
	structures.Fflat:  "変ヘ",
	structures.Gflat:  "変ホ",
	structures.Aflat:  "変イ",
	structures.Bflat:  "変ロ",
}

var JpMajorMinor = map[structures.MajorMinor]string{
	structures.Major: "長調",
	structures.Minor: "短調",
}
