package main

import (
    "fmt"
    "bytes"
    "encoding/binary"
    "os"
)

func ToNetworkByteOrder(val int32) ([]byte, error) {

    iobuf := new(bytes.Buffer)

    err := binary.Write(iobuf, binary.BigEndian, val)
    if err != nil {
        fmt.Println("toNetworkByteOrder:binary.Write failed:", err)
        return iobuf.Bytes(), err
    }
    // fmt.Printf("server version:%x:\n", iobuf.Bytes())

    return iobuf.Bytes(), nil

}

func FromNetworkByteOrder(iobuf []byte) (rv int32, err error) {

    bufReader := bytes.NewReader(iobuf)

    err = binary.Read(bufReader, binary.LittleEndian, &rv)
    if err != nil {
        fmt.Println("binary.Read failed:", err)
        return rv, err
    }    

    return rv, nil

}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        // os.Exit(1)
    }
}
