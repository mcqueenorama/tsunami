/* ThreadedIPEchoServer
*/
package main

import (
    "fmt"
    "net"
    "os"
    "runtime"
    "io/ioutil"

)

var cmdClient = &Command{
    Run:       runClient,
    UsageLine: "client",
    Short:     "start up a client",
    Long:      "start up a client",
}

// func main() {
// if len(os.Args) != 2 {
// fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
// os.Exit(1)
// }
// service := os.Args[1]
// tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
// checkError(err)
// conn, err := net.DialTCP("tcp", nil, tcpAddr)
// checkError(err)
// _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
// checkError(err)
// //result, err := readFully(conn)
// result, err := ioutil.ReadAll(conn)
// checkError(err)
// fmt.Println(string(result))
// os.Exit(0)
// }


// func main() {
func runClient(cmd *Command, args []string) bool {

    fmt.Fprintf(os.Stderr, "number of cpus:%d:\n", NCPU)
    fmt.Fprintf(os.Stderr, "number of cpus as reported by go:%d:\n", runtime.NumCPU())
    runtime.GOMAXPROCS(NCPU)

    tcpAddr, err := net.ResolveTCPAddr("tcp4", SERVICE)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    // _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    // checkError(err)

    _, err = TtpNegotiate(conn)
    if err != nil {
        checkError(err)
        return false
    }

    result, err := ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))

    // data_listener, err := net.ListenPacket("udp", service)
    // checkError(err)
    // defer data_listener.Close()

    // for i := 0; i < NCPU; i++ {
    //   go handleData(data_listener, i)
    // }

    var input string
    fmt.Scanln(&input)

    return true

}


//handle commands on the tcp connection
//this need something better than silent death upon read/write error
//for now just be an echo server


