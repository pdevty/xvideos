# xvideos

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
	xv, err := xvideos.Get("http://jp.xvideos.com/c/asian_woman-32/")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range xv {
		fmt.Println(v.Id, v.Url, v.Duration, v.Rating, v.Thumbnail, v.Title, v.Tags)
	}
}
```

Refer to [godoc documentation](https://affiliate.dmm.com/api/guide/) for more infomation.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
