# Goal counter

A simple goal counter for teams. Made with Go, Gale and WebSockets.

## HTTP Requests

Get teams `GET /`
```json
{
  "teams": ["red", "blue"]
}
```

Increment team goal `/teams/{team}/goal`
```json
{
  "team": "red",
  "goals": 1
}
```

Scoreboard `GET /scoreboard`
```json
{
  "red": 5,
  "blue": 3
}
```

Reset scoreboard `GET /reset`
```json
{
  "reset": true
}
```

## WS Events

### Ping (keep alive)

Client should send a `ping` message every 3 seconds to keep the connection alive.
The server will response with a `pong` message.

### Goal

When a goal is scored, the server will send a message with the team that scored.

```json
{
  "event": "increment",
  "goals": 2,
  "team": "red"
}
```

### Reset (scoreboard)

When a reset request is made, the server will send a message to all clients to reset the scoreboard.

```json
{
  "event": "reset",
}
```

# Credits

Made with ❤️ by [Martin Binder](https://mrtn.vip)
