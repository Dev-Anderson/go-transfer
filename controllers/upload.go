package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func PostUpload(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Erro ao ler corpo da solicitacao: %v", err)
	}

	fileName := fmt.Sprintf("arquivo_%d.txt", time.Now().UnixNano())
	f, err := os.Create(fileName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao criar o arquivo: %v", err)
	}

	if _, err := f.WriteString(string(body)); err != nil {
		c.String(http.StatusInternalServerError, "Erro ao gravar arquivo: %v", err)
	}

	c.String(http.StatusOK, "Arquivo recebido e salvo com sucesso!")
}
