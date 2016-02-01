package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	version = "0.0.1"
)

type task struct {
	user, method *string
}

func (t task) printinfo() (err error) {
	if !strings.Contains("sshttps", *t.method) {
		err = MethodError(*t.method)
	}
	fmt.Printf("Version %s of prepwd\nUser: %s\nMethod: %s\n", version, *t.user, *t.method)
	return
}

func (t task) prepwdwork() (num int, err error) {
	t.printinfo()

	return
}

func printusage() {
	var usagestr = `Usage:
prepwd <method> <user>
prepwd <user>

the latter will use https as default
`
	fmt.Print(usagestr)
}

func main() {
	var user string
	var method string
	switch len(os.Args) {
	case 1:
		// fmt.Print("Enter Username: ")
		// fmt.Scanln(&user)
		printusage()
		os.Exit(1)
	case 2:
		fmt.Println("Using HTTPS as default")
		user = os.Args[1]
		method = "https"
		break
	case 3:
		user = os.Args[2]
		method = strings.ToLower(os.Args[1])
		if !(method == "https" || method == "ssh") {
			panic(MethodError(method))
		}
		break
	default:
		printusage()
		os.Exit(1)
	}
	task{user: &user, method: &method}.prepwdwork()
}
