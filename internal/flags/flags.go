package flags

import "github.com/spf13/pflag"

func BindTo(flagset *pflag.FlagSet) {
	flagset.StringP(
		"my-flag",
		"f",
		"my-default-value",
		"An example flag",
	)
}
