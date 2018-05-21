package model

import (
	"github.com/robfig/cron"
	"sync"
)

// 管理所有用户的定时任务
type CronMap struct {
	Cmap map[int64]*cron.Cron
	sync sync.Mutex
}

func (c *CronMap) Set(id int64, cr *cron.Cron) {
	if len(c.Cmap) == 0 || c.Cmap == nil {
		c.Cmap = make(map[int64]*cron.Cron)
	}
	c.sync.Lock()
	c.Cmap[id] = cr
	c.sync.Unlock()
}

func (c *CronMap) Del(id int64) {
	if _, ok := c.Cmap[id]; ok {
		// 先停止，后删除
		c.Stop(id)
		c.sync.Lock()
		delete(c.Cmap, id)
		c.sync.Unlock()
	}
}

func (c *CronMap) Stop(id int64) {
	if cron, ok := c.Cmap[id]; ok {
		cron.Stop()
	}
}

func (c *CronMap) Start(id int64) {
	if cron, ok := c.Cmap[id]; ok {
		cron.Start()
	}
}
