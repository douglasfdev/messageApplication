# Messageria com Kafka e Golang

### É necessário a criação de uma Conexão ao banco de dados e de um tópico no Kafka

Para isso basta rodar o comando `docker-compose up -d` para iniciar a instalação dos containers e suas dependencias para rodar a aplicação

#### Para criar o banco de dados

Basta rodar o comando `docker exec -it mysql bash`

- - Logo após rodar o comando `mysql -uroot -p golang` para logar no mysql
- - E criar a tabela produtos:

```sql
  CREATE TABLE produtos (id varchar(255), name varchar(255), price float);
```

#### Para criar um tópico no kafka

Entre no container com comando shell: `docker-compose exec kafka bash`
E coloque o seguinte comando dentro do container: `kafka-topics --bootstrap-server=localhost:9092 --topic=products --create`

#### Para rodar a aplicação

Basta entrar no container da aplicação com o comando: `docker-compose exec goapp` e depois o comando `go run cmd/app/main.go`

##### Endpoints

```http
POST http://localhost:8000/products
Content-Type: application/json
```
| Parâmetro | Tipo     | Descrição                         |
| :-------- | :------- | :-------------------------------- |
| `name`    | `string` | **Obrigatório**. Nome do produto  |
| `price`   | `float`  | **Obrigatório**. Preço do produto |
Resposta do corpo: {Status 201 Created}
```json
{
  "name": "XPTO",
  "price": 100
}
```

```http
GET http://localhost:8000/products
Content-Type: application/json
```
Resposta é uma lista de produtos {Status 200 OK}
```json
[
  {
    "name": "XPTO",
    "price": 100
  }
]
```

----
> **Caso queira mandar uma mensagem via terminal do kafka para conferir se está sendo enviado as mensagens siga as etapas a baixo:**

##### Para envio de mensagens via Kafka

Entre no container com comando shell: `docker-compose exec kafka bash`
E coloque o seguinte comando dentro do container: `kafka-console-producer --bootstrap-server=localhost:9092 --topic=products`
Depois basta colar um json válido como esse de exemplo, em uma única linha:
```json
{"name": "XPTO","price": 100}
```