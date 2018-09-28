// Package builtin does nothing but import all builtin readers to execute their init functions.
package builtin

import (
	_ "github.com/puge/logkit/reader/autofile"
	_ "github.com/puge/logkit/reader/cloudtrail"
	_ "github.com/puge/logkit/reader/cloudwatch"
	_ "github.com/puge/logkit/reader/dirx"
	_ "github.com/puge/logkit/reader/elastic"
	_ "github.com/puge/logkit/reader/http"
	_ "github.com/puge/logkit/reader/kafka"
	_ "github.com/puge/logkit/reader/mockreader"
	_ "github.com/puge/logkit/reader/mongo"
	_ "github.com/puge/logkit/reader/redis"
	_ "github.com/puge/logkit/reader/script"
	_ "github.com/puge/logkit/reader/snmp"
	_ "github.com/puge/logkit/reader/socket"
	_ "github.com/puge/logkit/reader/sql"
	_ "github.com/puge/logkit/reader/tailx"
)
