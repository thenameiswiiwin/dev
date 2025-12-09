# Version Pinning and Security

This document outlines version pinning strategies and security best practices for the dev environment.

## Version Pinning Strategy

### Docker Base Images

All Dockerfiles use specific version tags instead of `latest`:

```dockerfile
# Good - pinned version
FROM python:3.12-slim

# Bad - unpredictable
FROM python:latest
```

### Binary Downloads

When downloading binaries (e.g., Neovim), use:
1. Specific version/tag
2. SHA256 checksum verification
3. HTTPS URLs

Example:
```dockerfile
ARG NVIM_VERSION=0.10.0
ARG NVIM_SHA256=abc123...

RUN curl -LO https://github.com/neovim/neovim/releases/download/v${NVIM_VERSION}/nvim-linux64.tar.gz \
    && echo "${NVIM_SHA256}  nvim-linux64.tar.gz" | sha256sum -c - \
    && tar -xzf nvim-linux64.tar.gz \
    && mv nvim-linux64 /usr/local/
```

### Language Toolchains

#### Python
```yaml
# manifests/presets/python.yaml
packages:
  - name: python3
    brew: python@3.12  # Pin minor version
```

#### Go
```yaml
# manifests/presets/go.yaml
packages:
  - name: go
    brew: go@1.24  # Pin minor version
```

#### Rust
Rust uses rustup which manages versions:
```dockerfile
RUN rustup default 1.82.0
```

#### Node.js
```yaml
# manifests/presets/web.yaml
packages:
  - name: node
    brew: node@22  # Pin major version
```

## Checksum Verification

### For Downloaded Scripts

Never pipe curl directly to bash without verification:

```bash
# Bad - No verification
curl -fsSL https://example.com/install.sh | bash

# Good - Verify checksum
curl -fsSL https://example.com/install.sh -o install.sh
echo "abc123...  install.sh" | sha256sum -c -
bash install.sh
```

### For Homebrew/APT

Package managers handle checksums automatically:
- Homebrew: Verifies SHA256 of downloaded bottles
- APT: Uses package signatures and checksums
- Pacman: Uses package signatures

## Security Best Practices

### 1. Minimal Base Images

Use `-slim` or `-alpine` variants:
```dockerfile
FROM python:3.12-slim  # Much smaller than python:3.12
```

### 2. Run as Non-Root (Production)

Development containers often run as root for convenience, but production should use non-root:

```dockerfile
RUN useradd -m -u 1000 devuser
USER devuser
```

### 3. Multi-Stage Builds

Use multi-stage builds to reduce final image size:

```dockerfile
FROM golang:1.24 AS builder
WORKDIR /build
COPY . .
RUN go build -o app .

FROM debian:bookworm-slim
COPY --from=builder /build/app /usr/local/bin/
CMD ["app"]
```

### 4. Scan Images

Use tools like:
- `docker scan` (Snyk)
- Trivy
- Grype

```bash
# Scan with Trivy
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
  aquasec/trivy image dev-env-python:latest
```

### 5. Keep Dependencies Updated

Regularly update:
- Base images
- System packages
- Language packages

```bash
# Rebuild with latest patches
dev build --no-cache
```

## Version Matrix

| Component | Current Version | Pin Strategy | Update Frequency |
|-----------|----------------|--------------|------------------|
| Python | 3.12.x | Minor | Quarterly |
| Go | 1.24.x | Minor | Per release |
| Rust | 1.82.x | Latest stable | Monthly |
| Node.js | 22.x | Major | LTS cycle |
| Neovim | Latest | Track releases | Monthly |

## Updating Versions

### 1. Update Dockerfiles

```bash
# Edit containers/<preset>/Dockerfile
# Update FROM image versions
# Update ARG version variables
```

### 2. Update Manifests

```bash
# Edit manifests/presets/<preset>.yaml
# Update package versions
```

### 3. Test Locally

```bash
dev build <preset> --test
```

### 4. Update CI

Ensure CI builds and tests with new versions.

### 5. Document Changes

Update CHANGELOG.md with version changes.

## Automation

Consider automating version updates with:
- Dependabot (for Dockerfiles and manifests)
- Renovate Bot
- Custom scripts to check for updates

Example Dependabot config (`.github/dependabot.yml`):

```yaml
version: 2
updates:
  - package-ecosystem: "docker"
    directory: "/containers"
    schedule:
      interval: "weekly"
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
```

## References

- [Docker Security Best Practices](https://docs.docker.com/develop/security-best-practices/)
- [NIST Container Security Guide](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-190.pdf)
- [CIS Docker Benchmark](https://www.cisecurity.org/benchmark/docker)
