# Gator CLI

A simple RSS feed aggreGATOR in Go

## Stack

- Go 1.24.3
- Postgres Latest

## Installation

`go install github.com/bulkashmak/gator-cli`

## Setup

1. Create a `.gatorconfig` config file in your `~` directory. Here's an example:
```json
{
  "db_url":"postgres://gator:gator@localhost:5432/gator?sslmode=disable",
  "current_user_name":""
}
```
2. Run Postgres locally. For docker check out `compose.yaml` file
3. Run DB migrations with `make migrate`

## Run

`gator-cli <command_name>`

### Commands

You can check out the available commands in `main.go` file

