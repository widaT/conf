#go 读取配置文件

##使用方法：

    configPath :="/conf/conf.ini"
    c := conf.NewConfig(configPath)
    localIp = c.Read("base", "localIp")
    port = c.Read("base", "port")
    bufLen =  c.GetInt("file", "bufLen")