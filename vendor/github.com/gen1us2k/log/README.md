#Usage
### Import logger package 
``` import "github.com/gen1us2k/log"```

### Set global log level. For example "debug" log level (maximum verbosity)
  * ```log.SetLevel(log.MustParseLevel("debug"))```
  * Other log levels: ``` "panic", "fatal", "error", "warning", "info", "debug"```
	    
### Create instance of named logger:
```
var buf bytes.Buffer
logger := NewCustomLogger("service or module name", &buf)
```
or
```
logger := NewLogger("service or module name")
```
**Important things:**
  * Every call of ```NewLogger``` return one instance by given name. If instance by name already created then this instance returned
  * Every call of ```NewCustomLogger``` returns one instance, but can change io.Writer for already created instance by name

### Use logger for log purposes
```
logger.Info(1,2,3)                           //2016/07/18 16:28:23 INF [service or module name] 1 2 3
logger.Infof("%s %d %s", "log", 1, "test")   //2016/07/18 16:28:23 INF [service or module name] log 1 test
```

# Running tests

```go test```

# log
Yet another logger with support different log levels
This project started by @digitalcrab for [meshbird](https://github.com/meshbird/meshbird) project on GopherGala 2016 hackathon