package writemessages

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func WriteMessages(conn io.Writer, source io.Reader){
	godotenv.Load()
	msgWritter := bufio.NewReader(source)

	go func(){
		for{
			msg,err := msgWritter.ReadString('\n')
			fmt.Println(msg)
			if err != nil{
				log.Fatalln("failed to read the message frome source", err)
			}
			fmt.Fprintf(conn, "PRIVMSG #%s: %s\r\n",os.Getenv("CHANNEL"), msg)
		}
	}()

} 