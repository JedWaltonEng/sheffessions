## 

- Run GH actions locally with https://github.com/nektos/act
    - `$ gh act -v`

- Start dev environment.
    - `$ cd docker && docker compose up`

- Migrations
    - `$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
    - `$ export POSTGRESQL_URL=postgres://postgres:mysecretpassword@localhost:5432/devdb?sslmode=disable`
    - https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
    - `$ sudo apt-get update & sudo apt-get install postgresql-client`
    - `$ migrate -database ${POSTGRESQL_URL} -path db/migrations up`
    - `$ psql -h localhost -U postgres -d devdb -p 5432 -c "\d confessions"`

Todo:
-----
- Build clone of Production, for staging.
- Implement logic to handle dev, staging, prod urls in the FE (likely a util function)
- Configure database/migrations/automated backups/test backups work/hellosql in go/sql tests.
- Develop basic schema to store confessions into database.
- Develop algorithm to randomly choose posts.
- Post health check with OpenAI.
- Generate Image from chosen posts.
- Schedule image posts & Upload image to instagram via graphql
