package main

import (
	_ "net/http/pprof"
	"os"
	"runtime"

	config "github.com/puge/logkit/conf"
	_ "github.com/puge/logkit/metric/all"
	"github.com/puge/logkit/mgr"
	"github.com/puge/logkit/parser"
	"github.com/puge/logkit/reader"
	"github.com/puge/logkit/samples"
	"github.com/puge/logkit/sender"
	utilsos "github.com/puge/logkit/utils/os"

	"github.com/qiniu/log"
	"github.com/puge/logkit/utils/models"
)

type Config struct {
	MaxProcs   int      `json:"max_procs"`
	DebugLevel int      `json:"debug_level"`
	ConfsPath  []string `json:"confs_path"`
	mgr.ManagerConfig
}

var conf Config

func main() {
	config.Init("f", "", "main.conf")
	if err := config.Load(&conf); err != nil {
		log.Fatal("config.Load failed:", err)
	}
	log.Printf("Config: %#v", conf)

	if conf.MaxProcs == 0 {
		conf.MaxProcs = runtime.NumCPU()
	}
	models.MaxProcs = conf.MaxProcs
	runtime.GOMAXPROCS(conf.MaxProcs)
	log.SetOutputLevel(conf.DebugLevel)

	rr := reader.NewRegistry()

	pr := parser.NewRegistry()
	// 注册你自定义的parser
	pr.RegisterParser("myparser", samples.NewMyParser)

	sr := sender.NewRegistry()
	// 注册你自定义的parser
	sr.RegisterSender("mysender", samples.NewMySender)

	m, err := mgr.NewCustomManager(conf.ManagerConfig, rr, pr, sr)
	if err != nil {
		log.Fatalf("NewManager: %v", err)
	}
	var paths []string
	for _, v := range conf.ConfsPath {
		_, err = os.Stat(v)
		if os.IsNotExist(err) {
			err = os.MkdirAll(v, os.ModePerm)
		}
		if err != nil {
			log.Fatalf("Cannot read or create ConfsPath %s: %v", conf.ConfsPath, err)
		}
		paths = append(paths, v)
	}

	if err = m.Watch(paths); err != nil {
		log.Fatalf("watch path error %v", err)
	}
	utilsos.WaitForInterrupt(func() { m.Stop() })
}
