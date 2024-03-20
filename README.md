# Combinator
A complete and open source implementation of Moses Schönfinkel's 1924 paper - [On the Building Blocks of Mathematical Logic](https://content.wolfram.com/uploads/sites/43/2020/12/Schonfinkel-OnTheBuildingBlocksOfMathematicalLogic.pdf).

## Guide

See the section-by-section guide to the paper [here](./GUIDE.md).

### Progress

- [X] [Introduction by Quine](./GUIDE.md#introduction-by-quine)
- [X] [Section 1](./GUIDE.md#section-1)
- [X] [Section 2](./GUIDE.md#section-2)
- [X] [Section 3](./GUIDE.md#section-3)
- [X] [Section 4](./GUIDE.md#section-4)
- [X] [Section 5](./GUIDE.md#section-5)
- [X] [Section 6](./GUIDE.md#section-6)

## Usage

```shell
go get github.com/planetlambert/combinator@latest
```

```go
import (
    "context"
    "fmt"

    "github.com/planetlambert/combinator"
)

func main() {
    ctx := context.Background()

    // Use a built-in basis (SKI in this example)
    transformedStatement, _ := combinator.SKI.Transform(ctx, "S(K(SI))Kab")

    // Prints "ba" - S(K(SI))K is the "reversal" combinator
    fmt.Println(transformedStatement)

}
```

[Go Package Documentation here.](https://pkg.go.dev/github.com/planetlambert/combinator)

## Testing

```shell
go test ./...
```