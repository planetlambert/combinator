# Combinator
A complete and open source implementation of Moses Schönfinkel's 1924 paper - [On the Building Blocks of Mathematical Logic](https://content.wolfram.com/uploads/sites/43/2020/12/Schonfinkel-OnTheBuildingBlocksOfMathematicalLogic.pdf).

## Guide

See the section-by-section guide to the paper [here](./GUIDE.md).

### Progress

- [X] Introduction by Quine
- [ ] Section 1
- [ ] Section 2
- [ ] Section 3
- [ ] Section 4
- [ ] Section 5
- [ ] Section 6

## Usage

```shell
go get github.com/planetlambert/combinator@latest
```

```go
import (
    "fmt"

    "github.com/planetlambert/combinator"
)

func main() {
    // Use a built-in basis (SKI in this example)
    transformedStatement, _ := combinator.SKI.Transform("S(K(SI))Kab")

    // Prints "ba" - S(K(SI))K is the "reversal" combinator
    fmt.Println(transformedStatement)

}
```

[Go Package Documentation here.](https://pkg.go.dev/github.com/planetlambert/combinator)

## Testing

```shell
go test ./...
```