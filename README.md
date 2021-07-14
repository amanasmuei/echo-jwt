<p align="center">
A Go Restful API for basic JWT authentication.
</p>

<br>

<p align="center">
   <a href="https://goreportcard.com/report/github.com/amanasmuei/echo-jwt"><img src="https://goreportcard.com/badge/github.com/amanasmuei/echo-jwt"></a>
</p>
<br>


## Table of Contents

-   [Installation](#installation)
-   [Examples](#examples)
-   [License](#license)


## 🚀 Quick Start

### Installation

```sh
$ go get https://github.com/amanasmuei/echo-jwt
```

### Instantiate API

Create `server.go`

```go

package main

import (

"github.com/amanasmuei/echo-jwt/server"

)

func main() {

    s := server.NewServer()
    s.Start(":8000")

}

```

Start server

```sh
$ go run server.go
```


## 📝 License

By contributing, you agree that your contributions will be licensed under its MIT License.

In short, when you submit code changes, your submissions are understood to be under the same [MIT License](http://choosealicense.com/licenses/mit/) that covers the project. Feel free to contact the maintainers if that's a concern.