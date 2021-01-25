# http-server-fault-emulator
HTTP server to emulate faults (close TCP connection, timeout etc)

This is small utility to emulate HTTP server failures that happens when an HTTP server is overloaded or faulty. The goal is to use this utility to reproduce responses that cannot be produced using a utility that is based on a conventional HTTP server like Gin.
A client uses the utility instead of real server could experience various issues:
 Closed connection - after establishing the TCP connection (there could be some delay) the server closes the TCP connection
 Empty response  - after establishing the TCP connection (there could be some delay) the server sends end of HTTP packet  closes the TCP connection  
