# yup-split

```
NAME:
   split - split a file into pieces

USAGE:
   split [OPTIONS] [FILE [PREFIX]]

      Output pieces of FILE to PREFIXaa, PREFIXab, ...; default size is 1000 lines,
      and default PREFIX is 'x'.
      With no FILE, or when FILE is -, read standard input.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --lines value, -l value          put NUMBER lines per output file (default: 1000)
   --bytes value, -b value          put SIZE bytes per output file (default: 0)
   --line-bytes value, -C value     put at most SIZE bytes of records per output file
   --numeric-suffixes, -d           use numeric suffixes starting at 0, not alphabetic (default: false)
   --suffix-length value, -a value  generate suffixes of length N (default 2) (default: 2)
   --verbose                        print a diagnostic just before each output file is opened (default: false)
   --help, -h                       show help
```
