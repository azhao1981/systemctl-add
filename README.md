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
```

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