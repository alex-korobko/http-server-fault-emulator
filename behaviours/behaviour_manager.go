package behaviours

import (
	configurable_response "github.com/alex-korobko/http-server-fault-emulator/behaviours/configurable_response"
	"net"
)

var supportedBehaviours = map[string] func(net.Conn, map[string]interface{}) {
	"configurable_response": configurable_response.ConfigurableResponse,
}

func GetBehaviourFunc(behName string)func(conn net.Conn, behavParams map[string]interface{} ){
 return supportedBehaviours[behName]
}
