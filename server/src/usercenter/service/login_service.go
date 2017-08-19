package service

import (
	"common/redisx"
	"common/sqlx"
	"common/vo"
	"github.com/carsonsx/log4g"
	"github.com/google/uuid"
	"strings"
	"time"
	"common/tokenutil"
)


func SaveLoginSession(token ,userId string, expiresIn int) error {
	now := time.Now()
	err := AddLoginLog(token, userId, now, expiresIn)
	if err != nil {
		return err
	}
	return SaveToken(token, userId, "self", now, expiresIn)
}

func SaveWxappLoginSession(token string, wxsession *vo.WxSession) error {
	now := time.Now()
	userId, err := saveUserInfo(wxsession.Openid, now)
	if err != nil {
		return err
	}
	err = AddLoginLog(token, userId, now, wxsession.ExpiresIn)
	if err != nil {
		return err
	}
	return SaveWxToken(token, userId, wxsession, now)
}

func saveUserInfo(openId string, now time.Time) (userId string, err error) {
	dao := sqlx.NewDao()
	err = dao.Query("select id from uc_user_info where open_id=?", openId).Scan(&userId)
	if err == sqlx.ErrNotFound {
		userId = strings.Replace(uuid.New().String(), "-", "", -1)
		err = dao.Exec("insert into uc_user_info (id,channel,open_id,insert_time) values (?,?,?,?)", userId, "weixin", openId, now)
	}
	return
}

func AddLoginLog(accessToken, userId string, now time.Time, expiresIn int) (err error) {

	dao := sqlx.NewDao()

	var _id *int64
	var _token *string
	err = dao.Query("select id, access_token from uc_login_log where status = ? and user_id=?", "1", userId).Scan(&_id, &_token)
	if err != nil && err != sqlx.ErrNotFound {
		return
	}

	dao.BeginTx()
	defer func() { dao.CommitTx(err) }()

	if err == nil { // 上次登录存在
		err = redisx.Client.Del(tokenutil.GetAccessTokenKey(*_token)).Err()
		if err != nil {
			log4g.Error(err)
			return
		}
		err = dao.Exec("update uc_login_log set release_time=?,status=? where id=?", now, "0", *_id)
	}

	expiresAt := now.Add(time.Duration(expiresIn)*time.Second)
	err = dao.Exec("insert into uc_login_log (user_id,access_token,login_time,expires_in,expires_at,status) values (?,?,?,?,?,?)",
		userId, accessToken, now, expiresIn, expiresAt, "1")
	if err != nil {
		return
	}


	return
}
