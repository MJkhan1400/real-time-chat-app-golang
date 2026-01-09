# Real-Time Chat App with Golang and SvelteKit

A real-time chat application built with Golang backend and SvelteKit frontend using WebSockets for communication. This project is adapted from the original tutorial that used Angular, now modernized with SvelteKit for learning purposes.

## Overview

This application demonstrates real-time communication between multiple clients through a Golang WebSocket server. The server manages client connections and broadcasts messages to all connected clients, while the SvelteKit frontend provides a modern, responsive chat interface.

## Features

- Real-time messaging using WebSockets
- Multi-client chat support
- Connection status notifications
- Clean, modern UI with SvelteKit
- Unique client identification
- System messages for connection events

## Architecture

### Backend (Golang)
- WebSocket server using Gorilla WebSocket
- Client connection management
- Message broadcasting system
- JSON-based message format

### Frontend (SvelteKit)
- Modern reactive UI
- WebSocket client integration
- Real-time message display
- Responsive design

## Prerequisites

- [Golang](https://golang.org/) 1.19+
- [Node.js](https://nodejs.org/) 18+
- [npm](https://www.npmjs.com/) or [pnpm](https://pnpm.io/)

## Project Structure

```
real-time-chat-app-golang/
├── server/                 # Golang WebSocket server
│   └── main.go
├── src/                    # SvelteKit frontend
│   ├── lib/
│   │   └── websocket.js   # WebSocket service
│   ├── routes/
│   │   └── +page.svelte    # Main chat page
│   └── app.css             # Global styles
├── package.json
└── README.md
```

## Installation & Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd real-time-chat-app-golang
```

### 2. Backend Setup

```bash
# Install Go dependencies
go get github.com/gorilla/websocket
go get github.com/google/uuid

# Navigate to server directory
cd server

# Run the WebSocket server
go run main.go
```

The server will start on `http://localhost:12345`

### 3. Frontend Setup

```bash
# Install frontend dependencies
npm install

# Start the development server
npm run dev
```

The frontend will be available at `http://localhost:5173`

## Usage

1. Start the Golang server first (port 12345)
2. Start the SvelteKit development server (port 5173)
3. Open the application in multiple browser tabs to simulate multiple users
4. Send messages and see them appear in real-time across all connected clients

## API Documentation

### WebSocket Connection

**Endpoint:** `ws://localhost:12345/ws`

### Message Format

Messages are exchanged in JSON format with the following structure:

```json
{
  "sender": "client-uuid",
  "content": "message content"
}
```

### System Messages

System messages start with `/` and indicate:
- `/A new socket has connected.` - New client joined
- `/A socket has disconnected.` - Client left

## Technical Implementation

### Golang Server Components

- **ClientManager**: Manages all connected clients
- **Client**: Represents individual WebSocket connections
- **Message**: JSON message structure
- **Goroutines**: Handle concurrent client communication

### SvelteKit Features

- **Reactive variables** for real-time updates
- **WebSocket API** integration
- **Component-based architecture**
- **Modern CSS** for styling

## Configuration

### Server Configuration
- **Port**: 12345 (configurable in `main.go`)
- **CORS**: Enabled for all origins (development setting)

### Client Configuration
- **WebSocket URL**: `ws://localhost:12345/ws`
- **Reconnection**: Automatic on connection loss

## Development

### Adding New Features

1. **Backend**: Extend `Message` struct and update business logic
2. **Frontend**: Modify Svelte components and WebSocket service
3. **Styling**: Update global CSS in `src/app.css`

### Testing

- Test WebSocket connections using browser developer tools
- Monitor server logs for connection events
- Verify real-time updates across multiple clients

## Learning Objectives

This project helps in understanding:
- WebSocket communication patterns
- Concurrent programming with Go goroutines
- Real-time web application architecture
- Modern frontend development with SvelteKit
- Client-server message handling
- JSON data exchange

## Original Tutorial

This project is based on the tutorial by Nic Raboy: [Create A Real Time Chat App With Golang, Angular, And Websockets](https://www.thepolyglotdeveloper.com/2016/12/create-real-time-chat-app-golang-angular-2-websockets/)

## License

This project is for educational purposes only.

## Contributing

Feel free to submit issues and enhancement requests as this is a learning project.