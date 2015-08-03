package main

var AppName string = "simple-cmd"

const usageTemplate = `Usage: simple-cmd command [arguments]

The commands are:
{{range .}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "simple-cmd help [command]" for more information.
`

const helpTemplate = `Usage: simple-cmd {{.UsageLine}}
{{.Long}}
`

const header = `-----
-- simple-cmd
-----
`
