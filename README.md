# belt-conveyor
Converts image files placed on the watch point to another format.

## Usage
```
USAGE:
   go run cmd/webp.go [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --watchpoint Directory Path, --wp Directory Path  Directory Path to watch.
   --destination Directory Path, -d Directory Path   Directory Path to store conversion result files.
   --quality value, -q value                         For WEBP, it can be a quality from 1 to 100 (the higher is the better). By default (without any parameter) and for quality above 100 the lossless compression is used. (default: 100)
   --help, -h                                        show help
```

## Required
[An OpenCV installation environment that meets GoCV requirements is required.](https://github.com/hybridgroup/gocv?tab=readme-ov-file#how-to-install)
Therefore, it cannot be built at this time and must be run in a Go installation environment.
