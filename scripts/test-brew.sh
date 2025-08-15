#!/bin/bash

# Script to test Homebrew formula with local binaries
# This ensures we test the current code instead of remote releases

set -e

echo "🧹 Cleaning up any hanging processes..."
# Kill any existing HTTP servers from previous runs
pkill -f "python3 -m http.server" 2>/dev/null || true
pkill -f "go run server.go" 2>/dev/null || true

echo "🔨 Building local binary for current platform..."

# Build only for current platform
BINARY_NAME="jwtd-$(go env GOOS)-$(go env GOARCH)"
go build -o "$BINARY_NAME" .

echo "🌐 Starting local HTTP server..."

# Find an available port
PORT=8000
while lsof -Pi :$PORT -sTCP:LISTEN -t >/dev/null ; do
    PORT=$((PORT + 1))
done

echo "Using port $PORT"

# Start Python HTTP server in background (suppress warnings)
python3 -m http.server $PORT 2>/dev/null &
SERVER_PID=$!
sleep 2

# Function to cleanup on exit
cleanup() {
    echo "🧹 Cleaning up..."
    kill $SERVER_PID 2>/dev/null || true
    rm Formula/jwtd-test.rb 2>/dev/null || true
    rm "$BINARY_NAME" 2>/dev/null || true
}
trap cleanup EXIT

echo "📝 Creating test formula..."

# Create test formula pointing to local server
cp Formula/jwtd.rb Formula/jwtd-test.rb

# Replace URLs with local server
sed -i.bak "s|https://github.com/dandehoon/jwtd/releases/latest/download/|http://localhost:$PORT/|g" Formula/jwtd-test.rb

# Calculate checksum of the binary and replace the sha256 line
CHECKSUM=$(shasum -a 256 "$BINARY_NAME" | cut -d' ' -f1)
sed -i.bak "s|sha256 :no_check|sha256 \"$CHECKSUM\"|" Formula/jwtd-test.rb

# Fix class name and formula name to match filename exactly
sed -i.bak 's|class Jwtd < Formula|class JwtdTest < Formula|' Formula/jwtd-test.rb

rm Formula/jwtd-test.rb.bak

echo "🍺 Testing Homebrew formula..."

echo "Validating formula syntax..."
ruby -c Formula/jwtd-test.rb

echo "Testing download URL accessibility..."
if curl -I -s "http://localhost:$PORT/$BINARY_NAME" | head -n1 | grep -q "200 OK"; then
    echo "✅ Download URL is accessible"
else
    echo "❌ Download URL failed"
    exit 1
fi

echo "Testing binary functionality directly..."
jwtd_path="./$BINARY_NAME"
chmod +x "$jwtd_path"
echo "Testing version command..."
"$jwtd_path" --version

echo "Testing JWT decoding..."
output=$("$jwtd_path" <<< "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
if echo "$output" | grep -q "John Doe"; then
    echo "✅ JWT decoding works correctly"
    echo "$output"
else
    echo "❌ JWT decoding failed"
    echo "$output"
    exit 1
fi

echo "✅ All tests passed!"
