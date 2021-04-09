# start -> 2021年3月16日22:08:32

# use this project

* **git clone https://github.com/axing42/siginIn** **or**
* **go get -u github.com/axing42/siginIn**
* **cd siginin**

* **go build**
* 在`config.json`配置文件里根据json对应格式配置好账号密码后(这里多个账户我使用了一个密码,有需求提issue)
* 执行exe文件
# 2021/4/9下午4:39:48 更新：
* 增加命令行参数 动态读取账号密码并进行签到
* **cd siginin**

* **go build** 之后运行exe文件：

* 声明 **-m** 参数 0为默认 可以不用添加 , 直接回车执行

* `hlx.exe`

* 1：单账户签到

* `hlx.exe -u xx.qq.com -p 密码写这里`

* 2: 多账号签到 以逗号分割

* ```go
  hlx.exe -u xx.qq.com,xx.qq.com -p 123,123
  ```

## 待完善 :

* 获取分类所有ID时,天天酷跑分类怎样都获取不到
* 细细回忆这个软件 竟没有一丝美好回忆
* end......
