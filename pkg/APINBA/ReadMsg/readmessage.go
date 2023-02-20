package readmsg

import (
	"bufio"
	"io"
	"log"
)

func ReadMessages(conn io.Reader)<- chan string {
	chatReader := bufio.NewReader(conn)
	msgsChan := make(chan string)

	go func(){
		for {
			msg, err := chatReader.ReadString('\n')
			if err != nil{
				log.Fatalln("Failed to read the msg", err)
			}

			msgsChan <- msg
		}
	}()

	return msgsChan
}

