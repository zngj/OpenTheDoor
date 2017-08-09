package controller

import (
	"common/errcode"
	"github.com/gin-gonic/gin"
	"smartgate/codec"
	"smartgate/service"
	"smartgate/vo"
	"usercenter/token"
)

func RouterStatus(c *gin.Context) {
	var rs vo.RouterStatusVO
	rs.Status = 0
	errcode.WriteDataResponse(c.Writer, &rs)
}

func RouterEvidenceIn(c *gin.Context) {
	_routerEvidence(c, 0)

}

func RouterEvidenceOut(c *gin.Context) {
	_routerEvidence(c, 1)
}

func RouterEvidence(c *gin.Context) {
	_routerEvidence(c, 2)
}

func _routerEvidence(c *gin.Context, typ int8) {
	userId, err := token.GetUserIdFromHeader(c.Request.Header)
	if err != nil {
		errcode.WriteErrorResponse(c.Writer, err)
		return
	}
	evidence, err := service.CreateEvidence(userId, typ)
	if err != nil {
		errcode.WriteErrorResponse(c.Writer, err)
		return
	}

	evidenceKey, err := codec.Encrypt(evidence.EvidenceId)
	if err != nil {
		errcode.WriteErrorResponse(c.Writer, err)
		return
	}

	var evidenceVo vo.EvidenceVO
	evidenceVo.EvidenceKey = evidenceKey
	evidenceVo.ExpiresAt = evidence.ExpiresAt.Unix()
	errcode.WriteDataResponse(c.Writer, &evidenceVo)
}
