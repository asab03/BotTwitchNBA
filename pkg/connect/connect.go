package connect

import (
	"fmt"
	"io"
	"log"
	"net"
)

func Connect() io.ReadWriter{
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
    if err != nil{
        log.Fatalln("Failed to connect to the IRC", err)
    }


    fmt.Fprintf(conn, "PASS oauth:mpim2eer3kkt0lqbom0r9gypiajefn \r\n")
    fmt.Fprintf(conn, "NICK BotName \r\n")
    fmt.Fprintf(conn, "JOIN #channel  \r\n")
    fmt.Fprintf(conn, "PRIVMSG #channel :Hello !\r\n")

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

