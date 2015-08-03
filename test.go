package main

import (
	"fmt"
	"strconv"
)

var cmdTest = &Command{
	UsageLine: "test [time of repeats] [string]",
	Short: "Simple Example - Echo function",
	Long: `
This is the long description to be shown when you help [commands].
`,
}

var resp string = "Hello, World!"

func init () {
	cmdTest.Run = test
}

func test (args []string) {

	if len (args) == 0 {
		errorf ("No argument has given")
	}

	if len (args) == 1 {
		if cout, err := strconv.Atoi (args [0]); err == nil {
			fmt.Printf ("The program will repeat %d times\n", cout) 
			for i := 0; i < cout; i++ {
				fmt.Printf ("%s\n", resp)
			}
		}
	}
	if len (args) == 2 {
		resp = args [1]
		if cout, err := strconv.Atoi (args [0]); err == nil {
			fmt.Printf ("The program will repeat %s for %d times\n", resp, cout)
			for i := 0; i < cout; i++ {
				fmt.Printf ("%s\n", resp)
			}
		}
	}
}
