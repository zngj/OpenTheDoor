package service

import (
	"common/redisx"
	"github.com/carsonsx/log4g"
	"github.com/google/uuid"
	"strings"
	"time"
	"usercenter/util"
	"common/vo"
	"usercenter/token"
	"common/dbx"
)

func SaveLoginSession(token string, wxsession *vo.WxappSession) error {
	userId, err := saveUserInfo(wxsession.Openid)
	if err != nil {
		return err
	}
	return saveSessionInfo(token, userId, wxsession)
}

func saveUserInfo(openId string) (userId string, err error) {
	dao := dbx.NewDao()
	err = dao.Query("select id from uc_user_info where open_id=?", openId).One(&userId)
	if err == dbx.ErrNotFound {
		userId = strings.Replace(uuid.New().String(), "-", "", -1)
		err = dao.Exec("insert into uc_user_info (id,channel,open_id,insert_time) values (?,?,?,?)", userId, "weixin", openId, time.Now())
	}
	return
}

func saveSessionInfo(accessToken, userId string, wxsession *vo.WxappSession) (err error) {

	now := time.Now()

	dao := dbx.NewDao()

	var _id *int64
	var _token *string
	err = dao.Query("select id, access_token from uc_login_log where status = ? and user_id=?", "1", userId).One(&_id, &_token)
	if err != nil && err != dbx.ErrNotFound {
		return
	}

	dao.BeginTx()
	defer func() {dao.CommitTx(err)}()

	if err == nil { // 上次登录存在
		err = redisx.Client.Del(util.GetAccessTokenKey(*_token)).Err()
		if err != nil {
			log4g.Error(err)
			return
		}
		err = dao.Exec("update uc_login_log set release_time=?,status=? where id=?", now, "0", *_id)
	}

	expiresAt := now.Add(time.Duration(wxsession.ExpiresIn) * time.Second).AddDate(0,0,1)
	err = dao.Exec("insert into uc_login_log (user_id,access_token,login_time,expires_in,expires_at,status) values (?,?,?,?,?,?)",
		userId, accessToken, now, wxsession.ExpiresIn, expiresAt, "1")
	if err != nil {
		return
	}

	err = token.Save(userId, accessToken, wxsession.Session_key, wxsession.Unionid, now, expiresAt)

	return
}
