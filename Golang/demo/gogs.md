# gogs

```go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"

	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/sdk/gitea"
	"github.com/astaxie/beego/logs"
)

// RandString 生成随机字符串
func RandString(l int) string {
	if l < 1 {
		return ""
	}
	t := time.Now().Nanosecond()
	h := md5.New()
	io.WriteString(h, "HelloWorld")
	io.WriteString(h, strconv.Itoa(t))
	hash := fmt.Sprintf("%x", h.Sum(nil))

	str := hash
	bytes := []byte(str)
	var result []byte
	for i := 0; i < l; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func main() {
	//NewGitServiceClient()
	user := "test1"
	email := "1222@qq.com"
	client, err := DefaultGitService()
	if err != nil {
		logs.Error(err)
		return
	}
	uid, err := client.CreateUser(user, "12345678", "test", email)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info(uid)

	pid, err := client.ImportProject(uid, "solutions")
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info(pid)

	hash, err := client.GetLatestCommitID(user, "solutions", "1000010004")
	if err != nil {
		logs.Error(err)
		return
	}

	logs.Info(hash)
	key := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCyCBahDTXU6MM4/wxHY3lxllxFBizNbRkI6ghs8nsFbFdc8bq5x8HZTB+/mqUmHKG/v6cBaEFjFWR3BFXcg/LqczphM0KWsL09Ayt891crHaKRoci0/OjCMMM8JL1X8yWRmeximPFk/8AfhAHvuDGWhL2T8RARzgwRsaZF8vWaHW39R3UY6Nf9MOgG2AiwOX1SpCNfnvxP7l8LkUloyjbhMl8LPkCl1jcc2eY/Bht+3/XyTg27zrk2PgrQ35DeaFgZHeV21x97ITBOofWAsU4UIT7EeMW/WWaGH5zCvYZ7RMW+f1srgCwqxWaXulhPNmhudWH6B6zFtVF1eaVD/CrtSWjj/Xkd5KMR1YNZ0v28aPe+fdIWdi4hwbGMNkhLmceWWfKOkhspORrdCgSrCQg5zBHCGpnUrlPk9N4LpQvDUdJlZuw0m1akChu/GAFaB1sdcY/Dqgy3+MSkRZu8l1ZkNMl+MFac7Vz4+ECtM6eU3DGUMkUVFsgYNH1rq4eARGyV8u5kA9/flFb43gB5cdhh8NHb3lBTLE7bi0LoqBzsoNpgpm4N/9jl0jBrPzgSQiLF2xHsOtCj8sWFx4LhIDYKKaUoOBUy5qtRr3C72Gr500mToECkJszXNZpSLeb+DflsTRoxIu4q9duX9ZQZdjNufIlxxY6u+vnQTpBzkLOefw== email@example.com"
	_, err = client.AddSshKey(user, "aaa", key)
	if err != nil {
		logs.Error(err)
		return
	}

}

type GitService struct {
	User        string        `json:"username"` //管理员账号
	Password    string        `json:"password"` //管理员密码
	AccessToken string        `json:"access_token"`
	BaseUrl     string        `json:"base_url"`   //gitlab 主机地址
	ImportURL   string        `json:"import_url"` //用于初始化仓库的仓库地址
	Client      *gitea.Client `json:"git"`
}

var gitServiceCli *GitService

//根据配置文件创建的默认git服务操作对象
func DefaultGitService() (g *GitService, err error) {
	if gitServiceCli != nil {
		return gitServiceCli, nil
	}
	//if Conf == nil {
	//	err = fmt.Errorf("configure was not init")
	//	logs.Error(err)
	//	return
	//}
	//c := Conf.GitService
	//g, err = NewGitServiceClient(c.User, c.Password, c.BaseUrl, c.ImportUrl)
	g, err = NewGitServiceClient("root", "ards_pwd",
		"http://localhost:10080/", "http://cst.gzhu.edu.cn:10080/root/solutions-init.git")

	if err != nil {
		logs.Error(err)
		return
	}
	if g == nil {
		err = fmt.Errorf("init gitlab client failed")
		logs.Error(err)
		return
	}
	gitServiceCli = g
	return
}

// 初始化服务客户端
func NewGitServiceClient(username, password, baseUrl, importUrl string) (g *GitService, err error) {
	if username == "" || password == "" || baseUrl == "" || importUrl == "" {
		err = fmt.Errorf("illegal arguments of init git service client ")
		logs.Error(err)
		return
	}
	g = &GitService{
		User:      username,
		Password:  password,
		BaseUrl:   baseUrl,
		ImportURL: importUrl,
	}
	// 每次启动后端重新生成 access token
	client := gitea.NewClient(g.BaseUrl, "")
	token, err := client.CreateAccessToken(g.User, g.Password, structs.CreateAccessTokenOption{
		Name: RandString(10),
	})
	if err != nil {
		err = fmt.Errorf("create access token for user %s failed: %s", g.User, err.Error())
		logs.Error(err)
		return nil, err
	}
	if token == nil {
		err = fmt.Errorf("create access token for user %s failed", g.User)
		logs.Error(err)
		return nil, err
	}

	// 用access token重新初始化客户端
	g.AccessToken = token.Token
	client = gitea.NewClient(g.BaseUrl, token.Token)
	g.Client = client

	return g, nil
}

//创建用户返回gitlab id
func (g *GitService) CreateUser(username, password, nickname, email string) (userId int, err error) {
	if username == "" || password == "" || nickname == "" || email == "" {
		err = fmt.Errorf("illegal arguments of create gitlab user ")
		logs.Error(err)
		return
	}
	if g.Client == nil {
		err = fmt.Errorf("git service client was not inited")
		logs.Error(err)
		return
	}
	user, err := g.Client.AdminCreateUser(structs.CreateUserOption{
		//SourceID:   0,
		//LoginName:  username,
		Username: username,
		FullName: nickname,
		Email:    email,
		Password: password,
		//SendNotify: false,
	})
	if err != nil {
		logs.Error(err)
		return
	}
	if user == nil {
		err = fmt.Errorf("create gitlab user failed ")
		logs.Error(err)
		return
	}
	userId = int(user.ID)
	return
}

func (g *GitService) DeleteUser(username string) (err error) {
	if username == "" {
		err = fmt.Errorf("illegal username for get delete user ")
		logs.Error(err)
		return
	}

	if g.Client == nil {
		err = fmt.Errorf("gitlab client was not inited")
		logs.Error(err)
		return
	}
	err = g.Client.AdminDeleteUser(username)
	if err != nil {
		logs.Error(err)
		return
	}
	return nil
}

// 导入仓库方式初始化用户仓库
func (g *GitService) ImportProject(userId int, projectName string) (projectId int, err error) {
	if userId < 0 || projectName == "" {
		err = fmt.Errorf("illgal argument for crteate git project ")
		logs.Error(err)
		return
	}
	if g.Client == nil {
		err = fmt.Errorf("gitlab client was not inited")
		logs.Error(err)
		return
	}
	repository, err := g.Client.MigrateRepo(structs.MigrateRepoOption{
		CloneAddr: g.ImportURL,
		//AuthUsername: "",
		//AuthPassword: "",
		UID:      userId,
		RepoName: projectName,
		//Mirror:       false,
		Private: true,
		//Description:  "",
	})
	if err != nil {
		logs.Error(err)
		return
	}
	if repository == nil {
		err = fmt.Errorf("create project for user failed")
		logs.Error(err)
		return
	}
	projectId = int(repository.ID)
	return
}

//获取仓库对应分支的最新commit id
func (g *GitService) GetLatestCommitID(username, projectName, branch string) (commitID string, err error) {
	if username == "" || projectName == "" || branch == "" {
		err = fmt.Errorf("illgal project id for get commit id ")
		logs.Error(err)
		return
	}
	if g.Client == nil {
		err = fmt.Errorf("git service client was not inited")
		logs.Error(err)
		return
	}
	b, err := g.Client.GetRepoBranch(username, projectName, branch)
	if err != nil {
		logs.Error(err)
		return
	}
	if b == nil {
		err = fmt.Errorf("add ssh key for user %s failed", username)
		logs.Error(err)
		return
	}
	return b.Commit.ID, nil
}

func (g *GitService) AddSshKey(username, title, sshKey string) (keyID int, err error) {
	if username == "" || title == "" || sshKey == "" {
		err = fmt.Errorf("illgal argument for add ssh key ")
		logs.Error(err)
		return
	}
	if g.Client == nil {
		err = fmt.Errorf("gitlab client was not inited")
		logs.Error(err)
		return
	}

	key, err := g.Client.AdminCreateUserPublicKey(username, structs.CreateKeyOption{
		Title: title,
		Key:   sshKey,
	})
	if err != nil {
		logs.Error(err)
		return
	}
	if key == nil {
		err = fmt.Errorf("add ssh key for user %s failed", username)
		logs.Error(err)
		return
	}
	keyID = int(key.ID)
	return
}
```
