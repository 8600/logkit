// Package builtin does nothing but import all builtin parsers to execute their init functions.
package builtin

import (
	_ "github.com/puge/logkit/parser/csv"
	_ "github.com/puge/logkit/parser/empty"
	_ "github.com/puge/logkit/parser/grok"
	_ "github.com/puge/logkit/parser/json"
	_ "github.com/puge/logkit/parser/kafkarest"
	_ "github.com/puge/logkit/parser/logfmt"
	_ "github.com/puge/logkit/parser/mysql"
	_ "github.com/puge/logkit/parser/nginx"
	_ "github.com/puge/logkit/parser/qiniu"
	_ "github.com/puge/logkit/parser/raw"
	_ "github.com/puge/logkit/parser/syslog"
)
