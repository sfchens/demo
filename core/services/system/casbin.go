package system

import (
	"demo/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"sync"
)

type CasbinService struct{}

func NewCasbin() *CasbinService {
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
		p = sub, obj, act, eft
		 
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
