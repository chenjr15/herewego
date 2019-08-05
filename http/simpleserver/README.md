 # Simple Go HTTP Server

这里是一个GO编写的HTTP程序。

## Features

- [作为文件服务器](doc/)
- [显示时间](time)    
- [显示用户UA](echo)
- TOML 格式的配置文件 [example](config.toml)

## 作为文件服务器

可以自动列出当前文件夹的内容， 并且可以通过模板指定如何渲染

```toml
[TemplateConfig]
    # 直接指定模板内容， 仅在ListDirTemplateFile为空时有效
    ListDirTemplate = "<html>\n\t\t\t\t<head><title>{{.Welcome}}</title></head>\n\t\t\t\t<body>\n\t\t\t\t<h1>{{.Welcome}}</h1>\n\t\t\t\t<p>Path: {{.Path}}</p>\n\t\t\t\t<table>\n\t\t\t\t<tr>\n\t\t\t\t\t\t<th>Name</th>\n\t\t\t\t\t\t<th>Size</th>\n\t\t\t\t</tr>\n\t\t\t\t{{range .Files}}\n\t\t\t\t<tr>\n\t\t\t\t\t\t<td><a href=\"{{.LinkName}}\" >{{.Name}}</a></td>\n\t\t\t\t\t\t<td>{{.Size}} bytes</td>\n\t\t\t\t</tr>\n\t\t\t\t{{end}}\n\t\t\t\t</table>\n\t\t\t\t</body>\n\t\t\t\t<html>\n\t\t\t\t"
    # 指定模板文件
    ListDirTemplateFile = "dir.template.html"
```
