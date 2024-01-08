# systemctl-add

a simple command that converts commands into systemd services.

## build

```bash
make
or
go build -o systemctl-add
```

## usage

```bash
sudo ./systemctl-add freeswitch '/usr/local/freeswitch/bin/freeswitch -ncwait -nonat'

SERVICE_DIR=./temp/ ./systemctl-add freeswitch '/usr/local/freeswitch/bin/freeswitch -ncwait -nonat'
```

## TODO

+ [ ] remote ssh support 
+ [ ] release

## ref


In a `systemctl .service` file:

- `PermissionsStartOnly=true`: Only the start operation is affected by the user and group settings.
- `TimeoutSec=45s`: Service operations must complete within 45 seconds.
- `Restart=always`: The service will continually restart on failure.
- `UMask=0022`: The service process will use a default umask of 0022, which results in new files having permissions of 644 and directories 755.

- `PermissionsStartOnly=true`: 仅启动操作受用户和组设置的影响。
- `TimeoutSec=45s`: 服务操作必须在45秒内完成。
- `Restart=always`: 服务在失败后将不断重启。
- `UMask=0022`: 服务进程将使用0022的默认umask，这导致新文件的权限为644，目录为755。
- `Type=forking`: 表示服务启动时会产生一个子进程（fork），父进程随即结束，而实际的服务则由子进程继续运行。这通常用于那些自己进行双进程管理的传统UNIX服务。在服务文件中设置`Type=forking`，systemd会等待父进程结束后认为服务已启动。


在GitHub上进行release的步骤通常如下：

1. 确保所有代码已合并到主分支。
2. 在仓库页面，点击“Releases”或“Tags”部分。
3. 点击“Create a new release”或“Draft a new release”按钮。
4. 输入release版本号，通常遵循语义化版本控制规则。
5. 填写release的标题和描述。
6. 如果需要，可以选择包含预发布代码的分支。
7. 可以选择上传编译好的二进制文件或其他相关资料。
8. 点击“Publish release”发布正式版本或“Save draft”保存草稿。

GitHub Actions中的go build命令生成的文件会在运行该命令的GitHub Actions工作流的当前工作目录中。如果你需要在工作流结束后获取这些文件，可以将它们上传到工作流的"artifacts"中，或者部署到一个服务器或发布为GitHub Release的一部分。

It’s common practice to prefix your version names with the letter v. Some good tag names might be v1.0.0 or v2.3.4.

If the tag isn’t meant for production use, add a pre-release version after the version name. Some good pre-release versions might be v0.2.0-alpha or v5.9-beta.3.

Semantic versioning
If you’re new to releasing software, we highly recommend to [learn more about semantic versioning](https://semver.org/).

A newly published release will automatically be labeled as the latest release for this repository.

If 'Set as the latest release' is unchecked, the latest release will be determined by higher semantic version and creation date. [Learn more about release settings.](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository)