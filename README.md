<h1 align="center">Verify My Test API 🍺</h1>
<p>
    <img src="https://shields.io/badge/Go-%5E16.6-red?logo=go">

</p>

> VerifyMyAge Code Challenge

## Prerequisites

<li><a href="https://golang.org/doc/install">Go</a></li>
<li><a href="https://www.docker.com/get-started">Docker</a></li>
<li><a href="https://docs.docker.com/compose/">Docker Compose</a></li>

## Local Setup

Clone this repo and shell to its path and run docker compose. Make sure 3306 and 8989 ports are available locally

```sh
git clone https://github.com/olivic9/verifyMyTest.git
cd verifyMyTest
mkdir "logs"
docker-compose up -d
```

Build

```sh
go build -o verify-my-test
```

Setup

```sh
./verify-my-test migrate

Or

make migrate
```

## Usage

All endpoints are under http://localhost:8989 check <a href="http://localhost:8989/swagger/index.html/">Swagger</a> docs
for further info.

## Run tests

```sh
go test ./test

Or

make test

```

## Author

👤 **Clayson Oliveira**

* Linkedin: [Clayson Oliveira](https://www.linkedin.com/in/clayson-oliveira-603a853b/)
* Github: [@olivic9](https://github.com/olivic9)