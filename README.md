# AI Joke Generator

A simple Go program that generates jokes using the Gemini AI API. The program includes a nice terminal loading spinner while waiting for the API response.

## Features

- Uses Google's Gemini AI API
- Terminal loading spinner animation
- Concurrent execution using goroutines
- Graceful error handling
- Clean and modern Go code structure

## Prerequisites

- Go 1.16 or later
- Gemini API key

## Setup

1. Clone the repository
2. Set your Gemini API key as an environment variable:
   ```bash
   export GEMINI_API_KEY='your-api-key-here'
   ```

## Usage

Run the program:
```bash
go run main.go
```

## License

MIT 