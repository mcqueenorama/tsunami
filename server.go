/* ThreadedIPEchoServer
*/
package main

import (
    "fmt"
    "net"
    "os"
    "runtime"
)

const (
    NCPU = 8 
    //I can test this version with echo %%%% | netcat localhost 1200
    // TTP_VERSION = int32(0x25252525)
    //this is the correct version
    //TTP_VERSION = int32(0x20061025)
    BUF_SIZE = 512
    SERVICE = ":1200"
    PASSWORD = "kitten"
)

var cmdServer = &Command{
    Run:       runServer,
    UsageLine: "server",
    Short:     "start up a server",
    Long:      "start up a server",
}

// func main() {
func runServer(cmd *Command, args []string) bool {

    fmt.Fprintf(os.Stderr, "number of cpus:%d:\n", NCPU)
    fmt.Fprintf(os.Stderr, "number of cpus as reported by go:%d:\n", runtime.NumCPU())
    runtime.GOMAXPROCS(NCPU)

    command_listener, err := net.Listen("tcp", SERVICE)
    checkError(err)
    defer command_listener.Close()

    go listenTCPCommands(command_listener)

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


func listenTCPCommands(listener net.Listener) {

    for {

        conn, err := listener.Accept()
        if err != nil {
            fmt.Fprintf(os.Stderr, "failed to accept:retrying:%t:\n", err)
            continue
        }

        rv, err := protocol.TtpNegotiate(conn)
        if err != nil {
            checkError(err)
            return
        }

        if rv != true {
            fmt.Fprintf(os.Stderr, "protocol mismatch:continuing anyway:\n")
            // return
        }

        go handleTCPCommands(conn)
    }

}

//handle commands on the tcp connection
//this need something better than silent death upon read/write error
//for now just be an echo server
func handleTCPCommands(conn net.Conn) {

    var buf [BUF_SIZE]byte

    for {

        //go into echo server mode
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return
        }
    }

}

func handleData(conn net.PacketConn, cpu int) {

    var buf [BUF_SIZE]byte

    for {

    fmt.Fprintf(os.Stderr, "about to read:cpu:%d:\n", cpu)
    n, addr, err := conn.ReadFrom(buf[0:])
    fmt.Fprintf(os.Stderr, "read bytes:%d:cpu:%d:buf:%s:\n", n, cpu, buf)
    if err != nil {
        return
    }

    wrote, err2 := conn.WriteTo(buf[0:n], addr)
    fmt.Fprintf(os.Stderr, "wrote bytes:%d:cpu:%d:\n", wrote, cpu)
    if err2 != nil {
        return
    }
  }

}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

