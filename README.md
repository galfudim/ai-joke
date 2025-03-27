# AI Joke Generator

A simple Go program that generates jokes using the Gemini AI API.

## Features

- Uses Google's Gemini AI API
- Terminal loading spinner animation
- Concurrent execution using goroutines

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

You can install in your environment:
```bash
go install ./...
```

And then run it like:
```bash
ai-joke
```