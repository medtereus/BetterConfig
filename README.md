# BetterConfig

**A small JSON-based configuration library for Go.**

BetterConfig — небольшая библиотека для хранения пользовательских настроек в JSON-файле. Предоставляет простой API с типизированными геттерами и сеттерами, автоматическим сохранением изменений и поддержкой значений по умолчанию.

Проект вдохновлён подходом к работе с настройками в [Fyne Preferences](<https://docs.fyne.io/>), но представляет собой самостоятельную, упрощённую реализацию на основе JSON.

> **Personal project.** This library was created for my own use to replace an earlier configuration storage solution that no longer suited my needs. It is not intended to be a universal configuration framework.

## Features

- JSON-based configuration storage.
- Simple, typed getters and setters.
- Support for `bool`, `string`, `int`, `uint` and `float64`.
- Fallback values when a key is missing or its value has an unexpected type.
- Explicit removal of configuration values.
- Persistent storage between application runs.
- No third-party dependencies required by the core implementation.

## Project structure

text

```
BetterConfig/
├── 0_InitConfig.go
├── getValueTypeWithFallbackGeneric.go
├── utils.go
├── v_bool.go
├── v_float.go
├── v_generic.go
├── v_int.go
├── v_string.go
├── v_uint.go
├── v__RemoveValue.go
└── go.mod
```

The implementation is intentionally straightforward:

- `0_InitConfig.go` — initialization, JSON loading and saving.
- `getValueTypeWithFallbackGeneric.go` — shared generic helper for retrieving values with fallback support.
- `utils.go` — configuration path utilities and initialization state.
- `v_bool.go` — boolean getters and setters.
- `v_float.go` — floating-point getters and setters.
- `v_int.go` — signed integer getters and setters.
- `v_string.go` — string getters and setters.
- `v_uint.go` — unsigned integer getters and setters.
- `v_generic.go` — generic value access.
- `v__RemoveValue.go` — removal of configuration entries.

## Installation

This is primarily a personal library. If you want to use it in another Go project, clone the repository and configure the module dependency accordingly.

For local development, a Go module `replace` directive can point to a local checkout:

go

```
require BetterConfig v0.0.0

replace BetterConfig => ./BetterConfig
```

Adjust the paths and module name to match your repository layout.

## Usage

### Initialization

Initialize BetterConfig before accessing configuration values:

go

```
package main

import (
    "log"
    "path"

    "BetterConfig"
)

func main() {
    err := BetterConfig.InitBetterConfig(
        path.Join("Medtereus", "TestingMyConfigLibrary"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Work with configuration values here.
}
```

The argument identifies the application-specific configuration folder. The library resolves the configuration directory and uses `config.json` for persistent storage.

### Setting and getting values

go

```
BetterConfig.SetBool("ENABLE_FEATURE", true)
BetterConfig.SetString("USERNAME", "Bobr Kurwa")
BetterConfig.SetInt("RETRY_COUNT", 5)
BetterConfig.SetUInt("MAX_ITEMS", 100)
BetterConfig.SetFloat("SCALE", 3.141592)
```

Read the values back using the corresponding getters:

go

```
enabled := BetterConfig.GetBool("ENABLE_FEATURE")
username := BetterConfig.GetString("USERNAME")
retries := BetterConfig.GetInt("RETRY_COUNT")
maxItems := BetterConfig.GetUInt("MAX_ITEMS")
scale := BetterConfig.GetFloat("SCALE")
```

Setters save the updated configuration to disk.

### Fallback values

The `WithFallback` variants let you specify a value to return when a key is missing or the stored value cannot be read as the requested type.

go

```
timeout := BetterConfig.GetIntWithFallback("TIMEOUT", 30)
username := BetterConfig.GetStringWithFallback("USERNAME", "Guest")
enabled := BetterConfig.GetBoolWithFallback("ENABLED", false)
scale := BetterConfig.GetFloatWithFallback("SCALE", 1.0)
maxItems := BetterConfig.GetUIntWithFallback("MAX_ITEMS", 100)
```

This is useful when an option has not been configured yet.

### Removing a value

Remove an entry using:

go

```
BetterConfig.RemoveValue("OLD_SETTING")
```

## Storage format

Configuration is stored as a JSON object whose keys correspond to setting names.

For example:

json

```
{
  "ENABLE_FEATURE": true,
  "USERNAME": "Bobr Kurwa",
  "RETRY_COUNT": 5,
  "MAX_ITEMS": 100,
  "SCALE": 3.141592
}
```

This format is human-readable and can be inspected or edited manually.

### Numeric values and JSON decoding

BetterConfig uses `json.Decoder.UseNumber()` when loading configuration data into `map[string]any`.

By default, Go's JSON decoder converts JSON numbers into `float64` when decoding into an interface value. This can cause precision loss for large integers, especially values above $2^{53}$.

With `UseNumber()`, JSON numbers are represented as `json.Number`, preserving their decimal representation until they are explicitly converted.

Numeric getters handle `json.Number` and convert values to the requested Go type. Integer conversions use range-aware parsing and return the configured fallback if conversion fails.

This allows the library to preserve large integer values, including 64-bit unsigned values, across JSON save-and-load cycles without passing them through `float64`.


## Testing

The project includes tests for writing and reading configuration values.

The test cases cover:

- Boolean values: `true` and `false`.
- Positive and negative floating-point values.
- Strings containing newlines, tabs, Cyrillic characters and German umlauts.
- Signed integer boundary values.
- Unsigned integer boundary values, including the maximum 64-bit unsigned integer where supported.

The write and read tests are separated:

- `set_test.go` writes the test values.
- `get_test.go` reads the previously saved values and checks them against the expected results.

To run the tests:

bash

```
go test -v
```

To run the write and read test functions separately:

bash

```
go test -run '^Test_SetAll$' -count=1
go test -run '^Test_GetAll$' -count=1
```

Run the write test first, followed by the read test. The read test expects the configuration file to contain the values written by the write test.

## Limitations

BetterConfig is intentionally minimal and has not been designed as a general-purpose configuration framework.

- Configuration values are stored in a single JSON file.
- Setters persist changes immediately.
- The library uses package-level state.
- It is not designed for concurrent writes without additional synchronization.
- It does not provide configuration schema validation, migrations, change notifications or automatic recovery from corrupted files.
- Generic values are subject to JSON serialization and deserialization constraints. Arbitrary Go values are not guaranteed to survive a complete save-and-load cycle with their original types.
- Numeric values must be handled carefully when decoding JSON into `any`, particularly when preserving large integers.

Use it with these limitations in mind.

## Motivation

I created BetterConfig for myself after becoming dissatisfied with my previous approach to configuration storage.

I wanted something small and predictable: a JSON file, a few typed getters and setters, fallback values, and persistent storage without introducing a larger framework into my project.

The implementation is deliberately simple because it solves a personal problem. The goal was not to reproduce all of Fyne Preferences or build a universal configuration management system.

**This project is primarily intended for my own use.** It is shared publicly for reference and convenience, but it comes without any promise of suitability for every application or use case.
