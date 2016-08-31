package main
import (
	"net/rpc"
	"fmt"
	"os"
)
type Args struct {
  Name string;
  Price float64
}

func readArgs () Args {
	var a string
	var b float64
	fmt.Println("A: ")
	fmt.Scanln(&a)
	fmt.Println("B: ")
	fmt.Scanln(&b)
	return Args{a, b}
}

func checkError(st string, err error) {
    if err != nil {
        fmt.Println(st, err)
		os.Exit(1)
    }
}

func main () {
	service := "localhost:1234"
	client, err := rpc.Dial("tcp", service)
	defer client.Close()
	checkError("Dial: ", err)
	fmt.Println("* - multiplicação")
	fmt.Println("/ - divisão")
	var op byte
	fmt.Scanf("%c\n", &op)
	switch op {
		case '*':
			args := readArgs()
			var reply int
			mulCall := client.Go("Fruit.AddFruit",	args, &reply, nil)
			replyMul := <- mulCall.Done
			checkError("Multiply: ", replyMul.Error)
			// fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)
			os.Exit(0)
		case '/':
			args := readArgs()
			var reply int
			divCall := client.Go("Fruit.GetPriceKg", args, &reply, nil)
			replyDiv := <- divCall.Done
			checkError("Divide: ", replyDiv.Error)
			// fmt.Printf("%d / %d = (%d,%d)\n", args.A, args.B, reply.Q, reply.R)
			os.Exit(0)
	}
}