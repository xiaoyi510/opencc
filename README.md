# OpenCC for Go

[![Go](https://github.com/griffinqiu/opencc/workflows/Go/badge.svg)](https://github.com/griffinqiu/opencc/actions?query=workflow%3AGo)

This is a Go version of OpenCC([Open Chinese Convert 開放中文轉換](https://github.com/BYVoid/OpenCC/)) which is a project for conversion between Traditional and Simplified Chinese developed by [BYVoid](https://www.byvoid.com/).

此项目是基于 Native Go 方式实现 OpenCC，利用 OpenCC 官方的词典，减少 C 库对环境的依赖，同时，基于 Go Embed 特性，可以让我们编译的时候，直接将字典打包到 Go 的二进制里面，以获得较好的开发部署体验。

## Installation

```sh
go get github.com/griffinqiu/opencc
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/griffinqiu/opencc"
)

func main() {
    s2t, err := opencc.New("s2t")
    if err != nil {
        log.Fatal(err)
    }


    in := `自然语言处理是人工智能领域中的一个重要方向。`
    out, err := s2t.Convert(in)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n%s\n", in, out)
    //自然语言处理是人工智能领域中的一个重要方向。
    //自然語言處理是人工智能領域中的一個重要方向。
}
```

## Conversions

- `s2t` Simplified Chinese to Traditional Chinese
- `t2s` Traditional Chinese to Simplified Chinese
- `s2tw` Simplified Chinese to Traditional Chinese (Taiwan Standard)
- `tw2s` Traditional Chinese (Taiwan Standard) to Simplified Chinese
- `s2hk` Simplified Chinese to Traditional Chinese (Hong Kong Standard)
- `hk2s` Traditional Chinese (Hong Kong Standard) to Simplified Chinese
- `s2twp` Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom
- `tw2sp` Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom
- `t2tw` Traditional Chinese (OpenCC Standard) to Taiwan Standard
- `t2hk` Traditional Chinese (OpenCC Standard) to Hong Kong Standard

## License

Apache License
