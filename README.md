# Fake Server
![Project status](https://img.shields.io/badge/version-1.0.1-green.svg)
![License](https://img.shields.io/dub/l/vibe-d.svg)

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
##### 64-bit (more common architecture)
```bash
$ curl -s https://api.github.com/repos/budnieswski/fake-server/releases/latest \
| awk -F\" '/browser_download_url.*linux_amd64.tar.gz/{print $(NF-1)}' \
| xargs wget -O - -q \
| tar -xz

$ ./fake-server # Run application
```
##### 32-bit
```bash
$ curl -s https://api.github.com/repos/budnieswski/fake-server/releases/latest \
| awk -F\" '/browser_download_url.*linux_386.tar.gz/{print $(NF-1)}' \
| xargs wget -O - -q \
| tar -xz

$ ./fake-server # Run application
```
##### arm64
```bash
$ curl -s https://api.github.com/repos/budnieswski/fake-server/releases/latest \
| awk -F\" '/browser_download_url.*linux_arm64.tar.gz/{print $(NF-1)}' \
| xargs wget -O - -q \
| tar -xz

$ ./fake-server # Run application
```

#### MacOS
##### M1 architecture (arm64)
```bash
$ curl -s https://api.github.com/repos/budnieswski/fake-server/releases/latest \
| awk -F\" '/browser_download_url.*_darwin_arm64.tar.gz/{print $(NF-1)}' \
| xargs wget -O - -q \
| tar -xz

$ ./fake-server # Run application
```
##### Intel/AMD architecture (amd64)
```bash
$ curl -s https://api.github.com/repos/budnieswski/fake-server/releases/latest \
| awk -F\" '/browser_download_url.*_darwin_amd64.tar.gz/{print $(NF-1)}' \
| xargs wget -O - -q \
| tar -xz

$ ./fake-server # Run application
```

#### Windows
The build is not currently being generated, you need to download the source and compile yourself ;/
> See the [Build](#build) section

## Goals
- [ ] Create flags
    - [ ] Help
    - [ ] Custom port
- [ ] Create sleep request
- [ ] Create custom response code
- [X] Add workflow to build app and create a release
    - [X] Linux
    - [X] MacOS
    - [ ] Windows