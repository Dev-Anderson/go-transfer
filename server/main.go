package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func socket() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Erro ao criar servidor:", err.Error())
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Erro ao aceitar conexao:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Conexao estabelecida.")

	//abre o arquivo para escrito
	file, err := os.Create("arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	//copia o conteudo recebido para o arquivo
	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Erro ao gravar o arquivo:", err)
		return
	}

	//envia mensagem de confirmacao para o remetente
	mensagem := "Arquivo recebido com sucesso!"
	_, err = conn.Write([]byte(mensagem))
	if err != nil {
		fmt.Println("Erro ao enviar mensagem de confirmacao:", err)
		return
	}
	fmt.Println("Arquivo recebido e gravado com sucesso.")
}

func main() {
	for {
		time.Sleep(time.Duration(time.Second * 5))
		fmt.Println("Iniciando o servidor")
		socket()
	}
}
