# Docflag
Docflag implements command-line flag parsing with simple and intuitive style.
The code is inspired by the revel/cmd package.

```
import "github.com/Zypeh/docflag"
func main() {
	docflag.Parse()
}
```

It offers:

* both short and long description
use "help [command]" for more information
here is how we include our command's description
```
var cmdTest = &Command{
	UsageLine: "test [time of repeats] [string]",
	Short "Simple Example - Echo function",
	Long: `
This is the long description to be shown when you help [command].
`
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
package docflag

func init () {
	cmdBuild.Run = build
}

func Build (args []string) { ... }
```

* Usage template file.
simply edit the template file in the App.cfg
```
[App]
AppName = "Echo-go"
Header = "-----[[ Simple Echo ]]-----"

[Template]
usageTemplate = "Usage: echo-go command [arguments]"
// SNIP...
```


	
