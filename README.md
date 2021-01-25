# http-server-fault-emulator
HTTP server to emulate faults (close TCP connection, timeout etc)
##Overview
This is small utility to emulate HTTP server failures that happens when an HTTP server is overloaded or faulty. The goal is to use this utility to reproduce responses that cannot be produced using a utility that is based on a conventional HTTP server like Gin.

A client uses the utility instead of real server could experience various issues:
 * Closed connection - after establishing the TCP connection (there could be some delay) the server closes the TCP connection
 * Empty response  - after establishing the TCP connection (there could be some delay) the server sends end of HTTP packet  closes the TCP connection
 * Return any binary response (sample stored in file) and closes the connection, to emulate interrupted process of data transmission.
 
 ## How to use
 
 To start the server run `` go run main.go``. The configuration file `config.yaml` is in the root of the project directory. The config file can be edited while server is running and changes to all properties (except the `emulationport` property, change to this property requires restarting the service) are applied on fly.
 
 The server behaviour to request to any path on the service is configured in the `behavior` section of the config file.  Currently only one behaviour is supported, but as it covers my needs it is good enough. 
 
 The following configuration  
   ```
  {
    behaviour: configurable_response,
    emulationport: "9000",
    configurable_response: {
        delay_before_millisec: 2000,
    }
  }
  ```
   
  starts server listening to connections on port `9000`. The active behaviour is `configurable_response` that is configured after accepting a connection the service to wait for 2 sec and close the connection.
  
 
 The following configuration
 
 ```
{
  behaviour: configurable_response,
  emulationport: "9000",
  configurable_response: {
      delay_before_millisec: 100,
      response_file: "response_sample2.txt",
      delay_after_millisec: 200,
  }
}
```
 
 starts server listening to connections on port `9000`. The active behaviour is `configurable_response` that is configured to delay after accpting a connection for 100 milliseconds, then read content for the `response_sample2.txt` file in binary mode and send it via socket, then sleep for another `200` milliseconds and close the connection.
 As it was mentioned, the config could be changed, for example, modifying the delays or using another file and the changes will be applied to the next request to the service. The file is not cached, so any modification in the file also applied in the next request to the service.
 
 
 ##TODO 
 * to update the response that was loaded from file replacing by template some values, like current date and time
 * to support different behaviours\behaviour params for different paths
 * to log what exactly happens for the request
 
 
 
 
 
