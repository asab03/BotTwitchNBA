package parser

import (
	"fmt"
	"strings"
)

func Parse(msgsChan <-chan string) <-chan string {
	parsedChan := make(chan string)
	//conn := connect.Connect()

	go func() {
		for msg := range msgsChan {

			//fmt.Println(msg)
			/*regex := `:(.*)!(.*)@(.*).tmi.twitch.tv PRIVMSG #narvalo03 :${(.*)}`
			re,err := regexp.Compile(regex)
			if err != nil{
			    log.Fatalln("failed to compile the regex")
			}
			result := re.FindStringSubmatch(msg)
			if len(result) == 0 {
				fmt.Println(result)
			}*/

			numberOfColons := strings.Count(msg, ":")
			if numberOfColons >= 2 {
				s := strings.Split(msg, ":")
				user := strings.Split(s[1], "!")[0]
				messageData := s[2]
				
				
				parsedChan <- fmt.Sprintf("viewer: %s  , message: %s \n", user, messageData)
						
			}
			
		}
	}()

	return parsedChan
}