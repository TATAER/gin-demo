package Entity

import (
	"ginApi/Dao"
)

type WechatUsers struct {
	Id        int    `json:"id"`
	NickName  string `json:"nick_name"`
	HeaderImg string `json:"header_img"`
	Phone     string `json:"phone"`
}

type WechatUserDetail struct {
	WechatUsers
	Level string `json:"level"`
}

/**
获取所有用的列表

*/
func GetUserList() ([]WechatUsers, error) {
	list := []WechatUsers{}
	err := Dao.DB.Find(&list).Error
	return list, err
}

/**
获取用户的详情
*/
func GetUserDetailById(id int) (WechatUsers, error) {
	var wc WechatUsers
	err := Dao.DB.Where("id = ?", id).First(&wc).Error
	return wc, err
}

/**
更新用户信息
*/
func UpdateDetailById(data WechatUsers) error {
	err := Dao.DB.Table("wechat_users").Where("id = ?", data.Id).Update(data).Error
	return err
}

func Auth(where map[string]interface{}) (int, bool) {
	var wc WechatUsers
	err := Dao.DB.Table("wechat_users").Where(where).Find(&wc).Error
	if err != nil {
		return 0, false
	}
	return wc.Id, true
}
