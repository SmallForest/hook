# hook
golang写的处理gitea Hook更新数据的web服务
## git服务
git服务使用的是gitea，在后台-系统web钩子，添加本应用程序的url:port/hook即可，比如  http://121.121.121.121:8888/hook
## 运行
可以直接端口访问也可以使用NGINX做反向代理，更好的隐藏端口。