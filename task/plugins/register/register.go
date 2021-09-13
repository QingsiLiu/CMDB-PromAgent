package register

import (
	"crypto/tls"
	"fmt"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"net/http"
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

	// 跳过不安全的验证（https证书）
	request := req.New()
	transport, _ := request.Client().Transport.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	for {
		// 先执行，后阻塞
		response, err := request.Post(api, req.BodyJSON(params), headers)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("request register--->req response err!!!")
		} else {
			body, _ := response.ToString()
			logrus.WithFields(logrus.Fields{
				"response": body,
			}).Debug("request register--->req response success!!!")
		}
		<-ticker.C
	}
}
