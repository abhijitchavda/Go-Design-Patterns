package main

import "fmt"

var succ1 handler
var succ2 handler
var succ3 handler

type handler interface {
	handleRequest(request string)
	setSuccessor(next handler)
}

type controllerone struct {
	R1 string
}

type controllertwo struct {
	R2 string
}

type controllerthree struct {
	R3 string
}

func (con1 controllerone) handleRequest(request string) {
	fmt.Println("R1 got the request...")
	if request == "R1" {
		fmt.Println("controller one : This one is mine!")
	} else {
		if succ1 != nil {
			succ1.handleRequest(request)
		}
	}
}

func (con1 controllerone) setSuccessor(next handler) {

	succ1 = next
}

func (con2 controllertwo) handleRequest(request string) {
	fmt.Println("R2 got the request...")
	if request == "R2" {
		fmt.Println("controller two : This one is mine!")
	} else {
		if succ2 != nil {
			succ2.handleRequest(request)
		}
	}
}
func (con2 controllertwo) setSuccessor(next handler) {

	succ2 = next

}

func (con3 controllerthree) handleRequest(request string) {
	fmt.Println("R3 got the request...")
	if request == "R3" {
		fmt.Println("controller three : This one is mine!")
	} else {
		if succ3 != nil {
			succ3.handleRequest(request)
		} else {
			fmt.Println("no processes could handle this request")
		}
	}
}
func (con3 controllerthree) setSuccessor(next handler) {

	succ3 = next
}

func runTest(h1 handler, h2 handler, h3 handler) {
	h1.setSuccessor(h2)
	h2.setSuccessor(h3)
	h3.setSuccessor(nil)

	fmt.Println("Req2 sent >>>")
	h1.handleRequest("R2")

	fmt.Println("Req3 sent >>>")
	h1.handleRequest("R3")

	fmt.Println("Req1 sent >>>")
	h1.handleRequest("R1")

	fmt.Println("Sending RX...")
	h1.handleRequest("RX")

}

func main() {
	con1 := controllerone{"R1"}
	con2 := controllertwo{"R2"}
	con3 := controllerthree{"R3"}
	fmt.Print(succ3)
	runTest(con1, con2, con3)
}
