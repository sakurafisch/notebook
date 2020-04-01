# docker

```go
package main

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

//func main() {
//	err := InitClient()
//	if err != nil {
//		log.Fatal(err)
//	}
//	ch1 := make(chan string)
//	ch2 := make(chan string)
//	ch3 := make(chan string)
//
//	go pull(ch1, "alpine")
//	go pull(ch2, "alpine")
//	go pull(ch3, "alpine")
//	go pull(ch3, "postgres:11-alpine")
//
//
//	fmt.Println(<-ch1)
//	fmt.Println(<-ch2)
//	fmt.Println(<-ch3)
//
//}

func main() {
	err := InitClient()
	if err != nil {
		log.Fatal(err)
	}
	Check()
}

func pull(ch1 chan string, name string) {
	err := PullImage(name)
	if err != nil {
		log.Fatal(err)
	}
	ch1 <- "pull ok" + name
}

func Check() {
	config := &container.Config{
		Image: "ng:v1",
	}
	hostConfig := &container.HostConfig{}

	normalOut, errOut, err := RunByDocker(config, hostConfig, "test")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("测试结果", string(normalOut), string(errOut))
	}
}

var Cli *client.Client

func InitClient() (err error) {
	// 得到某个特定版本的客户端,并远程连接端口
	Cli, err = client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

//传入创建及启动容器的相关配置，返回容器输出结果
func RunByDocker(config *container.Config, hostConfig *container.HostConfig, containername string) (normalOut, errOut []byte, err error) {
	if config == nil {
		return nil, nil, errors.New("config参数不能为空")
	}
	if hostConfig == nil {
		return nil, nil, errors.New("hostConfig参数不能为空")
	}
	if containername == "" {
		err = fmt.Errorf("Call RunByDocker with a empty containername")
		logs.Error(err)
		return nil, nil, err
	}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logs.Error(err)
		return nil, nil, err
	}
	defer cli.Close()
	cli.NegotiateAPIVersion(ctx)
	logs.Info(containername)
	//创建容器
	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, containername)
	if err != nil {
		logs.Error(err)
		return nil, nil, err
	}
	//启动容器
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		logs.Error(err)
		return nil, nil, err
	}
	timeoutInf := make(chan string, 1)
	sigkill := make(chan bool, 1)
	go MonitorAndKill(resp.ID, 25*time.Second, sigkill, timeoutInf)
	//等待运行
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			logs.Error(err)
			return nil, nil, err
		}
	case <-statusCh:
	}
	sigkill <- false
	tmp := <-timeoutInf
	if tmp != "" {
		normalOut = []byte(tmp)
		return normalOut, nil, nil
	}
	//读取docker正常输出
	ShowStdout, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: false})
	if err != nil {
		logs.Error(err)
		return nil, nil, err
	}
	defer ShowStdout.Close()

	normalOut, err = ioutil.ReadAll(ShowStdout)
	if err != nil {
		logs.Error(err)
		return nil, nil, err
	}
	//读取docker异常输出
	ShowStderr, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: false, ShowStderr: true})
	if err != nil {
		logs.Error(err)
		return nil, nil, err
	}
	defer ShowStderr.Close()

	errOut, err = ioutil.ReadAll(ShowStderr)
	if err != nil {
		logs.Error(err)
		return nil, nil, err
	}
	err = StopContainer(resp.ID)
	if err != nil {
		logs.Error(err)
		return normalOut, errOut, err
	}
	err = RemoveContainer(resp.ID, cli)
	if err != nil {
		logs.Error(err)
		return normalOut, errOut, err
	}
	return normalOut, errOut, err
}

//协程处理docker内应用程序超时问题
func MonitorAndKill(containerID string, limit time.Duration, sigkill <-chan bool, timeoutInf chan<- string) {
	var err error
	if containerID == "" {
		err = fmt.Errorf("Call MonitorAndKill with empty containerID")
		logs.Error(err)
		return
	}
	if sigkill == nil {
		err := fmt.Errorf("sigkill cannot be empty")
		logs.Error(err)
		return
	}
	if timeoutInf == nil {
		err := fmt.Errorf("timeoutInf cannot be empty")
		logs.Error(err)
		return
	}
	time.Sleep(limit)
	select {
	case <-sigkill:
		timeoutInf <- ""
		return
	default:
		timeoutInf <- "程序运行超时,请检查代码是否有误或者重试"
		err = StopContainer(containerID)
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Info(fmt.Errorf("%s stop successful", containerID))
		err := Cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Info(fmt.Errorf("%s kill successful", containerID))
	}
}

//创建Docker连接客户端
func NewDockerClient() (*client.Client, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	cli.NegotiateAPIVersion(ctx)

	return cli, nil
}

//利用dockerfile创建镜像
func ImageBuild(dir string, ImageName string) {
	cli, err := NewDockerClient()
	if err != nil {
		panic(nil)
	}

	//将dockerfile文件压缩
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	dockerFile := "myDockerfile"
	dockerFileReader, err := os.Open(dir)
	if err != nil {
		log.Fatal(err, " :unable to open Dockerfile")
	}

	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		log.Fatal(err, " :unable to read dockerfile")
	}

	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Fatal(err, " :unable to write tar header")
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Fatal(err, " :unable to write tar body")
	}

	//利用dockerfile创建镜像
	dockerFileTarReader := bytes.NewReader(buf.Bytes())
	imageBuildResponse, err := cli.ImageBuild(
		context.Background(),
		dockerFileTarReader,
		types.ImageBuildOptions{
			Tags:       []string{ImageName},
			Context:    dockerFileTarReader,
			Dockerfile: dockerFile,
			Remove:     true})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
}

//拉取镜像
func PullImage(imageName string) (err error) {
	if imageName == "" {
		err = fmt.Errorf("imageName is null")
		return err
	}
	out, err := Cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		logs.Error(err)
		return err
	}
	defer out.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, out)
	if err != nil {
		logs.Error(err)
		return err
	}
	logs.Info(buf.String())
	return nil
}

//列出全部镜像
func ListImages() ([]types.ImageSummary, error) {
	cli, err := NewDockerClient()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return images, nil
}

//列出全部容器
func ListContainers() ([]types.Container, error) {
	cli, err := NewDockerClient()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if cli == nil {
		return nil, fmt.Errorf("get docker cli failed ")
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	//for _, container := range containers {
	//	fmt.Println(container.ID, container.Names)
	//}
	return containers, nil
}

//创建容器并返回id
func CreateContainer(config *container.Config, hostConfig *container.HostConfig) (containerID string, err error) {
	if config == nil {
		return "", errors.New("config参数不能为空")
	}
	if hostConfig == nil {
		return "", errors.New("hostConfig参数不能为空")
	}

	cli, err := NewDockerClient()
	if err != nil {
		logs.Error(err)
		return "", err
	}
	if cli == nil {
		return "", fmt.Errorf("get docker cli failed ")
	}

	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, nil, "")
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return resp.ID, nil
}

// 启动容器
func StartContainer(containerID string) error {
	if containerID == "" {
		return errors.New("containerID不能为空")
	}
	cli, err := NewDockerClient()
	if err != nil {
		logs.Error(err)
		return err
	}
	if cli == nil {
		return fmt.Errorf("get docker cli failed ")
	}
	return cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
}

// 停止容器
func StopContainer(containerID string) error {
	if containerID == "" {
		return errors.New("containerID不能为空")
	}
	cli, err := NewDockerClient()
	if err != nil {
		logs.Error(err)
		return err
	}
	if cli == nil {
		return fmt.Errorf("get docker cli failed ")
	}
	timeout := time.Second * 10
	return cli.ContainerStop(context.Background(), containerID, &timeout)
}

//删除容器
func RemoveContainer(containerID string, cli *client.Client) error {
	if containerID == "" {
		return errors.New("containerID不能为空")
	}
	if cli == nil {
		return fmt.Errorf("docker cli cant not nil")
	}
	return cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{})
}

//获取容器输出日志
func GetLogs(containerID string) (string, error) {
	if containerID == "" {
		return "", errors.New("containerID不能为空")
	}
	cli, err := NewDockerClient()
	if err != nil {
		logs.Error(err)
		return "", err
	}
	if cli == nil {
		return "", fmt.Errorf("get docker cli failed ")
	}
	//读取docker正常输出
	normalOut, err := cli.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: false})
	if err != nil {
		logs.Error(err)
		return "", err
	}
	defer normalOut.Close()

	stdout, err := ioutil.ReadAll(normalOut)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	output := string(stdout)

	//读取docker异常输出
	errorOut, err := cli.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{ShowStdout: false, ShowStderr: true})
	if err != nil {
		logs.Error(err)
		return "", err
	}
	defer errorOut.Close()

	stdout, err = ioutil.ReadAll(errorOut)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	output = output + string(stdout)

	return output, nil
}
```
