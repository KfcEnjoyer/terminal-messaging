# TermChat

A lightweight, real-time terminal-based messaging system built in Go using net/rpc. TermChat allows users to communicate in a distributed environment through a simple terminal interface.


## Features

- **Simple Authentication**: Quick username-based registration
- **Real-Time Messaging**: Instant message delivery to users
- **Multiple Communication Channels**:
  - Direct messaging between users
  - Global chat for all connected users
  - Private chat rooms for group conversations
- **User Management**:
  - See who's online
  - Create and manage chat rooms
  - Join and leave rooms as needed
- **Privacy Controls**:
  - Mute/unmute global chat
  - View muted users list

## Installation

### Prerequisites

- Go 1.18 or higher
- Terminal with ANSI color support (recommended)

### Building from source

```bash
# Clone the repository
git clone https://github.com/KfcEnjoyer/terminal-messaging.git
cd terminal-messaging

# Build the server
go build -o termchat-server ./cmd/server

# Build the client
go build -o termchat-client ./cmd/client
```

## Quick Start

### Starting the server

```bash
./termchat-server
```

This starts the server on the default address `localhost:8080`.


### Connecting as a client

```bash
./termchat-client
```

This connects to the server running on `localhost:8080`.


## Usage

### Basic Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `show` | List all online users | `show` |
| `send` | Send a direct message to a user | `send [username] [message]` |
| `global` | Send a message to all users | `global [message]` |
| `logout` | Disconnect from the server | `logout` |

### Room Management

| Command | Description | Usage |
|---------|-------------|-------|
| `create` | Create a new chat room | `create [room_name]` |
| `join` | Join an existing chat room | `join [room_name]` |
| `leave` | Leave a chat room | `leave [room_name]` |
| `sendroom` | Send a message to a chat room | `sendroom [room_name] [message]` |
| `showroomus` | Show users in a specific room | `showroomus [room_name]` |

### Privacy Controls

| Command | Description | Usage |
|---------|-------------|-------|
| `mute` | Mute global chat messages | `mute` |
| `unmute` | Unmute global chat messages | `unmute` |
| `showm` | Show list of muted users | `showm` |

## Architecture

TermChat uses a client-server architecture built on Go's net/rpc package:

- **Server**: Manages user connections, message routing, and room management
- **Client**: Handles user input, message display, and maintains connection to the server

### Components

- **ServerService**: Core server-side logic for managing users and message distribution
- **UserService**: User-facing application that connects to the server
- **Message Queue**: Server-side queue that stores messages for each user
- **Rooms**: Virtual spaces for group conversations
  

### Project Structure

```
termchat/
├── cmd/
│   ├── client/       # Client executable
│   └── server/       # Server executable
├── internal/
│   ├── messaging/    # Core messaging protocols
│   └── utils/        # Utility functions
└── README.md
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Go's net/rpc package for making distributed systems easier
- All contributors and users of the project

## Future Improvements

- End-to-end encryption for private messages
- Message persistence between sessions
- File sharing capabilities
- Command-line flag support for configuring server and client
- Multiple server instances with interconnection capabilities
