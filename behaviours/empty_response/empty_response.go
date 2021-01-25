package empty_response

import (
	"net"
	"time"
)

const DELAY_MILLISEC = "delay_millisec"
const empty_http_response=`HTTP/1.1 200 OK
Date: Mon, 23 May 2005 22:38:34 GMT
Content-Type: application/json; charset=UTF-8
Content-Length: 0
Last-Modified: Wed, 08 Jan 2003 23:11:55 GMT
Server: Apache/1.3.3.7 (Unix) (Red-Hat/Linux)
ETag: "3f80f-1b6-3e1cb03b"
Accept-Ranges: bytes
Connection: close

`

func EmptyResponse(conn net.Conn, behaviourConfig map[string]interface{} ) {
	if behaviourConfig != nil {
		delayMillisec := behaviourConfig[DELAY_MILLISEC].(int)
		if  delayMillisec > 0 {
			time.Sleep(time.Duration(delayMillisec) * time.Millisecond)
		}
	}
	conn.Write([]byte(empty_http_response))
}

