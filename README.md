# logger

[![Build Status](https://secure.travis-ci.org/goforgery/logger.png?branch=master)](http://travis-ci.org/goforgery/logger)

Request logger with custom format support for Forger2.

## Use

By default writes Apache style logs to `stdout`.

```javascript
package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/logger"
)

func main() {
	app := f.CreateApp()
	app.Use(logger.Create())
	app.Listen(3000)
}
```

## Test

    go test
