// Package service ...
//
// Author: go_developer@163.com<张德满>
//
// File:  project.go
//
// Description: project 管理内存的项目信息
//
// Date: 2020/10/13 12:37 下午
package service

import (
	"sync"

	"github.com/go-developer/ginx/define"
)

var (
	// Project 项目数据实例
	//
	// Author : go_developer@163.com<张德满>
	Project *project
)

func init() {
	Project = &project{
		lock:         &sync.RWMutex{},
		projectTable: make(map[string]*define.Project),
	}
	if err := Project.init(); nil != err {
		panic(err.Error())
	}
}

type project struct {
	lock         *sync.RWMutex
	projectTable map[string]*define.Project
}

//  Init 初始化项目信息 ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/10/16 16:15:49
func (p *project) init() error {
	return nil
}

// loadProject 加载项目信息...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/10/16 16:18:39
func (p *project) loadProject() error {
	return nil
}
