package connect

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)







func Connect() io.ReadWriter{
    godotenv.Load()
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
    if err != nil{
        log.Fatalln("Failed to connect to the IRC", err)
    }



    fmt.Fprintf(conn, "PASS %s \r\n", os.Getenv("TWITCH_AUTH") )
    fmt.Fprintf(conn, "NICK %s \r\n", os.Getenv("BOTNAME"))
    fmt.Fprintf(conn, "JOIN #%s \r\n", os.Getenv("CHANNEL"))
    fmt.Fprintf(conn, "PRIVMSG #%s :Hello !\r\n", os.Getenv("CHANNEL"))


	return conn
}

func Disconnect() io.ReadWriteCloser{
    conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
    if err != nil{
        log.Fatalln("Failed to connect to the IRC", err)
    }
    fmt.Fprintf(conn, "JOIN #  \r\n")

    return conn
}

