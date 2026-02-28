#!/bin/bash
# Hook: SessionStart
# Fires once when a fresh Claude Code session begins (not on resume).
#
# Installs Go and golangci-lint if not present.
# In Claude Code Web, each session starts from an ephemeral container,
# so tools must be installed here.

set +e  # Never exit on error in session-start

GO_VERSION="1.26.0"
GOLANGCI_LINT_VERSION="v2.10.1"
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64)  GOARCH="amd64" ;;
  aarch64) GOARCH="arm64" ;;
  arm64)   GOARCH="arm64" ;;
  *)       GOARCH="amd64" ;;
esac

if [ "${CLAUDE_CODE_REMOTE:-}" = "true" ]; then
  echo "[session-start] Running in Claude Code Web (ephemeral container)" >&2

  # Install Go if not present
  if ! command -v go &>/dev/null; then
    echo "[session-start] Installing Go ${GO_VERSION}..." >&2
    curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${GOARCH}.tar.gz" \
      | tar -C /usr/local -xz
    export PATH="/usr/local/go/bin:$PATH"
    echo 'export PATH="/usr/local/go/bin:/root/go/bin:$PATH"' >> ~/.bashrc
  fi

  # Ensure GOPATH/bin is on PATH
  export GOPATH="${GOPATH:-/root/go}"
  export PATH="$GOPATH/bin:$PATH"

  # Install golangci-lint if not present
  if ! command -v golangci-lint &>/dev/null; then
    echo "[session-start] Installing golangci-lint ${GOLANGCI_LINT_VERSION}..." >&2
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh \
      | sh -s -- -b "$GOPATH/bin" "$GOLANGCI_LINT_VERSION" >/dev/null 2>&1
  fi

  # Download module dependencies
  if [ -f "${CLAUDE_PROJECT_DIR}/go.mod" ]; then
    echo "[session-start] Downloading Go modules..." >&2
    cd "${CLAUDE_PROJECT_DIR}" && go mod download 2>/dev/null || true
  fi
else
  echo "[session-start] Running in local devcontainer — tools pre-installed" >&2
fi

echo "[session-start] Done" >&2
exit 0
