<h1 align="center">
  <p align="center">Konstraint</p>
</h1>

<p align="center">
  <a href="https://github.com/raspbernetes/konstraint/actions" alt="Build"><img src="https://github.com/xUnholy/konstraint/workflows/build/badge.svg" /></a>
  <a href="https://golangci.com" alt="GolangCI"><img src="https://golangci.com/badges/github.com/xUnholy/konstraint.svg" /></a>
  <a href="https://codeclimate.com/github/xUnholy/konstraint/maintainability" alt="Maintainability"><img src="https://api.codeclimate.com/v1/badges/5f5fcb94169e53faf77d/maintainability" /></a>
  <a href="https://github.com/xUnholy/konstraint/blob/master/LICENSE" alt="License"><img src="https://img.shields.io/badge/license-Apache_2.0-blue.svg" /></a>
</p>

## Intro

Konstraint is a tool to auto generate Kubernetes Gatekeeper `constrainttemplate` resources by parsing the Rego files directly.

## Installation

Install `konstraint` using your preferred method.

```bash
go get github.com/xUnholy/konstraint
```

## Usage

Use the CLI to generate Gatekeeper `constrainttemplate` resources:

```bash
konstraint template -p example/ -l example/libs -o yaml
```

Alternatively run the CLI using a Docker image for CI/CD:

```bash
docker run --rm --workdir /tmp/workspace \
  -v "$(pwd)":/tmp/workspace \
  xunholy/konstraint:latest \
  template -p example -l example/libs \
  -o yaml
```
