package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/astaxie/beego/logs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

func main() {

	//url := os.Args[1]
	//directory := os.Args[2]
	url := "localhost:10022/tmp/test/aaa"
	directory := "/tmp/t/ccc012"
	//
	//currentUser, err := user.Current()
	//if err != nil {
	//	logs.Error(err)
	//}

	// Clone the given repository to the given directory
	logs.Info("git clone %s %s --recursive", url, directory)

	// Assuming id_rsa is available at ~/.ssh/id_rsa
	//sshAuth, err := ssh.NewPublicKeysFromFile("git", currentUser.HomeDir+"/.ssh/id_rsa", "admin")
	//if err != nil {
	//	logs.Error(err)
	//}

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: "/tmp/test/aaa0",
		//RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		//Auth:     sshAuth,
		Progress: os.Stdout,
	})

	if err != nil {
		logs.Error(err)
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		logs.Error(err)
	}
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		logs.Error(err)
	}

	fmt.Println(commit)
}

func ssh1() {
	//url := os.Args[1]
	//directory := os.Args[2]
	url := "git@cst.gzhu.edu.cn:9922/ded883227663adb608a07dd633a8/solutions.git"
	directory := "/tmp/ccc01"

	currentUser, err := user.Current()
	if err != nil {
		logs.Error(err)
	}

	// Clone the given repository to the given directory
	logs.Info("git clone %s %s --recursive", url, directory)

	// Assuming id_rsa is available at ~/.ssh/id_rsa
	sshAuth, err := ssh.NewPublicKeysFromFile("git", currentUser.HomeDir+"/.ssh/id_rsa", "admin")
	if err != nil {
		logs.Error(err)
	}

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
		//RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:     sshAuth,
		Progress: os.Stdout,
	})

	if err != nil {
		logs.Error(err)
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		logs.Error(err)
	}
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		logs.Error(err)
	}

	fmt.Println(commit)
}
