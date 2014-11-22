/* ThreadedIPEchoServer
*/
package main

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


