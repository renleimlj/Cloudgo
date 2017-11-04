# Cloudgo

Cloudgo in Martini

## 框架

        go get github.com/go-martini/martini

选择使用Martini是因为其在github上有非常详尽的中文文档，并且以此构建Web应用相对来说更加容易，语言上也有许多类似于JS的地方，适合入门。

## 测试

* 启动服务器

        go run main.go -p9090

![](https://s1.ax1x.com/2017/11/04/0DFKg.png)

* 在浏览器中输入`http://localhost:9090/hello/renlei`

    页面上出现：`Hello renlei`

* 在终端上用curl

        curl -v http://localhost:9090/hello/renlei

![](https://s1.ax1x.com/2017/11/04/0DkrQ.png)

* 在终端上用ApacheBench

        ab -n 1000 -c 100 http://localhost:9090/hello/renlei

![](https://s1.ax1x.com/2017/11/04/0DAbj.png)
![](https://s1.ax1x.com/2017/11/04/0DVVs.png)

一共发送了1000个请求，失败0个，成功1000个，耗时0.525秒。其中，百分之五十的请求可以在49毫秒内得到响应，百分之九十的请求可以在69毫秒内得到响应，最长的一个请求得到响应的时间为103毫秒。