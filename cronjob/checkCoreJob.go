package cronjob

import (
	"github.com/weijinpro/my-s-ui-1.2.2/service"
)

type CheckCoreJob struct {
	service.ConfigService
}

func NewCheckCoreJob() *CheckCoreJob {
	return &CheckCoreJob{}
}

func (s *CheckCoreJob) Run() {
	s.ConfigService.StartCore("")
}
