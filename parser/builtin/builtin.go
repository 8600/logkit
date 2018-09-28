// Package builtin does nothing but import all builtin parsers to execute their init functions.
package builtin

import (
	_ "github.com/PUGE/logkit/parser/csv"
	_ "github.com/PUGE/logkit/parser/empty"
	_ "github.com/PUGE/logkit/parser/grok"
	_ "github.com/PUGE/logkit/parser/json"
	_ "github.com/PUGE/logkit/parser/kafkarest"
	_ "github.com/PUGE/logkit/parser/logfmt"
	_ "github.com/PUGE/logkit/parser/mysql"
	_ "github.com/PUGE/logkit/parser/nginx"
	_ "github.com/PUGE/logkit/parser/qiniu"
	_ "github.com/PUGE/logkit/parser/raw"
	_ "github.com/PUGE/logkit/parser/syslog"
)
