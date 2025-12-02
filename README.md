# MatterTUI

MatterTUI is a Slack-style team chat platform built using Go for both the backend server and terminal user interface (TUI) client.

## Overview

This project is intented as my take on building a Slack-style team chat platform using Go for both the backend server and terminal user interface (TUI) client. MatterMost serves as the basis for development, but everything in this project is built from scratch.

I think this would be a good starting point for learning Go and building a modern chat system from the ground up. It can also be used to learn about system design, architecture, implementation, and even TUI applications.

> ⚠️ This project is in no capacity to be production-ready yet. It is still in the early stages of development.

## Key Features

- **User Authentication**: Secure login and registration flows.
- **Teams & Channels**: Organize conversations into teams and topic-based channels.
- **Real-time Messaging**: Instant communication updates.
- **Terminal User Interface (TUI)**: A rich, interactive terminal client.

## Technology Stack

- **Language**: Go (Golang)
- **Backend**: Go standard library and ecosystem.
- **Frontend**: Go-based TUI (Terminal User Interface) using Bubbletea and Bubble.
- **Database**: PostgreSQL (running in a local Docker container).
- **Authentication**: JWT-based authentication.
- **WebSocket**: Real-time messaging using WebSocket.
