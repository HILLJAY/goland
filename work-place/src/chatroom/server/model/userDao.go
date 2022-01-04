package model

import (
	"encoding/json"
	"work-place/src/github.com/garyburd/redisgo/redis"
)

var SingleUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

/**
通过工厂模式获取UserDao示例
*/
func UserDaoFactory(pool *redis.Pool) *UserDao {

	userDao := &UserDao{pool: pool}
	return userDao
}

/**
根据UserId查询数据库中用户对象
*/
func (userDao *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {

	res, err := conn.Do("hget", "users", id)

	if err != nil {

		err = ERROR_USER_NOTEXISTS
		return
	}

	err = json.Unmarshal([]byte(res.(string)), user)
	if err != nil {
		return
	}

	return
}

/**
根据Id和密码进行登陆验证
*/
func (userDao *UserDao) ChargeLogin(userId int, userPwd string) (user *User, err error) {
	//从连接池获取一根连接
	conn := userDao.pool.Get()

	user, err = userDao.GetUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {

		err = ERROR_USER_PWD
		return
	}

	return
}
