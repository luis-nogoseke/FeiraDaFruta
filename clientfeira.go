package main
import (
    "net/rpc"
    "fmt"
    "bufio"
    "os"
)

type Args struct {
  Name string;
  Price float64
}

var in *bufio.Reader

func readFruit () Args {
    var a string
    var b float64
    fmt.Printf("Name: ")
    bytes, _, _ := in.ReadLine()
    a = string(bytes)
    fmt.Printf("Price: ")
    fmt.Scanln(&b)
    return Args{a, b}
}

func readName () Args {
    var a string
    fmt.Printf("Name: ")
    bytes, _, _ := in.ReadLine()
    a = string(bytes)
    fmt.Printf("Price: ")
    return Args{a, 0}
}

func readKg () Args {
    var a string
    var b float64
    fmt.Printf("Name: ")
    bytes, _, _ := in.ReadLine()
    a = string(bytes)
    fmt.Printf("Price: ")
    fmt.Printf("Kg: ")
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
    in = bufio.NewReader(os.Stdin)
    var op byte
    for {
        fmt.Println("\n MENU")
        fmt.Println("1 - Adicionar Fruta")
        fmt.Println("2 - Calcular Preço")
        fmt.Println("3 - Preço por Quilo")
        fmt.Println("4 - Remover Fruta")
        fmt.Println("5 - Atualizar Preço")
        fmt.Println("6 - Sair")
        fmt.Scanf("%c\n", &op)
        switch op {
        case '1':
                args := readFruit()
                Call := client.Go("Fruit.AddFruit",	args, nil, nil)
                replyerr := <- Call.Done
                checkError("AddFruit: ", replyerr.Error)
                // fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)
                break
            case '2':
                args := readKg()
                var reply float64
                Call := client.Go("Fruit.GetPrice", args, &reply, nil)
                replyerr := <- Call.Done
                checkError("GetPrice: ", replyerr.Error)
                fmt.Printf("R$ %.2f\n", reply)
                break
            case '3':
                args := readName()
                var reply float64
                Call := client.Go("Fruit.GetPriceKg", args, &reply, nil)
                replyerr := <- Call.Done
                checkError("GetPriceKg: ", replyerr.Error)
                fmt.Printf("R$ %.2f\n", reply)
                break
            case '4':
                args := readName()
                Call := client.Go("Fruit.RemoveFruit", args, nil, nil)
                replyerr := <- Call.Done
                checkError("Remove: ", replyerr.Error)
                break
            case '5':
                args := readFruit()
                Call := client.Go("Fruit.UpdatePrice", args, nil, nil)
                replyerr := <- Call.Done
                checkError("Update: ", replyerr.Error)
                break
            case '6':
                os.Exit(0)
        }
    }
    os.Exit(0)
}
