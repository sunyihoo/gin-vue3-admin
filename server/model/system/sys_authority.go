package system

import "time"

type SysAuthority struct {
	CreatedAt       time.Time       // 创建时间
	UpdatedAt       time.Time       // 更新时间
	DeletedAt       *time.Time      `sql:"index"`
	AuthorityId     uint            `json:"authorityId" gorm:"not nil;unique;primary_key"` // 角色ID
	AuthorityName   string          `json:"authorityName" gorm:"comment:角色名"`              // 角色名
	ParentId        *uint           `json:"parentId" gorm:"comment:父角色ID"`                 // 父角色ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	SysBaseMens     []SysBaseMenu   `json:"menus" gorm:"many2many:sys_user_authority_menus"`
}
