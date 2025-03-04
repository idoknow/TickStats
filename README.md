# 📈 TickStats

> [!TIP]
> This project is under construction and improvement.

One-stop & lightweight solution for metric collecting and analytics.

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
