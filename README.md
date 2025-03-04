# 📈 TickStats

One-stop & lightweight solution for metric collecting and analytics.

> [!TIP]
> This project is under construction and improvement. Do not use it in production environment and please wait for the first release! ❤️

## Features

- Lightweight - Docker image size < 20 MB
- One RESTful API collect all your projects' datas!
  
  ```json
  {
    "metrics_data": {
        "usage_cnt": 1,
        "os": "linux",
        "...", "..."
    }
  }
  ```

## Demo

![image](https://github.com/user-attachments/assets/2d8e6267-a3ad-40cb-9957-3310687a5f27)

## Development

1. Start docker-compose stack

```bash
cd docker
docker compose -f docker-compose.middleware.yaml -p tickstats up -d
```

2. Copy .env.example to .env

```bash
cp .env.example .env
```

3. Start the application

```bash
make dev
```
