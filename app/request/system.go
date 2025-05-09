package request

type AddRouteReq struct {
	Id     int    `form:"id" json:"id" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
	Url    string `form:"url" json:"url" binding:"required"`
	Status int    `json:"status" json:"status" binding:"required"`
	Type   int    `json:"type" json:"type" binding:"required"`
}

type ListRouteReq struct {
	PageInfo
	Name   string `json:"name" form:"name"`
	Type   int    `json:"type" form:"type"`
	Status int    `json:"status" json:"status"`
}

type ListRoleReq struct {
	PageInfo
	Name   string `json:"name" form:"name"`
	Type   int    `json:"type" form:"type"`
	Status int    `json:"status" json:"status"`
}

type AddOrEditRoleCasbinReq struct {
	UserIds []int  `json:"user_ids" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
}

type AddOrEditRolePermissionReq struct {
	Name     string `json:"name" form:"name" binding:"required"`
	RouteIds []int  `json:"route_ids" binding:"required"`
}

type MenuInfo struct {
	Id        int64      `json:"id"`
	Status    int64      `json:"status"`
	Type      string     `json:"type"`
	Path      string     `json:"path"`
	Name      string     `json:"name"`
	RouteName string     `json:"routeName"`
	Component string     `json:"component"`
	Meta      MenuMeta   `json:"meta"`
	ParentId  int64      `json:"parentId"`
	Children  []MenuInfo `json:"children"`
	Perm      string     `json:"perm"`
	CreatedAt int64      `json:"createdAt"`
}

type MenuMeta struct {
	Authority          []string `json:"authority"`
	AffixTab           int64    `json:"affixTab"`
	HideChildrenInMenu int64    `json:"hideChildrenInMenu"`
	HideInBreadcrumb   int64    `json:"hideInBreadcrumb"`
	HideInMenu         int64    `json:"hideInMenu"`
	HideInTab          int64    `json:"hideInTab"`
	Icon               string   `json:"icon"`
	KeepAlive          int64    `json:"keepAlive"`
	Sort               int64    `json:"sort"`
	Name               string   `json:"name"`
}

type SysMenuTreeReq struct {
	Keywords   string `json:"keywords" form:"keywords"`
	Status     int64  `json:"status" form:"status"`
	OnlyParent bool   `json:"onlyParent" form:"onlyParent"`
}

type AddUserReq struct {
	Username string `json:"username" binding:"required,max=16,min=2" label:"用户名"`
	Password string `json:"password" binding:"required,max=16,min=6" label:"密码"`
	Nickname string `json:"nickname" binding:"required,max=16,min=2" label:"昵称"`
	Status   int64  `json:"status" binding:"required" label:"状态"`
	RoleId   int64  `json:"roleId" binding:"required" label:"角色ID"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Remark   string `json:"remark"`
}

type UpsertApiReq struct {
	ParentId    int64  `json:"parentId"`
	Description string `json:"description"`
	Method      string `json:"method"`
	Path        string `json:"path"`
}

type ApiListReq struct {
	Page        int64  `json:"page" form:"page"`
	PageSize    int64  `json:"pageSize" form:"pageSize"`
	Description string `json:"description" form:"description"`
	Path        string `json:"path"  form:"path"`
	OnlyParent  bool   `json:"onlyParent" form:"onlyParent"`
}
type ApiListResp struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResp struct {
	Id          uint     `json:"id"`
	Password    string   `json:"password"`
	RealName    string   `json:"realName"`
	Roles       []string `json:"roles"`
	Username    string   `json:"username"`
	AccessToken string   `json:"accessToken"`
}

type LogoutReq struct {
	WithCredentials bool `json:"withCredentials"`
}

type DictListReq struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"pageSize" form:"pageSize"`
	DictName string `json:"dictName" form:"dictName"`
	DictType string `json:"dictType" form:"dictType"`
	Status   int64  `json:"status" form:"status"`
}
type DictListResp struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

type UpsertDictReq struct {
	DictName  string `json:"dictName"`
	DictType  string `json:"dictType"`
	ItemKey   string `json:"itemKey"`
	ItemValue string `json:"itemValue"`
	Sort      int64  `json:"sort"`
	Remark    string `json:"remark"`
	Status    int64  `json:"status"`
}

type MenuTreeReq struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"pageSize" form:"pageSize"`
	Name     string `json:"name" form:"name"`
	Status   int64  `json:"status" form:"status"`
}

type RecordListReq struct {
	Page       int64    `json:"page" form:"page"`
	PageSize   int64    `json:"pageSize" form:"pageSize"`
	Username   string   `json:"username" form:"username"`
	CreateTime []string `form:"createTime[]" json:"createTime[]"`
}
type RecordListResp struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

type RoleListReq struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"pageSize" form:"pageSize"`
	Name     string `json:"name" form:"name"`
	Status   int64  `json:"status" form:"status"`
}
type RoleListResp struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

type UpsertRoleReq struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Sort   int64  `json:"sort"`
	Remark string `json:"remark"`
	Status int64  `json:"status"`
}

type AssignRoleReq struct {
	AuthId []uint `json:"authId" binding:"required"`
	ApiId  []uint `json:"apiId" binding:"required"`
}

type UpsertUserReq struct {
	Username string `json:"username" binding:"required,max=16,min=2" label:"用户名"`
	Password string `json:"password" label:"密码"`
	Nickname string `json:"nickname" binding:"required,max=16,min=2" label:"昵称"`
	Status   int64  `json:"status" label:"状态"`
	RoleId   uint   `json:"roleId" binding:"required" label:"角色ID"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Remark   string `json:"remark"`
}

type UserListReq struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"pageSize" form:"pageSize"`
	Username string `json:"username" form:"username"`
	Status   int64  `json:"status" form:"status"`
}
type UserListResp struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

type ApiInfoResp struct {
	Id          int64          `json:"id"`
	ParentId    int64          `json:"parentId"`
	Description string         `json:"description"`
	Method      string         `json:"method"`
	Path        string         `json:"path"`
	CreatedAt   int64          `json:"createdAt"`
	Children    []*ApiInfoResp `json:"children"`
}

type DictInfoResp struct {
	Id        int64  `json:"id"`
	DictName  string `json:"dictName"`
	DictType  string `json:"dictType"`
	ItemKey   string `json:"itemKey"`
	ItemValue string `json:"itemValue"`
	Sort      int64  `json:"sort"`
	Remark    string `json:"remark"`
	Status    int64  `json:"status"`
	CreatedAt int64  `json:"createdAt"`
}

type MenuResp struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

type RouterMetaResp struct {
	Icon     string `json:"icon"`
	Sort     int64  `json:"sort,omitempty"`
	Title    string `json:"title"`
	AffixTab bool   `json:"affixTab,omitempty"`
}

type RouterResp struct {
	Meta      RouterMetaResp `json:"meta"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Component string         `json:"component,omitempty"`
	Children  []RouterResp   `json:"children,omitempty"`
}

type RecordInfoResp struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	UserId      int64  `json:"userId"`
	Description string `json:"description"`
	Method      string `json:"method"`
	Path        string `json:"path"`
	StatusCode  int64  `json:"statusCode"`
	Elapsed     string `json:"elapsed"`
	Msg         string `json:"msg"`
	Request     string `json:"request"`
	Response    string `json:"response"`
	Platform    string `json:"platform"`
	Ip          string `json:"ip"`
	Address     string `json:"address"`
	CreatedAt   int64  `json:"createdAt"`
}

type RoleInfoResp struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Sort      int64  `json:"sort"`
	Remark    string `json:"remark"`
	Status    int64  `json:"status"`
	CreatedAt int64  `json:"createdAt"`
	AuthId    []uint `json:"authId"`
	ApiId     []uint `json:"apiId"`
}

type UserInfoResp struct {
	Id            uint     `json:"id"`
	RealName      string   `json:"realName"`
	Roles         []string `json:"roles"`
	Username      string   `json:"username" label:"用户名"`
	Nickname      string   `json:"nickname" label:"用户昵称"`
	Mobile        string   `json:"mobile" label:"手机号"`
	Gender        int64    `json:"gender" label:"性别(1-男 2-女 0-保密)"`
	Email         string   `json:"email" label:"邮箱"`
	Avatar        string   `json:"avatar" label:"头像"`
	Status        int64    `json:"status" label:"状态 1:禁用,2正常"`
	Remark        string   `json:"remark" label:"备注"`
	LastLoginTime int64    `json:"lastLoginTime" label:"最后一次登录的时间"`
	LastLoginIp   string   `json:"lastLoginIp" label:"最后一次登录的IP"`
	RoleId        uint     `json:"roleId" label:"角色ID"`
	RoleName      string   `json:"roleName" label:"角色名称"`
	Permissions   []string `json:"permissions"`
	CreatedAt     int64    `json:"createdAt"`
}
