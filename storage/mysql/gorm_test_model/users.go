package gorm_test_model

// 用户表
type Users struct {
	ID        int32   `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`             // 主键
	Name      *string `gorm:"column:name;type:varchar(255);is null;" json:"name"`                    // 名称
	Age       int32   `gorm:"column:age;type:int(11);not null;DEFAULT '18'" json:"age"`              // 年龄
	CardNo    string  `gorm:"column:card_no;type:varchar(18);not null;DEFAULT ''" json:"card_no"`    // 身份证
	HeadImg   string  `gorm:"column:head_img;type:varchar(255);not null;DEFAULT ''" json:"head_img"` // 头像
	CreatedAt int32   `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`            // 创建时间
	UpdatedAt int32   `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`            // 更新时间
	DeletedAt int32   `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"` // 删除时间
}

// 获取结构体对应的表名方法
func (m *Users) TableName() string {
	return "users"
}

// 实例化结构体对象
func NewUsers() *Users {
	return &Users{}
}

// 表字段的映射
var UsersColumns = struct {
	ID        string
	Name      string
	Age       string
	CardNo    string
	HeadImg   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	Age:       "age",
	CardNo:    "card_no",
	HeadImg:   "head_img",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// 包含所有表字段的切片
var UsersAllColumns = []string{
	UsersColumns.ID,        // 主键
	UsersColumns.Name,      // 名称
	UsersColumns.Age,       // 年龄
	UsersColumns.CardNo,    // 身份证
	UsersColumns.HeadImg,   // 头像
	UsersColumns.CreatedAt, // 创建时间
	UsersColumns.UpdatedAt, // 更新时间
	UsersColumns.DeletedAt, // 删除时间

}

// 设置：主键
func (m *Users) SetID(v int32) {
	m.ID = v
}

// 设置：名称
func (m *Users) SetName(v *string) {
	m.Name = v
}

// 设置：年龄
func (m *Users) SetAge(v int32) {
	m.Age = v
}

// 设置：身份证
func (m *Users) SetCardNo(v string) {
	m.CardNo = v
}

// 设置：头像
func (m *Users) SetHeadImg(v string) {
	m.HeadImg = v
}

// 设置：创建时间
func (m *Users) SetCreatedAt(v int32) {
	m.CreatedAt = v
}

// 设置：更新时间
func (m *Users) SetUpdatedAt(v int32) {
	m.UpdatedAt = v
}

// 设置：删除时间
func (m *Users) SetDeletedAt(v int32) {
	m.DeletedAt = v
}
