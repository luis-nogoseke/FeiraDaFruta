package main
import (
 	"fmt"
	"net/rpc"
	"net"
	"os"
	"encoding/csv"
)

var reader *csv.Reader
var writer *csv.Writer

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
	s := []string{args.Name, fmt.Sprintf("%.2f", args.Price)}
	err := writer.Write(s)
	writer.Flush()
    checkError("Cannot write to file", err)
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

	file, err := os.Open("feira.csv")
	if err != nil {
		// se nao existir cria
		file, err = os.Create("feira.csv")
		checkError("Cannot create file", err)
	}
	defer file.Close()
	reader = csv.NewReader(file)
	writer = csv.NewWriter(file)
	reader.Comma = ';'
	writer.Comma = ';'
	
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

