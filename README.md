### 介绍

本项目是 [projectdiscovery/gologger: A simple layer for leveled logging in go (github.com)](https://github.com/projectdiscovery/gologger) 的改编，修改了一部分功能，让日志颜色更加舒服



### 安装

```sh
go get -u github.com/Zcentury/gologger
```



### 使用方法

> 基础语法

```go
gologger.Info().Msg("验证成功")
gologger.Debug().Msg("身份验证已通过")
gologger.Warning().Msg("此版本已淘汰，请及时更新")
gologger.Error().Msg("参数校验失败！")
gologger.Fatal().Msg("程序中断！")
```

> 添加时间

```go
gologger.Info().TimeStamp().Msg("验证成功")
gologger.Debug().TimeStamp().Msg("身份验证已通过")
gologger.Warning().TimeStamp().Msg("此版本已淘汰，请及时更新")
gologger.Error().TimeStamp().Msg("参数校验失败！")
gologger.Fatal().TimeStamp().Msg("程序中断！")
```

或者全局设置

```go
gologger.LoggerOptions.SetTimestamp(true)
gologger.Info().Msg("验证成功")
gologger.Debug().Msg("身份验证已通过")
gologger.Warning().Msg("此版本已淘汰，请及时更新")
gologger.Error().Msg("参数校验失败！")
gologger.Fatal().Msg("程序中断！")
```

> 添加格外参数

```go
gologger.Error().Str("错误位置", "18行").Msg("程序中断！")
```

> msgf

```go
gologger.Error().Msgf("程序中断！%s", "语法错误")
```



### 高级用法

> 自定义输出格式

定义一个结构体，实现 **Format** 接口

```go
type FormatDemo struct{}

func (f *FormatDemo) Format(event *LogEvent) ([]byte, error) {
    result := ""
    
	for k, v := range event.Metadata {
        result += fmt.Sprintf("%s=%s", k, v)
	}

    return []byte(result)
}
```

配置全局Format

```go
gologger.LoggerOptions.SetFormatter(&FormatDemo{})
```

然后正常使用

```go
gologger.Info().Msg("验证成功")
gologger.Debug().Msg("身份验证已通过")
gologger.Warning().Msg("此版本已淘汰，请及时更新")
gologger.Error().Msg("参数校验失败！")
gologger.Fatal().Msg("程序中断！")
```

