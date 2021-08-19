# memory-eater
Eats memory at a specified constant rate

![build workflow](https://github.com/kishaningithub/memory-eater/actions/workflows/build.yml/badge.svg)

## Installation

```shell
$ brew install kishaningithub/tap/memory-eater
```

## Upgrading

```shell
$ brew upgrade memory-eater
```

## Examples

Eat 100MB per second
```shell
memory-eater --step-size=100MB --step-duration=1s
```

## Usage

```shell
$ memory-eater -h
Eats memory at a constant rate

Usage:
  memory-eater [flags]

Examples:
memory-eater --step-size=100MB --step-duration=1s

Flags:
  -h, --help                   help for memory-eater
      --step-duration string    (default "1s")
      --step-size string        (default "1MB")
```