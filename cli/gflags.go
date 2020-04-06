package cli

import "flag"

// Flags : Parses out global flags
type GFlags struct {
	Version bool
}

// Parse : ...
func (f GFlags) Parse() GFlags {
	version := flag.Bool("v", false, "prints version")

	flag.Parse()

	return GFlags{
		Version: *version,
	}
}
