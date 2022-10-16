
# Sequence Validator

  
O sequence validator é uma API que valida uma tabela quadrada (**NXN**) de caracteres,  e retorna se a tabela existe alguma sequência igual ou maior que 4, podendo essa sequência ser horizontal, vertical ou diagonal. Ele também tem uma rota que informa a quantidade de tabelas que possuem uma sequência válida, inválida, e a proporção entre as tabelas.

  ## Iniciar a aplicação

Após ter feito o clone do projeto na sua máquina, é necessário que esteja instalado o **Docker** e o **Docker Compose**, após isso execute o comando na raiz do projeto:

```bash

$ docker compose up --scale app=10

```

Espere criar a rodar todos os container e já pode acessar a *api.

*Caso os container não tenham inicializado normalmente, pare toda a aplicação com o comando abaixo, e tente executar novamente a aplicação:*

```bash
$ docker compose stop
``` 


# API Rotas 

  

## Validar tabelas  

### Request  

`POST localhost:8890/sequence`

  `Request Body:` 

```json

{
"letters": []string,
}

```
  

### Response
  

HTTP/1.1 200 OK

 ```json 
{"is_value": boolean}
```


  
 

**Exemplo: **

    `Request Body:` 

```json
{
"letters": ["DUHBHB", "DUBUHD", "UBUUHU", "BHBDHH", " DDDDUB", "UDBDUH"]
}
```
`Response:`

 ```json
{ "is_valid": true}
```


## Pegar informações sobre tabelas

### Request  

`GET localhost:8890/stats`  

### Response
  

HTTP/1.1 200 OK

 ```json 
{"count_valid": int, "count_invalid": int, "ratio": float}
``` 
 

**Exemplo: **


`Response:`

 ```json
{"count_valid": 40, "count_invalid": 60, "ratio": 0.4}
```

