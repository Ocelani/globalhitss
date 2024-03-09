# User API com RabbitMQ

Uma API RESTful CRUD em Golang, utilizando um banco de dados Postgres e a RabbitMQ para processamento assíncrono de inserções de registros relacionados a dados pessoais de clientes.

### Execução

Pré requisitos:
- Docker
- Golang

O sistema depende da execução de um banco de dados PostgreSQL e uma fila RabbitMQ.
Para isso, é possível executar essas dependências através do `docker compose`.

É disponibilizado um arquivo makefile com alguns scripts facilitadores de desenvolvimento.
Para levantar os containers, execute o comando:
```sh
make run
```

Como alternativa, é possível executar os comandos manualmente:
```sh
docker compose up -d
go mod tidy
go run ./cmd/userapi
```

#### Portas

São utilizadas as seguintes portas no sistema:
- API: `:3000`
- PostgreSQL `:5432`
- RabbitMQ `:5672`
- RabbitMQ monitor `:15672`

### API

#### Endpoints

São disponibilizado os seguintes endpoints:
- `POST /user/queue`: Criar usuário por meio da fila
- `POST /user`: Criar usuário
- `GET /user/:id`: Ler usuário
- `PUT /user/:id`: Atualizar usuário
- `DELETE /user/:id`: Deletar usuário

O envio dos dados no corpo da requisição segue o seguinte exemplo de modelo:
```json
{
  "nome":"Nome",
  "sobrenome":"Sobrenome",
  "contato":"11 12345-6789",
  "endereço":"rua Exemplo",
  "nascimento":"02/01/2006",
  "cpf":"012.345.678-99",
}
```

O formato padrão de resposta da API segue o modelo abaixo:
```json
{
  "user": {
    "id":1,
    "nome":"Nome",
    "sobrenome":"Sobrenome",
    "contato":"11 12345-6789",
    "endereço":"rua Exemplo",
    "nascimento":"02/01/2006",
    "cpf":"012.345.678-99",
  },
  "error":"caso ocorra algum erro"
}
```

### Testes

Para a execução de testes manuais, execute o script na pasta `./scripts`, ou o comando abaixo:
```sh
make tests-script
```

Para a execução de testes unitários, execute o comando abaixo:
```sh
make tests
```

Em relação ao volume de testes unitários, gostaria de ter agregado mais, porém, não foi possível devido ao tempo disponível.

### Problemas conhecidos

A execução da imagem do container correspondente à aplicação não ocorre devidamente.
Dessa forma, é indicado que a execução do programa ocorra manualmente, como indicado previamente:
```sh
docker compose up -d
go mod tidy
go run ./cmd/userapi
```

As credenciais de conexão estão escritas diretamente no código.
Compreendo que idealmente deveriam ser atribuídas através de variáveis de ambiente.

### Grato!
