package gclone

import (
	"fmt"
	"os/exec"
)

const (
	sshGithubTemplate   = "git@github.com:%s/%s"
	httpsGithubTemplate = "https://github.com/%s/%s"
)

// CloneGithub clones github repo from user using method [https|ssh]
func CloneGithub(user, repo, method *string) {
	exec.Command(fmt.Sprintf("git clone "))
}

func sshCloneGithub(user, repo *string) {

}

func httpsCloneGithub(user, repo *string) {

}
