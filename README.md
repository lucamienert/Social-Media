# Go Next Social Media

## Build

### Go API

```bash
cd services
go run rsa.go
```

Paste the corresponding keys into the

```bash
api/api.env
```

file:

```env
PORT=
ClientOrigin=

DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=

ACCESS_TOKEN_PRIVATE_KEY=
ACCESS_TOKEN_EXPIRED_IN=
ACCESS_TOKEN_MAXAGE=

REFRESH_TOKEN_PRIVATE_KEY=
REFRESH_TOKEN_EXPIRED_IN=
REFRESH_TOKEN_MAXAGE=
```

### Next Frontend

```env
NEXT_PUBLIC_API_URL=
```

### Start Docker

```bash
docker compuse up --build
```
