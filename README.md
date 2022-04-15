# Fake Server
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/budnieswski/fake-server?color=green&style=flat-square)](https://github.com/budnieswski/fake-server/releases/latest)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/budnieswski/fake-server?style=flat-square)](https://github.com/budnieswski/fake-server/tags)
![License](https://img.shields.io/dub/l/vibe-d.svg?style=flat-square)

The purpose of FakeServer is to respond to a request regardless of the route or method (POST, GET...) called, being mainly used as a server for tests during the development of an application.

## Contents
- [Features](#features)
- [Example](#example)
- [Install](#install)
    - [Build](#build)
    - [Download](#download)
        - [Linux](#linux)
        - [MacOS](#macos)
        - [Windows](#windows)
- [Goals](#goals)

## Features
- HTTP verb agnostic
- Sleep request by passing queryString `fs-sleep` with time in ms
  - Eg: `http://localhost:8088/my-super-test?fs-sleep=5000`
- Status code agnostic by passing queryString `fs-status`
  - Eg: `http://localhost:8088/my-super-test?fs-status=404`
- Response contents:
  - body
  - errors
  - headers
  - method
  - path
  - query
  - time

## Example
> Inside a Docker, localhost is called: *http://host.docker.internal:8088*

**Request:**
```bash
curl --request POST \
  --url http://localhost:8088/cart/62582df20e2cd6a770aa9e5b \
  --header 'Content-Type: application/json' \
  --data '{
    "cart_id": 123,
    "items": [
      { "name": "Smartphone", "price": 155.99, "quantity": 1 },
      { "name": "Mouse", "price": 23.50, "quantity": 2 }
    ]}'
```

**Response:**
```json
{
  "body": {
    "cart_id": 123,
    "items": [
      {
        "name": "Smartphone",
        "price": 155.99,
        "quantity": 1
      },
      {
        "name": "Mouse",
        "price": 23.5,
        "quantity": 2
      }
    ]
  },
  "errors": {},
  "headers": {
    "Accept": [ "*/*" ],
    "Content-Length": [ "170" ],
    "Content-Type": [ "application/json" ],
    "User-Agent": [ "insomnia/2022.2.1" ]
  },
  "method": "POST",
  "path": "/cart/62582df20e2cd6a770aa9e5b",
  "query": {},
  "time": "2022-04-14T19:13:21-03:00"
}
```

## Install
You can access the [latest releases](https://github.com/budnieswski/fake-server/releases/latest) page and download the compressed file according to your operating system.

### Build
> You must have Golang installed to be able to make the build
```shell
$ git clone https://github.com/budnieswski/fake-server.git
$ cd fake-server
$ go build .
$ ./fake-server # Run application
```

### Download
#### Linux
```bash
$ wget -q -O - https://github.com/budnieswski/fake-server/releases/latest/download/fake-server_Linux_x86_64.tar.gz \
| tar -xz

$ ./fake-server # Run application
```

#### MacOS
##### M1 architecture (arm64)
```bash
$ wget -q -O - https://github.com/budnieswski/fake-server/releases/latest/download/fake-server_macOS_arm64.tar.gz \
| tar -xz

$ ./fake-server # Run application
```
##### Intel/AMD architecture (x86_64)
```bash
$ wget -q -O - https://github.com/budnieswski/fake-server/releases/latest/download/fake-server_macOS_x86_64.tar.gz \
| tar -xz

$ ./fake-server # Run application
```

#### Windows
You can download the latest version [here](https://github.com/budnieswski/fake-server/releases/latest/download/fake-server_Windows_x86_64.zip)

## Goals
- [ ] Create flags
    - [ ] Help
    - [ ] Version
    - [ ] Run in background
    - [ ] Custom port
- [X] Create sleep request
- [X] Create custom status code
- [X] Add workflow to build app and create a release
    - [X] Linux
    - [X] MacOS
    - [X] Windows