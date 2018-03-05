package main

import (
	"fmt"
)

var cursor int

type QueryResults interface {
	fetchData()
}

type QueryResultsImpl struct {
}

func (qr QueryResultsImpl) fetchData() [5]int {
	var arry [5]int
	arry[0] = 5
	arry[1] = 15
	arry[2] = 25
	arry[3] = 35
	arry[4] = 45
	return arry
}

type QueryResultsIteratorImpl struct {
	max int
	ar  [5]int
}

func (qry QueryResultsIteratorImpl) fst() int {
	fmt.Println("FIRST : ", qry.ar[0])
	return 0
}

func (qry QueryResultsIteratorImpl) nxt() int {
	if qry.Done() {
		return 1
	} else {
		cursor = qry.current() + 1
		fmt.Println("NEXT  : ", qry.ar[cursor])
		return 0
	}
}

func (qry QueryResultsIteratorImpl) current() int {
	return cursor
}

func (qry QueryResultsIteratorImpl) Done() bool {
	return cursor == (qry.max)-1
}

func main() {
	var f int
	var newarry [5]int
	var qr QueryResultsImpl
	newarry = qr.fetchData()
	queryimpl := QueryResultsIteratorImpl{len(newarry), newarry}
	queryimpl.fst()
	for {
		f = queryimpl.nxt()
		if f == 1 {
			fmt.Println("<end>")
			break
		}
	}
}
