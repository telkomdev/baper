## Baper

Go library which helps you read `CPU Stat` and convert it to `JSON`. You can also use `Baper` as `Unix Daemon`

### Getting started

Running `CLI Mode`
```shell
$ make build
```

Show help and options
```shell
$ baper -h
```

Running with specific `interval` 
```shell
$ baper -i 2
```

Output 
```shell
process running on PID =  73863
{"disk":{"kbPerTime":40.96,"tps":30,"mbPerSecond":1.2},"cpu":{"user":4,"system":2,"idle":94},"loadAverage":{"oneMinute":2.01,"fiveMinute":1.58,"fifteenMinute":1.56}}{"disk":{"kbPerTime":6.81,"tps":45,"mbPerSecond":0.3},"cpu":{"user":3,"system":1,"idle":96},"loadAverage":{"oneMinute":2.01,"fiveMinute":1.58,"fifteenMinute":1.56}}{"disk":{"kbPerTime":128,"tps":1,"mbPerSecond":0.12},"cpu":{"user":3,"system":1,"idle":96},"loadAverage":{"oneMinute":1.93,"fiveMinute":1.57,"fifteenMinute":1.56}}
```