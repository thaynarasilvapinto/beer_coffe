## Rodando o projeto 

* Primeiro tenha o docker instalado na sua maquina

* Para criar um conteiner utilize o comando:

         docker run --name beer-coffee -p 5432:5432 -e POSTGRES_PASSWORD=12345678 -e POSTGRES_USER=beer-coffee -e POSTGRES_DB=beer-coffee -d postgres


## Comandos úteis do golang 

dentro da pasta /src

`go build`

`go run main.go`