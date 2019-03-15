apollo-auto apollo自动化测试平台 V2.0
====

apollo-auto是基于apollo的自动化测试平台，采用golang开发，安装方便，资源消耗少，支持大并发。

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


数据库
CREATE DATABASE Apollo_Auto;


访问方式
----
前台访问：http://your_host:8080
用户名：admin 密码：123456


目前支持比对的数据结构定义如下(对比的数据结构以json形式展示)
----

ApolloData

字段 | 注释 
--- | ----
id  |  id
currentNode | 当前节点id
end | 是否结束 0-不是 1-是


ApolloBorrow

字段 | 注释
---- | ----
id | id
auditState | 进件状态
loanState | 放款状态
amount | 借款金额 (分)
realAmount | 实际放款金额(分)
loanTime | 实际放款时间

ApolloRepay(暂无)

TradeBorrow
字段 | 注释
---- | ----
id | id
state | 状态
amount | 金额


TradeBill(暂无)

TradeRepayTask(暂无)

TradeRepayTaskRelation(暂无)



预期结果与实际结果的对比



