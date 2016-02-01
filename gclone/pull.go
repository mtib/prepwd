package gclone

import (
	"fmt"
	"os/exec"
)

const (
	sshGithubTemplate   = "git@github.com:%s/%s"
	httpsGithubTemplate = "https://github.com/%s/%s"
	sshGistTemplate     = "git@gist.github.com:%v"
	httpsGistTemplate   = "https://gist.github.com/%v"
	defaultDepth        = 10
)

type downloadError string

func (d downloadError) Error() string {
	return fmt.Sprintf("Download Error: %v", string(d))
}

// CloneGithub clones github repo from user using method [https|ssh]
func CloneGithub(user, repo, method *string) (err error) {
	switch *method {
	case "https":
		err = httpsCloneGithub(user, repo)
	case "ssh":
		err = sshCloneGithub(user, repo)
	default:
		err = downloadError("unknown method")
	}
	return
}

// CloneGithubGist clones gists using method [https|ssh]
func CloneGithubGist(gist, name, method *string) (err error) {
	var clone *exec.Cmd
	var remote string
	switch *method {
	case "ssh":
		remote = fmt.Sprintf(sshGistTemplate, *gist)
	case "https":
		remote = fmt.Sprintf(httpsGistTemplate, *gist)
	default:
		err = downloadError("unknown method")
		return
	}
	clone = exec.Command("git", "clone", remote, *name)
	err = clone.Start()
	if err != nil {
		return
	}
	fmt.Print("Cloning Gist")
	clone.Wait()
	fmt.Println(" ... done")
	return
}

func sshCloneGithub(user, repo *string) error {
	return bareClone(sshGithubTemplate, user, repo)
}

func httpsCloneGithub(user, repo *string) error {
	return bareClone(httpsGithubTemplate, user, repo)
}

func bareClone(templ string, user, repo *string) (err error) {
	clone := exec.Command("git", "clone", fmt.Sprintf("--depth=%v", defaultDepth), fmt.Sprintf(templ, *user, *repo))
	err = clone.Start()
	if err == nil {
		fmt.Printf("Cloning %v/%v ", *user, *repo)
		clone.Wait()
		fmt.Println("... done")
	}
	return
}
