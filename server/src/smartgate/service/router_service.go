package service

import (
	"smartgate/dao"
	"common/util"
	"github.com/carsonsx/log4g"
	"common/model"
	"time"
)

func CreateEvidence(userId string, typ int8) (evidence *model.RouterEvidence, err error) {
	evidence = new(model.RouterEvidence)
	evidence.EvidenceId = util.NewUuid()
	evidence.UserId = userId
	evidence.Direction = typ
	evidence.CreateTime = time.Now()
	evidence.ExpiresAt = evidence.CreateTime.AddDate(0, 0, 1)
	evidence.Status = 1
	err = dao.NewRouterDao().InsertEvidence(evidence)
	if err == nil {
		log4g.Info("create new evidence %s for user %s", evidence.EvidenceId, evidence.UserId)
	}
	return
}
