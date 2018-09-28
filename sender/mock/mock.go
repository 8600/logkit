package mock

import (
	"sync"

	"github.com/json-iterator/go"

	"github.com/PUGE/logkit/conf"
	"github.com/PUGE/logkit/sender"
	. "github.com/PUGE/logkit/utils/models"
	"github.com/qiniu/pandora-go-sdk/base/reqerr"
)

var _ sender.SkipDeepCopySender = &Sender{}

// mock sender is used for debug
type Sender struct {
	name     string
	Datas    []Data
	count    int
	mux      sync.Mutex
	isReqErr bool
}

func init() {
	sender.RegisterConstructor(sender.TypeMock, NewSender)
}

// NewMockSender 测试用sender
func NewSender(c conf.MapConf) (sender.Sender, error) {
	name, _ := c.GetStringOr(sender.KeyName, "mockSender")
	isReqErr, _ := c.GetBoolOr("is_req_err", false)
	ms := &Sender{
		name:     name,
		count:    0,
		mux:      sync.Mutex{},
		isReqErr: isReqErr,
	}
	return ms, nil
}

//Name function will return the name and datas recieved as string
func (mock *Sender) Name() string {
	mock.mux.Lock()
	defer mock.mux.Unlock()
	if mock.isReqErr {
		return mock.name
	}
	raw, err := jsoniter.Marshal(mock.Datas)
	if err != nil {
		raw = []byte(err.Error())
	}
	return mock.name + " " + string(raw)
}

func (mock *Sender) Send(d []Data) error {
	mock.mux.Lock()
	defer mock.mux.Unlock()
	if mock.isReqErr && len(d) > 0 {
		failedDatas := sender.ConvertDatasBack([]Data{d[0]})
		return &StatsError{
			StatsInfo: StatsInfo{
				Success: int64(len(d) - len(failedDatas)),
				Errors:  int64(len(failedDatas)),
			},
			ErrorDetail: reqerr.NewSendError(
				"mock failed",
				failedDatas,
				reqerr.TypeBinaryUnpack,
			),
		}
	}
	mock.Datas = append(mock.Datas, d...)
	mock.count++
	return nil
}

func (mock *Sender) Close() error {
	return nil
}
func (mock *Sender) SendCount() int {
	mock.mux.Lock()
	defer mock.mux.Unlock()
	return mock.count
}

func (_ *Sender) SkipDeepCopy() bool { return true }
