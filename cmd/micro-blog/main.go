package main

func main() {
	// 建立一個新Http服務
	server := InitalizeServer()
	server.Run()
}
