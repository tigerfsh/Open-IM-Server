package gate

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/statistics"
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	rwLock       *sync.RWMutex
	validate     *validator.Validate
	ws           WServer
	rpcSvr       RPCServer
	sendMsgCount uint64
	userCount    uint64
)

func Init(rpcPort, wsPort int) {
	//log initialization
	log.NewPrivateLog(config.Config.ModuleName.LongConnSvrName)
	rwLock = new(sync.RWMutex)
	// TODO 读写锁的使用方式，实现细节（次优先级），RWMutex vs Lock+Mutex
	validate = validator.New()
	statistics.NewStatistics(&sendMsgCount, config.Config.ModuleName.LongConnSvrName, fmt.Sprintf("%d second recv to msg_gateway sendMsgCount", sendMsgCount), 300)
	statistics.NewStatistics(&userCount, config.Config.ModuleName.LongConnSvrName, fmt.Sprintf("%d second add user conn", userCount), 300)
	// TODO 了解statistics（次优先级）
	ws.onInit(wsPort)
	rpcSvr.onInit(rpcPort)
}

func Run() {
	go ws.run()
	go rpcSvr.run()
}
