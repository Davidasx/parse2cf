# Parse2CF

将 [Competitive Companion](https://github.com/jmerle/competitive-companion) 得到的样例保存为 [cf-tool](https://github.com/xalanq/cf-tool) 支持的格式。

## 编译

```bash
go build .
mkdir -p /usr/local/bin/
cp parse2cf /usr/local/bin/
```

## 运行

在终端中运行 `parse2cf`，然后点击浏览器中的 Competitive Companion 插件，就会在当前目录创建名称为 `in#.txt` 和 `ans#.txt` 的样例文件。

