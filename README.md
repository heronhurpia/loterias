# Loterias

## Ações
- [ ] Criar e testar migrações
- [ ] Fazer teste de conexão com dB
- [ ] Criar pacote de acesso ao banco de dados
- [ ] Exibir erro se conexão não for do tipo pgsql
- [ ] Criar arquivo .env e .env-backup
- [ ] Criar teste unitário para executar migrações
  
## Megasena

## Quina

## Loto fácil

## Notas
- Como determinar endereço ip do container:
  - docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' container_name_or_id
- Conectar com banco de dados:
  - psql -U postgres -h 172.20.0.2

## migrações
- Criar ou alterar uma tabela:
migrate create -ext sql -dir database/migrations -seq create_users_table
- Para rodar as migrações:
migrate -path database/migrations -database "postgresql://postgres:bolao_password@172.20.0.2:5432/postgres?sslmode=disable" -verbose up

export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'
