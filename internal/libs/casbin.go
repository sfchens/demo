package libs

import (
	"demo/global"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"sync"
)

var (
	once        sync.Once
	casbinLogic *CasbinLogic
)

type CasbinLogic struct {
	syncedEnforcer *casbin.SyncedEnforcer
}

func NewCasbinLogic() *CasbinLogic {
	return casbinLogic
}

func (c *CasbinLogic) init() {
	once.Do(func() {
		adapter, err := gormadapter.NewAdapterByDB(global.MysqlDB)
		if err != nil {
			panic(err)
		}

		//通过 keyMatch2 实现了 RESTful 路径的匹配
		//支持通配符 * 进行批量授权
		text := `
				[request_definition]
				r = sub, obj, act
		
				[policy_definition]
				p = sub, obj, act
		
				[role_definition]
				g = _, _
		
				[policy_effect]
				e = some(where (p.eft == allow))
		
				[matchers]
				m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
				`
		m, err := model.NewModelFromString(text)
		if err != nil {
			fmt.Printf("err1: %+v\n", err.Error())
			return
		}
		c.syncedEnforcer, err = casbin.NewSyncedEnforcer(m, adapter)
		if err != nil {
			fmt.Printf("err2: %+v\n", err.Error())
		}
	})

	c.syncedEnforcer.AddFunction("isAdmin", func(arguments ...interface{}) (interface{}, error) {
		// 获取用户名
		username := arguments[0].(string)
		// 检查用户名的角色是否为超级管理员
		return c.syncedEnforcer.HasRoleForUser(username, "role_1")
	})
	err := c.syncedEnforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
}

//Enforce 校验权限
func (c *CasbinLogic) Enforce(user, uri, action string) (bool, error) {
	return c.syncedEnforcer.Enforce(user, uri, action)
}

// AddPolicy 添加策略
func (c *CasbinLogic) AddPolicy(roleId int, uri, method string) (bool, error) {
	return c.syncedEnforcer.AddPolicy(c.MakeRoleName(roleId), uri, method)
}

// 拼接角色ID，为了防止角色与用户名冲突
func (c *CasbinLogic) MakeRoleName(roleId int) string {
	return fmt.Sprintf("role_%d", roleId)
}

// AddPolicies 批量添加策略
func (c *CasbinLogic) AddPolicies(rules [][]string) (bool, error) {
	return c.syncedEnforcer.AddPolicies(rules)
}

// DeleteRole 删除角色对应的用户和权限
func (c *CasbinLogic) DeleteRole(roleId int) (bool, error) {
	return c.syncedEnforcer.DeleteRole(c.MakeRoleName(roleId))
}

// DeleteRolePolicy 删除角色下的权限
func (c *CasbinLogic) DeleteRolePolicy(roleId int) (bool, error) {
	return c.syncedEnforcer.RemoveFilteredNamedPolicy("p", 0, c.MakeRoleName(roleId))
}

// DeleteRoleUser 删除添加用户
func (c *CasbinLogic) DeleteRoleUser(roleId int) (bool, error) {
	return c.syncedEnforcer.RemoveFilteredNamedGroupingPolicy("g", 1, c.MakeRoleName(roleId))
}

// AddUserRole 添加角色和用户对应关系
func (c *CasbinLogic) AddUserRole(uid string, roleId int) (bool, error) {
	return c.syncedEnforcer.AddGroupingPolicy(uid, c.MakeRoleName(roleId))
}

// AddUserRoles 批量添加角色和用户对应关联
func (c *CasbinLogic) AddUserRoles(uidS []string, roleIds []int) (bool, error) {
	rules := make([][]string, 0)
	for _, u := range uidS {
		for _, r := range roleIds {
			rules = append(rules, []string{u, c.MakeRoleName(r)})
		}
	}
	return c.syncedEnforcer.AddGroupingPolicies(rules)
}

// DeleteUserRole 删除用户的角色信息
func (c *CasbinLogic) DeleteUserRole(user string) (bool, error) {
	return c.syncedEnforcer.RemoveFilteredNamedGroupingPolicy("g", 0, user)
}
