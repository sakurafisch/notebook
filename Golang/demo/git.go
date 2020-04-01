package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"strings"
)

func main() {

	uid := "off364i2R67h6AgSnvKthIiNqwrs"
	branch := "1000010022"

	gitLabBaseUrl := "http://cst.gzhu.edu.cn:9980/"
	userPath := uid + "/solutions.git"

	url := gitLabBaseUrl + userPath //仓库地址
	localPath := "/tmp/1"
	username := "root"
	password := "cst17@admin"

	g, err := NewGitClient(url, localPath, branch, username, password)
	if err != nil {
		logs.Error(err)
	}

	//err = g.Clone()
	//if err != nil {
	//	logs.Error(err)
	//}

	err = g.Checkout()
	if err != nil {
		logs.Error(err)
	}

	err = g.Pull()
	if err != nil {
		logs.Error(err)
	}
}

type Git struct {
	Username   string
	Password   string
	Url        string          //远程仓库地址
	Branch     string          //分支
	LocalPath  string          //本地仓库文件夹地址
	Repository *git.Repository //初始化仓库对象
}

func Te() (err error) {
   url :=`ssh://localhost:10022/tmp/test/aaa`

	_, err = git.PlainClone("/tmp/bb", false, &git.CloneOptions{
		URL:  url,
		Auth: &http.BasicAuth{"root", "admin"},
		//ReferenceName: plumbing.NewBranchReferenceName(g.Branch),
	})
	if err != nil {
		logs.Error(err)
		return err
	}
	return
}

func NewGitClient(url, localPath, branch, username, password string) (*Git, error) {
	if url == "" || localPath == "" {
		err := fmt.Errorf("url  or localPath should not be blank")
		logs.Error(err)
		return nil, err
	}
	if branch == "" {
		branch = "master"
	}
	g := &Git{
		Username:  username,
		Password:  password,
		Url:       url,
		Branch:    branch,
		LocalPath: localPath,
	}
	return g, nil
}

func (g *Git) Clone() (err error) {
	if g.Url == "" || g.LocalPath == "" {
		err := fmt.Errorf("url or localPath should not be blank")
		logs.Error(err)
		return err
	}
	if g.Branch == "" {
		g.Branch = "master"
	}
	var auth = &http.BasicAuth{}
	if g.Username != "" || g.Password != "" {
		auth = &http.BasicAuth{g.Username, g.Password}
	}

	_, err = git.PlainClone(g.LocalPath, false, &git.CloneOptions{
		URL:           g.Url,
		Auth:          auth,
		ReferenceName: plumbing.NewBranchReferenceName(g.Branch),
	})
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (g *Git) Pull() (err error) {
	if g.Url == "" || g.LocalPath == "" {
		err := fmt.Errorf("url or localPath should not be blank")
		logs.Error(err)
		return err
	}
	if g.Branch == "" {
		g.Branch = "master"
	}
	var auth = &http.BasicAuth{}
	if g.Username != "" || g.Password != "" {
		auth = &http.BasicAuth{g.Username, g.Password}
	}

	r, err := git.PlainOpen(g.LocalPath)
	if err != nil {
		logs.Error(err)
		return err
	}
	//记录仓库信息，方便使用
	g.Repository = r
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		logs.Error(err)
		return
	}

	err = w.Pull(&git.PullOptions{
		Auth:          auth,
		RemoteName:    "origin",
		ReferenceName: plumbing.NewBranchReferenceName(g.Branch),
		Force:         true,
	})
	if err != nil {
		//忽略already up-to-date的错误
		if !strings.Contains(err.Error(), "already up-to-date") && !strings.Contains(err.Error(), "non-fast-forward update") {
			logs.Error(err)
			return
		}
		logs.Debug(err)
		return nil
	}
	return
}

func (g *Git) Checkout0() (err error) {
	if g.LocalPath == "" {
		err := fmt.Errorf("url or localPath should not be blank")
		logs.Error(err)
		return err
	}
	if g.Branch == "" {
		g.Branch = "master"
	}
	r, err := git.PlainOpen(g.LocalPath)
	if err != nil {
		logs.Error(err)
		return err
	}
	g.Repository = r
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		logs.Error(err)
		return
	}

	//var auth = &http.BasicAuth{}
	//if g.Username != "" || g.Password != "" {
	//	auth = &http.BasicAuth{g.Username, g.Password}
	//}
	//更新仓库信息
	//err = r.Fetch(&git.FetchOptions{
	//	Auth:       auth,
	//	RemoteName: "origin",
	//	Force:      true,
	//})
	//if err != nil {
	//	//忽略already up-to-date的错误
	//	if !strings.Contains(err.Error(), "already up-to-date") {
	//		logs.Error(err)
	//		return
	//	}
	//	logs.Debug(err)
	//	return nil
	//}

	//查看所有远程分支
	//refs, err := r.Storer.IterReferences()
	//if err != nil {
	//	return err
	//}
	//bs := storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
	//	return ref.Name().IsRemote()
	//}, refs)
	//err = bs.ForEach(func(b *plumbing.Reference) error {
	//	fmt.Println(b)
	//	return nil
	//})

	//判断远程分支是否存在，存在才创建本地分支
	s := r.Storer
	remoteRef, err := s.Reference(plumbing.ReferenceName("refs/remotes/origin/" + g.Branch))
	if err != nil {
		err = fmt.Errorf("找不到远程分支 %s:%s", g.Branch, err.Error())
		logs.Error(err)
		return
	}
	logs.Info("代码版本：", remoteRef.Hash())
	//判断本地分支是否存在，不存在才创建
	create := false
	_, err = s.Reference(plumbing.ReferenceName("refs/heads/" + g.Branch))
	if err != nil {
		create = true
	}

	err = w.Checkout(&git.CheckoutOptions{
		Create: create,
		Force:  true,
		//Hash:   remoteRef.Hash(), //切换到远程分支对应的最新hash，无需在本地创建分支
		Branch: plumbing.NewBranchReferenceName(g.Branch),
	})
	if err != nil {
		logs.Error(err)
		return
	}
	return
}

func (g *Git) Checkout() (err error) {
	if g.LocalPath == "" {
		err := fmt.Errorf("url or localPath should not be blank")
		logs.Error(err)
		return err
	}
	if g.Branch == "" {
		g.Branch = "master"
	}
	r, err := git.PlainOpen(g.LocalPath)
	if err != nil {
		logs.Error(err)
		return err
	}
	g.Repository = r
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		logs.Error(err)
		return
	}

	//var auth = &http.BasicAuth{}
	//if g.Username != "" || g.Password != "" {
	//	auth = &http.BasicAuth{g.Username, g.Password}
	//}
	//更新仓库信息
	//err = r.Fetch(&git.FetchOptions{
	//	Auth:       auth,
	//	RemoteName: "origin",
	//	Force:      true,
	//})
	//if err != nil {
	//	//忽略already up-to-date的错误
	//	if !strings.Contains(err.Error(), "already up-to-date") {
	//		logs.Error(err)
	//		return
	//	}
	//	logs.Debug(err)
	//	return nil
	//}

	//查看所有远程分支
	refs, err := r.Storer.IterReferences()
	if err != nil {
		return err
	}
	bs := storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
		return ref.Name().IsRemote()
	}, refs)
	err = bs.ForEach(func(b *plumbing.Reference) error {
		fmt.Println(b)
		return nil
	})

	//判断远程分支是否存在，存在才创建本地分支
	s := r.Storer
	remoteRef, err := s.Reference(plumbing.ReferenceName("refs/remotes/origin/" + g.Branch))
	if err != nil {
		err = fmt.Errorf("找不到远程分支 %s:%s", g.Branch, err.Error())
		logs.Error(err)
		return
	}

	//判断本地分支是否存在，不存在才创建
	//create := false
	locRef, err := s.Reference(plumbing.ReferenceName("refs/heads/" + g.Branch))
	if err != nil {
		logs.Error("bendi", err)
	}
	logs.Info(locRef)

	ref := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+g.Branch), remoteRef.Hash())
	err = r.Storer.SetReference(ref)
	//err = s.CheckAndSetReference(locRef, remoteRef)
	if err != nil {
		logs.Error(err)
	}

	err = s.SetReference(ref)
	if err != nil {
		logs.Error(err)
	}

	err = w.Checkout(&git.CheckoutOptions{
		//Create: true,
		Force: true,
		//Hash:  remoteRef.Hash(), //切换到远程分支对应的最新hash，无需在本地创建分支
		//Branch: ref.Name(),
		Branch: remoteRef.Name(),
	})
	err = w.Checkout(&git.CheckoutOptions{
		//Create: true,
		Force: true,
		//Hash:  remoteRef.Hash(), //切换到远程分支对应的最新hash，无需在本地创建分支
		//Branch: ref.Name(),
		Branch: remoteRef.Name(),
	})
	if err != nil {
		logs.Error(err)
		return
	}
	return
}
