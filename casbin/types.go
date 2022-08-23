package casbin

type Policy struct {
	Sub         string //用户名或角色
	ControlType string //控制类型
	Res         string //资源地址
	Action      string //访问方式
}
