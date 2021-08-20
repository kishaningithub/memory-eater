# memory-eater
Eats memory at a specified constant rate

![build workflow](https://github.com/kishaningithub/memory-eater/actions/workflows/build.yml/badge.svg)

## Installation

### Home brew

```shell
$ brew install kishaningithub/tap/memory-eater
```

### Yum

```shell
yum install -y https://github.com/kishaningithub/memory-eater/releases/download/v0.0.8/memory-eater_0.0.8_linux_amd64.rpm
```

### Systemd

Do the yum installation as above and do the following
```shell
curl https://raw.githubusercontent.com/kishaningithub/memory-eater/main/memory-eater.service --output /etc/systemd/system/memory-eater.service
```

For a detailed list see [releases page](https://github.com/kishaningithub/memory-eater/releases)

## Examples

Eat 100MB per second
```shell
memory-eater --step-size=100MB --step-duration=1s
```

## Usage

```shell
$ memory-eater -h
Eats memory at a given constant rate

Usage:
  memory-eater [flags]

Examples:
memory-eater --step-size=100MB --step-duration=1s

Flags:
  -h, --help                   help for memory-eater
      --profile                profile this cli tool
      --step-duration string   Duration of every step. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h" (default "1s")
      --step-size string       Size of memory to eat during every step. Valid byte units are "B", "KB", "MB", "GB", "TB", "PB" and "EB" (default "1MB")
```

## Advanced examples

## Profiling the utility

1. To profile just add `--profile` flag to any command you want to profile
2. To view the results run `go tool pprof -http=:1235 mem.pprof`
3. Open url `http://localhost:1235` in your browser