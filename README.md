# no-as-a-discord-bot

*no-as-a-discord-bot* is a simple Discord bot that returns various negative responses, based on [No-as-a-service API](https://github.com/hotheadhacker/no-as-a-service). 

## Requirements
- Go 1.23.4

## Usage and installation

> [! WARNING]
> This part is under improvements.

Compile it in your machine with the command below.

```
go build -o main
```

Run with the command below. 
NOTE: Make sure to run with `reasons.json` exists in current working directory.
```
export DISCORD_TOKEN=<your-discord-bot-token-here>
./main
```

