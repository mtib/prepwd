package gclone

import (
	"os"
	"testing"
)

type repoTest struct {
	user, repo string
}

type tests []repoTest

var successfull = tests{repoTest{"mtib", "prepwd"}}

func TestHttpsClone(m *testing.T) {
	mClone(m, "https")
}

func TestSshClone(m *testing.T) {
	mClone(m, "ssh")
}

func mClone(m *testing.T, method string) {
	os.Chdir("/tmp")
	for index := 0; index < len(successfull); index++ {
		err := CloneGithub(&successfull[index].user, &successfull[index].repo, &method)
		if err != nil {
			m.FailNow()
		}
		os.RemoveAll(successfull[index].repo)
	}
}
