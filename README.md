# goluent

Override standard log library by using fluend-logger https://github.com/fluent/fluent-logger-golang

All log message print stdout and fluentd with tag goluent.<your hostname>.*.

This is inspired by http://12factor.net/ 

# fluentd configuration

add next

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

