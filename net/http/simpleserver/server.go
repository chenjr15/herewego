package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pelletier/go-toml"
)

// GetGeneralHandler Return the func mathchs "/"
func GetGeneralHandler(config *Config) func(http.ResponseWriter, *http.Request) {
	var tmpl = template.New("dirinfo")
	if config.TemplateConfig.ListDirTemplateFile != "" {
		bufs, err := ioutil.ReadFile(config.TemplateConfig.ListDirTemplateFile)
		if err != nil {
			log.Printf("Fail to open templte file %s,%v\n", config.TemplateConfig.ListDirTemplateFile, err)
			log.Printf("Try to use default template file %s\n", config.TemplateConfig.ListDirTemplate)
			tmpl = template.Must(tmpl.Parse(config.TemplateConfig.ListDirTemplate))

		}
		template.Must(tmpl.Parse(string(bufs)))

	} else {
		tmpl = template.Must(tmpl.Parse(config.TemplateConfig.ListDirTemplate))
	}

	generalHandler := func(w http.ResponseWriter, r *http.Request) {

		epath, err := url.PathUnescape("." + r.URL.EscapedPath())
		raddr := r.RemoteAddr

		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			raddr = fmt.Sprintf("XFF|%s", xff)
		}
		log.Printf("From: %s\tRequest: %s\nUA:%s", raddr, epath, r.UserAgent())

		w.Header().Set("Content-Type", "text/html")

		writer := bufio.NewWriterSize(w, config.EchoBufSize)
		defer func() {

			err := writer.Flush()
			// n, err := w.Write([]byte(msg))
			if err != nil {
				log.Println(err)
			}
			// log.Printf("Write %d bytes", n)
		}()

		// 尝试匹配index 文件
		_, err = os.Stat(epath + config.RouteConfig.Index)
		if os.IsNotExist(err) {
			//Index 文件不存在的情况
			// Just do noting
		} else if err == nil {
			log.Printf("Redirecting ")
			// 这里如果是子目录的话要加上子目录的路径,否则永远都会跳到首页
			http.Redirect(w, r, config.RouteConfig.Index, http.StatusSeeOther)
			return

		} else {
			log.Printf("Err %v", err)
			http.Redirect(w, r, config.RouteConfig.Index, http.StatusSeeOther)
			return

		}

		//  去所访问的文件信息
		fi, err := os.Stat(epath)
		if os.IsNotExist(err) {
			// 文件不存在的情况
			w.WriteHeader(404)
			writer.WriteString("404")
			log.Printf("epath:%s, epath: %s, %v", epath, epath, err)
			return
		} else if err != nil {
			log.Println(err)
		}
		handleRegularFile := func() {

			// 返回文件内容
			//w.Header().Set("Content-Type", "raw")
			// + 应该返回特定的文件内容
			mimetype := mime.TypeByExtension(path.Ext(epath))
			w.Header().Set("Content-Type", mimetype)
			// w.Header().Set("Content-Type", "application/octet-stream")
			// w.Header().Set("Content-Disposition", "attachment; filename="+fi.Name())
			f, err := os.OpenFile(path.Join(path.Dir(epath), fi.Name()), os.O_RDONLY, 0644)
			if err != nil {
				writer.WriteString(err.Error())
			}
			buf := make([]byte, 2048)
			for {
				n, err := f.Read(buf)
				if n > 0 && err != nil {
					log.Printf("Failed to read file. %v", err)
				}
				if n == 0 {
					break
				}
				writer.Write(buf[:n])
			}
		}
		var handleDir func()
		handleDir = func() {

			dir, err := os.Open(epath)
			// 尝试读取文件内容
			files, err := dir.Readdirnames(100)
			if err != nil {
				log.Print(err)
			}
			var dirInfo DirInfo
			dirInfo.Welcome = config.DefaultWelcomeMsg
			dirInfo.Path = epath
			dirInfo.Files = make([]FileInfo, len(files)+1)
			dirInfo.Files[0] = FileInfo{
				Name:     "../",
				LinkName: "../",
				Size:     0,
			}

			for i, filename := range files {
				i++
				filllFileInfo(path.Join(epath, filename), &dirInfo.Files[i])

			}
			if err := tmpl.Execute(w, dirInfo); err != nil {
				log.Printf("Failed to render template, %v", err)
			}

		}
		handleSymLink := func() {
			writer.WriteString("This is a symbolic link \n")

		}
		handleNoMatch := func() {
			writer.WriteString(fmt.Sprintf("unmatched filetype %d", fi.Mode()))
			log.Printf("Not matched %s ,filemode: %v", epath, fi.Mode())
		}

		mode := fi.Mode()
		switch {
		case mode&os.ModeSymlink != 0:
			//
			handleSymLink()
		case mode.IsRegular():
			handleRegularFile()

		case mode.IsDir():
			handleDir()

		default:
			handleNoMatch()
			log.Printf("Not matched %s ,filemode: %d", epath, fi.Mode())

		}

		return
	}
	return generalHandler

}

// ShowTime return current time
func ShowTime(w http.ResponseWriter, r *http.Request) {
	n, err := w.Write([]byte(time.Now().String()))
	if err != nil {
		log.Println(err)
	}
	log.Printf("Write %d bytes", n)
	return
}

// Echo return what remote says in http body, limit size is set by config file
func Echo(w http.ResponseWriter, r *http.Request) {
	writer := bufio.NewWriter(w)
	buf := make([]byte, config.EchoBufSize)
	defer writer.Flush()

	writer.WriteString(fmt.Sprintf("Method: %v\n", r.Method))
	writer.WriteString(fmt.Sprintf("Remote: %v\n", r.RemoteAddr))
	writer.WriteString(fmt.Sprintf("URL: %v\n", r.URL.EscapedPath()))
	writer.WriteString(fmt.Sprintf("Query: %v\n", r.URL.RawQuery))
	writer.WriteString(fmt.Sprintf("UA: %v\n", r.UserAgent()))
	writer.WriteString(fmt.Sprintf("Body:\n"))

	count, readfail := r.Body.Read(buf)
	if readfail != nil {
		log.Printf("Fail to read from request %v", readfail)
	}
	n, err := writer.Write(buf[:count])
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Write %d bytes", n)
	return
}

// ListDir return json format ls result
func ListDir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// 获取ls参数, 即目标路径
	lspath := path.Clean(r.URL.Query().Get("path"))
	lspath = path.Clean(lspath)
	log.Println(lspath)
	for strings.HasPrefix(lspath, "../") || strings.HasPrefix(lspath, "./") {
		lspath = strings.ReplaceAll(lspath, "../", "")
		lspath = strings.ReplaceAll(lspath, "./", "")
	}
	if lspath == ".." {
		lspath = ""
	}
	if lspath == "." {
		lspath = ""
	}
	if lspath == "" {
		lspath = "/404"
	}

	// no /
	// lspath = url.PathEscape(lspath)
	log.Println("listing:", lspath)

	fi, err := os.Stat(lspath)
	if err != nil {

		w.Write([]byte(err.Error()))
		return
	}
	var filelist []FileInfo
	mode := fi.Mode()
	switch {
	case mode&os.ModeSymlink != 0:
	case mode.IsRegular():
		fileinfo := FileInfo{
			fi.Name(),
			0,
			fi.Name(),
		}

		filelist = []FileInfo{fileinfo}

	case mode.IsDir():
		dir, err := os.Open(lspath)
		// 尝试读取文件内容
		files, err := dir.Readdirnames(100)
		if err != nil {
			log.Print(err)
			return
		}

		filelist = make([]FileInfo, len(files)+1)
		filelist[0] = FileInfo{
			Name:     "../",
			LinkName: "../",
			Size:     0,
		}

		for i, filename := range files {
			i++
			filllFileInfo(path.Join(lspath, filename), &filelist[i])

		}

	default:
		log.Printf("Not matched %s ,filemode: %d", lspath, fi.Mode())

	}
	bytes, err := json.Marshal(filelist)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(bytes)
	return
}

func filllFileInfo(filepath string, fileinfo *FileInfo) {
	file, err := os.Stat(filepath)
	if err != nil {
		log.Printf("Fail to stat file %s, %v", filepath, err)
		return
	}
	if fileinfo.Name == "" {
		fileinfo.Name = file.Name()
		fileinfo.Size = file.Size()
	}

	// 对于软链接，读取链接地址并递归当前函数
	if file.Mode()&os.ModeSymlink != 0 {
		fileinfo.Name = fileinfo.Name + "@"
		realpath, err := os.Readlink(filepath)
		if err != nil {
			log.Printf("Failed to read link %s, %v", filepath, err)
			return

		}

		filllFileInfo(realpath, fileinfo)
	}

	if file.IsDir() {
		fileinfo.LinkName = fileinfo.Name + "/"

	} else {
		fileinfo.LinkName = fileinfo.Name
	}

}

// GetDefaultConfig return pointer of default configuration sturct
func GetDefaultConfig() (defaultConfig *Config) {

	return &Config{
		HTTPConfig: HTTPConfig{
			Host: "0.0.0.0:60066",
		},
		DefaultWelcomeMsg: "Simple Go server!",
		EchoBufSize:       8 * 1024,
		RouteConfig: RouteConfig{
			Index:   "index.html",
			Echo:    "/echo",
			Time:    "/time",
			ListDir: "/listdir",
		},
		TemplateConfig: TemplateConfig{
			ListDirTemplate: `<html>
				<head><title>{{.Welcome}}</title></head>
				<body>
				<h1>{{.Welcome}}</h1>
				<p>Path: {{.Path}}</p>
				<table>
				<tr>
						<th>Name</th>
						<th>Size</th>
				</tr>
				{{range .Files}}
				<tr>
						<td><a href="{{.LinkName}}" >{{.Name}}</a></td>
						<td>{{.Size}} bytes</td>
				</tr>
				{{end}}
				</table>
				</body>
				<html>
				`,
			ListDirTemplateFile: "",
		},
	}
}

// GetDefaultContents return the byte slice of toml default configuration
func GetDefaultContents() []byte {
	bytes, err := toml.Marshal(GetDefaultConfig())
	if err != nil {
		return []byte{}
	}
	return bytes

}
