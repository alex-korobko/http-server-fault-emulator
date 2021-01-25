package behaviours

import (
	close_connection "github.com/alex-korobko/http-server-fault-emulator/behaviours/close_connection"
	empty_response "github.com/alex-korobko/http-server-fault-emulator/behaviours/empty_response"
	"net"
)

var supportedBehaviours = map[string] func(net.Conn, map[string]interface{}) {
	"close_connection": close_connection.CloseConnection,
	"empty_response": empty_response.EmptyResponse,
}

func GetBehaviourFunc(behName string)func(conn net.Conn, behavParams map[string]interface{} ){
 return supportedBehaviours[behName]
}
