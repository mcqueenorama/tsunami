Build and run it:

go build
tsunami server
tsunami client
tsunami help


This is a golang implementation of tsunami-udp:

http://tsunami-udp.sourceforge.net/

I found a good c-based implementation, but it didn't look like a good fit for simply adding go bindings:

https://github.com/sebsto/tsunami-udp

Here's a good description with some psuedo-code:

http://tsunami-udp.cvs.sourceforge.net/viewvc/tsunami-udp/docs/howTsunamiWorks.txt?revision=1.5&view=markup

Here's the psuedo-code from that page:

         **Server**
	 start
	 while(running) {
	   wait(new incoming client TCP connection)
	   fork server process:
	   [
	     check_authenticate(MD5, "kitten");
	     exchange settings and values with client;
	     while(live) {
	       wait(request, nonblocking)
	       switch(request) {
	          case no request received yet: { send next block in sequence; }
	          case request_stop:            { close file, clean up; exit; }
	          case request_retransmit:      { send requested blocks; }
	       }
	       sleep(throttling)
	     }
	   ]
	 }
	

	**Client**
	 start, show command line
	 while(running) {
	    read user command;
	    switch(command) {
	       case command_exit:    { clean up; exit; }
	       case command_set:     { edit the specified parameter; }
	       case command_connect: { TCP connect to server; auth; protocol version compare;
	                               send some parameters; }
	       case command_get && connected:  { 
	           send get-file request containing all transfer parameters;
	           read server response - filesize, block count;
	           initialize bit array of received blocks, allocate retransmit list;
	           start separate disk I/O thread;
	           while (not received all blocks yet) {
	              receive_UDP();
	              if timeout { send retransmit request(); }
	
	              if block not marked as received yet in the bit array {
	                 pass block to I/O thread for later writing to disk;
	                 if block nr > expected block { add intermediate blocks to retransmit list; }
	              }
	
	              if it is time { 
	                 process retransmit list, send assembled request_retransmit to server;
	                 send updated statistics to server, print to screen;
	              } 
	           }
	           send request_stop;
	           sync with disk I/O, finalize, clean up;
	       }
	       case command_help:    { display available commands etc; }
	    }
	 }

Some testing commands:

time seq -f "%g" 1 20000 | parallel    echo {} \> /dev/udp/localhost/1200
for qqq in `seq 1 20`; do echo a $qqq ; done | parallel --pipe --line-buffered -N 1 -L 0 --round nc -u localhost 1200
for qqq in `seq 1 20`; do echo a $qqq ; done | parallel --pipe cat - \> /dev/udp/localhost/1200
