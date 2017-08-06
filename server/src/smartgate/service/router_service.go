package service

import (
	"smartgate/dao"
	"common/util"
	"smartgate/model"
	"time"
)

func CreateEvidence(userId string) (evidence *model.RouterEvidence, err error) {
	evidence = new(model.RouterEvidence)
	evidence.EvidenceId = util.NewUuid()
	evidence.UserId = userId
	evidence.CreateTime = time.Now()
	evidence.ExpiresAt = evidence.CreateTime.AddDate(0, 0, 1)
	evidence.Status = 1
	err = dao.NewRouterDao().InsertEvidence(evidence)
	return
}
