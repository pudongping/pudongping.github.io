---
title: 如何给已经运行中的 docker 容器添加或者修改端口映射？
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Docker
tags:
  - Docker
  - Linux
abbrlink: 2a105f69
date: 2021-07-02 23:55:35
img:
coverImg:
password:
summary:
---

# 如何给已经运行中的 docker 容器添加或者修改端口映射？

## 容器还没有构建

如果你的容器还没有构建时，想添加端口映射时，你只需要在创建容器的时候添加 `-p` 参数，想添加几个端口映射就追加几个 `-p` 参数。类似于如下示例：

```shell

docker run --name api_dfo_hyperf_ws \
-v /Users/pudongping/glory/codes/dfo/api_dfo_hyperf:/api_dfo_hyperf \
-p 9502:9502 \
-p 9503:9503 \
-p 9504:9504 \
-p 9505:9505 -it \
--entrypoint /bin/sh \
alex/alex_api_dfo:v1.0

```

## 容器已经构建，但是想修改或者添加端口时

### 先停止掉正在运行的容器。

> 以下内容都是以容器 id 为 `cbe26510c276` 进行操作的，请务必将容器 id 换成你自己需要修改的容器 id。

```shell

docker stop {容器的名称或者 id }

# 比如：
docker stop cbe26510c276

```

### 查看容器完整的 `hash_of_the_container` 数值

```shell

docker inspect {容器的名称或者 id } | grep Id

# 比如：
docker inspect cbe26510c276 | grep Id
# 会得到如下结果：
# "Id": "cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00"

```

### 打开 `hostconfig.json` 配置文件

```shell

vim /var/lib/docker/containers/{hash_of_the_container}/hostconfig.json

# 比如：
vim /var/lib/docker/containers/cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00/hostconfig.json

```

> 如果你不想先查看完整的容器 `hash_of_the_container` 数值，你也可以直接先切换到当前容器相关目录中 `cd /var/lib/docker/containers/{hash_of_the_container}*` ， 然后再去编辑 `hostconfig.json` 配置文件。

```shell

cd /var/lib/docker/containers/{hash_of_the_container}*

# 比如：
cd /var/lib/docker/containers/cbe26510c276*
# 然后再去编辑 `hostconfig.json` 配置文件
vim hostconfig.json

```


### 修改 `hostconfig.json` 配置文件

在 `hostconfig.json` 配置文件中，找到 `"PortBindings":{}` 这个配置项，然后进行修改。我这里添加了两个端口映射，分别将宿主机的 `8502` 端口以及 `8505` 端口映射到容器的 `8502` 端口和 `8505` 端口。

`HostPort` 对应的端口代表 **宿主机** 的端口。

> 建议容器使用什么端口，宿主机就映射什么端口，方便以后管理。当然，具体情况，具体分析。

```json

{
    "Binds": [
        "/data/portal_api_dfo_hyperf/:/portal_api_dfo_hyperf"
    ],
    "ContainerIDFile": "",
    "LxcConf": [],
    "Memory": 0,
    "MemorySwap": 0,
    "CpuShares": 0,
    "CpusetCpus": "",
    "Privileged": false,
    "PortBindings": {
        "8502/tcp": [
            {
                "HostIp": "",
                "HostPort": "8502"
            }
        ],
        "8505/tcp": [
            {
                "HostIp": "",
                "HostPort": "8505"
            }
        ]
    },
    "Links": null,
    "PublishAllPorts": false,
    "Dns": null,
    "DnsSearch": null,
    "ExtraHosts": null,
    "VolumesFrom": null,
    "Devices": [],
    "NetworkMode": "bridge",
    "IpcMode": "",
    "PidMode": "",
    "CapAdd": null,
    "CapDrop": null,
    "RestartPolicy": {
        "Name": "no",
        "MaximumRetryCount": 0
    },
    "SecurityOpt": null,
    "ReadonlyRootfs": false,
    "Ulimits": null,
    "LogConfig": {
        "Type": "",
        "Config": null
    },
    "CgroupParent": ""
}

```

### 如果 `config.v2.json` 配置文件或者 `config.json` 配置文件中也记录了端口，也需要进行修改，如果没有，就不需要改。

> 只需要修改 `"ExposedPorts": {}` 相关之处。


```json

{
    "State": {
        "Running": false,
        "Paused": false,
        "Restarting": false,
        "OOMKilled": false,
        "Dead": false,
        "Pid": 0,
        "ExitCode": 137,
        "Error": "",
        "StartedAt": "2021-05-17T07:48:26.743090016Z",
        "FinishedAt": "2021-07-02T06:05:33.025441199Z"
    },
    "ID": "cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00",
    "Created": "2020-12-23T07:02:00.997803339Z",
    "Path": "/bin/sh",
    "Args": [],
    "Config": {
        "Hostname": "cbe26510c276",
        "Domainname": "",
        "User": "",
        "Memory": 0,
        "MemorySwap": 0,
        "CpuShares": 0,
        "Cpuset": "",
        "AttachStdin": true,
        "AttachStdout": true,
        "AttachStderr": true,
        "PortSpecs": null,
        "ExposedPorts": {
            "8502/tcp": {},
            "8505/tcp": {}
        },
        "Tty": true,
        "OpenStdin": true,
        "StdinOnce": true,
        "Env": [
            "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
            "SW_VERSION=v4.5.10",
            "COMPOSER_VERSION=2.0.8",
            "PHPIZE_DEPS=autoconf dpkg-dev dpkg file g++ gcc libc-dev make php7-dev php7-pear pkgconf re2c pcre-dev pcre2-dev zlib-dev libtool automake"
        ],
        "Cmd": null,
        "Image": "hyperf/hyperf:7.4-alpine-v3.11-swoole",
        "Volumes": null,
        "WorkingDir": "",
        "Entrypoint": [
            "/bin/sh"
        ],
        "NetworkDisabled": false,
        "MacAddress": "",
        "OnBuild": null,
        "Labels": {
            "license": "MIT",
            "maintainer": "Hyperf Developers <group@hyperf.io>",
            "version": "1.0"
        }
    },
    "Image": "f503e215646f50e5242a6bdf1c6a9a176910c99600b3472b390892f59c87c3a1",
    "NetworkSettings": {
        "IPAddress": "",
        "IPPrefixLen": 0,
        "MacAddress": "",
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "Gateway": "",
        "IPv6Gateway": "",
        "Bridge": "",
        "PortMapping": null,
        "Ports": null
    },
    "ResolvConfPath": "/var/lib/docker/containers/cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00/resolv.conf",
    "HostnamePath": "/var/lib/docker/containers/cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00/hostname",
    "HostsPath": "/var/lib/docker/containers/cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00/hosts",
    "LogPath": "/var/lib/docker/containers/cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00/cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00-json.log",
    "Name": "/portal_api_dfo_hyperf",
    "Driver": "aufs",
    "ExecDriver": "native-0.2",
    "MountLabel": "",
    "ProcessLabel": "",
    "AppArmorProfile": "",
    "RestartCount": 0,
    "UpdateDns": false,
    "Volumes": {
        "/portal_api_dfo_hyperf": "/data/portal_api_dfo_hyperf"
    },
    "VolumesRW": {
        "/portal_api_dfo_hyperf": true
    },
    "AppliedVolumesFrom": null
}

```

### 最后重启 docker，然后查看容器相关配置信息是否已经修改完毕

```shell

# 重启 docker
service docker restart 
# 或者
systemctl restart docker

# 查看容器相关配置信息
docker inspect {容器的名称或者 id }
# 比如：
docker inspect cbe26510c276

# 配置符合你的要求后，再次启动容器
docker start {容器的名称或者 id }
# 比如：
docker start cbe26510c276

```

## 如果是使用的 `Docker Desktop for Mac` 时

### 参考文献

> [How to login the VM of Docker Desktop for Mac](https://www.modb.pro/db/5458)

因为在 Docker for MacOS 中，容器的宿主机并不是 MacOS 本身，而是在 MacOS 中运行的一个 VM 虚拟机
。虚拟机的路径可以通过查看 Docker Desktop 的配置界面 `Disk image location` 配置获得。

### 那么我们如何进入这个虚拟机呢？

最简单的方式是采用 [justincormack/nsenter1](https://github.com/justincormack/nsenter1) 进入，这个镜像只有 101KB，已经非常小了。

```shell

# –rm 表示在退出的时候就自动删除该容器；
# –privileged 表示允许该容器访问宿主机（也就是我们想要登录的 VM ）中的各种设备；
# –pid=host 表示允许容器共享宿主机的进程命名空间（namespace），或者通俗点儿解释就是允许容器看到宿主机中的各种进程；
docker run -it --rm --privileged --pid=host justincormack/nsenter1

```

然后再进入 `/var/lib/docker/containers` 目录修改 `config.v2.json` 配置文件和 `hostconfig.json` 配置文件即可。整体来说，在 MacOS 上除了进入  `/var/lib/docker/containers` 目录时，进入方式有所不同以外，修改配置文件方式和上文一样。需要注意的是，修改的时候请使用 `vi` 编辑器，因为这个镜像没有安装 `vim` 编辑器的。

```shell

# 比如：
vi /var/lib/docker/containers/a7377587b9f08cfe87af9a8ffa4da0f90bf07fb0a1cd6833a5ffcd9c37b842d0/config.v2.json

vi /var/lib/docker/containers/a7377587b9f08cfe87af9a8ffa4da0f90bf07fb0a1cd6833a5ffcd9c37b842d0/hostconfig.json

```
