# xvideos [![GoDoc](https://godoc.org/github.com/pdevty/xvideos?status.svg)](https://godoc.org/github.com/pdevty/xvideos) [![Build Status](https://travis-ci.org/pdevty/xvideos.svg)](https://travis-ci.org/pdevty/xvideos) [![Coverage Status](https://coveralls.io/repos/pdevty/xvideos/badge.svg)](https://coveralls.io/r/pdevty/xvideos)

This package scrapes the HTML of xvideos.com and gives you information you can use in your go programs

## Installation

execute:

    $ go get github.com/pdevty/xvideos

## Usage

```go
package main

import (
	"fmt"
	"github.com/pdevty/xvideos"
	"log"
)

func main() {
	// url to be set xvideos list page
	xv, err := xvideos.Get("http://jp.xvideos.com/c/asian_woman-32/")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range xv {
		fmt.Println(v.Id, v.Url, v.Duration, v.Rating, v.Thumbnail, v.Title, v.Tags)
	}
}
```

Refer to [godoc](http://godoc.org/github.com/pdevty/xvideos) for more infomation.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
