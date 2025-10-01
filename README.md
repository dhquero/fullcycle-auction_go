## Lab - Concorrência com Golang - Leilão

**Objetivo:** Adicionar uma nova funcionalidade ao projeto já existente para o leilão fechar automaticamente a partir de um tempo definido.

Clone o seguinte repositório: [clique para acessar o repositório](https://github.com/devfullcycle/labs-auction-goexpert).

Toda rotina de criação do leilão e lances já está desenvolvida, entretanto, o projeto clonado necessita de melhoria: adicionar a rotina de fechamento automático a partir de um tempo.

Para essa tarefa, você utilizará o go routines e deverá se concentrar no processo de criação de leilão (auction). A validação do leilão (auction) estar fechado ou aberto na rotina de novos lançes (bid) já está implementado.

**Você deverá desenvolver:**
- Uma função que irá calcular o tempo do leilão, baseado em parâmetros previamente definidos em variáveis de ambiente;
- Uma nova go routine que validará a existência de um leilão (auction) vencido (que o tempo já se esgotou) e que deverá realizar o update, fechando o leilão (auction);
- Um teste para validar se o fechamento está acontecendo de forma automatizada;

**Dicas:**
- Concentre-se na no arquivo internal/infra/database/auction/create_auction.go, você deverá implementar a solução nesse arquivo;
- Lembre-se que estamos trabalhando com concorrência, implemente uma solução que solucione isso:
- Verifique como o cálculo de intervalo para checar se o leilão (auction) ainda é válido está sendo realizado na rotina de criação de bid;
- Para mais informações de como funciona uma goroutine, clique aqui e acesse nosso módulo de Multithreading no curso Go Expert;
 
**Entrega:**
- O código-fonte completo da implementação.
- Documentação explicando como rodar o projeto em ambiente dev.
- Utilize docker/docker-compose para podermos realizar os testes de sua aplicação.

## Execução no ambiente de desenvolvimento

### Execução do teste unitários
Para execução do teste, execute o comando:

```sh
$ go test -timeout 30s -run ^TestCreateAuction_GoroutineUpdatesStatus$ fullcycle-auction_go/internal/infra/database/auction
```

### Compilação e execução
Para execução em ambiente local, execute o comando:

```sh
$ docker-compose up -d
```

### Validação do fechamento do leilão

O arquivo `test/auction.http` realiza a chamada da API que cria e consulta os leilões.

A utilização do arquivo é realizada com a extensão do VS Code:

```
Name: REST Client
Id: humao.rest-client
Description: REST Client for Visual Studio Code
Version: 0.25.1
Publisher: Huachao Mao
VS Marketplace Link: https://marketplace.visualstudio.com/items/?itemName=humao.rest-client
```
