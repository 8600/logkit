package builtin

import (
	_ "github.com/puge/logkit/sender/discard"
	_ "github.com/puge/logkit/sender/elasticsearch"
	_ "github.com/puge/logkit/sender/file"
	_ "github.com/puge/logkit/sender/http"
	_ "github.com/puge/logkit/sender/influxdb"
	_ "github.com/puge/logkit/sender/kafka"
	_ "github.com/puge/logkit/sender/mock"
	_ "github.com/puge/logkit/sender/mongodb"
	_ "github.com/puge/logkit/sender/pandora"
)
