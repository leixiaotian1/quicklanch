# quicklanch
在个人主机上快速启动软件的小工具

cmd/main.go：程序的入口点。
internal/handler/：包含Echo的路由和控制器逻辑。
internal/service/：定义业务逻辑和用例。
internal/repository/：数据访问层，与数据库或外部服务（如Ollama）交互。
pkg/ollama/：封装与Ollama通信的客户端逻辑。
config/config.go：配置文件，包括环境变量和配置参数。