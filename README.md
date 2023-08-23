## 

- Run GH actions locally with https://github.com/nektos/act
    `$ gh act -v`

- Start dev environment.
    `$ cd docker && docker compose up`


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
