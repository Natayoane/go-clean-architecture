# go-clean-architecture

Desafio Full Cycle: Clean Architecture em Go

## Serviços Disponíveis

- **Web API REST**: CRUD de pedidos via HTTP  
  Porta padrão: **8080**
- **gRPC**: Serviço gRPC para pedidos  
  Porta padrão: **50051**
- **GraphQL**: Playground e API GraphQL para pedidos  
  Porta padrão: **8082**

## Como rodar o projeto

1. **Clone o repositório**
   ```sh
   git clone https://github.com/Natayoane/go-clean-architecture.git
   cd go-clean-architecture
   ```

2. **Configure o banco de dados e RabbitMQ**
   - O projeto espera um banco MySQL e RabbitMQ rodando (veja `docker-compose.yaml` para subir rapidamente).
   - Para subir com Docker:
     ```sh
     docker-compose up -d
     ```

3. **Configure as variáveis de ambiente**
   - Edite o arquivo `.env` ou configure as variáveis conforme o arquivo `configs/config.go`.

4. **Instale as dependências**
   ```sh
   go mod tidy
   ```

5. **Rode a aplicação**
   ```sh
   go run cmd/ordersystem/main.go
   ```

## Portas dos Serviços

| Serviço   | Porta | Endereço de acesso                        |
|-----------|-------|-------------------------------------------|
| Web REST  | 8080  | http://localhost:8080                     |
| gRPC      | 50051 | (use Evans)                               |
| GraphQL   | 8082  | http://localhost:8082/ (playground)       |

> As portas podem ser alteradas no arquivo de configuração.

## Exemplos de uso

### Web REST

- **Criar pedido**
  ```sh
  curl -X POST http://localhost:8080/order -d '{"id":"1","price":100,"tax":10}'
  ```
- **Listar pedidos**
  ```sh
  curl http://localhost:8080/order
  ```

### GraphQL

Acesse o playground em [http://localhost:8082/](http://localhost:8082/)

- **Query para listar pedidos**
  ```graphql
  query {
    listOrders {
      id
      Price
      Tax
      FinalPrice
    }
  }
  ```

- **Mutation para criar pedido**
  ```graphql
  mutation {
    createOrder(input: {id: "1", Price: 100, Tax: 10}) {
      id
      Price
      Tax
      FinalPrice
    }
  }
  ```

### gRPC

- Use o [Evans](https://github.com/ktr0731/evans) para consumir o serviço gRPC na porta **50051**.
  - Exemplo de uso:
    ```sh
    evans -r repl -p 50051
    ```
  - Dentro do Evans, você pode listar os serviços e métodos disponíveis e fazer chamadas interativamente.
