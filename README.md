# Lancer l'application :

go run .

# Choisir son joueur :

L'application vous demande de rentrer le prenom et le nom du joueur. 

```
  fmt.Println("Prenom du joueur :")
	fmt.Scan(&firstName)
	fmt.Println("Nom du joueur :")
	fmt.Scan(&player)
```

Puis elle fetch l'api afin de trouver l'id du joueur correspondant


# Choisir le match :

L'application vous propose les differents matchs de la nuit passée et vous demande de choisir l'id du match. 

```
  var gameId string
	fmt.Println("choisissez l'Id du match : ")
	fmt.Scan(&gameId)
```

# Connection a Twitch 

L'application se connect ensuite au compte Twitch du streamer, ecoute le Tchat puis parse les messages pour en ressortir le nom de ll'utilisateur ainsi que le contenue du message

```
    conn := connect.Connect()
    msgsChan := readmsg.ReadMessages(conn)
    parsedChan := parser.Parse(msgsChan)
```


Si un message correspond au nom du joueur choisi par le streamer, le bot envoi un message dans le tchat twitch avec le nom du viewer l'informant qu'il a gagné.

```
go func(){
        for msg := range parsedChan{
            numberOfColons := strings.Count(msg, ":")
            if numberOfColons >= 2 {
				s := strings.Split(msg, ":")
				user := strings.Split(s[1], ",")[0]
				messageData := s[2]
				
				if strings.Contains(messageData, playerLastName){
			        fmt.Fprintf(conn, "PRIVMSG #narvalo03 :@%s à gagner !\r\n", user)
				} else {
                    fmt.Println(msg)
                }			
				
				
			}
            
            
        }
    }()
```
