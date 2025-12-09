# Kubernetes Deployment Manifests

This directory contains Kubernetes deployment configurations for each development preset.

## Structure

Each preset has its own directory with the following files:
- `kustomization.yaml` - Kustomize configuration
- `deployment.yaml` - Pod deployment specification
- `service.yaml` - Service configuration
- `configmap.yaml` - Environment variables and configuration

## Available Presets

### Python (`deploy/python/`)
- **Image**: `dev-env-python:latest`
- **Ports**:  - 8000 (HTTP app)
  - 5678 (Debug)
- **Resource Limits**: 512Mi memory, 500m CPU

### Go (`deploy/go/`)
- **Image**: `dev-env-go:latest`
- **Ports**:
  - 8080 (HTTP app)
  - 2345 (Delve debugger)
- **Resource Limits**: 512Mi memory, 500m CPU

### Rust (`deploy/rust/`)
- **Image**: `dev-env-rust:latest`
- **Ports**:
  - 8080 (HTTP app)
  - 5000 (Debug)
- **Resource Limits**: 512Mi memory, 500m CPU

### Web (`deploy/web/`)
- **Image**: `dev-env-web:latest`
- **Ports**:
  - 3000 (Next.js/React)
  - 5173 (Vite)
  - 9229 (Node debugger)
- **Resource Limits**: 512Mi memory, 500m CPU

## Usage

### Apply with kubectl

```bash
# Apply Python preset
kubectl apply -k deploy/python/

# Apply Go preset
kubectl apply -k deploy/go/

# Apply Rust preset
kubectl apply -k deploy/rust/

# Apply Web preset
kubectl apply -k deploy/web/
```

### Render manifests without applying

Use the `dev k8s render` command:

```bash
# Render Python manifests
dev k8s render python

# Render all presets
for preset in python go rust web; do
  dev k8s render $preset
done
```

### Build with kustomize

```bash
# Build Python manifests
kubectl kustomize deploy/python/

# Build and save to file
kubectl kustomize deploy/python/ > python-manifests.yaml
```

## Customization

### Change namespace

Edit `kustomization.yaml`:
```yaml
namespace: my-namespace
```

### Change replicas

Edit `deployment.yaml`:
```yaml
spec:
  replicas: 3
```

### Add environment variables

Edit `configmap.yaml`:
```yaml
data:
  DATABASE_URL: "postgresql://..."
  API_KEY: "..."
```

### Adjust resource limits

Edit `deployment.yaml`:
```yaml
resources:
  requests:
    memory: "512Mi"
    cpu: "500m"
  limits:
    memory: "1Gi"
    cpu: "1000m"
```

## Health Checks

All deployments include:
- **Liveness Probe**: Checks `/healthz` endpoint
- **Readiness Probe**: Checks `/readyz` endpoint

Make sure your application implements these endpoints.

## Notes

- Images use `imagePullPolicy: IfNotPresent` - build locally first with `dev build`
- Services are `ClusterIP` type - use Ingress or port-forward for external access
- ConfigMaps are non-sensitive - use Secrets for sensitive data
- Resource limits are minimal defaults - adjust based on your needs
