package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Erro ao conectar com o servidor:", err)
		return
	}
	defer conn.Close()

	file, err := os.Open("arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Erro ao enviar arquivo:", err)
		return
	}

}
