## hy 



```bash 
# to get the packages
go get 

# to run server
go run main.go

```


```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> hy/routes.SetupPingRoutes.func1 (3 handlers)
[GIN-debug] POST   /merge                    --> hy/routes.MergePDFHandler (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2024/05/05 - 20:28:15 | 200 |    7.395417ms |       127.0.0.1 | POST     "/merge"
[GIN] 2024/05/05 - 20:28:17 | 200 |    2.234458ms |       127.0.0.1 | POST     "/merge"
[GIN] 2024/05/05 - 20:28:17 | 200 |    2.457167ms |       127.0.0.1 | POST     "/merge"
[GIN] 2024/05/05 - 20:28:17 | 200 |    1.614791ms |       127.0.0.1 | POST     "/merge"
[GIN] 2024/05/05 - 20:28:18 | 200 |    2.113958ms |       127.0.0.1 | POST     "/merge"

```