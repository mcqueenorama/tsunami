package main

import (
    "fmt"
    "net"
)

const (
    //I can test this version with echo %%%% | netcat localhost 1200
    TTP_VERSION = int32(0x25252525)
    //this is the correct version
    //TTP_VERSION = int32(0x20061025)
)

// negotiate the protocol version
// send out my version
//  get it in network byte order
//  this bit seems stupid, certainly there's a better way
// get their version
// return true if theirs matches mine
func TtpNegotiate(conn net.Conn) (rv bool, err error)  {

    serverVersion, err := ToNetworkByteOrder(TTP_VERSION)
    if err != nil {
        return rv, err
    }
    fmt.Printf("server version:%x:\n", serverVersion)

    n, err := conn.Write(serverVersion)
    if err != nil {
        return rv, err
    }
    fmt.Printf("wrote:%d:bytes to client\n", n)

    iobuf := make([]byte, 100)

    n, err = conn.Read(iobuf[0:])
    if err != nil {
        return rv, err
    }
    fmt.Printf("read:%d:bytes from client\n", n)
  
    clientVersion, err := FromNetworkByteOrder(iobuf)
    if err != nil {
        return rv, err
    }  
    fmt.Printf("clientVersion:%x:\n", clientVersion)

    if clientVersion != TTP_VERSION {
        return false, nil
    }

    return true, nil;

}