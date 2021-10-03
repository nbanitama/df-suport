package main

import "github.com/nbanitama/df-supoort/df"

var folders = []string{
	"/home/novandibanitama/Documents/agent/2 oct/toped-1/intents",
	"/home/novandibanitama/Documents/agent/2 oct/toped-2/intents",
	"/home/novandibanitama/Documents/agent/2 oct/toped-3/intents",
	"/home/novandibanitama/Documents/agent/2 oct/toped-4/intents",
}

func main() {
	df.FilterByActionAndLifespan(folders, "ConvertParametersToEventName", 3)

	df.Export("noba-30f2f", "files/service_account/sa.json", "/home/novandibanitama/Documents/b.zip")
}
