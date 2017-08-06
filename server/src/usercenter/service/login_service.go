package service

import (
	"common/mysqlx"
	"common/redisx"
	"database/sql"
	"github.com/carsonsx/log4g"
	"github.com/google/uuid"
	"strings"
	"time"
	"usercenter/util"
	"usercenter/vo"
	"usercenter/token"
)

func SaveLoginSession(token string, wxsession *vo.WxappSession) error {
	userId, err := saveUserInfo(wxsession.Openid)
	if err != nil {
		return err
	}
	return saveSessionInfo(token, userId, wxsession)
}

func saveUserInfo(openId string) (userId string, err error) {

	var rows *sql.Rows
	rows, err = mysqlx.GetDBConn().Query("select id from uc_user_info where open_id=?", openId)
	if err != nil {
		log4g.Error(err)
		return
	}
	defer rows.Close()

	if rows.Next() {
		var id *string
		err = rows.Scan(&id)
		if err != nil {
			log4g.Error(err)
			return "", err
		}
		if id != nil {
			userId = *id
		}
	}

	if userId == "" {
		_userId := strings.Replace(uuid.New().String(), "-", "", -1)
		err = mysqlx.Exec(nil, "insert into uc_user_info (id,channel,open_id,insert_time) values (?,?,?,?)", _userId, "weixin", openId, time.Now())
		if err != nil {
			return
		}
		userId = _userId
	}

	return
}

func saveSessionInfo(accessToken, userId string, wxsession *vo.WxappSession) error {

	now := time.Now()

	var rows *sql.Rows
	rows, err := mysqlx.GetDBConn().Query("select id, access_token from uc_login_log where status = ? and user_id=?", "1", userId)
	if err != nil {
		log4g.Error(err)
		return err
	}
	defer rows.Close()

	if rows.Next() {
		var id *int64
		var token *string
		err = rows.Scan(&id, &token)
		if err != nil {
			log4g.Error(err)
			return err
		}
		if *token != "" {
			err = redisx.Client.Del(util.GetAccessTokenKey(*token)).Err()
			if err != nil {
				log4g.Error(err)
				return err
			}
			mysqlx.Exec(nil, "update uc_login_log set release_time=?,status=? where id=?", now, "0", *id)
		}
	}

	expiresAt := now.Add(time.Duration(wxsession.ExpiresIn) * time.Second)

	err = mysqlx.Exec(nil, "insert into uc_login_log (user_id,access_token,login_time,expires_in,expires_at,status) values (?,?,?,?,?,?)",
		userId, accessToken, now, wxsession.ExpiresIn, expiresAt, "1")
	if err != nil {
		log4g.Error(err)
		return err
	}

	token.Save(userId, accessToken, wxsession.Session_key, wxsession.Unionid, now, expiresAt)

	return err
}
