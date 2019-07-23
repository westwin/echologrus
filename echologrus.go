package echologrus

import (
	"io"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Logger : implement logrus Logger
type Logger struct {
	*logrus.Logger
}

// Level delegate echo.Logger
func (l Logger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Panic("Invalid level")
	}

	return log.OFF
}

// SetHeader delegate echo.Logger
func (l Logger) SetHeader(_ string) {}

// SetPrefix delegate echo.Logger
func (l Logger) SetPrefix(s string) {}

// Prefix delegate echo.Logger
func (l Logger) Prefix() string { return "" }

// SetLevel delegate echo.Logger
func (l Logger) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	default:
		l.Panic("Invalid level")
	}
}

// Output delegate echo.Logger
func (l Logger) Output() io.Writer {
	return l.Out
}

// SetOutput delegate echo.Logger
func (l Logger) SetOutput(w io.Writer) {
	l.Logger.SetOutput(w)
}

// Printj delegate echo.Logger
func (l Logger) Printj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Print()
}

// Debugj delegate echo.Logger
func (l Logger) Debugj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Debug()
}

// Infoj delegate echo.Logger
func (l Logger) Infoj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Info()
}

// Warnj delegate echo.Logger
func (l Logger) Warnj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Warn()
}

// Errorj delegate echo.Logger
func (l Logger) Errorj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Error()
}

// Fatalj delegate echo.Logger
func (l Logger) Fatalj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Fatal()
}

// Panicj delegate echo.Logger
func (l Logger) Panicj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Panic()
}

func (l Logger) logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path

	//bytesIn := req.Header.Get(echo.HeaderContentLength)

	l.Logger.WithFields(map[string]interface{}{
		//"time_rfc3339": time.Now().Format(time.RFC3339),
		//"remote_ip":     c.RealIP(),
		"status": res.Status,
		"host":   req.Host,
		//"uri":  req.RequestURI,
		//"method":        req.Method,
		"path": p,
		//"referer":       req.Referer(),
		//"user_agent":    req.UserAgent(),
		//"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"latency": stop.Sub(start).String(),
		//"bytes_in":      bytesIn,
		//"bytes_out":     strconv.FormatInt(res.Size, 10),
	}).Info("Handled request")

	return nil
}

func (l Logger) logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return l.logrusMiddlewareHandler(c, next)
	}
}

// Hook is a function to process middleware.
func (l Logger) Hook() echo.MiddlewareFunc {
	return l.logger
}
