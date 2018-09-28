package all

import (
	_ "github.com/puge/logkit/metric/curl"
	_ "github.com/puge/logkit/metric/system"
	_ "github.com/puge/logkit/metric/telegraf"
	_ "github.com/puge/logkit/metric/telegraf/memcached"
)
