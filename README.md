# GoProBar

GoProBar is a simple rendering progress bar in terminal application / CLI

## Installation

Use the [Go Modules](https://go.dev/ref/mod) to install GoProBar.

```bash
go get github.com/seyuta/goprobar
```

## Usage

```go
import (
	probar "github.com/seyuta/goprobar"
)

func main() {
	total := 100000

	for i := 1; i <= total; i++ {
		probar.ProgressBar(i, total)
	}
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)