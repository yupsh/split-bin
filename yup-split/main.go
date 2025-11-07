package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/split"
)

const (
	flagLines        = "lines"
	flagBytes        = "bytes"
	flagSize         = "line-bytes"
	flagNumeric      = "numeric-suffixes"
	flagSuffixLength = "suffix-length"
	flagVerbose      = "verbose"
)

func main() {
	app := &cli.App{
		Name:  "split",
		Usage: "split a file into pieces",
		UsageText: `split [OPTIONS] [FILE [PREFIX]]

   Output pieces of FILE to PREFIXaa, PREFIXab, ...; default size is 1000 lines,
   and default PREFIX is 'x'.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagLines,
				Aliases: []string{"l"},
				Usage:   "put NUMBER lines per output file",
				Value:   1000,
			},
			&cli.IntFlag{
				Name:    flagBytes,
				Aliases: []string{"b"},
				Usage:   "put SIZE bytes per output file",
			},
			&cli.StringFlag{
				Name:    flagSize,
				Aliases: []string{"C"},
				Usage:   "put at most SIZE bytes of records per output file",
			},
			&cli.BoolFlag{
				Name:    flagNumeric,
				Aliases: []string{"d"},
				Usage:   "use numeric suffixes starting at 0, not alphabetic",
			},
			&cli.IntFlag{
				Name:    flagSuffixLength,
				Aliases: []string{"a"},
				Usage:   "generate suffixes of length N (default 2)",
				Value:   2,
			},
			&cli.BoolFlag{
				Name:  flagVerbose,
				Usage: "print a diagnostic just before each output file is opened",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (file and optional prefix)
	for i := 0; i < c.NArg(); i++ {
		arg := c.Args().Get(i)
		// First arg is file, second is prefix
		if i == 0 {
			params = append(params, yup.File(arg))
		} else if i == 1 {
			params = append(params, Prefix(arg))
		}
	}

	// Add flags based on CLI options
	if c.IsSet(flagLines) {
		params = append(params, Lines(c.Int(flagLines)))
	}
	if c.IsSet(flagBytes) {
		params = append(params, Bytes(c.Int(flagBytes)))
	}
	if c.IsSet(flagSize) {
		params = append(params, Size(c.String(flagSize)))
	}
	if c.Bool(flagNumeric) {
		params = append(params, Numeric)
	}
	if c.IsSet(flagSuffixLength) {
		params = append(params, SuffixLength(c.Int(flagSuffixLength)))
	}
	if c.Bool(flagVerbose) {
		params = append(params, Verbose)
	}

	// Create and execute the split command
	cmd := Split(params...)
	return yup.Run(cmd)
}
