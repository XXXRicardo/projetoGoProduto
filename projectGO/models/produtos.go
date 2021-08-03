package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Mercado struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	Nome       string  `json: "nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

type Mercado_id struct {
	Nome       string  `json: "nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

func BuscaTodosOsProdutosJason() []Mercado {
	resp, err := http.Get("http://localhost:5000/api/v1/mercado/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(body)
	}
	mercado1 := []Mercado{}
	content := []Mercado{}
	s := Mercado{}

	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range append(content) {
		s.ID = post.ID
		s.Nome = post.Nome
		s.Descricao = post.Descricao
		s.Quantidade = post.Quantidade
		s.Preco = post.Preco

		mercado1 = append(mercado1, s)
	}
	defer resp.Body.Close()

	return mercado1
}

func EditaProdutoJson(id string) []Mercado {

	fmt.Println(id)

	resp, err := http.Get("http://localhost:5000/api/v1/mercado/" + id)

	if err != nil {
		log.Fatalln(err)
	}
	var teste Mercado

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	s := Mercado{}
	mercado1 := []Mercado{}

	err = json.Unmarshal(bodyBytes, &teste)
	if err != nil {
		log.Fatal(err)
	}

	s.ID = teste.ID
	s.Nome = teste.Nome
	s.Descricao = teste.Descricao
	s.Quantidade = teste.Quantidade
	s.Preco = teste.Preco

	mercado1 = append(mercado1, s)

	return mercado1
}

func CriarNovoProdutoJson(nome, descricao string, precoConvertido float64, quantidade int) {
	todo := Mercado_id{nome, descricao, precoConvertido, quantidade}
	jsonReq, err := json.Marshal(todo)
	resp, err := http.Post("http://localhost:5000/api/v1/mercado/", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoStruct Mercado
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)
}

func DeletaProdutoJson(id string) {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/api/v1/mercado/"+id, nil)

	fmt.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

}

func AtualizaProdutoJson(nome string, descricao string, id int, preco float64, quantidade int) {
	todo := Mercado{id, nome, descricao, preco, quantidade}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/mercado/", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &todo)

}
