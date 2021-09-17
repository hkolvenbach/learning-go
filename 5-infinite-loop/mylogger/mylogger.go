package mylogger

import "log"

func ListenForLog(ch chan string) {
	for {
		log.Println(<-ch)
	}
}
