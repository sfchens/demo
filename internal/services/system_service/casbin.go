package system_service

import (
	"demo/app/request"
	"demo/global"
	"demo/internal/models"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"strconv"
	"sync"
)

type CasbinService struct{}

func NewCasbinLogic() *CasbinService {
	return &CasbinService{}
}

//func (casbinService *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
//	authorityId := strconv.Itoa(int(AuthorityID))
//	casbinService.ClearCasbin(0, authorityId)
//	rules := [][]string{}
//	for _, v := range casbinInfos {
//		rules = append(rules, []string{authorityId, v.Path, v.Method})
//	}
//	e := casbinService.Casbin()
//	success, _ := e.AddPolicies(rules)
//	if !success {
//		return errors.New("存在相同api,添加失败,请联系管理员")
//	}
//	return nil
//}

//func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
//	err := global.INTRA_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
//		"v1": newPath,
//		"v2": newMethod,
//	}).Error
//	return err
//}

//func (casbinService *CasbinService) GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []request.CasbinInfo) {
//	e := casbinService.Casbin()
//	authorityId := strconv.Itoa(int(AuthorityID))
//	list := e.GetFilteredPolicy(0, authorityId)
//	for _, v := range list {
//		pathMaps = append(pathMaps, request.CasbinInfo{
//			Path:   v[1],
//			Method: v[2],
//		})
//	}
//	return pathMaps
//}

func (s *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := s.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (s *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormAdapter.NewAdapterByDB(global.DB)
		text := `
		[request_definition]
		r = sub, obj, act
		 
		[policy_definition]
		p = sub, obj, act, eft, type, name
		 
		[role_definition]
		g = _,_
		 
		[policy_effect]
		e = some(where (p.eft == allow))
		 
		[matchers]
		m = (g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act) || g(r.sub,"superAdmin")
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			global.GetZapLog().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedEnforcer, err = casbin.NewSyncedEnforcer(m, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

func (s *CasbinService) Enforce(val ...interface{}) (bool, error) {
	return s.Casbin().Enforce(val...)
}

func (s *CasbinService) ListRole(input request.ListRoleReq) (out request.PageResult, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
	)

	var db = global.DB
	err = db.Count(&out.Total).Error
	if err != nil {
		return
	}

	var listTmp []models.CasbinRule
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Scan(&listTmp).Error
	if err != nil {
		return
	}
	return
}

// AddOrEditUserRole 添加用户角色
func (s *CasbinService) AddOrEditUserRole(input request.AddOrEditRoleCasbinReq) (err error) {
	var (
		name    = input.Name
		userIds = input.UserIds
	)
	if len(userIds) == 0 {
		err = errors.New("params exception")
		return
	}

	db := global.DB.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()

	for _, userId := range userIds {
		var userModel models.SysUsers
		err = db.Where("id=?", userId).First(&userModel).Error
		if err != nil {
			return
		}

		if userModel.Status != 1 {
			err = errors.New("user disable")
			return
		}

		_, err = global.Casbin.AddRoleForUser(userModel.Username, name)

		if err != nil {
			return
		}
	}
	return
}

func (s *CasbinService) DeleteRoleForUser(userId int, name string) (err error) {

	var userCount int64
	err = global.DB.Where("id=?", userId).Count(&userCount).Error
	if err != nil {
		return
	}
	_, _ = global.Casbin.DeleteRoleForUser(strconv.Itoa(userId), name)
	return
}

func (s *CasbinService) AddOrEditRolePermission(input request.AddOrEditRolePermissionReq) (err error) {
	//var (
	//	routeIds = input.RouteIds
	//	name     = input.Name
	//)
	//
	//db := global.DB.Begin()
	//defer func() {
	//	if err != nil {
	//		db.Rollback()
	//	} else {
	//		db.Commit()
	//	}
	//}()
	//
	//for _, routeId := range routeIds {
	//	var routeModel models.SystemRoute
	//	err = global.DB.Where("id=?", routeId).First(&routeModel).Error
	//	if err != nil {
	//		return
	//	}
	//
	//	if routeModel.Status != 1 {
	//		err = errors.New(fmt.Sprintf("route %v disable", routeModel.Url))
	//		return
	//	}
	//	_, err = global.Casbin.AddPolicy(name, routeModel.Url, routeModel.Method)
	//	if err != nil {
	//		return
	//	}
	//}

	return
}
