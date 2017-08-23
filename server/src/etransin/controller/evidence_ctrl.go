package controller

import (
	"github.com/gin-gonic/gin"
	"common/sg"
	"fmt"
	"common/vo"
	"github.com/carsonsx/log4g"
	"etransin/service"
	"common/codec"
	"etransin/dao"
	"common/tokenutil"
)

func GetEvidenceIn(c *gin.Context) {
	_getEvidence(c, service.GATE_DIRECTION_IN)
}

func GetEvidenceOut(c *gin.Context) {
	_getEvidence(c, service.GATE_DIRECTION_OUT)
}

func _getEvidence(c *gin.Context, typ int8) {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	evidence, err := service.CreateEvidence(userId, typ)
	if sgc.CheckError(err) {
		return
	}
	evidenceKey := fmt.Sprintf("%s%d", evidence.EvidenceId, evidence.ExpiresAt.Unix())
	evidenceEncryptKey, err := codec.PrivateEncrypt(evidenceKey)
	if sgc.CheckError(err) {
		return
	}

	var evidenceVo vo.EvidenceVO
	evidenceVo.EvidenceId = evidenceKey + " (debug)"
	evidenceVo.EvidenceKey = evidenceEncryptKey
	evidenceVo.ExpiresAt = evidence.ExpiresAt.Unix()
	sgc.WriteData(&evidenceVo)

	log4g.Info("send evidence  id: %s", evidenceVo.EvidenceId)
	log4g.Debug("send evidence key: %s", evidenceVo.EvidenceKey)
}

func DiscardEvidence(c *gin.Context) {
	sgc := sg.Context(c)
	sgc.WriteSuccessOrError(dao.NewEvidenceDao().Discard(""))
}

