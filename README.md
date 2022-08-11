# simple-chat-api

## Deprecating this repo

I've realized that the idea of solving this problem has become far more interesting to
me than solving it _in Go_. Despite how nice Fiber is and how simple Go is to work with,
I know that I'll be able to develop faster using simpler technologies. I may undeprecate
this repo and come back to make an implementation in Go once I've completed the initial
implementation.

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

## Technical & security model

- Uses each users password hash as their JWT signature key. This is slow for initial
lookups but the user enitities can be cached to increase performance. The benefit of 
this model is that all user's sessions can be invalidated by changing their password
(must notify cache), while not having to maintain a session ID table and lookups.
- And the next day I've realized that this^ is a terrible idea in terms of actual
client security. First off, this would make it so that anyone with database access
would be able to generate a valid token. Second, this would make it so that the real
password is actually the hash which is pretty stupid. That means if the DB leaks, bad
actors could generate tokens for users at will without even having their passwords.