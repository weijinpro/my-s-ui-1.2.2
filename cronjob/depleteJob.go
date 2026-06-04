package cronjob

import (
	"github.com/weijinpro/my-s-ui-1.2.2/database"
	"github.com/weijinpro/my-s-ui-1.2.2/logger"
	"github.com/weijinpro/my-s-ui-1.2.2/service"
)

type DepleteJob struct {
	service.ClientService
	service.InboundService
}

func NewDepleteJob() *DepleteJob {
	return new(DepleteJob)
}

func (s *DepleteJob) Run() {
	inboundIds, err := s.ClientService.DepleteClients()
	if err != nil {
		logger.Warning("Disable depleted users failed: ", err)
		return
	}
	if len(inboundIds) > 0 {
		err := s.InboundService.RestartInbounds(database.GetDB(), inboundIds)
		if err != nil {
			logger.Error("unable to restart inbounds: ", err)
		}
	}
}
