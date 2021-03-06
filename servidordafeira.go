package main
import (
    "fmt"
    "net/rpc"
    "net"
    "os"
    "encoding/csv"
    "io"
    "strconv"
    "errors"
)

var reader *csv.Reader
var writer *csv.Writer
var database map[string]float64

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

func printMap(m map[string]float64) {
    for key, value := range m {
    fmt.Println(key, "=", value)
    }
}


func UpdateCsv (m map[string]float64) {
    os.Remove("feira.csv")
    file, err := os.Create("feira.csv")
    checkError("Cannot create file", err)
    writer = csv.NewWriter(file)
    writer.Comma = ';'
    for name, price := range m {
        s := []string{name, fmt.Sprintf("%.2f", price)}
        err := writer.Write(s)
        writer.Flush()
        checkError("Cannot write to file", err)
  }
}

func (t *Fruit) AddFruit (args *Args, reply *int) error {
    fmt.Println("AddFruit: ", args.Name, args.Price)

    if _, ok := database[args.Name]; ok {
        return errors.New("Fruit already registered")
    }
    database[args.Name] = args.Price
    UpdateCsv(database)
    return nil
}

func (t *Fruit) UpdatePrice (args *Args, reply *int) error {
    fmt.Println("UpdatePrice", args.Name, args.Price)
    if _, ok := database[args.Name]; ok {
        database[args.Name] = args.Price
        UpdateCsv(database)
        return nil
    }
    return errors.New("Fruit not registered")
}

func (t *Fruit) RemoveFruit (args *Args, reply *int) error {
    fmt.Println("Remove", args.Name, args.Price)
    if _, ok := database[args.Name]; ok {
        delete(database, args.Name)
        UpdateCsv(database)
        return nil
    }
    return errors.New("Fruit not registered")
}

func (t *Fruit) GetPriceKg (args *Args, reply *float64) error {
    fmt.Println("GetPriceKg", args.Name)
    if _, ok := database[args.Name]; ok {
        *reply = database[args.Name]
        return nil
    }
    return errors.New("Fruit not registered")
}

func (t *Fruit) GetPrice (args *Args, reply *float64) error {
    fmt.Println("GetPrice", args.Name, args.Price)
    if _, ok := database[args.Name]; ok {
        *reply = database[args.Name] * args.Price
        return nil
    }
    return errors.New("Fruit not registered")
}



func main () {
    database = map[string]float64{}
    file, err := os.Open("feira.csv")
    if err == nil {
        reader = csv.NewReader(file)
        reader.Comma = ';'
        for {
            record, err := reader.Read()
            // Stop at EOF.
            if err == io.EOF {
            break
            }
        database[record[0]], _ = strconv.ParseFloat(record[1],64)
        }
        file.Close()
    }

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
