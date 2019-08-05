 # Simple Go HTTP Server

这里是一个GO编写的HTTP程序。

## Features

- [作为文件服务器](doc/)
- [显示时间](time)    
- [显示用户UA](echo)
- TOML 格式的配置文件 [example](config.html)

## 作为文件服务器

可以自动列出当前文件夹的内容， 并且可以通过模板指定如何渲染

```toml
[TemplateConfig]
    # 直接指定模板内容， 仅在ListDirTemplateFile为空时有效
    ListDirTemplate = """
    <html>
    <head><title>{{.Welcome}}</title></head>
    <body>
    <h1>{{.Welcome}}</h1>
    <p>Path: {{.Path}}</p>
    <table>
    <tr><th>Name</th><th>Size</th></tr>
    {{range .Files}}
        <tr>
            <td><a href=\"{{.LinkName}}\">{{.Name}}</a></td>
            <td>{{.Size}} bytes</td>
        </tr>
    {{end}}
    </table>
    </body>
    <html>
    """
    # 指定模板文件
    ListDirTemplateFile = "dir.template.html"
```
