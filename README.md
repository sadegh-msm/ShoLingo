# ShoLingo: A URL Shortener in Go

ShoLingo is a simple yet effective URL shortener service written in Go. This project showcases the power of Go in building high-performance web services with a minimal footprint.

## Table of Contents

1. [Features](#features)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Contributing](#contributing)
6. [License](#license)

## Features

- Generate short URLs for long URLs.
- High-performance handling of requests.
- Use of Go's strong concurrency model with Goroutines.
- Redis as the primary store for high-speed URL redirection.

## Prerequisites

- [Go](https://golang.org/dl/) (Version 1.16 or higher)
- [Redis](https://redis.io/) (Version 5.0 or higher)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/sadegh-msm/ShoLingo.git
    ```

2. Navigate into the cloned repository:

    ```bash
    cd ShoLingo
    ```

3. Install the required Go dependencies:

    ```bash
    go mod download
    ```

4. Compile and run the server:

    ```bash
    go run main.go
    ```

## Usage

After starting the server, you can use the following endpoints:

    ```go
    e.GET("/:url", routes.ResolveURL)
    e.POST("/api/v1", routes.ShortenURL)
    ```

You can also change the configuration in `.env` file.
