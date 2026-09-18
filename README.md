# DevOps CI/CD Pipeline Project

![CI/CD](https://github.com/YOUR_USERNAME/devops-cicd-project/actions/workflows/ci-cd.yml/badge.svg)

A Go microservice with a complete, automated CI/CD pipeline: every push to `main` triggers automated testing, building, containerization, and (optionally) deployment to Kubernetes.

**Tech stack:** Golang · Docker · Jenkins · GitHub Actions · Kubernetes · CI/CD

---

## What this project demonstrates

- Writing and testing a small Go web service
- Multi-stage Dockerfile for a minimal, production-style container image
- CI/CD pipeline defined as code — both a `Jenkinsfile` (for on-prem/self-hosted CI) and a GitHub Actions workflow (runs automatically on GitHub, no server needed)
- Kubernetes manifests for deployment with readiness/liveness probes and a Service
- Infrastructure and pipeline fully version-controlled (everything-as-code)

---

## Architecture

```
git push
   │
   ▼
GitHub Actions / Jenkins
   │
   ├─ 1. go test         (run unit tests)
   ├─ 2. go build         (compile binary)
   ├─ 3. docker build     (containerize app)
   ├─ 4. docker push      (push image to registry)
   └─ 5. kubectl apply    (deploy to Kubernetes)
   │
   ▼
Running app on Kubernetes (2 replicas, health-checked)
```

---

## Project structure

```
.
├── main.go                     # Go web server
├── main_test.go                 # Unit tests
├── go.mod
├── Dockerfile                    # Multi-stage build
├── Jenkinsfile                   # Pipeline-as-code (Jenkins)
├── .github/workflows/ci-cd.yml   # Pipeline-as-code (GitHub Actions)
└── k8s/
    ├── deployment.yaml            # K8s Deployment (2 replicas + health probes)
    └── service.yaml                # K8s Service (NodePort)
```

---

## Run it locally

### 1. Run the Go app directly
```bash
go run main.go
# visit http://localhost:8080
```

### 2. Run with Docker
```bash
docker build -t go-app:v1 .
docker run -p 8080:8080 go-app:v1
```

### 3. Deploy to Kubernetes (Minikube)
```bash
minikube start
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
minikube service go-app-service
```

---

## CI/CD setup

### Option A: GitHub Actions (zero setup, runs automatically)
Already configured in `.github/workflows/ci-cd.yml`. As soon as you push this repo to GitHub, it runs tests and builds the Docker image automatically — check the **Actions** tab.

To also push the image to Docker Hub automatically, add these repo secrets under **Settings → Secrets and variables → Actions**:
- `DOCKERHUB_USERNAME`
- `DOCKERHUB_TOKEN` (generate from Docker Hub → Account Settings → Security)

### Option B: Jenkins (self-hosted, more "traditional DevOps")
```bash
docker run -d --name jenkins -p 8081:8080 -p 50000:50000 \
  -v jenkins_home:/var/jenkins_home \
  -v /var/run/docker.sock:/var/run/docker.sock \
  jenkins/jenkins:lts
```
Then create a Pipeline job pointing to this repo — Jenkins will pick up the `Jenkinsfile` automatically.

---

