package all

import (
	_ "github.com/PUGE/logkit/metric/curl"
	_ "github.com/PUGE/logkit/metric/system"
	_ "github.com/PUGE/logkit/metric/telegraf"
	_ "github.com/PUGE/logkit/metric/telegraf/memcached"
)
