# temporal-polyglot-poc

# Installation

```sh
brew install temporal
cd golang
go mod download
cd ../python
python3.11 -m venv venv
. venv/bin/activate
pip install -r requirements.txt
```

# Running workflow
```sh
temporal server start-dev &
open http://localhost:8233
cd python
. venv/bin/activate
python run_worker.py
cd ../golang
go run worker/main.go
go run start/main.go
```
