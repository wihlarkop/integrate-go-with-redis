# Research Integration Golang with Redis

## Table Of Contents
- [Golang Template Project](#golang-template-project)
    - [About the project](#about-the-project)
    - [Getting started](#getting-started)
        - [Project Structure](#project-structure)
    - [Requirements](#requirements)
    - [How To Run](#how-to-run)
    - [Source](#source)


## About the project

The template is used to create golang project. All golang projects must follow the conventions in the
template. Calling for exceptions must be brought up in the engineering team.

## Getting started

Below we describe the conventions or tools specific to golang project.

### Project Structure

```tree
├── .gitignore
├── README.md
├── gateway
│   ├── gateway.go
├── infrastructure
│   ├── redis.go
├── main.go
├── docker-compose.yml
├── go.mod
├── .env-example


```

A brief description of the layout:

* `README.md` is a detailed description of the project.

## Requirements

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [go-redis/redis](https://redis.uptrace.dev/)

## How to Run

you need download package/library first

```shell
go mod download
```

now you can run 
```shell
go run main.go
```