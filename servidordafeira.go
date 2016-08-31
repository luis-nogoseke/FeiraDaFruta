package main
import (
 	"fmt"
	"net/rpc"
	"net"
	"os"
)

type Args struct {
  Name string;
  Price float64
}

func checkError(st string, err error) {
    if err != nil {
        fmt.Println(st)
		os.Exit(1)
    }
}


type Fruit int

func (t *Fruit) AddFruit (args *Args, reply *int) error {
	fmt.Println("AddFruit: ", args.Name, args.Price)
	return nil
}

func (t *Fruit) UpdatePrice (args *Args, reply *int) error {
	fmt.Println("UpdatePrice", args.Name, args.Price)
	return nil
}

func (t *Fruit) RemoveFruit (args *Args, reply *int) error {
	fmt.Println("", args.Name, args.Price)
	return nil
}

func (t *Fruit) GetPriceKg (args *Args, reply *int) error {
	fmt.Println("GetPriceKg", args.Name, args.Price)
	return nil
}

func (t *Fruit) GetPrice (args *Args, reply *int) error {
	fmt.Println("GetPrice", args.Name, args.Price)
	return nil
}



func main () {
	fruit := new(Fruit)
	rpc.Register(fruit)
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:1234")
	checkError("ResolveTCPAddr: ", err)
	listener, err :=	net.ListenTCP("tcp", tcpAddr)
	checkError("ListenTCP: ", err)
	for {
		conn, err := listener.Accept()
		checkError("Accept: ", err)
		rpc.ServeConn(conn)
	}
}

