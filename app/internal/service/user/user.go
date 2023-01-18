package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"oj/app/define"
	g "oj/app/global"
	"oj/app/internal/model"
	"strconv"
	"sync"
	"time"
)

type sUser struct{}

var (
	onceUser sync.Once
	insUser  *sUser
)

func NewUserService() *sUser {
	onceUser.Do(func() {
		insUser = &sUser{}
	})
	return insUser
}

func (s *sUser) CheckAdmin(id int) error {
	//goland:noinspection SqlResolve
	sqlStr := "select admin from user where id = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[CheckAdmin]  prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	var user model.User
	err = stmt.QueryRow(id).Scan(&user.Admin)
	if err != nil {
		g.Logger.Errorf("[CheckAdmin]  QueryRow failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	if user.Admin != define.Admin {
		return fmt.Errorf("not have this root")
	}
	return nil
}

func (s *sUser) ChangeUserAdmin(id int) error {
	//goland:noinspection SqlResolve
	sqlStr := "update user set admin= ? where id = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[ChangeUserAdmin]  prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(define.Admin, id)
	if err != nil {
		g.Logger.Errorf("[ChangeUserAdmin]  update failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}

func (s *sUser) CancelUserAdmin(id int) error {
	//goland:noinspection SqlResolve
	sqlStr := "update user set admin= ? where id = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[CancelUserAdmin]  prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(define.CancelAdmin, id)
	if err != nil {
		g.Logger.Errorf("[CancelUserAdmin]  update failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}

func (s *sUser) CheckUserIsExist(username string) error {
	//goland:noinspection SqlResolve
	sqlStr := "select id,identity,name,password,salt,created_at from user where name = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[CheckUserIsExist]  prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	var user model.User
	row := stmt.QueryRow(username)
	err = row.Scan(&user.Id, &user.Identity, &user.Name, &user.Password, &user.Salt, &user.CreatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil
		}
		g.Logger.Errorf("[CheckUserIsExist] queryrow failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return fmt.Errorf("username already exist")
}

func (s *sUser) EncryptPassword(salt, pwd string) string {
	m5 := md5.New()
	m5.Write([]byte(pwd))
	m5.Write([]byte(salt))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}

func (s *sUser) CreateUser(user model.User) error {
	//goland:noinspection SqlResolve
	sqlStr := "insert into user(identity,name,password,salt,created_at) values(?,?,?,?,?)"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[CreateUser]  prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(&user.Identity, &user.Name, &user.Password, &user.Salt, &user.CreatedAt)
	if err != nil {
		g.Logger.Errorf("[CreateUser]  insert failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}

func (s *sUser) CheckPassword(user *model.User) error {
	//goland:noinspection SqlResolve
	sqlStr := "select password,salt from user where name = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[CheckPassword]  prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	var sqlUser model.User
	err = stmt.QueryRow(user.Name).Scan(&sqlUser.Password, &sqlUser.Salt)
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("invalid username or password")
		}
		g.Logger.Errorf("[CheckPassword]  prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	pwd := s.EncryptPassword(sqlUser.Salt, user.Password)
	if pwd != sqlUser.Password {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}

func (s *sUser) GetUserByUserName(userName string) (*model.User, error) {
	//goland:noinspection SqlResolve
	sqlStr := "select id,identity,phone,mall,created_at from user where name = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[GetUserByUserName]  prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	defer stmt.Close()
	var user model.User
	err = stmt.QueryRow(userName).Scan(&user.Id, &user.Identity, &user.Phone, &user.Mall, &user.CreatedAt)
	if err != nil {
		g.Logger.Errorf("[GetUserByUserName]  prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	return &user, nil

}

func (s *sUser) GenerateToken(ctx context.Context, userSubject *model.User) (string, error) {
	user, err := s.GetUserByUserName(userSubject.Name)
	if err != nil {
		return "", err
	}
	return NewToken(user, ctx), nil
}

// NewToken 根据信息创建token
func NewToken(u *model.User, ctx context.Context) string {
	jwtConfig := g.Config.Auth.Jwt
	expiresTime := time.Now().Unix() + jwtConfig.ExpiresTime
	g.Logger.Info("expiresTime: %v\n", expiresTime)
	id := u.Id
	g.Logger.Info("id: %v\n", strconv.Itoa(id))
	claims := jwt.StandardClaims{
		Audience:  u.Name,
		ExpiresAt: expiresTime,
		Id:        strconv.Itoa(id),
		IssuedAt:  time.Now().Unix(),
		Issuer:    jwtConfig.Issuer,
		NotBefore: time.Now().Unix(),
	}
	var jwtSecret = []byte(jwtConfig.SecretKey)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		println("generate token success!\n")
		err = g.DbVerify.Set(ctx,
			fmt.Sprintf("jwt_%d", u.Id),
			token,
			time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
		if err != nil {
			g.Logger.Errorf("set [jwt] cache failed, %v", err)
			return "fail"
		}
		return token
	} else {
		println("generate token fail\n")
		return "fail"
	}
}

func (s *sUser) UUid() string {
	return uuid.NewV4().String()
}
