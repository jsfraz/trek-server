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
| TRACKER_TOKEN_SECRET  | ...long text...     |
| SUPERUSER_USERNAME    | root                |
| SUPERUSER_PASSWORD    | 12345678            |

## Deploying using Nginx reverse proxy

The server runs on `localhost:8080` by default.

First thing you need to do is copy the Nginx config.

```bash
sudo cp trek.conf /etc/nginx/conf.d/
sudo systemctl reload nginx
```

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d trek-backend.josefraz.cz
```

Make sure to replace the values with your own.
