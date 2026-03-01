FROM ubuntu:24.04

# Avoid interactive prompts during package installation
ENV DEBIAN_FRONTEND=noninteractive

ARG GO_VERSION=1.26.0
ARG GOLANGCI_LINT_VERSION=v2.10.1
ARG TARGETARCH=amd64

RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    curl \
    jq \
    ca-certificates \
    make \
    && rm -rf /var/lib/apt/lists/*

# Install Go
RUN GOARCH="${TARGETARCH:-amd64}" && \
    curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${GOARCH}.tar.gz" \
    | tar -C /usr/local -xz

ENV PATH="/usr/local/go/bin:/root/go/bin:$PATH"
ENV GOPATH="/root/go"

# Install golangci-lint
RUN curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANGCI_LINT_VERSION}/install.sh" \
    | sh -s -- -b /usr/local/bin "${GOLANGCI_LINT_VERSION}"

# Install yq (YAML processor)
ARG YQ_VERSION=v4.52.4
RUN case "${TARGETARCH:-amd64}" in \
        amd64) YQ_ARCH="amd64" ;; \
        arm64) YQ_ARCH="arm64" ;; \
        arm) YQ_ARCH="arm" ;; \
        386) YQ_ARCH="386" ;; \
        *) echo "Unsupported TARGETARCH: ${TARGETARCH}" >&2; exit 1 ;; \
    esac && \
    curl -fsSL "https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_${YQ_ARCH}" \
        -o /usr/local/bin/yq && \
    chmod +x /usr/local/bin/yq

WORKDIR /workspace
