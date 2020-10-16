package main

// HTTPConfig like listent host and port
type HTTPConfig struct {
	Host string `toml:"host" comment:"监听的IP和端口，格式host:port, 如0.0.0.0:60066"`
}

// RouteConfig shows the route for specify function
type RouteConfig struct {
	Index   string `toml:"index" comment:"默认的index文件， 如果访问的目录下有该文件， 则直接重定向至该文件"`
	Echo    string `toml:"echo"`
	Time    string `toml:"time"`
	ListDir string `toml:"listdir`
}

// TemplateConfig if ListDirTemplateFile was set , ListDirTemplate will be ignore
type TemplateConfig struct {
	ListDirTemplate     string `comment:"直接指定模板内容， 仅在ListDirTemplateFile为空时有效"`
	ListDirTemplateFile string `comment:"指定模板文件"`
}

// Config Struct for configuration
type Config struct {
	HTTPConfig        HTTPConfig `toml:"HTTPConfig"`
	DefaultWelcomeMsg string     `toml:"default_welcome_msg"`

	EchoBufSize    int            `toml:"echo_buf_size"`
	RouteConfig    RouteConfig    `toml:"RouteConfig"`
	TemplateConfig TemplateConfig `toml:"TemplateConfig"`
}

// FileInfo the info of file in current directory
type FileInfo struct {
	Name     string
	Size     int64
	LinkName string
}

// DirInfo the info of current directory
type DirInfo struct {
	// Welcome message
	Welcome string
	Path    string
	Files   []FileInfo
}
