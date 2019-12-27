package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/core/lager"
	"io/ioutil"
	"net/http"
)

// 与calculator服务不同，体质指数界面（web-app）在Go-chassis框架内是一个消费者类型的服务，故只需调用chassis.Init()对Go-chassis框架进行初始化
func main() {
	http.HandleFunc("/", BmiPageHandler)
	http.HandleFunc("/calculator/bmi", BmiRequestHandler)

	if err := chassis.Init(); err != nil {
		lager.Logger.Errorf("Init FAILED %s", err.Error())
		return
	}

	port := flag.String("port", "8889", "Port web-app will listen")
	address := flag.String("address", "0.0.0.0", "Address web-app will listen")
	fullAddress := fmt.Sprintf("%s:%s", *address, *port)
	http.ListenAndServe(fullAddress, nil)
}

// 体质指数界面（web-app）微服务收到前端界面发过来的请求时，通过core.NewRestInvoker()将请求转发到calculator服务。
// 在转发调用的过程中，用户并不需要感知calculator服务具体的地址和端口，服务发现的过程由 go-chassis 框架自动完成。
func BmiRequestHandler(w http.ResponseWriter, r *http.Request) {
	heightStr := r.URL.Query().Get("height")
	weightStr := r.URL.Query().Get("weight")

	// requestURI := fmt.Sprintf("cse://calculator/bmi?height=%s&weight=%s", heightStr, weightStr)
	requestURI := fmt.Sprintf("http://RESTServer/calculator/bmi?height=%s&weight=%s", heightStr, weightStr)
	restInvoker := core.NewRestInvoker()
	req, _ := rest.NewRequest("GET", requestURI, nil)
	resp, _ := restInvoker.ContextDo(context.TODO(), req)
	fmt.Println(fmt.Sprintf("%#v", resp))
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lager.Logger.Errorf("Read response body ERROR: %s", err.Error())
	}
	w.Write(body)
}

// 通过golang 官方库中http.ServeFile将前端静态页面展示出来
func BmiPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "external/index.html")
}
