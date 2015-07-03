# goluent

Override standard log library by using fluend-logger https://github.com/fluent/fluent-logger-golang

All log message print stdout and fluentd with tag goluent.<your hostname>.*.

This is inspired by http://12factor.net/ 

# fluentd configuration

add this

```
<match goluent.**>
  @type stdout
  @id stdout_output
</match>
```


# How to use

in your *.go file
```
import (
       log "github.com/drobune/goluent"
)
```

```
log.Info("goluent now")
// 2015/07/03 15:33:10 goluent now

log.Errorf("over %v sec", 10)
//2015/07/03 15:37:00 over 10 sec
```
