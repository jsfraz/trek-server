# trek-server

Trek backend for tracking your motorcycles or whatever.

## Environment variables

| ENV                   | Example             |
|-----------------------|---------------------|
| GIN_MODE              | debug/release       |
| POSTGRES_USER         | trek                |
| POSTGRES_PASSWORD     | 12345678            |
| POSTGRES_SERVER       | localhost           |
| POSTGRES_PORT         | 5432                |
| POSTGRES_DB           | trek                |
| ACCESS_TOKEN_SECRET   | ...long text...     |
| ACCESS_TOKEN_LIFESPAN | 111600              |
| DEVICE_TOKEN_SECRET   | ...long text...     |
| SUPERUSER_USERNAME    | root                |
| SUPERUSER_PASSWORD    | 12345678            |

<!--
## Testing PostgreSQL container

`sudo docker run -d --name postgres --restart always -e POSTGRES_USER=trek -e POSTGRES_PASSWORD=12345678 -v /home/razj/postgres:/var/lib/postgresql/data -p 5432:5432 postgres:alpine`
--->
