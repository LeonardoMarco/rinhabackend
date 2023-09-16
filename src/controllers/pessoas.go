package controllers

import (
	"net/http"
	"rinhabackendleo/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePessoa(c *gin.Context) {
	var RequestData models.Pessoas

	if err := c.ShouldBindJSON(&RequestData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Validation error"})
		return
	}

	if RequestData.Apelido == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Validation error"})
		return
	}

	pessoa, err := models.CreatePessoa(RequestData)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	t, _ := time.Parse("2006-01-02T00:00:00Z", pessoa.Nascimento)

	response := models.Pessoas{
		Id:         pessoa.Id,
		Apelido:    pessoa.Apelido,
		Nome:       pessoa.Nome,
		Nascimento: t.Format("2006-01-02"),
		Stack:      pessoa.Stack,
	}

	c.JSON(http.StatusCreated, response)
}

func GetPessoas(c *gin.Context) {
	term, _ := c.GetQuery("t")

	if term == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	pessoas, err := models.GetPessoas(term)

	if err != nil {
		c.Status(http.StatusBadGateway)
		return
	}

	var fmtPessoa []models.Pessoas

	for i := 0; i < len(pessoas); i++ {
		t, _ := time.Parse("2006-01-02T00:00:00Z", pessoas[i].Nascimento)

		fmtPessoa = append(fmtPessoa, models.Pessoas{
			Id:         pessoas[i].Id,
			Apelido:    pessoas[i].Apelido,
			Nome:       pessoas[i].Nome,
			Nascimento: t.Format("2006-01-02"),
			Stack:      pessoas[i].Stack,
		})
	}

	if fmtPessoa == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, fmtPessoa)
}

func GetPessoaById(c *gin.Context) {
	id := c.Params.ByName("id")

	pessoa, err := models.GetPessoasById(id)

	if err != nil {
		c.Status(http.StatusBadGateway)
	}

	var fmtPessoa []models.Pessoas

	t, _ := time.Parse("2006-01-02T00:00:00Z", pessoa.Nascimento)

	fmtPessoa = append(fmtPessoa, models.Pessoas{
		Id:         pessoa.Id,
		Apelido:    pessoa.Apelido,
		Nome:       pessoa.Nome,
		Nascimento: t.Format("2006-01-02"),
		Stack:      pessoa.Stack,
	})

	if fmtPessoa[0].Id == "" {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, fmtPessoa)
}

func GetPessoasCount(c *gin.Context) {
	count, err := models.GetPessoasCount()

	if err != nil {
		c.Status(http.StatusBadGateway)
	}

	c.JSON(http.StatusOK, count)
}
