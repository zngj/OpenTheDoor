package controller

import (
	"common/errcode"
	"common/vo"
	"github.com/gin-gonic/gin"
	"net/http"
	"common/codec"
	"smartgate/service"
	"usercenter/token"
	"fmt"
)

func GetUserId(header http.Header) (userId string, err error) {
	accessToken := header.Get(token.HEADER_ACCESS_TOKEN)
	return token.GetUserId(accessToken)
}

func RouterStatus(c *gin.Context) {
	var rs vo.RouterStatusVO
	rs.Status = 0
	errcode.WriteDataResponse(c.Writer, &rs)
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
	userId, err := GetUserId(c.Request.Header)
	if err != nil {
		errcode.WriteErrorResponse(c.Writer, err)
		return
	}
	evidence, err := service.CreateEvidence(userId, typ)
	if err != nil {
		errcode.WriteErrorResponse(c.Writer, err)
		return
	}
	evidenceKey := fmt.Sprintf("%s%d",evidence.EvidenceId,evidence.ExpiresAt.Unix())
	evidenceEncryptKey, err := codec.PrivateEncrypt(evidenceKey)
	if err != nil {
		errcode.WriteErrorResponse(c.Writer, err)
		return
	}

	var evidenceVo vo.EvidenceVO
	evidenceVo.EvidenceId = evidenceKey + " (debug)"
	evidenceVo.EvidenceKey = evidenceEncryptKey
	evidenceVo.ExpiresAt = evidence.ExpiresAt.Unix()
	errcode.WriteDataResponse(c.Writer, &evidenceVo)
}
