achilles
====

achilles 是支持在分布式环境下的异步任务流程执行引擎。

GitLab:https://192.168.0.121/apollo/apollo-auto


- 1、支持手动触发,流程图监控.
- 2、支持定时任务自动触发.
- 3、测试case的管理.
- 4、测试环境的治理.

感觉不错的话，给个星星吧 ：）


安装方法
----

方法一、 编译安装

- 按照依赖 go get -u github.com/astaxie/beego (192.30.253.113   github.com)

- git clone http://192.168.0.121/apollo/apollo-auto.git
- 创建mysql数据库，并将sql.txt导入
- 修改config 配置数据库
- 运行 go build
- 运行 ./run.sh start|stop


