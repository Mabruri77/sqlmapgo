# sqlmapgo

sqlmapgo is a custom SQL injection and database takeover tool written in Go.

## Features

- Lightweight and fast SQL injection detection and exploitation.
- Supports scanning and exploitation of SQL injection vulnerabilities.
- Customizable headers and request payloads.
- Works with both GET and POST requests.

## Installation

Before using sqlmapgo, ensure you have Go installed on your system. You can download and install Go from [golang.org](https://golang.org/dl/).

## Usage

### Basic Usage

To run sqlmapgo, navigate to the directory where the `main.go` file is located and use the `go run` command with the appropriate flags.

```bash
go run . -u <url>

```

Replace <url> with the target URL you want to scan for SQL injection vulnerabilities.
Using a Request File

sqlmapgo also supports using a request file (request.txt) containing the HTTP request you want to analyze for SQL injection.

bash

go run . -r request.txt

Flags

    -u <url>: Specify the target URL to scan.
    -r <file>: Specify the path to a file containing the HTTP request to analyze.

Example

bash

go run . -u "https://example.com/page?id=1"

or

bash

go run . -r ./requests/sample-request.txt

Contributing

Contributions are welcome! Feel free to fork the repository and submit pull requests.
