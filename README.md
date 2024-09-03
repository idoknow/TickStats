# TickStats

One stop solution for lightweight telemetry and analytics.

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
