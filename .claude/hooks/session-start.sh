#!/bin/bash
# Hook: SessionStart
# Fires once when a fresh Claude Code session begins (not on resume).
#
# Installs Go and golangci-lint if not present.
# In Claude Code Web, each session starts from an ephemeral container,
# so tools must be installed here.

set +e  # Never exit on error in session-start

GO_VERSION="1.26.0"
GOLANGCI_LINT_VERSION="v2.11.3"
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
    if ! grep -qF 'export PATH="/usr/local/go/bin:/root/go/bin:$PATH"' ~/.bashrc 2>/dev/null; then
      echo 'export PATH="/usr/local/go/bin:/root/go/bin:$PATH"' >> ~/.bashrc
    fi
  fi

  # Ensure GOPATH/bin is on PATH
  export GOPATH="${GOPATH:-/root/go}"
  export PATH="$GOPATH/bin:$PATH"

  # Install golangci-lint if not present (download binary + verify checksum)
  if ! command -v golangci-lint &>/dev/null; then
    echo "[session-start] Installing golangci-lint ${GOLANGCI_LINT_VERSION}..." >&2
    _lint_ver="${GOLANGCI_LINT_VERSION#v}"
    _lint_tar="golangci-lint-${_lint_ver}-linux-${GOARCH}.tar.gz"
    _lint_url="https://github.com/golangci/golangci-lint/releases/download/${GOLANGCI_LINT_VERSION}"
    _tmp="$(mktemp -d)"
    curl -fsSL "${_lint_url}/${_lint_tar}" -o "${_tmp}/${_lint_tar}"
    curl -fsSL "${_lint_url}/golangci-lint-${_lint_ver}-checksums.txt" -o "${_tmp}/checksums.txt"
    (cd "${_tmp}" && grep "${_lint_tar}" checksums.txt | sha256sum -c --status)
    if [ $? -eq 0 ]; then
      tar -C "${_tmp}" -xzf "${_tmp}/${_lint_tar}"
      mv "${_tmp}/golangci-lint-${_lint_ver}-linux-${GOARCH}/golangci-lint" "$GOPATH/bin/golangci-lint"
    else
      echo "[session-start] golangci-lint checksum verification failed — skipping install" >&2
    fi
    rm -rf "${_tmp}"
    unset _lint_ver _lint_tar _lint_url _tmp
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
