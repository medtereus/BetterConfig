# BetterConfig

A small JSON-based configuration library for Go.

**Personal project.** Created for my own use as a replacement for an earlier configuration storage solution. It is not intended to be a universal configuration framework.

Inspired by the configuration API of [Fyne Preferences](https://docs.fyne.io/api/v2/fyne/preferences.html), but implemented independently without Fyne dependencies.

## Features

- JSON-based persistent storage.
- Typed getters and setters: `bool`, `string`, `int`, `uint`, `float64`.
- Generic access with `GetGeneric[T]` and `SetGeneric[T]`.
- Fallback values for missing or incompatible entries.
- Automatic saving after modifications.
- Explicit removal of configuration entries.
- Large integer support through `json.Number` and `Decoder.UseNumber()`.

## Installation

```bash
go get github.com/medtereus/BetterConfig@latest
```

## Usage

### Initialization

Initialize the library before accessing configuration values:

```go
package main

import (
    "log"

    "github.com/medtereus/BetterConfig"
)

func main() {
    if err := BetterConfig.InitBetterConfig("MyApp"); err != nil {
        log.Fatal(err)
    }

    // Configuration is ready.
}
```

The argument specifies an application-specific configuration directory under the operating system's user configuration directory.

The library creates and uses `config.json` inside that directory.

### Typed values

```go
BetterConfig.SetBool("ENABLED", true)
BetterConfig.SetString("USERNAME", "Solar")
BetterConfig.SetInt("RETRIES", 5)
BetterConfig.SetUInt("MAX_ITEMS", 100)
BetterConfig.SetFloat("SCALE", 3.141592)

enabled := BetterConfig.GetBool("ENABLED")
username := BetterConfig.GetString("USERNAME")
retries := BetterConfig.GetInt("RETRIES")
```

### Fallback values

Use `WithFallback` methods to provide defaults when a key is missing or its value cannot be converted to the requested type.

```go
timeout := BetterConfig.GetIntWithFallback("TIMEOUT", 30)
username := BetterConfig.GetStringWithFallback("USERNAME", "Guest")
```

### Generic values

Store and retrieve custom Go structures:

```go
type User struct {
    Name string
    Age  uint8
}

BetterConfig.SetGeneric("USER", User{
    Name: "Solar",
    Age:  27,
})

user := BetterConfig.GetGeneric[User]("USER")
```

Generic values are serialized and deserialized through JSON.

### Removing values

```go
BetterConfig.RemoveValue("OLD_SETTING")
```

## Storage

Configuration is stored as a human-readable JSON file.

JSON numbers are decoded using `json.Decoder.UseNumber()` to preserve integer precision, including 64-bit values.

## Limitations
- Package-level global state.
- One configuration file per application.
- Setters save immediately.
- No built-in synchronization for concurrent writes.
- No schema validation, migrations or automatic recovery from corrupted configuration files.
- Generic values must be compatible with JSON serialization.
- Setters do not expose file-saving errors to the caller.

## Testing

Run tests with:

```bash
go test -v ./...
```

Some read/write tests depend on previously saved test data and must be executed in the appropriate order.
