package user_api

import (
	"gin-vue/api/utils"
	"gin-vue/global"
	"gin-vue/modles/form"
	"gin-vue/modles/models"
	"gin-vue/modles/res"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (UsersApi) UserCount(c *gin.Context) {
	var user []models.User
	rows := global.DB.Find(&user).RowsAffected
	res.OkWithData(rows, c)
}

func (UsersApi) PasswordModify(c *gin.Context) {
	var myUser models.User

	result := global.DB.First(&myUser, "name", "admin")
	if result.Error != nil {
		res.FailWithMsg("failed to find user", c)
		return
	}

	userName := "admin"
	password := c.PostForm("password")
	userRole := "1"
	myNewUser := models.User{
		Name:     userName,
		Password: password,
		Role:     userRole,
	}

	global.DB.Model(myUser).Updates(myNewUser)

	res.OkWithData(myNewUser, c)
}

func (UsersApi) AndroidResetPassword(c *gin.Context) {
	var user models.User
	userName := c.PostForm("username")
	global.DB.Where("name =?", userName).First(&user)
	if user.Name == "" {
		res.FailWithMsg("User.NotExist", c)
		return
	}
	if user.Status != "启用" {
		res.FailWithMsg("User.Disable", c)
		return
	}
	password := c.PostForm("oldpassword")
	if user.Password != password {
		res.FailWithMsg("User.PasswordIsWrong", c)
		return
	}
	newPassword := c.PostForm("newpassword")
	myNewUser := models.User{
		Password: newPassword,
	}
	global.DB.Model(user).Updates(myNewUser)
	res.OkWithData(userName, c)
}

func (UsersApi) UserAdd(c *gin.Context) {
	var myUser models.User
	var form form.AddUserForm
	if err := c.ShouldBind(&form); err != nil {
		global.Logger.Printf("add user failed:%v\n", err.Error())
		res.FailWithMsg("Common.InvalidParam", c)
		return
	}
	userName := form.Name
	password := form.Password
	userRole := form.Role
	status := form.Status

	global.DB.Where("name = ?", userName).First(&myUser)

	if myUser.Name != "" {
		res.FailWithMsg("User.Exists", c)
		return
	}
	myUser.ID = models.NewUUID()
	myUser.Name = userName
	myUser.Password = password
	myUser.Role = userRole
	myUser.Status = status
	global.DB.Create(&myUser)
	global.Logger.Println("user:", myUser)

	res.OkWithData(myUser, c)
}

func (UsersApi) UserUpdate(c *gin.Context) {
	var myUser models.User
	id := c.Param("id")
	result := global.DB.First(&myUser, "id", id)
	if result.Error != nil {
		res.FailWithMsg("User.NotExist", c)
		return
	}
	var form form.UpdateUserForm

	if err := c.ShouldBind(&form); err != nil {
		global.Logger.Printf("update user failed:%v\n", err.Error())
		res.FailWithMsg("Common.InvalidParam", c)
		return
	}

	password := form.Password
	userRole := form.Role
	status := form.Status

	if myUser.Name == "vmadmin" && userRole == "云桌面用户" {
		res.FailWithMsg("Common.InvalidParam", c)
		return
	}

	if myUser.Name == "vmadmin" && status == "禁用" {
		res.FailWithMsg("Common.InvalidParam", c)
		return
	}

	myNewUser := models.User{
		Name:     myUser.Name,
		Password: password,
		Role:     userRole,
		Status:   status,
	}

	global.DB.Model(myUser).Updates(myNewUser)
	global.Logger.Println("user:", myNewUser)
	res.OkWithData(myNewUser, c)
}

func (UsersApi) UserUpdateSelfPassword(c *gin.Context) {
	var myUser models.User
	userId, err := c.Cookie("userId")
	if err != nil {
		res.FailWithMsg("Common.InvalidParam", c)
		return
	}
	result := global.DB.First(&myUser, "id", userId)
	if result.Error != nil {
		res.FailWithMsg("User.NotExist", c)
		return
	}

	password := c.PostForm("password")

	myNewUser := models.User{
		Name:     myUser.Name,
		Password: password,
	}

	global.DB.Model(myUser).Updates(myNewUser)
	res.OkWithData(myNewUser, c)
}

func (UsersApi) UserGet(c *gin.Context) {
	var users []models.User
	var totalNum int64
	count, _ := strconv.Atoi(c.Query("count"))
	pageNum, _ := strconv.Atoi(c.Query("index"))
	name := c.Query("name")
	keyword := "%" + name + "%"
	offset := (pageNum - 1) * count
	if name == "" {
		global.DB.Model(&models.User{}).Count(&totalNum)
		global.DB.Limit(count).Offset(offset).Find(&users)
	} else {
		global.DB.Model(&models.User{}).Where("name LIKE ?", keyword).Count(&totalNum)
		global.DB.Limit(count).Offset(offset).Where("name LIKE ?", keyword).Find(&users)
	}
	data := make(map[string]interface{})
	data["users"] = users
	data["totalNum"] = totalNum
	res.OkWithData(data, c)
}

func (UsersApi) UserAllCountGet(c *gin.Context) {
	var users []models.User
	global.DB.Find(&users)
	data := make(map[string]interface{})
	data["num"] = len(users)
	res.OkWithData(data, c)
}

func (UsersApi) UserProfileGet(c *gin.Context) {
	userId, err := c.Cookie("userId")
	if err != nil {
		res.FailWithMsg("Common.InvalidParam", c)
		return
	}
	var user models.User
	global.DB.Where("id = ?", userId).First(&user)
	data := &res.GetUserProfileReponse{
		UserName: user.Name,
		Password: user.Password,
	}
	res.OkWithData(data, c)
}

func (UsersApi) UserBindDevices(c *gin.Context) {

	var user models.User
	var form form.BindUserForm
	if err := c.ShouldBind(&form); err != nil {
		global.Logger.Printf("user bind device failed:%v\n", err.Error())
		res.FailWithMsg("Common.InvalidParam", c)
		return
	}
	userId := form.UserId
	deviceName := form.StrDeviceName
	deviceNames := strings.Split(deviceName, ",")
	global.DB.Where("id = ?", userId).First(&user)
	if user.Name == "" {
		res.FailWithMsg("User.NotExist", c)
		return
	}

	for _, value := range deviceNames {
		device := utils.GetDeviceByName(value)
		if device == nil {
			continue
		}
		var template models.Template
		global.DB.Where("id =?", device.TemplateId).First(&template)
		err := utils.CreateVMLocalUser(user.Name, user.Password, value, template.UserName, template.UserPwd)
		if err != nil {
			continue
		}
		myBind := &models.Bind{
			ID:       models.NewUUID(),
			DeviceId: device.ID,
			UserId:   userId,
		}
		global.DB.Create(&myBind)
	}

	res.OkWithData(nil, c)
}

func (UsersApi) UserDel(c *gin.Context) {
	var myUser models.User
	var myBinds []*models.Bind
	id := c.Param("id")
	result := global.DB.First(&myUser, "id", id)
	if result.Error != nil {
		res.FailWithMsg("User.NotExist", c)
		return
	}

	global.DB.Find(&myBinds, "user_id", id)
	if len(myBinds) > 0 {
		res.FailWithMsg("User.HasBind", c)
		return
	}

	result = global.DB.Delete(&myUser)
	global.Logger.Println("userDel error:", result.Error)
	if result.Error != nil {
		res.FailWithMsg("User.NotExist", c)
		return
	}

	global.DB.Where("user_id = ?", myUser.ID).Delete(&models.Bind{})

	res.OkWithData(myUser, c)
}

func (UsersApi) UserLogin(c *gin.Context) {

	global.Logger.Println("UserAgent:", c.Request.UserAgent())

	var user models.User
	var token models.Token
	userName := c.PostForm("username")
	pwd := c.PostForm("password")
	global.DB.Where("name = ?", userName).First(&user)

	requestFrom := c.Request.UserAgent()
	if strings.Contains(requestFrom, "Windows") && user.Role == "云桌面用户" {
		res.FailWithMsg("User.NoPermission", c)
		return
	}

	if user.Password != pwd {
		res.FailWithMsg("User.PasswordIsWrong", c)
		return
	}

	if user.Status != "启用" {
		res.FailWithMsg("User.Disable", c)
		return
	}
	tokenValue := models.NewUUID()

	utils.UpdateOrCreateToken(user.ID, tokenValue)
	c.SetCookie("accessToken", tokenValue, 3600, "/", "", false, true)
	c.SetCookie("userId", user.ID, 3600, "/", "", false, true)
	global.DB.Where("user_id=?", user.ID).First(&token)

	res.Ok(token, "登录成功", c)
}

func (UsersApi) UserLogout(c *gin.Context) {

	userId, err := c.Cookie("userId")
	if err != nil {
		res.FailWithMsg("User.LogoutFailed", c)
		return
	}
	global.DB.Where("user_id = ?", userId).Delete(&models.Token{})
	res.Ok(userId, "退出成功", c)
}
