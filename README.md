# jwtd

[![CI](https://github.com/dandehoon/jwtd/workflows/CI/badge.svg)](https://github.com/dandehoon/jwtd/actions)

A simple command-line JSON Web Tokens decoder tool.

## Functionality

- Accepts JWT tokens via command-line arguments or standard input
- Cleans tokens (removes Bearer prefix, whitespace, newlines)
- Formats timestamps (`exp`, `nbf`, `iat`) as human-readable dates

## Installation

### Option 1: Quick Install

Download and install the latest release directly from GitHub:

```bash
curl -sSL https://raw.githubusercontent.com/dandehoon/jwtd/main/install.sh | sudo bash
```

### Option 2: Install with Go

If you have Go installed:

```bash
go install github.com/dandehoon/jwtd@latest
```

### Option 3: Download Binary Manually

1. Go to the [releases page](https://github.com/dandehoon/jwtd/releases)
2. Download the appropriate binary for your platform
3. Make it executable and move it to a directory in your PATH

## How to Use

2.  **Use the tool:**
    You can provide the JWT token in two ways:

    ### Method 1: Command-line argument (Direct)

    Pass the JWT token directly as a command-line argument:

    ```bash
    jwtd "your.jwt.token"
    ```

    The tool accepts tokens in various formats:

    ```bash
    # Standard JWT token
    jwtd "your.jwt.token"

    # With Bearer prefix
    jwtd "Bearer your.jwt.token"

    # With any custom prefix, extra spaces, or newlines
    jwtd "  JWT   your.jwt.token  "
    ```

    ### Method 2: Pipe input (Standard input)

    Pipe the JWT into the program via standard input:

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

    **Examples:**

    Using command-line argument:

    ```bash
    jwtd "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    ```

    Using pipe input:

    ```bash
    echo "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" | jwtd
    ```

    Both methods will output:

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

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
