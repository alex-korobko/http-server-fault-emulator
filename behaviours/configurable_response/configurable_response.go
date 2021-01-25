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
		configParam := behaviourConfig[DELAY_BEFORE_MILLISEC]
		if configParam!= nil {
			delayBeforeMillisec := configParam.(int)
			if delayBeforeMillisec > 0 {
				time.Sleep(time.Duration(delayBeforeMillisec) * time.Millisecond)
			}
		}

		configParam = behaviourConfig[RESPONSE_FILE]
		if configParam != nil {
			responseFile := configParam.(string)
			if len(responseFile) > 0 {
				file, openFileErr := os.Open(responseFile)
				if openFileErr != nil {
					log.Fatal("Cannot open response file : " + openFileErr.Error())
				}
				defer file.Close()

				stats, statsErr := file.Stat()
				if statsErr != nil {
					log.Fatal("Cannot get size of response file : " + statsErr.Error())
				}
				size := stats.Size()

				bytes := make([]byte, size)
				bufr := bufio.NewReader(file)
				_, readBuffErr := bufr.Read(bytes)
				if readBuffErr != nil {
					log.Fatal("Cannot read the response file to buffer : " + readBuffErr.Error())
				}
				//TODO add replacing some values like current date by template
				_, writeToSocketErr := conn.Write(bytes)
				if writeToSocketErr != nil {
					log.Fatal("Cannot read the response file to buffer : " + writeToSocketErr.Error())
				}
			}
		}
		configParam = behaviourConfig[DELAY_AFTER_MILLISEC]
		if configParam != nil {
			delayAfterMillisec := configParam.(int)
			if delayAfterMillisec > 0 {
				time.Sleep(time.Duration(delayAfterMillisec) * time.Millisecond)
			}
		}
	}

}

