// Package builtin does nothing but import all builtin readers to execute their init functions.
package builtin

import (
	_ "github.com/PUGE/logkit/reader/autofile"
	_ "github.com/PUGE/logkit/reader/cloudtrail"
	_ "github.com/PUGE/logkit/reader/cloudwatch"
	_ "github.com/PUGE/logkit/reader/dirx"
	_ "github.com/PUGE/logkit/reader/elastic"
	_ "github.com/PUGE/logkit/reader/http"
	_ "github.com/PUGE/logkit/reader/kafka"
	_ "github.com/PUGE/logkit/reader/mockreader"
	_ "github.com/PUGE/logkit/reader/mongo"
	_ "github.com/PUGE/logkit/reader/redis"
	_ "github.com/PUGE/logkit/reader/script"
	_ "github.com/PUGE/logkit/reader/snmp"
	_ "github.com/PUGE/logkit/reader/socket"
	_ "github.com/PUGE/logkit/reader/sql"
	_ "github.com/PUGE/logkit/reader/tailx"
)
