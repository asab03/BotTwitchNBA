# Presentation 

L'application est developpée en Golang

L'application a pour objectif de faire deviner à un chat twitch un joueur NBA en fonction de ses lignes de stats de son match de la nuit passée.
Pour cela, le programme récupère les informations grâce à un api et il affiche les stats du joueur demandé par le streamer.

Par la suite le programme se connecte au chat du stream et écoute les messages des viewers. Dès qu'un utilisateur donne la bonne réponse, le programme envoie un message dans le chat pour désigner le vainqueur.

# Avant de lancer l'application

```
git clone Github.com/asab03/BotTwitchNBA
```
se créer un compte sur le site https://rapidapi.com/hub pour pouvoir acceder à l'api

créer un fichier .env à la racine du projet et y stocker la clé API fournies apres l'inscription. 

Dans se fichier .env on retrouve egalement la clé pour l'API twitch, le channel et le nom du bot

# Lancer l'application :

```go run .```


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
