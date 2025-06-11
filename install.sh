#!/bin/bash

set -e

# Configuration
REPO="danztran/jwtd"
BINARY_NAME="jwtd"
INSTALL_DIR="${HOME}/.local/bin"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Helper functions
log_info() {
  echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
  echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
  echo -e "${RED}[ERROR]${NC} $1"
}

# Detect OS and architecture
detect_platform() {
  local os=$(uname -s | tr '[:upper:]' '[:lower:]')
  local arch=$(uname -m)

  case $os in
    linux*)
      OS="linux"
      ;;
    darwin*)
      OS="darwin"
      ;;
    msys*|mingw*|cygwin*)
      OS="windows"
      ;;
    *)
      log_error "Unsupported operating system: $os"
      exit 1
      ;;
  esac

  case $arch in
    x86_64|amd64)
      ARCH="amd64"
      ;;
    arm64|aarch64)
      ARCH="arm64"
      ;;
    *)
      log_error "Unsupported architecture: $arch"
      exit 1
      ;;
  esac

  # Construct binary name
  if [ "$OS" = "windows" ]; then
    BINARY_FILE="${BINARY_NAME}-${OS}-${ARCH}.exe"
  else
    BINARY_FILE="${BINARY_NAME}-${OS}-${ARCH}"
  fi

  log_info "Detected platform: ${OS}-${ARCH}"
}

# Get latest release version
get_latest_version() {
  log_info "Fetching latest release information..."

  if command -v curl >/dev/null 2>&1; then
    LATEST_VERSION=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
  elif command -v wget >/dev/null 2>&1; then
    LATEST_VERSION=$(wget -qO- "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
  else
    log_error "Neither curl nor wget is available. Please install one of them."
    exit 1
  fi

  if [ -z "$LATEST_VERSION" ]; then
    log_error "Failed to fetch latest version"
    exit 1
  fi

  log_info "Latest version: $LATEST_VERSION"
}

# Download binary
download_binary() {
  local download_url="https://github.com/${REPO}/releases/download/${LATEST_VERSION}/${BINARY_FILE}"
  local temp_dir=$(mktemp -d)
  local temp_file="${temp_dir}/${BINARY_FILE}"

  log_info "Downloading ${BINARY_FILE}..."
  log_info "URL: $download_url"

  if command -v curl >/dev/null 2>&1; then
    curl -sL "$download_url" -o "$temp_file"
  elif command -v wget >/dev/null 2>&1; then
    wget -q "$download_url" -O "$temp_file"
  fi

  if [ ! -f "$temp_file" ]; then
    log_error "Failed to download binary"
    exit 1
  fi

  echo "$temp_file"
}

# Install binary
install_binary() {
  local temp_file="$1"

  # Create install directory if it doesn't exist
  mkdir -p "$INSTALL_DIR"

  # Copy binary to install directory
  local install_path="${INSTALL_DIR}/${BINARY_NAME}"
  cp "$temp_file" "$install_path"
  chmod +x "$install_path"

  log_info "Installed ${BINARY_NAME} to ${install_path}"

  # Check if install directory is in PATH
  if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    log_warn "Warning: ${INSTALL_DIR} is not in your PATH"
    log_warn "Add the following line to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    log_warn "export PATH=\"\$PATH:${INSTALL_DIR}\""
    log_warn ""
    log_warn "Or run the following command to add it temporarily:"
    log_warn "export PATH=\"\$PATH:${INSTALL_DIR}\""
  fi

  # Clean up
  rm -rf "$(dirname "$temp_file")"
}

# Verify installation
verify_installation() {
  if [ -x "${INSTALL_DIR}/${BINARY_NAME}" ]; then
    log_info "Installation successful!"
    log_info "Version: $(${INSTALL_DIR}/${BINARY_NAME} --version 2>/dev/null || echo 'Unable to get version')"
    log_info ""
    log_info "You can now use '${BINARY_NAME}' command"
    log_info "Example: echo 'your.jwt.token' | ${BINARY_NAME}"
  else
    log_error "Installation verification failed"
    exit 1
  fi
}

# Main installation process
main() {
  log_info "Installing ${BINARY_NAME}..."

  detect_platform
  get_latest_version

  local temp_file=$(download_binary)
  install_binary "$temp_file"
  verify_installation

  log_info "Installation complete!"
}

# Handle command line arguments
case "${1:-}" in
  --help|-h)
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Install ${BINARY_NAME} from GitHub releases"
    echo ""
    echo "OPTIONS:"
    echo "  --help, -h  Show this help message"
    echo "  --version   Show version and exit"
    echo ""
    echo "Environment variables:"
    echo "  INSTALL_DIR   Installation directory (default: \$HOME/.local/bin)"
    echo ""
    exit 0
    ;;
  --version)
    echo "jwtd installer v1.0.0"
    exit 0
    ;;
  "")
    main
    ;;
  *)
    log_error "Unknown option: $1"
    echo "Use --help for usage information"
    exit 1
    ;;
esac
