package models

import "go-loja-master/database"

type Produto struct {
	Id         int
	Produto    string
	Validade   string
	Preco      float64
	Quantidade int
}

func BuscarProdutos() []Produto {

	db := database.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var produto, validade string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &produto, &quantidade, &preco, &validade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Produto = produto
		p.Validade = validade
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CreateProdict(produto string, quantidade int, preco float64, validade string) {
	db := database.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos(produto, quantidade, preco, validade) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(produto, quantidade, preco, validade)
	defer db.Close()
}
func DeleteProduct(id string) {
	db := database.ConectaComBancoDeDados()

	delete, err := db.Prepare("DELETE FROM produtos WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Produto {
	db := database.ConectaComBancoDeDados()

	productDB, err := db.Query("SELECT * FROM produtos WHERE id=" + id)

	if err != nil {
		panic(err.Error())
	}

	productUpdate := Produto{}

	for productDB.Next() {
		var id, quantidade int
		var produto, validade string
		var preco float64

		err = productDB.Scan(&id, &produto, &quantidade, &preco, &validade)

		if err != nil {
			panic(err.Error())
		}

		productUpdate.Id = id
		productUpdate.Produto = produto
		productUpdate.Validade = validade
		productUpdate.Preco = preco
		productUpdate.Quantidade = quantidade
	}

	defer db.Close()

	return productUpdate
}

func UpdateProduct(id int, produto, validade string, preco float64, quantidade int) {
	db := database.ConectaComBancoDeDados()

	updateProduct, err := db.Prepare("UPDATE produtos SET produto=?,quantidade=?, preco=?, validade=? WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(produto, quantidade, preco, validade, id)

	defer db.Close()
}

func Loggin(email string, senha string) bool {
	db := database.ConectaComBancoDeDados()

	usuariodb, err := db.Query("SELECT * FROM usuario ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}

	for usuariodb.Next() {
		var id int
		var emailtemporario string
		var senhatemporaria string

		err = usuariodb.Scan(&id, &emailtemporario, &senhatemporaria)
		if err != nil {
			panic(err.Error())
		}
		if emailtemporario == email && senhatemporaria == senha {
			return true
		}
	}
	defer db.Close()
	return false
}
