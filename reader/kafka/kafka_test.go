package kafka

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/puge/logkit/conf"
	"github.com/stretchr/testify/assert"

	"github.com/puge/logkit/reader"
	. "github.com/puge/logkit/reader/test"
	. "github.com/puge/logkit/utils/models"
)

func TestKafkaReader(t *testing.T) {
	logkitConf := conf.MapConf{
		reader.KeyMetaPath: MetaDir,
		reader.KeyFileDone: MetaDir,
		reader.KeyMode:     reader.ModeElastic,
	}
	meta, err := reader.NewMetaWithConf(logkitConf)
	assert.NoError(t, err)
	defer os.RemoveAll(MetaDir)
	er := &Reader{
		meta:             meta,
		ConsumerGroup:    "group1",
		Topics:           []string{"topic1"},
		ZookeeperPeers:   []string{"localhost:2181"},
		ZookeeperTimeout: time.Second,
		Whence:           "oldest",
		errChan:          make(chan error, 1000),
		lock:             new(sync.Mutex),
		statsLock:        new(sync.RWMutex),
	}
	assert.EqualValues(t, "KafkaReader:[topic1],[group1]", er.Name())

	assert.Equal(t, StatsInfo{}, er.Status())
}
