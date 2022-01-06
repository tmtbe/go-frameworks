# Golang项目脚手架
## 安装
* 先确保安装好golang1.16版本或以上
* 确定设置了环境变量 GOPATH和GOROOT并把其中的bin目录加入到了系统PATH中
* make install完成安装
* 本机需要安装Docker，测试时会利用docker启动相应infra依赖
## 运行
通过Makefile执行操作
* make install 安装
* make run 运行服务
* make test 执行检查、mock代码生成，测试
* make onlyTest 仅仅执行mock代码与测试
* make mock 创建mock
* make check 检查
* make build 执行构建
## 提交CommitLog的规范
参考下面的链接
https://www.conventionalcommits.org/zh-hans/v1.0.0/
## 项目结构
* cmd 
  * main.go 启动命令
  * wire.go wire声明
* internal 项目代码目录
  * app 应用代码目录
    * context app上下文，类似Spring Context，在此声明依赖注入 
    * module1 模块目录，如果项目需要分多个模块
      * application application层处理interfaces过来的请求
      * domain domain层，业务领域层
      * infrastructure infra层，主要通过基础设施实现domain层定义的接口
      * interfaces interfaces层处理API
    * application.go 应用的application，类似Spring Boot Application
  * pkg 基础设施外部依赖目录
    * app 应用级能力
    * cachestore cache缓存能力
    * config viper提供的config能力
    * context infra上下文
    * database gorm提供的database能力
    * log zap包提供的log能力
    * migrate 数据库 migrate
    * redis redis
    * telemetry open telemetry 提供链路追踪和指标收集
    * transports
      * http gin框架提供http能力
    * utils 工具包
    * provider.go wire需要的提供者
* tests 分层测试目录
  * mocks 根据Interface自动生成的mock文件
  * pkg 测试时需要替换（STUB，MOCK）的基础设施
  * sql 存放测试用的数据库数据声明
* resources
  * configs
    * application.yml 项目配置
  * db
    * migrations sql版本文件进行项目初始化db migration的
    
## 分层设计
所有层级设计围绕domain展开，domain是与环境，基础设施，框架无关的纯业务领域层
* Interfaces
  * 处理框架相关的API入口，比如使用Gin框架处理http请求，这层主要涉及与框架的交互，比如处理HTTP的参数获取和验证，集中的错误处理等等
  * 请注意传参给Application层的实体应当尽量避免携带框架相关的验证注释
  * 如有必要需要进行结构体的转换
* Application
  * 承接Interfaces后续的逻辑处理，接收经过Interfaces处理后的参数，如果参数过多应当创建command结构体进行传参
  * 引用Domain的Service完成传输Response实体的组装，Domain Model与Response等简单逻辑的转换
  * 请注意Application层的实体不应当包含Interfaces中框架的注释
  * 本层级设计应当避免引用关于框架的一切行为
  * Application层应当只引用Domain层的Service
* Domain
  * Domain层是独立完成存在的业务领域模型，它不应当依赖任何层级
  * Domain层将暴露Service接口给Application层（这里使用接口暴露目的是易于分层测试）
  * Domain层将定义Repos，Clients等接口和模型交由infra层实现
  * 请注意Domain层定义的模型不应当有Infra相关的注解
* Infrastructure
  * 通过基础设施实现domain层定义的Repos，Clients等接口
  * 本层级主要关注与基础设施间的直接交互
  * 请注意本层级出现的与基础设施相关的注释（比如与Gorm相关的结构体注释）不应当与Domain层耦合，如有必要进行模型转换

# 分层测试
* API+Application层进行测试，主要用于Happy Path，快速的，简单的测试
* Domain+Infra+基础设施层进行测试，主要用于业务逻辑验证，本测试将通过Docker拉起基础设施环境