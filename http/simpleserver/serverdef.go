package main

// HTTPConfig like listent host and port
type HTTPConfig struct {
	Host string `toml:"host"`
}

// RouteConfig shows the route for specify function
type RouteConfig struct {
	Index string `toml:"index"`
	Echo  string `toml:"echo"`
	Time  string `toml:"time"`
}
// TemplateConfig if ListDirTemplateFile was set , ListDirTemplate will be ignore
type TemplateConfig struct {
	ListDirTemplate     string
	ListDirTemplateFile string
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
	Name string
	Size int64
	LinkName string
}
// DirInfo the info of current directory
type DirInfo struct {
	// Welcome message
	Welcome string
	Path    string
	Files   []FileInfo
}
