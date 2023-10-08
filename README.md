# 种子杯2023 go-sdk

## 创建人

- chenxuan

## 使用说明

- 需要安装go

## 开发说明
> 核心参考demo/main.cpp,这是一个简单的demo

1. 引入包`github.com/chenxuan520/seedcup-gosdk`,以及`github.com/chenxuan520/seedcup-gosdk/elements`

2. 创建seedcup对象以及初始化`seedcup.Init({config_path})`(需要传入配置文件路径)

	- Init函数返回值为error

3. 调用RegisterCallBack函数,创建消息回调处理

	- 第一个函数参数为`func(msg *elements.GameMsg, game *seedcup.Game) error`,当服务器下发地图时候会触发

		- 第一个参数传入地图中的信息

		- 第二个参数传入本体seedcup,可以使用seedcup的TakeAction函数向服务器发送你的操作

	- 第二个函数参数为`func(playerID int32, winners []int32, _ []elements.Scores) error`,游戏结束时候下发

	- 函数参数的返回值均为error,如果不等于nil,`seedcup.Run()`会返回

4. 调用`seedcup.Run()`,该函数不会返回,除非游戏结束或者发生内部错误(网络错误或者,回调函数发出错误)
