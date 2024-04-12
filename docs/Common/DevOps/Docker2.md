## 一、Docker介绍

[原文链接](https://www.lixian.fun/3812.html)

### 1.下载Dcoker依的赖环境

```
想安装Docker，需要先将依赖的环境全部下载下来，就像Maven依赖JDK一样yum -y install yum-utils device-mapper-persistent-data lvm2
```

### 2.指定Docker镜像源

```
默认下载Docker会去国外服务器下载，速度较慢，可以设置为阿里云镜像源，速度更快yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```

### 3.安装Docker

```bash
yum makecache fastyum -y install docker-ce
```

### 4.启动Docker并测试

```
安装成功后，需要手动启动，设置为开机启动，并测试一下 Docker#启动docker服务systemctl start docker#设置开机自动启动systemctl enable docker#测试docker run hello-world
```

## 二、Docker的中央仓库

```
1.Docker官方的中央仓库：这个仓库是镜像最全的，但是下载速度较慢。https://hub.docker.com/2.国内的镜像网站：网易蜂巢，daoCloud等，下载速度快，但是镜像相对不全。https://c.163yun.com/hub#/home http://hub.daocloud.io/ （推荐使用）3.在公司内部会采用私服的方式拉取镜像（添加配置）#需要创建 /etc/docker/daemon.json，并添加如下内容{	"registry-mirrors":["https://registry.docker-cn.com"],	"insecure-registries":["ip:port"]}#重启两个服务systemctl daemon-reloadsystemctl restart docker
```

## 三、镜像的操作

### 1.拉取镜像

```
从中央仓库拉取镜像到本地docker pull 镜像名称[:tag]#举个栗子:docker pull daocloud.io/library/tomcat:8.5.15-jre8
```

### 2.查看本地全部镜像

```
查看本地已经安装过的镜像信息，包含标识，名称，版本，更新时间，大小docker images
```

### 3.删除本地镜像

```
镜像会占用磁盘空间，可以直接手动删除，标识通过查看获取docker rmi 镜像的标识
```

### 4.镜像的导入导出

```
如果因为网络原因可以通过硬盘的方式传输镜像，虽然不规范，但是有效，但是这种方式导出的镜像名称和版本都是null，需要手动修改#将本地的镜像导出docker save -o 导出的路径 镜像id#加载本地的镜像文件docker load -i 镜像文件#修改镜像文件docker tag 镜像id 新镜像名称：版本
```

## 四、容器的操作

### 1.运行容器

```
运行容器需要定制具体镜像，如果镜像不存在，会直接下载#简单操作docker run 镜像的标识|镜像的名称[:tag]#常用的参数docker run -d -p 宿主机端口:容器端口 --name 容器名称 镜像的标识|镜像名称[:tag]#-d:代表后台运行容器#-p 宿主机端口:容器端口：为了映射当前Linux的端口和容器的端口#--name 容器名称:指定容器的名称
```

### 2.查看正在运行的容器

```
查看全部正在运行的容器信息docker ps [-qa]#-a 查看全部的容器，包括没有运行#-q 只查看容器的标识
```

### 3.查看容器日志

```
查看容器日志，以查看容器运行的信息docker logs -f 容器id#-f：可以滚动查看日志的最后几行
```

### 4.进入容器的内部

```
可以进入容器的内部进行操作docker exec -it 容器id bash
```

### 5.复制内容到容器

```
将宿主机的文件复制到容器内部的指定目录docker cp 文件名称 容器id:容器内部路径
```

### 6.重启&启动&停止&删除容器

```
容器的启动，停止，删除等操作，后续会经常使用到#重新启动容器docker restart 容器id#启动停止运行的容器docker start 容器id #停止指定的容器(删除容器前，需要先停止容器)docker stop 容器id#停止全部容器docker stop $(docker ps -qa)#删除指定容器docker rm 容器id#删除全部容器docker rm $(docker ps -qa)
```

## 五、Docker应用

### 1.docker安装tomcat

```
运行Tomcat容器，为部署ssm工程做准备 docker run -d -p 8080:8080 --name tomcat  daocloud.io/library/tomcat:8.5.15-jre8#或者已经下载了tomcat镜像docker run -d -p 8080:8080 --name tomcat 镜像的标识
```

### 2.运行MySQL容器

```
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root daocloud.io/library/mysql:5.7.4
```

### 3.部署ssm工程

```
修改SSM工程环境，设置为Linux中Docker容器的信息通过Maven的package重新打成war包将Windows下的war包复制到Linux中通过docker命令将宿主机的war包复制到容器内部 docker cp 文件名称 容器id:容器内部路径 测试访问SSM工程
```

### 六、数据卷


为了部署SSM的工程，需要使用到cp的命令将宿主机内的ssm.war文件复制到容器内部。



数据卷：将宿主机的一个目录映射到容器的一个目录中。

可以在宿主机中操作目录中的内容，那么容器内部映射的文件，也会跟着一起改变。

### 1.创建数据卷

```
#创建数据卷后，默认会存放在一个目录下/var/lib/docker/volumes/数据卷名称/_datadocker volume create 数据卷名称
```

### 2.查看全部数据卷

```
#查看全部数据卷信息docker volume ls
```

### 3.查看数据卷详情

```
#查看数据卷的详细信息，可以查询到存放的路径，创建时间等等docker volume inspect 数据卷名称
```

### 4.删除数据卷

```
#删除指定的数据卷docker volume rm 数据卷名称
```

### 5.容器映射数据卷

```
#通过数据卷名称映射，如果数据卷不存在。Docker会帮你自动创建，会将容器内部自带的文件，存储在默认的存放路径中。docker run -d -p 8080:8080 --name tomcat -v 数据卷名称:容器内部的路径 镜像id #通过路径映射数据卷，直接指定一个路径作为数据卷的存放位置。但是这个路径下是空的。docker run -d -p 8080:8080 --name tomcat -v 路径(/root/自己创建的文件夹):容器内部的路径 镜像id
```

## 七、Dockerfile自定义镜像

### 1.Dockerfile

```
创建自定义镜像就需要创建一个Dockerfiler,如下为Dockerfile的语言 from：指定当前自定义镜像依赖的环境copy：将相对路径下的内容复制到自定义镜像中workdir：声明镜像的默认工作目录run：执行的命令，可以编写多个cmd：需要执行的命令（在workdir下执行的，cmd可以写多个，只以最后一个为准） #示例：from daocloud.io/library/tomcat:8.5.15-jre8copy ssm.war /usr/local/tomcat/webapps
```

### 2.通过Dockerfile制作镜像

```
#编写完Dockerfile后需要通过命令将其制作为镜像，并且要在Dockerfile的当前目录下，之后即可在镜像中查看到指定的镜像信息，注意最后的 .docker build -t 镜像名称[:tag] ./
```

## 八、Docker-Compose

### 1.下载并安装Docker-Compose

### 1.1下载Docker-Compose

```
#去github官网搜索docker-compose，下载1.24.1版本的Docker-Compose下载路径：https://github.com/docker/compose/releases/download/1.24.1/docker-compose-Linux-x86_64
```

### 1.2设置权限

```
#需要将DockerCompose文件的名称修改一下，给予DockerCompose文件一个可执行的权限mv docker-compose-Linux-x86_64 docker-composechmod 777 docker-compose
```

### 1.3配置环境变量

```
#方便后期操作，配置一个环境变量#将docker-compose文件移动到了/usr/local/bin，修改了/etc/profile文件，给/usr/local/bin配置到了PATH中 mv docker-compose /usr/local/binvi /etc/profile#添加内容：export PATH=$JAVA_HOME:/usr/local/bin:$PATHsource /etc/profile
```

### 1.4测试

```
在任意目录下输入docker-compose
```

### 2.Docker-Compose管理MySQL和Tomcat容器

```
yml文件以key:value方式来指定配置信息多个配置信息以换行+缩进的方式来区分在docker-compose.yml文件中，不要使用制表符 version: '3.1'services:  mysql:           # 服务的名称    restart: always   # 代表只要docker启动，那么这个容器就跟着一起启动    image: daocloud.io/library/mysql:5.7.4  # 指定镜像路径    container_name: mysql  # 指定容器名称    ports:      - 3306:3306   #  指定端口号的映射    environment:      MYSQL_ROOT_PASSWORD: root   # 指定MySQL的ROOT用户登录密码      TZ: Asia/Shanghai        # 指定时区    volumes:     - /opt/docker_mysql_tomcat/mysql_data:/var/lib/mysql   # 映射数据卷  tomcat:    restart: always    image: daocloud.io/library/tomcat:8.5.15-jre8    container_name: tomcat    ports:      - 8080:8080    environment:      TZ: Asia/Shanghai    volumes:      - /opt/docker_mysql_tomcat/tomcat_webapps:/usr/local/tomcat/webapps      - /opt/docker_mysql_tomcat/tomcat_logs:/usr/local/tomcat/logs
```

### 3.使用docker-compose命令管理容器

```
在使用docker-compose的命令时，默认会在当前目录下找docker-compose.yml文件 #1.基于docker-compose.yml启动管理的容器docker-compose up -d #2.关闭并删除容器docker-compose down #3.开启|关闭|重启已经存在的由docker-compose维护的容器docker-compose start|stop|restart #4.查看由docker-compose管理的容器docker-compose ps #5.查看日志docker-compose logs -f
```

### 4.docker-compose配合Dockerfile使用


使用docker-compose.yml文件以及Dockerfile文件在生成自定义镜像的同时启动当前镜像，并且由docker-compose去管理容器

### 4.1docker-compose文件

```
编写docker-compose文件 # yml文件version: '3.1'services:  ssm:    restart: always    build:            # 构建自定义镜像      context: ../      # 指定dockerfile文件的所在路径      dockerfile: Dockerfile   # 指定Dockerfile文件名称    image: ssm:1.0.1    container_name: ssm    ports:      - 8081:8080    environment:      TZ: Asia/Shanghai
```

### 4.2 Dockerfile文件

```
编写Dockerfile文件 from daocloud.io/library/tomcat:8.5.15-jre8copy ssm.war /usr/local/tomcat/webapps
```

### 4.3 运行

```
#可以直接基于docker-compose.yml以及Dockerfile文件构建的自定义镜像docker-compose up -d# 如果自定义镜像不存在，会帮助我们构建出自定义镜像，如果自定义镜像已经存在，会直接运行这个自定义镜像#重新构建自定义镜像docker-compose build#运行当前内容，并重新构建docker-compose up -d --build
```

## 九、CI、CD介绍及准备

### 1.CI、CD引言


项目部署
1.将项目通过maven进行编译打包
2.将文件上传到指定的服务器中
3.将war包放到tomcat的目录中
4.通过Dockerfile将Tomcat和war包转成一个镜像，由DockerCompose去运行容器
项目更新后，需要将上述流程再次的从头到尾的执行一次，如果每次更新一次都执行一次上述操作，很费时，费力。我们就可以通过CI、CD帮助我们实现持续集成，持续交付和部署

### 2.CI介绍


CI（continuous intergration）持续集成
持续集成：编写代码时，完成了一个功能后，立即提交代码到Git仓库中，将项目重新的构建并且测试。



1.快速发现错误。
2.防止代码偏离主分支。

### 3.搭建Gitlab服务器

### 3.1.准备环境


实现CI，需要使用到Gitlab远程仓库，先通过Docker搭建Gitlab
创建一个全新的虚拟机，并且至少指定4G的运行内存，4G运行内存是Gitlab推荐的内存大小。
并且安装Docker以及Docker-Compose

### 3.2 修改ssh的22端口

```
#将ssh的默认22端口，修改为60022端口，因为Gitlab需要占用22端口 vi /etc/ssh/sshd_config  PORT 22 -> 60022systemctl restart sshd
```

### 3.3 编写docker-compose.yml

```
docker-compose.yml文件去安装gitlab（下载和运行的时间比较长的） version: '3.1'services: gitlab:  image: 'twang2218/gitlab-ce-zh:11.1.4'  container_name: "gitlab"  restart: always  privileged: true  hostname: 'gitlab'  environment:   TZ: 'Asia/Shanghai'   GITLAB_OMNIBUS_CONFIG: |    external_url 'http://192.168.199.110'    gitlab_rails['time_zone'] = 'Asia/Shanghai'    gitlab_rails['smtp_enable'] = true    gitlab_rails['gitlab_shell_ssh_port'] = 22  ports:   - '80:80'   - '443:443'   - '22:22'  volumes:   - /opt/docker_gitlab/config:/etc/gitlab   - /opt/docker_gitlab/data:/var/opt/gitlab   - /opt/docker_gitlab/logs:/var/log/gitlab
```

## 十、搭建GitlabRunner

### 1.准备文件

```
daemon.json {“registry-mirrors”: [“https://registry.docker-cn.com”],“insecure-registries”: [ip:ports]} 文件夹 environment里面准备maven安装包，jdk1.8安装包，Dockerfile，daemon.json以及docker-compose
```

### 2.开始搭建


创建工作目录 /usr/local/docker_gitlab-runner
将docker-compose.yml文件以及environment目录全部复制到上述目录中
在宿主机启动docker程序后先执行 sudo chown root:root /var/run/docker.sock (如果重启过 docker,重新执行)
在/usr/local/docker_gitlab-runner 目录中执行docker-compose up -d –build 启动容器
添加容器权限，保证容器可以使用宿主机的dockerdocker exec -it gitlab-runner usermod -aG root gitlab-runner
注册Runner信息到gitlab

### 3.进入后续步骤

```
docker exec -it gitlab-runner gitlab-runner register # 输入 GitLab 地址Please enter the gitlab-ci coordinator URL (e.g. https://gitlab.com/):http://192.168.199.109/ # 输入 GitLab TokenPlease enter the gitlab-ci token for this runner:1Lxq_f1NRfCfeNbE5WRh # 输入 Runner 的说明Please enter the gitlab-ci description for this runner:可以为空 # 设置 Tag，可以用于指定在构建规定的 tag 时触发 ciPlease enter the gitlab-ci tags for this runner (comma separated):deploy # 这里选择 true ，可以用于代码上传后直接执行（根据版本，也会没有次选项）Whether to run untagged builds [true/false]:true # 这里选择 false，可以直接回车，默认为 false（根据版本，也会没有次选项）Whether to lock Runner to current project [true/false]:false # 选择 runner 执行器，这里我们选择的是 shellPlease enter the executor: virtualbox, docker+machine, parallels, shell, ssh, docker-ssh+machine, kubernetes, docker, docker-ssh:shell
```

## 十一、整合项目入门测试

### 1.创建项目


创建maven工程，添加web.xml文件，编写HTML页面

### 2.编写.gitlab-ci.yml文件

```
stages:  - test test:  stage: test  script:    - echo first test ci   # 输入的命令
```

### 3.将maven工程推送到gitlab中

```
执行git命令推送到Gitlab git push origin master
```

### 4.查看效果


可以在gitlab中查看到gitlab-ci.yml编写的内容

## 十二、完善项目配置


添加Dockerfile以及docker-compose.yml， 并修改.gitlab-ci.yml文件

### 1.创建Dockerfile

```
# DockerfileFROM daocloud.io/library/tomcat:8.5.15-jre8COPY testci.war /usr/local/tomcat/webapps
```

### 2.创建docker-compose.yml

```
# docker-compose.ymlversion: "3.1"services:  testci:    build: docker    restart: always    container_name: testci    ports:      - 8080:8080
```

### 3.修改.gitlab-ci.yml

```
# ci.ymlstages:  - test test:  stage: test  script:    - echo first test ci    - /usr/local/maven/apache-maven-3.6.3/bin/mvn package    - cp target/testci-1.0-SNAPSHOT.war docker/testci.war    - docker-compose down    - docker-compose up -d --build    - docker rmi $(docker images -qf dangling=true)
```