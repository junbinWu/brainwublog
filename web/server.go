package web

import (
	"log"
	"regexp"
	"net/http"
	"path"
	"time"
	"strings"
	"os"
)

type Server struct {
	routes []route
	Logger *log.Logger
	Config ServerConfig
}

func (this *Server) SetStaticDir(dir string) {
	this.Config.StaticDir = dir
}

func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.dispatchHandler(w, r)
}

func (this *Server) dispatchHandler(w http.ResponseWriter, r *http.Request) {
	resquestPath := r.URL.Path
	ctx := Context{r, this, make(map[string]string), w}
	r.ParseForm()

	for i, v := range r.Form {
		ctx.Params[i] = v[0]
	}

	//如果是静态文件自动处理
	if r.Method == "GET" && this.ServeStaticFile(resquestPath, w, r) {
		return
	}

	t := time.Now()
	defer this.logInfo(ctx, t)
	for _, fRoute := range this.routes {
		//如果请求的方法与路由注册时的方法不同 则继续循环
		if r.Method != fRoute.method {
			continue
		}

		reg := fRoute.reg
		//如果注册路由不匹配继续循环
		if !reg.MatchString(resquestPath) {
			continue
		}

		matchStr := reg.FindStringSubmatch(resquestPath)
		//如果匹配出来的结果与请求路径的长度不符 则说明不是一个请求继续匹配
		if len(matchStr) == 0 || len(matchStr[0]) != len(resquestPath) {
			continue
		}
		fRoute.handler(&ctx)
		return
	}

	ctx.NotFound("无此页面")
}

func (this *Server) AddRoute(r string, handler HandleFunc, method string) {
	if len(method) == 0 {
		method = "GET"
	}
	if reg, err := regexp.Compile(r); err == nil {
		rt := route{
			reg : reg,
			path : r,
			handler:handler,
			method:method,
		}
		this.routes = append(this.routes, rt)
	}
}


func (this *Server) logInfo(ctx Context, t time.Time) {
	duration := time.Now().Sub(t)
	req := ctx.Request
	client := req.RemoteAddr
	pos := strings.LastIndex(client, ":")
	if pos > 0 {
		client = client[0:pos]
	}
	this.Logger.Println("from:", client, req.Method, req.URL.Path, duration)
}

func (this *Server) ServeStaticFile(name string, w http.ResponseWriter, r *http.Request) bool {
	var fp string
	if len(os.Args) >= 2 {
		fp = path.Join(os.Args[1],this.Config.StaticDir, name)
	} else {
		fp = path.Join(this.Config.StaticDir, name)
	}
	//先对路径进行判断 如果文件存在则返回
	if fileExists(fp) {
		//会自己做304处理  灰常智能
		http.ServeFile(w, r, fp)
		return true
	} else {
		//如果是以/开头的路径则去除/ 这里寻找的是相对路径
		if has := strings.HasPrefix(fp,"/"); has {
			fp = fp[1:]
			if fileExists(fp) {
				http.ServeFile(w, r, fp)
				return true
			}
		}
		return false
	}
}

func (this *Server) Run(addr string) {
	this.Logger.Println("Listen", addr)
	err := http.ListenAndServe(addr, this)
	if err != nil {
		log.Fatal("listenerServe: ", err)
	}
}

type route struct {
	reg *regexp.Regexp
	path        string
	handler     HandleFunc
	httpHandler http.Handler
	method      string
}

type ServerConfig struct {
	StaticDir string
	Addr      string
	Port      string
}

type HandleFunc func(ctx *Context)
