# grafana-tools

Grafana 小工具，目前只实现了以下功能：

- 自动导出所有 Dashboard 到本地文件，以方便备份和迁移



## 1.编译安装

执行编译命令会后在当前目录的 `./bin` 下生成对应平台的可执行文件

```sh
## 编译当前平台可执行文件
make
# 或
make build

## 编译全平台可执行文件
make all

## 清除现有可执行文件
make clean
```



## 2.使用帮助

```ini
usage: chat [<flags>] <command> [<args> ...]

Grafana 管理工具

Flags:
  -h, --help     Show context-sensitive help (also try --help-long and --help-man).
      --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  bak [<dir>]
    备份所有 Dashboard,（可选）指定备份路径，默认为当前目录
```



## 3.备份所有 Dashboard

执行此操作后会将所有具有 Dashboard 备份到指定目录，所备份的文件可以直接在 Grafana 上导入

备份的文件名格式为：`<仪表板所在文件夹名>-<仪表板名称>.json`

```sh
## 根据实际情况修改 API Key 和 Grafana URI
export grafanaAPI="文档默认的附录部分有获取方法"
export grafanaURL="http://127.0.0.1:3000"

## 使用帮助
## [path] 为可选备份路径，默认为在当前目录下以当前时间生成 `grafana-dashboard-bak-20220907-150130` 格式的目录
bin/grafana-tools bak [path]

## 使用示例（自动以日期格式生成备份目录）
bin/grafana-tools bak

## 使用示例（指定备份路径）
bin/grafana-tools bak /home/bak
```



## 4.手动恢复 Dashboard

- 1.登录到 Grafana
- 2.鼠标移动到页面左侧 + 图标，然后选择 Import
  - 恢复方法1：点击 `Upload JSON file`, 然后选择备份的某个 JSON 文件即可
  - 恢复方法2：直接复制备份文件内容并将其粘贴到 `Import via panel json` 框内即可




## 附录：获取 Grafana API Token

- 1.登录 Grafana
- 2.点击页面左侧齿轮图标以转到 Configuration 页面
- 3.转到 [API Keys] 选项卡
- 4.点击 [Add API Key]
	
	- Key name: 为此 API Key 起个名字，方便记忆和管理
	- Role: 为此角色选择组，需要选择 Admin
	- Time to live: 此 API 有效期，比如 `1d`
- 5.点击 Add 完成添加
- 6.**记录此页面显示的 Key**，此页面只显示一次，关闭后将无法再查看

  - 同时此页面会给出一个类似以下格式的 `curl` 测试命令，以测试是否可以正常访问

    ```sh
    curl -H "Authorization: Bearer eyJrIjoiTEd4ZTJLV3JWVVEwQm1nNkFEbnRiSnNnbjRtV1hxT2siLxxxxx" http://10.100.0.23:3000/api/dashboards/home
    ```



