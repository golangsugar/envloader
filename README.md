# EnvLoader

[![Go Report Card](https://goreportcard.com/badge/github.com/golangsugar/envloader)](https://goreportcard.com/report/github.com/golangsugar/envloader)
[![GoDoc](https://godoc.org/github.com/golangsugar/envloader?status.svg)](https://godoc.org/github.com/golangsugar/envloader)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](

`EnvLoader` is a Go package that provides a convenient way to load environment variable values from a text file into a map.

## Installation

Use the following command to install `envloader`:

```bash
go get github.com/your-username/envloader
```

## Usage

Import the `envloader` package in your Go code:

```go
import "github.com/golangsugar/envloader"
```

### Load From File

To load environment variables from a text file, use the `LoadFromFile` function:

```go
err := envloader.LoadFromFile("example.env", true)
if err != nil {
    log.Fatal(err)
}
```

The `LoadFromFile` function takes the following parameters:
- `configFile` (string): The name of the text file containing the environment variable values.
- `errorIfFileDoesntExist` (bool): If set to `true`, an error will be returned if the file doesn't exist. If set to `false`, the function will not return an error and continue execution.

The text file should contain valid environment variable lines. Each line should adhere to the format: `KEY=VALUE`. Lines starting with `#` are considered comments and will be ignored.

Here is an example of a valid `config.txt` file:

```plaintext
DATABASE_HOST=localhost
DATABASE_PORT=5432
API_KEY=abc123
```

### Logging File Closing Errors

By default, `envloader` does not log errors when closing the file after reading. However, you can enable logging of file closing errors by calling the `LogFileClosingError` function before using `LoadFromFile`:

```go
envloader.LogFileClosingError()
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/your-username/envloader).

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). See the [LICENSE](LICENSE) file for details.