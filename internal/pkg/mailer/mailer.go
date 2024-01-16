package mailer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hifat/con-q-api/internal/pkg/zlog"
)

type ReqSendEmail struct {
	From          string `json:"From"`
	To            string `json:"To"`
	TemplateID    string `json:"TemplateID"`
	TemplateModel any    `json:"TemplateModel"`
}

func SendEmail(req ReqSendEmail) (res *http.Response, err error) {
	fmt.Printf("\n reset password info: %+v \n", req.TemplateModel)
	time.Sleep(3 * time.Second)
	zlog.Info("let's say the email has been sent")
	return nil, nil
}
