# echologrus
echo logrus middleware

## Usage

```go
import (
	  "github.com/westwin/echologrus"
  	"github.com/sirupsen/logrus"
)

e := echo.New()
e.Logger = echologrus.Logger{Logger: logrus.StandardLogger()}
e.Use(echologrus.Hook())
```
