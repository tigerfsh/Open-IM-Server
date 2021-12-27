package friend

import (
	api "Open_IM/pkg/base_info"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/token_verify"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/friend"
	"Open_IM/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AddBlacklist(c *gin.Context) {
	params := api.AddBlacklistReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		log.NewError("0", "BindJSON failed ", err.Error())
		return
	}
	req := &rpc.AddBlacklistReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(params.OperationID, "AddBlacklist args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.AddBlacklist(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "AddBlacklist failed ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call add blacklist rpc server failed"})
		return
	}
	resp := api.AddBlacklistResp{CommResp: api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg}}
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.CommID.OperationID, "AddBlacklist api return ", resp)
}

func ImportFriend(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		log.NewError("0", "BindJSON failed ", err.Error())
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, params)
	var ok bool
	ok, req.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.OperationID, "ImportFriend args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.ImportFriend(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "ImportFriend failed", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "cImportFriend failed " + err.Error()})
		return
	}

	resp := api.ImportFriendResp{CommResp: api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg}, Data: RpcResp.FailedFriendUserIDList}
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.OperationID, "ImportFriend api return ", resp)
}

func AddFriend(c *gin.Context) {
	params := api.AddFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.AddFriendReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo("AddFriend args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.AddFriend(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "AddFriend failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call AddFriend rpc server failed"})
		return
	}

	resp := api.AddFriendResp{CommResp: api.CommResp{ErrCode: RpcResp.ErrCode, ErrMsg: RpcResp.ErrMsg}}
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.CommID.OperationID, "AddFriend api return ", resp)
}

func AddFriendResponse(c *gin.Context) {
	params := api.AddFriendResponseReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.AddFriendResponseReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	utils.CopyStructFields(&req, params)
	log.NewInfo(req.CommID.OperationID, "AddFriendResponse args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.AddFriendResponse(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "AddFriendResponse failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call add_friend_response rpc server failed"})
		return
	}

	resp := api.AddFriendResponseResp{CommResp: api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg}}
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.CommID.OperationID, "AddFriendResponse api return ", resp)
}

func DeleteFriend(c *gin.Context) {
	params := api.DeleteFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.DeleteFriendReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "DeleteFriend args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.DeleteFriend(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "DeleteFriend failed ", err, req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call delete_friend rpc server failed"})
		return
	}

	resp := api.DeleteFriendResp{CommResp: api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg}}
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.CommID.OperationID, "DeleteFriend api return ", resp)
}

func GetBlacklist(c *gin.Context) {
	params := api.GetBlackListReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.GetBlacklistReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "GetBlacklist args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.GetBlacklist(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "GetBlacklist failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call get blacklist rpc server failed"})
		return
	}

	resp := api.GetBlackListResp{CommResp: api.CommResp{ErrCode: RpcResp.ErrCode, ErrMsg: RpcResp.ErrMsg}}
	utils.CopyStructFields(&resp.BlackUserInfoList, RpcResp.BlackUserInfoList)
	log.NewInfo(req.CommID.OperationID, "GetBlacklist api return ", resp)

}

func SetFriendComment(c *gin.Context) {
	params := api.SetFriendCommentReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.SetFriendCommentReq{}
	utils.CopyStructFields(req.CommID, params)
	req.Remark = params.Remark
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "SetFriendComment args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.SetFriendComment(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "SetFriendComment failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call set friend comment rpc server failed"})
		return
	}
	resp := api.SetFriendCommentResp{CommResp: api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg}}
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.CommID.OperationID, "SetFriendComment api return ", resp)
}

func RemoveBlacklist(c *gin.Context) {
	params := api.RemoveBlackListReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.RemoveBlacklistReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}

	log.NewInfo(req.CommID.OperationID, "RemoveBlacklist args ", req.String())
	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.RemoveBlacklist(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "RemoveBlacklist failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call remove blacklist rpc server failed"})
		return
	}
	resp := api.RemoveBlackListResp{CommResp: api.CommResp{ErrCode: RpcResp.CommonResp.ErrCode, ErrMsg: RpcResp.CommonResp.ErrMsg}}
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.CommID.OperationID, "RemoveBlacklist api return ", resp)
}

func IsFriend(c *gin.Context) {
	params := api.IsFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.IsFriendReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "IsFriend args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.IsFriend(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "IsFriend failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call add friend rpc server failed"})
		return
	}
	resp := api.IsFriendResp{CommResp: api.CommResp{ErrCode: RpcResp.ErrCode, ErrMsg: RpcResp.ErrMsg}, Response: RpcResp.Response}
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.CommID.OperationID, "IsFriend api return ", resp)
}

//
//func GetFriendsInfo(c *gin.Context) {
//	params := api.GetFriendsInfoReq{}
//	if err := c.BindJSON(&params); err != nil {
//		log.NewError("0", "BindJSON failed ", err.Error())
//		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
//		return
//	}
//	req := &rpc.GetFriendsInfoReq{}
//	utils.CopyStructFields(req.CommID, params)
//	var ok bool
//	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
//	if !ok {
//		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
//		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
//		return
//	}
//	log.NewInfo(req.CommID.OperationID, "GetFriendsInfo args ", req.String())
//
//	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
//	client := rpc.NewFriendClient(etcdConn)
//	RpcResp, err := client.GetFriendsInfo(context.Background(), req)
//	if err != nil {
//		log.NewError(req.CommID.OperationID, "GetFriendsInfo failed ", err.Error(), req.String())
//		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call search friend rpc server failed"})
//		return
//	}
//
//	resp := api.GetFriendsInfoResp{CommResp:api.CommResp{ErrCode: RpcResp.ErrCode, ErrMsg: RpcResp.ErrMsg}}
//	utils.CopyStructFields(&resp, RpcResp)
//	c.JSON(http.StatusOK, resp)
//	log.NewInfo(req.CommID.OperationID, "GetFriendsInfo api return ", resp)
//}

func GetFriendList(c *gin.Context) {
	params := api.GetFriendListReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.GetFriendListReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "GetFriendList args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.GetFriendList(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "GetFriendList failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call get friend list rpc server failed"})
		return
	}

	resp := api.GetFriendListResp{CommResp: api.CommResp{ErrCode: RpcResp.ErrCode, ErrMsg: RpcResp.ErrMsg}}
	utils.CopyStructFields(&resp, RpcResp)
	log.NewInfo(req.CommID.OperationID, "GetFriendList api return ", resp)
}

func GetFriendApplyList(c *gin.Context) {
	params := api.GetFriendApplyListReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.GetFriendApplyListReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "GetFriendApplyList args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)

	RpcResp, err := client.GetFriendApplyList(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "GetFriendApplyList failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call get friend apply list rpc server failed"})
		return
	}

	resp := api.GetFriendApplyListResp{CommResp: api.CommResp{ErrCode: RpcResp.ErrCode, ErrMsg: RpcResp.ErrMsg}}
	utils.CopyStructFields(&resp, RpcResp)
	log.NewInfo(req.CommID.OperationID, "GetFriendApplyList api return ", resp)

}

func GetSelfApplyList(c *gin.Context) {
	params := api.GetSelfApplyListReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.GetSelfApplyListReq{}
	utils.CopyStructFields(req.CommID, params)
	var ok bool
	ok, req.CommID.OpUserID = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"))
	if !ok {
		log.NewError(req.CommID.OperationID, "GetUserIDFromToken false ", c.Request.Header.Get("token"))
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "GetUserIDFromToken failed"})
		return
	}
	log.NewInfo(req.CommID.OperationID, "GetSelfApplyList args ", req.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImFriendName)
	client := rpc.NewFriendClient(etcdConn)
	RpcResp, err := client.GetSelfApplyList(context.Background(), req)
	if err != nil {
		log.NewError(req.CommID.OperationID, "GetSelfApplyList failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "call get self apply list rpc server failed"})
		return
	}

	resp := api.GetSelfApplyListResp{CommResp: api.CommResp{ErrCode: RpcResp.ErrCode, ErrMsg: RpcResp.ErrMsg}}
	utils.CopyStructFields(resp, RpcResp)
	c.JSON(http.StatusOK, resp)
	log.NewInfo(req.CommID.OperationID, "GetSelfApplyList api return ", resp)

}