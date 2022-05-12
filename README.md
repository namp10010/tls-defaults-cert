# tls-defaults-cert

### Run server

```bash
cd server
go run main.go
```

### Run client 

```bash
cd client
go run main.go
```

When not use the env

```bash
go run main.go
main.go:19: x509: certificate signed by unknown authority
```

When use the env will load the server cert using the env for location

```bash
go run main.go
main.go:23: tls conn opened
here is what we got
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
Connection: close

400 Bad Requ
```