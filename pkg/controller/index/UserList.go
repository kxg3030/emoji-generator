package index

import (
	"emoji/pkg/database"
	"emoji/pkg/model/entity"
	"emoji/pkg/model/logic"
	"emoji/pkg/model/service"
	"emoji/pkg/system"
	"emoji/pkg/unity"
	"emoji/pkg/validate"
	"github.com/gin-gonic/gin"
	"html"
)

type UserList struct {

}

func NewUserList()*UserList  {
	return &UserList{

	}
}

func (this *UserList)Login(ctx *gin.Context)  {
	var userList entity.UserList
	code := html.EscapeString(ctx.Query("code"))
	if validate.WhetherParamExist(code) {
		userList.NickName  = html.EscapeString(ctx.Query("nickName"))
		userList.Avatar    = html.EscapeString(ctx.Query("avatar"))
		userInfoFromWeChat:= service.NewWeChat().GetUserOpenId(code)
		if len(userInfoFromWeChat.OpenId) > 0{
			result:= logic.NewUserListLogic(database.GetOrm()).FindUserRecord(userInfoFromWeChat.OpenId);
			ctx.Writer.Header().Set("token",unity.GetToken("openId",userInfoFromWeChat.OpenId))
			if len(result) != 0{
				go func() {
					logic.NewUserListLogic(database.GetOrm()).UpdateUserColumn(userInfoFromWeChat.OpenId,userList)
				}()
				system.PrintSuccess(ctx,202,"",map[string]interface{}{})
				return
			}
			if logic.NewUserListLogic(database.GetOrm()).InsertUserRecord(userList){
				system.PrintSuccess(ctx,202,"",map[string]interface{}{})
			}else{
				system.PrintException(ctx,105,"",map[string]interface{}{})
			}
		}else{
			system.PrintException(ctx,109,"",map[string]interface{}{})
		}
	}else{
		system.PrintException(ctx,110,"",map[string]interface{}{})
	}
}