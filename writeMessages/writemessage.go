package writemessages

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

func WriteMessages(conn io.Writer, source io.Reader){
	msgWritter := bufio.NewReader(source)

	go func(){
		for{
			msg,err := msgWritter.ReadString('\n')
			fmt.Println(msg)
			if err != nil{
				log.Fatalln("failed to read the message frome source", err)
			}
			fmt.Fprintf(conn, "PRIVMSG #narvalo03: %s\r\n", msg)
		}
	}()

} 