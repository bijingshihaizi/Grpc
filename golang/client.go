package main

import (
	"fmt"
	"github.com/goGrpc/user_protoc"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
)

func main() {
	// 连接服务端接口
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := user_protoc.NewUserServiceClient(conn)

//# GetUserAll接口
	replyGetUserAll, errGetUserAll := client.GetUserAll(context.Background(), &user_protoc.Params{
		Condition:[]*user_protoc.Condition{{ParentId:"1",Id:"2"}},
		Orders:[]*user_protoc.Params_Orders{{Orders:""}},
		Start:1,
		Limit:10})
	if errGetUserAll != nil {
		log.Fatal(errGetUserAll)
	}
	fmt.Println(replyGetUserAll)

//# Count接口 	success
	//replyCount, errCount := client.Count(context.Background(), &user_protoc.Condition{ParentId:"1",Id:"2"})
	//if errCount != nil {
	//	log.Fatal(errCount)
	//}
	//fmt.Println(replyCount)

//# CheckUserPhone 接口
	//replyCheckUserPhone, errCheckUserPhone := client.CheckUserPhone(context.Background(), &user_protoc.Add{
	//	CheckedRole:[]*user_protoc.CheckedRole{Num:16},
	//	Code:1,
	//	Mobile:"15711258458",
	//	Password:"123456",
	//	Username:"test",
	//})
	//if errCheckUserPhone != nil {
	//	log.Fatal(errCheckUserPhone)
	//}
	//fmt.Println(replyCheckUserPhone)

//# AddUser 接口
	//replyAddUser, errAddUser := client.AddUser(context.Background(), &user_protoc.AddData{
	//	AddData_Data:[]*user_protoc.AddData_Data{Create_time:16,Password:"",Reg_user_id:"",Reg_ip:"",Username:"",Real_name:"",Mobile:"",Parent_id:""},
	//	Role_id:"",
	//})
	//if errAddUser != nil {
	//	log.Fatal(errAddUser)
	//}
	//fmt.Println(replyAddUser)

//# getUserInfo 接口
	//replyGetUserInfo, errGetUserInfo := client.GetUserInfo(context.Background(), &user_protoc.UserId{Id:1})
	//if errGetUserInfo != nil {
	//	log.Fatal(errGetUserInfo)
	//}
	//fmt.Println(replyGetUserInfo)

//# UpdateUser 接口
	//replyUpdateUser, errUpdateUser := client.UpdateUser(context.Background(), &user_protoc.UpdateDatas{Id:[]*user_protoc.UserId{Id:187},UpdateData:[]*user_protoc.Add{CheckedRole:[]*user_protoc.CheckedRole{Num:16},Code:1,Mobile:"15711258458",Password:"123456",Username:"test"}})
	//if errUpdateUser != nil {
	//	log.Fatal(errUpdateUser)
	//}
	//fmt.Println(replyUpdateUser)

//# DeleteUser 接口  	success
//	replyDeleteUser, errDeleteUser := client.DeleteUser(context.Background(), &user_protoc.UserId{Id:187})
//	if errDeleteUser != nil {
//		log.Fatal(errDeleteUser)
//	}
//	fmt.Println(replyDeleteUser)

//# IsBind 接口 success
//	replyIsBind, errIsBind := client.IsBind(context.Background(), &user_protoc.Open{Id:"o1Lxo5CHeyjy1Yy5xAphdPDMyRX4"})
//	if errIsBind != nil {
//		log.Fatal(errIsBind)
//	}
//	fmt.Println(replyIsBind)

//# GetOpenUser 接口
//	replyGetOpenUser, errGetOpenUser := client.GetOpenUser(context.Background(), &user_protoc.Key{Key:"8a408d7a7fa25fe832821e8b23494b89",AppId:"2CCPRRGGCQ"})
//	if replyGetOpenUser != nil {
//		log.Fatal(replyGetOpenUser)
//	}
//	fmt.Println(replyGetOpenUser)

//# GetAllAcl 接口 返回数据过长可能有中文
//	replyGetAllAcl, errGetAllAcl := client.GetAllAcl(context.Background(), &user_protoc.Num{Count:"1"})
//	if errGetAllAcl != nil {
//		log.Fatal(errGetAllAcl)
//	}
//	fmt.Println(replyGetAllAcl)

//# GetAllAcl 接口 返回数据过长可能有中文
//	replyGetAllRole, errGetAllRole := client.GetAllRole(context.Background(), &user_protoc.Open{Id:"1"})
//	if errGetAllRole != nil {
//		log.Fatal(errGetAllRole)
//	}
//	fmt.Println(replyGetAllRole)

//# AddRoleAcl 接口		success
//	replyAddRoleAcl, errAddRoleAcl := client.AddRoleAcl(context.Background(), &user_protoc.AddRole{RoleName:"测试账号",RoleDesc:"测试描述",AclId:"1, 2, 3, 4",CreateUserId:"39",UserParentId:"39"})
//	if errAddRoleAcl != nil {
//		log.Fatal(errAddRoleAcl)
//	}
//	fmt.Println(replyAddRoleAcl)

//# UpdateRoleAcl 接口		success
//	replyUpdateRoleAcl, errUpdateRoleAcl := client.UpdateRoleAcl(context.Background(), &user_protoc.UpdataRole{
//		Id:193,
//		Update:[]*user_protoc.UpdataRole_Update{{Id:193,AclId:"2"}},
//	})
//	if errUpdateRoleAcl != nil {
//		log.Fatal(errUpdateRoleAcl)
//	}
//	fmt.Println(replyUpdateRoleAcl)

//# DelRoleAcl 接口 		success
//	replyDelRoleAcl, errDelRoleAcl := client.DelRoleAcl(context.Background(), &user_protoc.UserId{Id:193})
//	if errDelRoleAcl != nil {
//		log.Fatal(errDelRoleAcl)
//	}
//	fmt.Println(replyDelRoleAcl)

//# GetBindUser 接口		编码问题 有返回数据
//	replyGetBindUser, errGetBindUser := client.GetBindUser(context.Background(), &user_protoc.BindOpenId{BindOpenid:"o1Lxo5CHeyjy1Yy5xAphdPDMyRX4"})
//	if errGetBindUser != nil {
//		log.Fatal(errGetBindUser)
//	}
//	fmt.Println(replyGetBindUser)

//# UpdateUserByOpenId 接口		success
//	replyUpdateUserByOpenId, errUpdateUserByOpenId := client.UpdateUserByOpenId(context.Background(), &user_protoc.OpenData{OpenId:"o1Lxo5CHeyjy1Yy5xAphdPDMyRX4",BindPlateNo:"1-京A66666,1-13333333335-1583807660"})
//	if errUpdateUserByOpenId != nil {
//		log.Fatal(errUpdateUserByOpenId)
//	}
//	fmt.Println(replyUpdateUserByOpenId)

//# UpdateBind 接口				success
//	replyUpdateBind, errUpdateBind := client.UpdateBind(context.Background(), &user_protoc.UpdateBindData{UserId:"63",BindType:3,BindOpenid:"a33b4c5d6e7f,1",BindName:"",BindAvatar:"C",BindMobile:"15711258458",CreateTime:1598522555,IsDelete:1})
//	if errUpdateBind != nil {
//		log.Fatal(errUpdateBind)
//	}
//	fmt.Println(replyUpdateBind)

//# AddCustUser 接口		MYSQL报错
//	replyAddCustUser, errAddCustUser := client.AddCustUser(context.Background(), &user_protoc.CustUser{
//		CustUser:[]*user_protoc.UpdateBindData{{UserId:"12",BindType:3,BindOpenid:"a33b4c5d6e7f,1",BindName:"",BindAvatar:"C",BindMobile:"15711258458",CreateTime:1598522555,IsDelete:0}},
//	})
//	if errAddCustUser != nil {
//		log.Fatal(errAddCustUser)
//	}
//	fmt.Println(replyAddCustUser)

//# GetList 接口		success
//	replyGetList, errGetList := client.GetList(context.Background(), &user_protoc.WalletCondition{IsDelete:0,PlateNo:"京A66666,1",WxOpenid:"o1Lxo5CHeyjy1Yy5xAphdPDMyRX4",Field:"*",Order:"id",Limit:"0",Start:"0"})
//	if errGetList != nil {
//		log.Fatal(errGetList)
//	}
//	fmt.Println(replyGetList)

//# SaveData 接口		success
//	replySaveData, errSaveData := client.SaveData(context.Background(), &user_protoc.SaveDatas{DeviceNo:"A1:1B:1C:1D:1E:1F",DeviceType:"1",IsDelete:0,PlateNo:"京A66666,1",WalletOpenid:"a11b1c1d1e1f,1-19",MyOpenid:"",WxOpenid:"o1Lxo5CHeyjy1Yy5xAphdPDMyRX4",IsDefault:"1"})
//	if errSaveData != nil {
//		log.Fatal(errSaveData)
//	}
//	fmt.Println(replySaveData)

//# UpdateData 接口		success
//	replyUpdateData, errUpdateData := client.UpdateData(context.Background(), &user_protoc.WalletUpdate{
//		Condition:[]*user_protoc.WalletUpdate_Condition{{DeviceNo:"A1:1B:1C:1D:1E:1F",DeviceType:"2",WxOpenid:"o1Lxo5CHeyjy1Yy5xAphdPDMyRX4"}},
//		Data:[]*user_protoc.WalletUpdate_Data{{Data:1,WalletOpenid:"a11b1c1d1e1f,1-19"}},
//	})
//	if errUpdateData != nil {
//		log.Fatal(errUpdateData)
//	}
//	fmt.Println(replyUpdateData)
}
