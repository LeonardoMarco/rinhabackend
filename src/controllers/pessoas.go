package controllers

import (
	"net/http"
	"rinhabackendleo/src/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func CreatePessoa(c *gin.Context) {
	var RequestData models.Pessoas
	db := c.MustGet("DB_CONTEXT").(*gorm.DB)
	cache := c.MustGet("CACHE_CONTEXT").(*redis.Client)

	if err := c.ShouldBindJSON(&RequestData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Validation error"})
		return
	}

	if RequestData.Apelido == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Validation error"})
		return
	}

	pessoa, err := models.CreatePessoa(RequestData, db, cache)

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

	db := c.MustGet("DB_CONTEXT").(*gorm.DB)

	pessoas, err := models.GetPessoas(term, db)

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

	db := c.MustGet("DB_CONTEXT").(*gorm.DB)
	cache := c.MustGet("CACHE_CONTEXT").(*redis.Client)

	pessoa, err := models.GetPessoasById(id, db, cache)

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
	db := c.MustGet("DB_CONTEXT").(*gorm.DB)
	count, err := models.GetPessoasCount(db)

	if err != nil {
		c.Status(http.StatusBadGateway)
	}

	c.JSON(http.StatusOK, count)
}
