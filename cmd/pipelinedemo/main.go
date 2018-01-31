package main

import (
	"test/newProject/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	//mergeDemo()

	const filename = "large.in"
	const n = 100000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	//defer声明该语句在函数执行完成之前执行
	defer file.Close()
	p := pipeline.RandomSource(n)
	write:=bufio.NewWriter(file)
	pipeline.WriteSink(write, p)
	write.Flush()
	//file.Close()
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReaderSource(bufio.NewReader(file),-1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count>=100 {
			break
		}
	}
}

func mergeDemo() {
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 8, 13)))
	//for {
	//	if num,ok := <- p;ok{
	//		fmt.Println(num)
	//	}else {
	//		break
	//	}
	//}
	for v := range p {
		fmt.Println(v)
	}
}
