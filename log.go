[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /css/*filepath            --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (1 handlers)
[GIN-debug] HEAD   /css/*filepath            --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (1 handlers)
[GIN-debug] Loaded HTML Templates (4): 
	- 
	- footer.html
	- header.html
	- index.html

[GIN-debug] GET    /api/videos               --> main.main.func1 (4 handlers)
[GIN-debug] POST   /api/videos               --> main.main.func2 (4 handlers)
[GIN-debug] GET    /view/videos              --> learn-gin/controller.VideoController.ShowAll-fm (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
::1 - [20 Jul 24 19:04 +05] POST /api/videos 200 580.765µs 
::1 - [20 Jul 24 19:04 +05] GET /view/videos 200 639.973µs 
