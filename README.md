# simple-chat-api

## Tooling overview

- **Language**: [Go](https://go.dev/)
- **Framework**: [Fiber](https://gofiber.io/)
- **Deployment**: [Digital Ocean App Platform](https://www.digitalocean.com/go/app-platform)
- **Database**: [PostgreSQL in Digital Ocean](https://try.digitalocean.com/managed-databases/)
- **Environment**: [direnv](https://direnv.net/)

## Running locally

### Adding environment variables

Copy the contents of `.envrc.example` into a new file `.envrc` and then fulfill those variables. 
If you want automatic loading, be sure to install [direnv](https://direnv.net/).

### via `go run`

```bash
go run main.go
```

### via executable

```bash
go build
simple-chat-api
```
