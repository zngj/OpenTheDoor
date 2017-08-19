package controller

import (
	"common/model"
	"common/sg"
	"common/vo"
	"github.com/gin-gonic/gin"
	"etransin/dao"
	"etransin/service"
	"common/sgconst"
	"common/tokenutil"
	"strconv"
	"github.com/carsonsx/log4g"
)

func RouterStatus(c *gin.Context) {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	var rs vo.RouterStatusVO
	rs.Status, err = service.GetRouterStatus(userId)
	sgc.WriteDataOrError(&rs, err)
}

func RouterInList(c *gin.Context)  {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	var routers []*model.RouterInfo
	if sgc.CheckError(dao.NewRouterDao().FindIn(userId, &routers)) {
		return
	}
	sgc.WriteData(convertRouters(routers))
}

func RouterOutList(c *gin.Context)  {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	var routers []*model.RouterInfo
	if sgc.CheckError(dao.NewRouterDao().FindOut(userId, &routers)) {
		return
	}
	sgc.WriteData(convertRouters(routers))
}


func MyRouters(c *gin.Context)  {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	strLastId := c.Query("last_id")
	log4g.Debug("query last router id %s", strLastId)
	var lastId int64
	if strLastId != "" {
		lastId, err = strconv.ParseInt(strLastId, 10, 64)
	}
	var routers []*model.RouterInfo
	if sgc.CheckError(dao.NewRouterDao().FindByUser(userId, lastId, 10, &routers)) {
		return
	}
	sgc.WriteData(convertRouters(routers))
}

func convertRouters(routers []*model.RouterInfo) []*vo.RouterInfoVO {
	routersVo := make([]*vo.RouterInfoVO, len(routers))
	for i := range routersVo {
		routersVo[i] = convertRouterInfo(routers[i])
	}
	return routersVo
}

func convertRouterInfo(m *model.RouterInfo) *vo.RouterInfoVO {
	v := new(vo.RouterInfoVO)
	v.Id = m.Id
	v.UserId = m.UserId
	v.AtDate = m.AtDate.Unix()
	v.InStationId = m.InStationId.String()
	v.InStationName = m.InStationName.String()
	v.InGateId = m.InGateId.String()
	if m.InTime != nil {
		v.InTime = m.InTime.Unix()
	}
	v.OutStationId = m.OutStationId.String()
	v.OutStationName = m.OutStationName.String()
	v.OutGateId = m.OutGateId.String()
	if m.OutTime != nil {
		v.OutTime = m.OutTime.Unix()
	}
	v.Status = m.Status.Int8()
	v.StatusName = sgconst.GetRouterStatusString(m.Status.Int8())
	v.Money = m.Money.Float32()
	return v
}

