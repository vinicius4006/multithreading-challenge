package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("#### Desafio Multithreading ####")
	time.Sleep(time.Second)
	fmt.Print("Digite seu CEP: ")
	var cep string
	fmt.Scanln(&cep)
	ch1 := make(chan string)
	ch2 := make(chan string)
	go requestBrasilAPI(cep, ch1)
	go requestViaCEP(cep, ch2)
	select {
	case cep1 := <-ch1:
		println("Resultado mais rápido recebido da API: BrasilAPI\n" + cep1)
	case cep2 := <-ch2:
		println("Resultado mais rápido recebido da API: ViaCEP\n" + cep2)
	case <-time.After(time.Second):
		println("timeout")
	}
}

func requestBrasilAPI(cep string, ch chan string) {
	ht := http.Client{}

	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/cep/v1/"+cep, nil)
	if err != nil {
		log.Panicln(err.Error())
	}

	res, err := ht.Do(req)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		println("CEP encontrado!")
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err.Error())
	}
	ch <- string(bytes)
}

func requestViaCEP(cep string, ch chan string) {
	ht := http.Client{}

	req, err := http.NewRequest("GET", "http://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		log.Panicln(err.Error())
	}

	res, err := ht.Do(req)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		println("CEP encontrado!")
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err.Error())
	}
	ch <- string(bytes)
}
