# Go Redis-like Server

This is a simple Redis-like server written in Go that supports basic RESP protocol commands like `SET` and `GET`.

## How to Run

1. **Install Go** (if not installed): https://go.dev/doc/install  
2. **Install Make** (Linux: `sudo apt install make`, macOS: `brew install make`)  
3. **Clone this repo** and run:
```bash
make run
```

Server runs on localhost:5001.

## Example (using redis-cli)

```bash
redis-cli -p 5001
SET hello world
GET hello
```
