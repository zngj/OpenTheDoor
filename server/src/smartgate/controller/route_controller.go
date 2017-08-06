package controller

import (
	"github.com/gin-gonic/gin"
	"usercenter/token"
	"common/cmnmsg"
	"smartgate/service"
	"smartgate/vo"
	"smartgate/codec"
)

func RouterStatus(c *gin.Context) {
	var rs vo.RouterStatusVO
	rs.Status = 0
	cmnmsg.WriteDataResponse(c.Writer, &rs)
}

func RouterEvidence(c *gin.Context) {
	userId, err := token.GetUserIdFromHeader(c.Request.Header)
	if err != nil {
		cmnmsg.WriteErrorResponse(c.Writer, err)
		return
	}
	evidence, err := service.CreateEvidence(userId)
	if err != nil {
		cmnmsg.WriteErrorResponse(c.Writer, err)
		return
	}

	evidenceKey, err := codec.Encrypt(evidence.EvidenceId)
	if err != nil {
		cmnmsg.WriteErrorResponse(c.Writer, err)
		return
	}

	var evidenceVo vo.EvidenceVO
	evidenceVo.EvidenceKey = evidenceKey
	evidenceVo.ExpiresAt = evidence.ExpiresAt.Unix()
	cmnmsg.WriteDataResponse(c.Writer, &evidenceVo)
}