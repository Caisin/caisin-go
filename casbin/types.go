package casbin

type Policy struct {
	Sub         string `json:"sub"`         //用户名或角色
	ControlType string `json:"controlType"` //控制类型
	Res         string `json:"res"`         //资源地址
	Action      string `json:"action"`      //访问方式
}
