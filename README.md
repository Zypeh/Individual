# Indivudual
simple-cmd implements command-line flag parsing with simple and intuitive style.
The code is stole from the revel/cmd package.

It offers:

* simple flag with a short description
```
./simple-cmd
-----
- simple-cmd
-----
Usage: simple-cmd command [arguments]

The commands are:

    test	Simple Example - Echo function

Use "simple-cmd help [command]" for more information.
```

* both short and long description
use "simple-cmd help [command]" for more information
here is how we include our command's description
```
var cmdTest = &Command{
	UsageLine: "test [time of repeats] [string]",
	Short "Simple Example - Echo function",
	Long: `
This is the long description to be shown when you help [command].
`,
}
```

* Simple adding commands
adding your command in main.go, for example, cmdBuild.
```
// Add your command here
var commands = []*Command {
	cmdTest,
	// .... Add your command here ...
	cmdBuild,
}
```
and start writing your function under the command.
```
package main

func init () {
	cmdBuild.Run = build
}

func Build (args []string) { ... }
```

* Usage template file.
simply edit the template file in the template.go
```
//...
var AppName string = "simple-cmd"
const usageTemplate = `Usage: simple-cmd command [argument]
...
....
.....
`
```


	
