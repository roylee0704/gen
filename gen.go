package main

import "fmt"

func main() {
	//br := testing.Benchmark(BenchmarkFunction)
	//brn := testing.Benchmark(BenchmarkNormal)

	//fmt.Println(br)
	//fmt.Println("har?", brn)
	fmt.Println(next([]int{0, 9, 9, 9, 9}, action))
}

// func BenchmarkNormal(b *testing.B) {
// 	values := make([]int, 6)
// 	for i := 0; i < b.N; i++ {
// 		gen(values, 100000)
// 	}
// }
//
// func BenchmarkFunction(b *testing.B) {
// 	values := make([]int, 6)
// 	for i := 0; i < b.N; i++ {
// 		gen(values, 100000)
// 	}
// }

func next(values []int, action func([]int) []int) []int {
	return action(values)
}

func action(values []int) []int {
	n := len(values)
	datas := make([]int, n)
	curPt := n - 1

	copy(datas, values)

	if !isMax(datas[curPt]) {
		incr(&datas[curPt])
	} else {
		clock(curPt, datas)
	}
	return datas
}

func gen(values []int, max int) {

	curPt := len(values) - 1

	for i := 0; i < max; i++ {
		if !isMax(values[curPt]) {
			incr(&values[curPt])
		} else {
			if clock(curPt, values) {
				break
			}
		}
		//fmt.Printf("%v\n", values)
	}
}

func incr(digit *int) {
	(*digit)++
}

func isMax(digit int) bool {
	return digit == 9
}

func reset(digit *int) {
	*digit = 0
}

func clock(curPt int, values []int) (stop bool) {

	stop = true
	for p := curPt; p > 0; p-- {
		// incr front unit and leave
		if isMax(values[p-1]) {
			continue
		} else {
			incr(&values[p-1])
			for k := p; k <= curPt; k++ {
				reset(&values[k])
			}
			stop = false
			break
		}
	}
	return
}
