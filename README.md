# echologrus

echo logrus middleware

## Usage

```go
import (
    "github.com/westwin/echologrus"
  	"github.com/sirupsen/logrus"
)

e := echo.New()
el := echologrus.Logger{Logger: logrus.StandardLogger()}
e.Logger = el
e.Use(el.Hook())
```
