package controller

import (
	"common/vo"
	"github.com/gin-gonic/gin"
	"common/codec"
	"smartgate/service"
	"usercenter/token"
	"fmt"
	"github.com/carsonsx/log4g"
	"common/sg"
)

func GetUserId(c *gin.Context) (userId string, err error) {
	accessToken := c.Request.Header.Get(token.HEADER_ACCESS_TOKEN)
	return token.GetUserId(accessToken)
}

func RouterStatus(c *gin.Context) {
	sgc := sg.Context(c)
	userId, err := GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	var rs vo.RouterStatusVO
	rs.Status, err = service.GetRouterStatus(userId)
	sgc.WriteDataOrError(&rs, err)
}

func RouterEvidenceIn(c *gin.Context) {
	_routerEvidence(c, service.GATE_DIRECTION_IN)

}

func RouterEvidenceOut(c *gin.Context) {
	_routerEvidence(c, service.GATE_DIRECTION_OUT)
}

func RouterEvidence(c *gin.Context) {
	_routerEvidence(c, 2)
}

func _routerEvidence(c *gin.Context, typ int8) {
	sgc := sg.Context(c)
	userId, err := GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	evidence, err := service.CreateEvidence(userId, typ)
	if sgc.CheckError(err) {
		return
	}
	evidenceKey := fmt.Sprintf("%s%d",evidence.EvidenceId,evidence.ExpiresAt.Unix())
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
