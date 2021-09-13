package plugins

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"os"
	"promagent/config"
	"strings"
	"time"
)

type Register struct {
	config *config.AgentConfig
}

func (r *Register) Init(config *config.AgentConfig) {
	r.config = config
}

func (r *Register) Run() {
	ticker := time.NewTicker(r.config.TaskConfig.Register.Interval)
	defer ticker.Stop()

	api := fmt.Sprintf("%s/v1/prometheus/register", strings.TrimRight(r.config.ServerConfig.Addr, "/"))
	fmt.Println(api)
	// UUID, hostname, Addr
	hostname, _ := os.Hostname()
	params := req.Param{
		"uuid":     r.config.UUID,
		"addr":     r.config.Addr,
		"hostname": hostname,
	}
	headers := req.Header{
		"Authorization": fmt.Sprintf("Token %s", r.config.ServerConfig.Token),
	}

	for {
		// 先执行，后阻塞
		response, err := req.Post(api, params, headers)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("request register--->req response err!!!")
		} else {
			logrus.WithFields(logrus.Fields{
				"response": response.Dump(),
			}).Debug("request register--->req response success!!!")
		}
		<-ticker.C
	}
}
