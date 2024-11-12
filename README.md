# Goal counter

A simple goal counter for teams. Made with Go, Gale and WebSockets.

## Deploy and run

Clone the repository:
```bash
git clone https://github.com/bottyaneu/goalcounter.git
```

Build the docker image:
```bash
docker build -t bottyaneu/goalcounter .
```

Run the app (the app will be available at `http://localhost:8000`):
```bash
docker run -p 8000:3000 bottyaneu/goalcounter
```
Note: Add `-d` to run the container in detached mode.

Run the `goalcounter` image:


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

### Connect to the WS Server

```bash
GET ws://localhost:8000/ws
```
(Replace `localhost` with the server's IP address and `8000` with the port)
With JS (client side):
```javascript
const ws = new WebSocket("ws://localhost:8000/ws")
```
**Note: Use something similar to vueUse `useWebsocket` to handle the connection.**

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
