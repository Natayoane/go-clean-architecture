# go-clean-architecture

Desafio Full Cycle: Clean Architecture em Go

## Serviços Disponíveis

- **Web API REST**: CRUD de pedidos via HTTP  
  Porta padrão: **8000**
- **gRPC**: Serviço gRPC para pedidos  
  Porta padrão: **50051**
- **GraphQL**: Playground e API GraphQL para pedidos  
  Porta padrão: **8080**

## Como rodar o projeto

1. **Clone o repositório**
   ```sh
   git clone https://github.com/Natayoane/go-clean-architecture.git
   cd go-clean-architecture
   ```

2. **Suba todos os serviços com Docker Compose**
   ```sh
   docker-compose up --build
   ```
   
   Ou para rodar em background:
   ```sh
   docker-compose up --build -d
   ```

3. **Verifique se os serviços estão rodando**
   ```sh
   docker-compose ps
   ```

**Nota**: O Docker Compose já configura automaticamente:
- Banco de dados MySQL na porta 3306
- RabbitMQ na porta 5672 (management em 15672)
- Aplicação Go com todos os serviços (REST, gRPC, GraphQL)
- Todas as dependências e configurações necessárias

## Portas dos Serviços

| Serviço   | Porta | Endereço de acesso                        |
|-----------|-------|-------------------------------------------|
| Web REST  | 8000  | http://localhost:8000                     |
| gRPC      | 50051 | (use Evans)                               |
| GraphQL   | 8080  | http://localhost:8080/ (playground)       |
| MySQL     | 3306  | localhost:3306                            |
| RabbitMQ  | 5672  | localhost:5672                            |
| RabbitMQ Management | 15672 | http://localhost:15672 (guest/guest) |

> As portas podem ser alteradas no arquivo `docker-compose.yaml`.

## Exemplos de uso

### Web REST

- **Criar pedido**
  ```sh
  curl -X POST http://localhost:8000/order \
    -H "Content-Type: application/json" \
    -d '{"id":"1","price":100,"tax":10}'
  ```
- **Listar pedidos**
  ```sh
  curl http://localhost:8000/order
  ```

### GraphQL

Acesse o playground em [http://localhost:8080/](http://localhost:8080/)

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

## Comandos Úteis do Docker

- **Ver logs da aplicação**
  ```sh
  docker-compose logs app
  ```

- **Ver logs de todos os serviços**
  ```sh
  docker-compose logs
  ```

- **Parar todos os serviços**
  ```sh
  docker-compose down
  ```

- **Recriar e iniciar os serviços**
  ```sh
  docker-compose up --build --force-recreate
  ```

- **Ver status dos containers**
  ```sh
  docker-compose ps
  ```
