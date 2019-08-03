# toml 格式文件的读取

TOML 格式简介: [toml](https://github.com/toml-lang/toml)

Go 读取toml可以使用: [go-toml](https://github.com/pelletier/go-toml)

## Import

下载库
```bash
go get https://github.com/pelletier/go-toml
```

```go
import "github.com/pelletier/go-toml"
```

## Load

可以从`string`,`[]byte`中读取数据
```go
content := `
key=val
[section]
    key=inside section
`
tree, err :=toml.Load(content)
tree2, err :=toml.LoadBytes([]bytes(content))

```

得到tree之后用Get即可

```
value := tree.Get("key").(String)
sectionValue:= tree.Get("scection.key").(String)

```


## Marshal

传一个`*struct`给他即可

```go
func Marshal(v interface{}) ([]byte, error)
```

再写入文件中即可


go-toml 支持使用类似于json格式的tag
```
toml:"Field"      Overrides the field's name to output.
omitempty         When set, empty values and groups are not emitted.
comment:"comment" Emits a # comment on the same line. This supports new lines.
commented:"true"  Emits the value as commented.
```

## Unmarshal



```go
type Postgres struct {
    User     string
    Password string
}
type Config struct {
    Postgres Postgres
}

doc := []byte(`
	[postgres]
	user = "pelletier"
	password = "mypassword"`)

config := Config{}
toml.Unmarshal(doc, &config)
fmt.Println("user=", config.Postgres.User)
```
