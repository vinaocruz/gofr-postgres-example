# Gofr com Postgres

Este projeto apresenta uma breve API para manejar biblioteca com framework Gofr, Postgres e testes com K6.io. O intuito é por em prática essas ferramentas.

A ideia de usar o [gofr.dev](https://gofr.dev) é conhecê-lo, considerando principalmente seus recursos integrados de log, métricas usando Prometheus e tracing, pilares importantes para sistemas robustos e corporativos.

## Setup
Execute através de containers, executando o docker-compose:

```
docker-compose up -d
```

Para facilitar as alterações do ambiente de desenvolvimento, este projeto utiliza a função de live reloading com [Air](https://github.com/air-verse/air).

## Migrate
No projeto foi utilizado o golang-migrate para versionamento da estrutura. Considere a [instalação do CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) para funcionamento. Você pode instalar com Go toolchain com comando abaixo:

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Valide que os binários em `$GOPATH/bin` estejam acessíveis.

### Criando estrutura do banco
Execute os atalhos definidos no Makefile para executar os migrations:
```
make migrate_up
```
Caso queira criar um novo migrate, execute:
```
make create_migration name=create_test_table
```

Para rollback de todas as migrations:
```
make migrate_down
```

## API

Acesse o endereço [http://localhost:8000/.well-known/swagger](http://localhost:8000/.well-known/swagger) para verificar a documentação dos endpoints disponíveis. Você pode executar a api através da própria página.

## Testes
Para fim didático, foi implementado teste de um caso de uso com mais cenários e sua cobertura. Para executá-lo:

```
cd internal/app/library/domain/usecase/
go test
```

## K6 - teste de carga

TODO