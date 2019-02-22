package main

import (
	"math/rand"
	"time"
	"fmt"
	"os"
	"strconv"
)

//测试I/O密集型计算的性能

//const NUM = 10000000



func swap(a int, b int) (int, int) {
	return b, a
}

func partition(aris []int, begin int, end int) (int) {
	pvalue := aris[begin]
	i := begin
	j := begin + 1
	for j < end {
		if aris[j] < pvalue {
			i++
			aris[i], aris[j] = swap(aris[i], aris[j])
		}
		j++
	}
	aris[i], aris[begin] = swap(aris[i], aris[begin])
	return i
}

func qsort(aris []int, begin int, end int) {
	if begin+1 < end {
		mid := partition(aris, begin, end)
		qsort(aris, begin, mid)
		qsort(aris, mid+1, end)
	}
}

func randArr(aris []int, len int) {
	for i := 0; i < len; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		aris[i] = r.Intn(1000)
	}
}

func main() {
	arg := "100000"

	if len(os.Args) == 2 {
		arg = os.Args[1]
	}
	NUM,error := strconv.Atoi(arg)
	if error != nil{
		fmt.Println(" error")
	}
	var intas []int
	for i := 0; i < NUM; i++ {
		intas = append(intas,rand.Intn(100000))
	}
	fmt.Println("BEGIN SORT")
	beginTime := time.Now()
	qsort(intas, 0, NUM)
	endTime := time.Now()
	fmt.Println("END SORT")
	fmt.Println(endTime.Sub(beginTime))
}
