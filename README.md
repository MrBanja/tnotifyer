# Telegram Notifier

A simple command-line tool to send messages to Telegram channels or chats using the Telegram Bot API.

## Features

- Send messages to Telegram channels or chats
- Configurable request timeout
- Simple command-line interface

## Prerequisites

- Go 1.20 or higher
- A Telegram Bot Token (obtain from [@BotFather](https://t.me/botfather))
- Chat ID of the target channel or chat

## Usage

```bash
./tnotify -token "YOUR_BOT_TOKEN" -chat-id "YOUR_CHAT_ID" -msg "Your message"
```

### Command-line Arguments

| Flag     | Description                       | Required | Default |
|----------|-----------------------------------|----------|---------|
| -token   | Telegram bot token from BotFather | Yes      | ""      |
| -chat-id | Target chat or channel ID         | Yes      | ""      |
| -msg     | Message to send                   | Yes      | ""      |
| -timeout | Request timeout duration          | No       | 5s      |

### Example

```bash
./tnotify \
  -token "1234567890:ABCdefGHIjklMNOpqrsTUVwxyz" \
  -chat-id "-1001234567890" \
  -msg "Hello from Telegram Notifier!" \
  -timeout 10s
```

## Installation

### Using `go install`

```bash
go install github.com/MrBanja/tnotifyer@latest
```

## Building from Source

### Linux

#### x86_64

```bash
GOOS=linux GOARCH=amd64 go build -o tnotify
```

#### ARM64 (e.g., Raspberry Pi)

```bash
GOOS=linux GOARCH=arm64 go build -o tnotify
```

### macOS

#### Intel

```bash
GOOS=darwin GOARCH=amd64 go build -o tnotify
```

#### Apple Silicon (M1/M2)

```bash
GOOS=darwin GOARCH=arm64 go build -o tnotify
```

### Windows

#### 64-bit

```bash
set GOOS=windows
set GOARCH=amd64
go build -o tnotify.exe
```

Or using PowerShell:

```powershell
$env:GOOS = "windows"; $env:GOARCH = "amd64"; go build -o tnotify.exe
```

#### 32-bit

```bash
set GOOS=windows
set GOARCH=386
go build -o tnotify.exe
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.