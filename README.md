# Hexlet Path Size

hexlet-path-size is a small CLI utility written in Go that calculates the size of files and directories.

The application supports:
- recursive directory traversal
- human-readable output
- hidden files filtering

## Installation

Clone the repository:
```bash
git clone <repository-url>
cd <repository-name>
```

Build the application:

```bash
make build
```
The binary will be available at:

```bash
./bin/hexlet-path-size
```

## Usage

Basic usage:
```bash
./bin/hexlet-path-size <path>
```
Example:

```bash
./bin/hexlet-path-size ./testdata
```

Output:
```text
4096B    ./testdata
```

### Flags

#### Recursive directory traversal

Use -r or --recursive to calculate directory size recursively:
```bash
./bin/hexlet-path-size -r ./testdata
```

#### Human-readable format

Use -H or --human to print sizes in a human-readable format:
```bash
./bin/hexlet-path-size -H ./testdata
```

Example output:
```bash
4.2MB    ./testdata
```
#### Include hidden files

By default, hidden files and directories are skipped.
Use -a or --all to include them:

```bash
./bin/hexlet-path-size -a ./testdata
```

## Extra

[Asciinema Demo](https://asciinema.org/a/gU2oLsY1fdSpZu43)

### Hexlet tests and linter status:

[![Actions Status](https://github.com/Fr0stFree/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Fr0stFree/go-project-242/actions)

