package main

type Config struct{
	listenAddr string
}

type Message struct{
	cmd Command
	peer *Peer
}

func main() {
	
}