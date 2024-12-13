package src

import "flag"

type Command struct {
	flags   *flag.FlagSet
	Execute func(cmd *Command, args []string)
}
