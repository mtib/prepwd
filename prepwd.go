package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/mtib/prepwd/gclone"
)

const (
	version = "0.0.1"
	repoURL = "https://api.github.com/users/%v/repos"
	gistURL = "https://api.github.com/users/%v/gists"
	starURL = "https://api.github.com/users/%v/starred"
)

type task struct {
	user, method *string
}

type repos []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type gists []struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type stars []struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	SSHURL   string `json:"ssh_url"`
	CloneURL string `json:"clone_url"`
}

func (g *gists) escapeNames() {
	repl := func(s *string, o, n string) {
		*s = strings.Replace(*s, o, n, -1)
	}
	for k, v := range *g {
		v.Description = html.EscapeString(v.Description)
		repl(&v.Description, " ", "-")
		repl(&v.Description, "/", "")
		repl(&v.Description, ".", "_")
		(*g)[k].Description = v.Description
	}
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
	os.Mkdir(*t.user, os.ModePerm)
	os.Chdir(*t.user)
	os.Mkdir("repos", os.ModePerm)
	os.Chdir("repos")
	repoList, err := getUnmarshalRepos(*t.user)
	if err != nil {
		panic(err)
	}
	for _, v := range *repoList {
		gclone.CloneGithub(t.user, &v.Name, t.method)
	}
	os.Chdir("..")
	os.Mkdir("gists", os.ModePerm)
	os.Chdir("gists")
	gistList, err := getUnmarshalGists(*t.user)
	if err != nil {
		panic(err)
	}
	gistList.escapeNames()
	for k, v := range *gistList {
		fmt.Printf("[%v/%v] ", k+1, len(*gistList))
		gclone.CloneGithubGist(&v.ID, &v.Description, t.method)
	}
	os.Chdir("..")
	os.Mkdir("stars", os.ModePerm)
	os.Chdir("stars")
	starList, err := getUnmarshalStars(*t.user)
	if err != nil {
		// stars not very important
		return
	}
	for _, v := range *starList {
		owner := v.FullName[0:strings.Index(v.FullName, "/")]
		if owner != *t.user {
			gclone.CloneGithub(&owner, &v.Name, t.method)
		}
	}
	return
}

func getUnmarshalRepos(user string) (*repos, error) {
	resp, err := http.Get(fmt.Sprintf(repoURL, user))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var repoList repos
	err = json.Unmarshal(body, &repoList)
	if err != nil {
		return nil, err
	}
	return &repoList, nil
}

func getUnmarshalGists(user string) (*gists, error) {
	resp, err := http.Get(fmt.Sprintf(gistURL, user))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var gistList gists
	err = json.Unmarshal(body, &gistList)
	if err != nil {
		return nil, err
	}
	return &gistList, nil
}

func getUnmarshalStars(user string) (*stars, error) {
	resp, err := http.Get(fmt.Sprintf(starURL, user))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var starList stars
	err = json.Unmarshal(body, &starList)
	if err != nil {
		return nil, err
	}
	return &starList, nil
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
