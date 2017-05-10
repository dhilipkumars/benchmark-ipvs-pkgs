package main

import (
	"fmt"
	"net"
	"syscall"

	libnet "github.com/docker/libnetwork/ipvs"
	"github.com/docker/libnetwork/ns"

	seesaw "github.com/google/seesaw/ipvs"
	"github.com/vishvananda/netlink/nl"

)

type Common struct {
	LN *libnet.Handle
}

var C Common

func Init() error {

	if err := InitSeeSaw(); err != nil {
		return err
	}

	if err := InitLibNet(); err != nil {
		return err
	}

	return nil

}

func InitSeeSaw() error {

	err := seesaw.Init()
	if err != nil {
		fmt.Printf("Error: Initalizing Seesaw err=%v\n", err)
		return err
	}
	return nil
}

func InitLibNet() error {

	var err error
	ns.Init()
	C.LN, err = libnet.New("")
	if err != nil {
		fmt.Printf("Error: Libnetwork intialization failed err=%v\n", err)
		return err
	}

	return nil

}

func SingleServiceSeeSaw() error {
	var S seesaw.Service
	S.Address = net.ParseIP("1.2.3.4")
	S.Port = 8181
	S.Protocol = syscall.IPPROTO_TCP
	S.Scheduler = "rr"

	err := seesaw.AddService(S)
	if err != nil {
		fmt.Printf("Error: SeeSaw AddService err=%v\n", err)
		return err
	}

	var D seesaw.Destination

	D.Address = net.ParseIP("1.2.3.4")
	D.Port = 63333
	err = seesaw.AddDestination(S, D)
	if err != nil {
		fmt.Printf("Error: SeeSaw AddDestination err=%v\n", err)
		return err
	}

	_, err = seesaw.GetService(&S)
	if err != nil {
		fmt.Printf("Error: SeeSaw GetService err=%v\n", err)
		return err
	}

	S.Scheduler = "lc"
	err = seesaw.UpdateService(S)
	if err != nil {
		fmt.Printf("Error: SeeSaw UpdateService err=%v\n", err)
		return err
	}

	D.Flags = seesaw.DFForwardTunnel
	err = seesaw.UpdateDestination(S, D)
	if err != nil {
		fmt.Printf("Error: SeeSaw UpdateDestination err=%v\n", err)
		return err
	}

	err = seesaw.DeleteDestination(S, D)
	if err != nil {
		fmt.Printf("Error: SeeSaw DeleteDestination err=%v\n", err)
		return err
	}

	err = seesaw.DeleteService(S)
	if err != nil {
		fmt.Printf("Error: SeeSaw DeleteService err=%v\n", err)
		return err
	}

	return nil

}

func SingleServiceLibNet() error {

	var S libnet.Service

	S.AddressFamily = nl.FAMILY_V4

	S.Protocol = syscall.IPPROTO_TCP
	S.Port = 80
	S.Address = net.ParseIP("1.2.3.4")
	S.Netmask = 0xFFFFFFFF
	S.SchedName = libnet.RoundRobin

	err := C.LN.NewService(&S)

	if err != nil {
		fmt.Printf("Error: Libnet NewService(%v) err=%v\n", S, err)
		return err
	}

	var D libnet.Destination

	D.Address = net.ParseIP("10.1.1.2")
	D.AddressFamily = nl.FAMILY_V4
	D.Weight = 4
	D.ConnectionFlags = libnet.ConnectionFlagDirectRoute
	D.Port = 5000

	err = C.LN.NewDestination(&S, &D)
	if err != nil {
		fmt.Printf("Error: Libnet NewDesination err=%v\n", err)
		return err
	}

	if !C.LN.IsServicePresent(&S) {
		fmt.Printf("Error: Faild to fetch the service")
		return fmt.Errorf("Failed to Lookup Service")
	}

	S.SchedName = libnet.LeastConnection

	err = C.LN.UpdateService(&S)
	if err != nil {
		fmt.Printf("Error: Libnet UpdateService err=%v\n", err)
		return err
	}

	D.Weight = 5
	err = C.LN.UpdateDestination(&S, &D)
	if err != nil {
		fmt.Printf("Error: Libnet UpdateDestination err=%v\n", err)
		return err
	}

	err = C.LN.DelDestination(&S, &D)
	if err != nil {
		fmt.Printf("Error: Libnet DelDestination err=%v\n", err)
		return err
	}

	err = C.LN.DelService(&S)
	if err != nil {
		fmt.Printf("Error: Libnet DelService err=%v\n", err)
		return err
	}

	return nil

}

func main() {

	//Initialize both the libraries
	if Init() != nil {
		fmt.Printf("Error Initalization failed\n")
		return
	}


	if SingleServiceLibNet() != nil {
		fmt.Printf("Single service on libnetwork failed\n")
	}

	if SingleServiceSeeSaw() != nil {
		fmt.Printf("Single Service for SeeSaw failed\n")
	}

	return
}
