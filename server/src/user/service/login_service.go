package service

import (
	"common/mysqlx"
	"common/redisx"
	"github.com/carsonsx/log4g"
	"github.com/google/uuid"
	"strings"
	"time"
	"user/vo"
)

func SaveLoginSession(token string, wxsession *vo.WxappSession) error {
	if err := saveUserInfo(wxsession.Openid); err != nil {
		return err
	}
	return saveSessionInfo(token, wxsession)
}

func saveUserInfo(openId string) error {
	exist, err := mysqlx.Exists("select id from t_user where open_id=?", openId)
	if err != nil {
		log4g.Error(err)
		return err
	}
	if !exist {
		userId := strings.Replace(uuid.New().String(), "-", "", -1)
		err = mysqlx.Exec(nil, "insert into t_user (id,channel,open_id,insert_time) values (?,?,?,?)", userId, "weixin", openId, time.Now())
		if err != nil {
			return err
		}
	}
	return nil
}

func saveSessionInfo(token string, wxsession *vo.WxappSession) error {
	fields := make(map[string]interface{})
	fields["openid"] = wxsession.Openid
	fields["session_key"] = wxsession.Session_key
	fields["unionid"] = wxsession.Unionid
	fields["create_time"] = time.Now()
	return redisx.Client.HMSet("token_"+token, fields).Err()
}
