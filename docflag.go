package docflag

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"
	
	"github.com/agtorre/gocolorize"
	"gopkg.in/gcfg.v1"
)

// Use wrapper to differentiate logged panics from unexpected ones.
type LoggedError struct { error }

// Read from configuration file
type Config struct {

	// Apps Configurations
	App struct {
		AppName string
		Header string
	}

	// Template string
	Template struct {
		usageTemplate string
		helpTemplate string
	}
}

// Steal from the revel/cmd
type Command struct {
	Run                    func (args []string)
	UsageLine, Short, Long string
}

func (cmd *Command) Name () string {
	name := cmd.UsageLine
	i := strings.Index (name, " ")
	if i >= 0 {
		name = name [:i]
	}
	return name
}

// Add your command here
var commands = []*Command {
	cmdTest,
}

var AppCfg Config

func Parse () {
	if runtime.GOOS == "windows" {
		gocolorize.SetPlain (true)
	}

	var err = gcfg.ReadFileInto (&AppCfg, "App.cfg")
	if err != nil {
		errorf ("Missing App.cfg")
		os.Exit(1)
	}
	
	fmt.Fprintf (os.Stdout, gocolorize.NewColor ("Blue").Paint (AppCfg.App.Header))
	flag.Usage = func () { usage (1) }
	flag.Parse ()
	args := flag.Args ()

	if len (args) < 1 || args [0] == "help" {
		if len (args) == 1 {
			usage (0)
		}
		if len (args) > 1 {
			for _, cmd := range commands {
				if cmd.Name () == args [1] {
					tmpl (os.Stdout, AppCfg.Template.helpTemplate, cmd)
					return
				}
			}
		}
		usage (2)
	}

	// Commands use panic to abort execution when something goes wrong.
	// Panics are logged at the point of error. Ignore those.
	defer func () {
		if err := recover (); err != nil {
			if _, ok := err.(LoggedError); !ok {
				// This panic was not expecred/logged.
				panic (err)
			}
			os.Exit (1)
		}
	} ()

	for _, cmd := range commands {
		if cmd.Name () == args [0] {
			cmd.Run (args [1:])
			return
		}
	}

	errorf ("Unknown command.\n Run '%q help' for usage", args [0])
}

func errorf (format string, args ...interface{}) {
	// Ensure the user's command prompt starts on the next line.
	if !strings.HasSuffix (format, "\n") {
		format += "\n"
	}
	fmt.Fprintf (os.Stderr, format, args...)
	panic (LoggedError{}) // Panic instead of os.Exit so that the defer will run.
}

func usage (exitCode int) {
	gcfg.ReadFileInto (&AppCfg, "App.cfg")
	tmpl (os.Stderr, AppCfg.Template.usageTemplate, commands)
	os.Exit (exitCode)
}

func tmpl (w io.Writer, text string, data interface{}) {
	t := template.New ("top")
	template.Must (t.Parse (text))
	if err := t.Execute (w, data); err != nil {
		panic (err)
	}
}

func init () {
	rand.Seed (time.Now ().UnixNano ())
}
