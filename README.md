# Multithreading and API Challenge

## Descrição

Neste desafio, você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:
- [Brasil API](https://brasilapi.com.br/api/cep/v1/01153000 + cep)
- [Via CEP](http://viacep.com.br/ws/ + cep + /json/)

## Requisitos

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Executando o Programa

1. Certifique-se de ter o Go instalado em sua máquina. Caso não tenha, você pode instalá-lo [aqui](https://golang.org/dl/).

2. Clone este repositório e navegue até a pasta onde o arquivo `main.go` está localizado.

3. No terminal, execute o seguinte comando para rodar o programa:
    ```sh
    go run main.go
    ```

4. O programa solicitará que você digite um CEP:
    ```sh
    Digite seu CEP:
    ```

5. Digite o CEP desejado e pressione Enter.

## Exemplo de Saída

Após a execução bem-sucedida, a saída no terminal será semelhante a:

```sh

Desafio Multithreading
Digite seu CEP: 01153000
CEP encontrado!
Resultado mais rápido recebido da API: BrasilAPI
{"cep":"01153000","state":"SP","city":"São Paulo","neighborhood":"Barra Funda","street":"Rua Vitorino Carmilo","service":"open-cep"}


