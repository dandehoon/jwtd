# jwtd

A simple command-line JSON Web Tokens decoder tool.

## Installation

### Option 1: Quick Install (Recommended)

Download and install the latest release directly from GitHub:

```bash
curl -sSL https://raw.githubusercontent.com/danztran/jwtd/main/install.sh | bash
```

### Option 2: Install with Go

If you have Go installed:

```bash
go install github.com/danztran/jwtd@latest
```

### Option 3: Download Binary Manually

1. Go to the [releases page](https://github.com/danztran/jwtd/releases)
2. Download the appropriate binary for your platform
3. Make it executable and move it to a directory in your PATH

## How to Use

2.  **Use the tool:**
    Pipe the JWT into the program via standard input.

    ```bash
    echo "your.jwt.token" | jwtd
    ```

    The tool accepts tokens in various formats:

    ```bash
    # Standard JWT token
    echo "your.jwt.token" | jwtd

    # With Bearer prefix
    echo "Bearer your.jwt.token" | jwtd

    # With any custom prefix, extra spaces, or newlines
    echo "  JWT   your.jwt.token  " | jwtd
    ```

    **Tip for macOS users:** You can use `pbpaste` to pipe the content of your clipboard directly:

    ```bash
    pbpaste | jwtd
    ```

    **Example:**

    Running:

    ```bash
    echo "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" | jwtd
    ```

    Will output:

    ```
    alg: HS256
    typ: JWT
    ---
    iat: 1516239022 (2018-01-18T01:30:22Z)
    name: John Doe
    sub: 1234567890
    ```

## Building from Source

If you prefer to build from source:

```bash
go build -o jwtd main.go
```

Then you can run it using `./jwtd`.

## Functionality

- Reads a JWT from standard input.
- Automatically cleans up tokens with:
  - "Bearer" or any other prefix (e.g., "JWT ", "Token ")
  - Leading or trailing whitespace
  - Additional newlines
- Decodes the Base64Url-encoded header and payload of the JWT.
- Prints the key-value pairs from the header and payload.
- Sorts the keys alphabetically for consistent output.
- Formats Unix timestamp values for `exp` (expiration time), `nbf` (not before), and `iat` (issued at) claims into a human-readable RFC3339 timestamp.
- Separates the decoded header and payload sections with "---".
- The signature part of the JWT is not verified or decoded, only the header and payload are processed.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
