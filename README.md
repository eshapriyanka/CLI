# Go Network Status Checker (CLI)

A lightweight Command Line Interface (CLI) tool built with **Go** to check the connectivity status of websites. This tool is containerized using **Docker** (Multi-stage build) to ensure a minimal footprint and consistent runtime environment.

## Features

- **Written in Go:** High-performance compiled binary.
- **Dockerized:** Uses Multi-stage builds for an ultra-lightweight image (Alpine Linux).
- **Unix Philosophy:** Handles standard flags (`--url`) and exit codes (`0` for success, `1` for error).

## Tech Stack

- **Language:** Go (Golang) 1.23
- **Container:** Docker (Alpine Linux)
- **Architecture:** CLI Application

## How to Run (Using Docker)

1. **Build the Image**
   ```bash
   docker build -t netcheck .
   ```
2. **Run the Tool**
   ```bash
   docker run netcheck --url "[https://google.com](https://google.com)"
   ```
