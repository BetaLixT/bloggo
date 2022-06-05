package config

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
)

type OptionTest struct {
	SomeConfig  string
	OtherConfig int
}

func TestInitializeConfig(t *testing.T) {
	os.Setenv("PORT", "43623")
	os.Setenv("BLOGGO_OptionTest__SomeConfig", "value00")
	os.Setenv("BLOGGO_OptionTest__OtherConfig", "435")
	lgr, err := zap.NewProduction()
	if err != nil {
		fmt.Println("failed to create logger.... why...?")
		t.FailNow()
	}
	cfg := NewConfig(lgr)
	if cfg.GetString("PORT") != "43623" {
		fmt.Println("port value is invalid")
		t.Fail()
	}
	opt := OptionTest{}
	if err := cfg.UnmarshalKey("OptionTest", &opt); err != nil {
		fmt.Printf("failed to unmarshal option %v\n", err)
		t.FailNow()
	}
	if opt.SomeConfig != "value00" {
		fmt.Printf("value SomeConfig is invalid %s\n", opt.SomeConfig)
		t.Fail()
	}
	if opt.OtherConfig != 435 {
		fmt.Printf("value OtherConfig is invalid %d\n", opt.OtherConfig)
		t.Fail()
	}
}
