# temporal-polyglot-poc

# Installation

```sh
brew install temporal
cd golang
go mod download
```

# Running workflow
```sh
temporal server start-dev &
open http://localhost:8233
cd golang
go run start/main.go
go run worker/main.go
```
