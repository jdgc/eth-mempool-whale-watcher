package utils

import "github.com/jdgc/eth-mempool-whale-watcher/constants"

func RedString(s string) string {
	return constants.COLOR_RED + s + constants.COLOR_RESET
}

func GreenString(s string) string {
	return constants.COLOR_GREEN + s + constants.COLOR_RESET
}

func YellowString(s string) string {
	return constants.COLOR_YELLOW + s + constants.COLOR_RESET
}

func BlueString(s string) string {
	return constants.COLOR_BLUE + s + constants.COLOR_BLUE
}
