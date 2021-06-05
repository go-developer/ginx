// Package define ...
//
// Author: go_developer@163.com<白茶清欢>
//
// File:  core.go
//
// Description: core 定义核心的数据结构
//
// Date: 2020/10/13 12:38 下午
package define

import "sync"

// Project 定义内存项目的数据结构
//
// Author : go_developer@163.com<白茶清欢>
//
// Date : 12:39 下午 2020/10/13
type Project struct {
	Lock   *sync.Mutex // 数据锁
	Flag   string      `json:"flag"`   // 项目标示
	Domain string      `json:"domain"` // 项目域名
	Port   string      `json:"port"`   // 项目端口
}

// Route 定义内存路由的数据结构
//
// Author : go_developer@163.com<白茶清欢>
//
// Date : 12:40 下午 2020/10/13
type Route struct {
	URI        string        `json:"uri"`         // uri 地址(网关对外暴露的地址)
	Method     string        `json:"method"`      // 请求方法(调用网关接口的请求方式)
	FilterList []interface{} `json:"filter_list"` // 启用的过滤器 TODO: 待设计
	Backend    *BackendAPI   `json:"backend"`     // 代理到后端api的真实配置
}

// BackendAPI 后端真实api的配置
//
// Author : go_developer@163.com<白茶清欢>
//
// Date : 12:43 下午 2020/10/13
type BackendAPI struct {
	Scheme     string            `json:"scheme"`      // 请求方法(get/post)
	Method     string            `json:"method"`      // 请求方法
	Param      []*Param          `json:"param"`       // 参数列表
	OutputRule map[string]string `json:"output_rule"` // 数据输出的规则
}

// Param ...
//
// Author : go_developer@163.com<白茶清欢>
//
// Date : 2020/10/16 15:31:23
type Param struct {
	Name                string                       `json:"name"`                   // 参数名
	BindRouteParam      bool                         `json:"bind_route_param"`       // 是否绑定在路由上的参数 /test/:id/:name
	CheckFunc           func(data interface{}) error `json:"check_func"`             // 参数的校验函数
	IsRequired          bool                         `json:"is_required"`            // 是否必传
	DefaultValue        interface{}                  `json:"default_value"`          // 非比传时候的默认值
	MapBackendParamName string                       `json:"map_backend_param_name"` // 映射到后端接口的参数名
}
