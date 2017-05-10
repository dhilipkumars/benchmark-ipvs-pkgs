package main

import (
	"testing"
	//"fmt"

	//"github.com/docker/libnetwork/testutils"

)



func BenchmarkSingleServiceLibNet(b *testing.B) {

	//defer testutils.SetupTestOSContext(t)()
	InitLibNet()
	//Run the benchmark for libnet
	for n:=0; n< b.N; n++ {
		err := SingleServiceLibNet()
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkSingleServiceSeeSaw(b *testing.B) {

	//defer testutils.SetupTestOSContext(t)()
	InitSeeSaw()
	//Run the benchmark for libnet
	for n:=0; n< b.N; n++ {
		err := SingleServiceSeeSaw()
		if err != nil {
			b.Fail()
		}
	}
}