package main

import (
	"fmt"
	"github.com/lxyuma/score-training/questions/scales"
	"github.com/lxyuma/score-training/structures"
	//	"os"
)

func main() {
	num := structures.IntervalNum(structures.C, structures.E)
	name := structures.IntervalName(num)

	fmt.Println(name)

	fmt.Println(structures.Cflat.IsSharp())

	problem := scales.NewProblem()
	fmt.Println(problem.Question)
	fmt.Println(problem.Answer)

	//buffer := make([]byte, 256)
	//for {
	//	byte_count, _ := os.Stdin.Read(buffer)
	//	if byte_count == 0 {
	//		break
	//	}
	//	os.Stdout.Write(buffer[:byte_count])
	//}
}
