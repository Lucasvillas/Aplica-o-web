package controllers

import (
	"go-loja-master/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var checar string

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscarProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "criar-produto", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := r.FormValue("produto")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")
		validade := r.FormValue("validade")

		precoConv, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro ao converter o preço", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro ao converter a quantidade", err)
		}

		models.CreateProdict(produto, quantidadeConv, precoConv, validade)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeleteProduct(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditProduct(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		produto := r.FormValue("produto")
		validade := r.FormValue("validade")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro ao converter o Id para int: ", err)
		}

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter o preco em float: ", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter a quantidade em int: ", err)
		}

		models.UpdateProduct(idConv, produto, validade, precoConv, quantidadeConv)

	}

	http.Redirect(w, r, "/", 301)
}

func Loggar(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "conect", nil)
}

func Rodar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("Email")
		senha := r.FormValue("Senha")

		if models.Loggin(email, senha) {
			checar = email
			http.Redirect(w, r, "/", 301)
		} else {
			http.Redirect(w, r, "loggin", 301)
		}
	}
}
