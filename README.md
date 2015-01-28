# Go Fuck

A Brainfuck VM written in Go for funsies

## Setup

` go get github.com/adamveld12/gofuck `


## Testing

Run `goconvey` in the project directory and browse to [localhost:8080](localhost:8080)
This also works with `go test ./vm`

## How to use

``` go
import (
  "fmt"
  "github.com/adamveld12/gofuck/vm"
)

func main(){
   input, output := vm.Execute(",>,<[->+<]>.")
   input <- 2
   input <- 3
   fmt.Println(<-output)
}
```

You can also go install it and use this at the command line like so:

``` sh
 $ gofuck <file>
```

or

``` sh
 $ gofuck
 $ <type or paste BF code here>
```
and on new line the code is executed and the output printed

## Contributing changes

- [Git Guildines](https://github.com/thoughtbot/guides/tree/master/protocol/git)
- [How I Review Code](https://github.com/thoughtbot/guides/tree/master/code-review)
- [My Coding Practices](https://github.com/thoughtbot/guides/tree/master/best-practices)
- [My Style Guide](https://github.com/thoughtbot/guides/tree/master/style)
- "Please open github issues" and other words of encouragement

## License

MIT
