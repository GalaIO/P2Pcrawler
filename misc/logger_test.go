package misc

import "testing"

func TestConsoleLogger(t *testing.T) {
	type family struct {
		Addr  string
		Email string
	}

	logger := GetLogger().SetPrefix("main")
	logger.Debug("test debug", Dict{"name": "xiaoguo", "grade": List{100, 99, 98}, "family": family{
		Addr:  "wall stresst",
		Email: "gala@163.com",
	}})
	logger.Info("test info", nil)
	logger.Warn("test warn", nil)
	logger.Error("test error", nil)
	logger.Trace("test trace", nil)
}
