package configurable_response

import (
	"bufio"
	"net"
	"os"
	"time"
	"log"
)

const DELAY_BEFORE_MILLISEC = "delay_before_millisec"
const RESPONSE_FILE="response_file"
const DELAY_AFTER_MILLISEC = "delay_after_millisec"

func ConfigurableResponse(conn net.Conn, behaviourConfig map[string]interface{} ) {
	if behaviourConfig != nil {
		delayBeforeMillisec := behaviourConfig[DELAY_BEFORE_MILLISEC].(int)
		if  delayBeforeMillisec > 0 {
			time.Sleep(time.Duration(delayBeforeMillisec) * time.Millisecond)
		}
		responseFile := behaviourConfig[RESPONSE_FILE].(string)
		if  len(responseFile) > 0 {
			file, err := os.Open(responseFile)
			if err != nil {
				log.Fatal("Cannot open response file : " + err.Error())
			}
			defer file.Close()

			stats, statsErr := file.Stat()
			if statsErr != nil {
				log.Fatal("Cannot get size of response file : " + err.Error())
			}
			size := stats.Size()

			bytes := make([]byte, size)
			bufr := bufio.NewReader(file)
			_,err = bufr.Read(bytes)
			//TODO add replacing some values like current date by template
			conn.Write(bytes)
		}
		delayAfterMillisec := behaviourConfig[DELAY_AFTER_MILLISEC].(int)
		if  delayAfterMillisec > 0 {
			time.Sleep(time.Duration(delayAfterMillisec) * time.Millisecond)
		}
	}

}

