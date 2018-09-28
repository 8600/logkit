package builtin

import (
	_ "github.com/PUGE/logkit/sender/discard"
	_ "github.com/PUGE/logkit/sender/elasticsearch"
	_ "github.com/PUGE/logkit/sender/file"
	_ "github.com/PUGE/logkit/sender/http"
	_ "github.com/PUGE/logkit/sender/influxdb"
	_ "github.com/PUGE/logkit/sender/kafka"
	_ "github.com/PUGE/logkit/sender/mock"
	_ "github.com/PUGE/logkit/sender/mongodb"
	_ "github.com/PUGE/logkit/sender/pandora"
)
