package close_connection

import (
	"net"
	"time"
)

const DELAY_MILLISEC = "delay_millisec"

func CloseConnection(conn net.Conn, behaviourConfig map[string]interface{} ) {
	if behaviourConfig != nil {
		delayMillisec := behaviourConfig[DELAY_MILLISEC].(int)
		if  delayMillisec > 0 {
			time.Sleep(time.Duration(delayMillisec) * time.Millisecond)
		}
	}
}
