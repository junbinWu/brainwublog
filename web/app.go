package web

import (
	"log"
	"net/http"
	"os"
)

type Context struct {
	Request *http.Request
	server  *Server
	Params  map[string]string
	http.ResponseWriter
}

func (this *Context) NotFound(msg string) {
	this.ResponseWriter.WriteHeader(404)
	this.ResponseWriter.Write([]byte(msg))
}

func (this *Context) WriteString(msg string) {
	this.ResponseWriter.Header().Add("Content-Type","text/html")
	this.ResponseWriter.Write([]byte(msg))
}

func (this *Context) NotModified() {
	this.ResponseWriter.WriteHeader(304)
}

func (this *Context) Redirect(status int, url string) {
	this.ResponseWriter.Header().Set("Location", url)
	this.ResponseWriter.WriteHeader(status)
}

func (this *Context) SetCookie(c *http.Cookie) {
	http.SetCookie(this.ResponseWriter,c)
//	this.ResponseWriter.Header().Add("Set-Cookie", c.String())
}

func (this *Context) GetCookie(name string) (*http.Cookie, error) {
	return this.Request.Cookie(name)
}

func (this *Context) Forbidden() {
	this.ResponseWriter.WriteHeader(403)
}

func (this *Context) Unauthorized() {
	this.ResponseWriter.WriteHeader(401)
}

func Run(addr string, server *Server) {
	if server == nil {
		mainServer.Run(addr)
	} else {
		server.Run(addr)
	}
}

func SetLogger(logger *log.Logger) {
	mainServer.Logger = logger
}

func AddRoute(r string, handler HandleFunc, method string) {
	if len(method) == 0 {
		mainServer.AddRoute(r, handler, "GET")
	} else {
		mainServer.AddRoute(r, handler, method)
	}
}

func SetStaticDir(dir string) {
	mainServer.Config.StaticDir = dir
}

var mainServer = NewServer()

func NewServer() *Server {
	return &Server{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
		Config: ServerConfig{StaticDir:""},
	}
}
