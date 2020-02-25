package core

import (
	"fmt"
	"sync"
)

//自带锁，并发安全的Map
type CMap struct {
	contextMap map[string]interface{}
	*sync.RWMutex
}

func NewCMap() *CMap {
	return &CMap{
		contextMap: make(map[string]interface{}),
		RWMutex:    new(sync.RWMutex),
	}
}

/*
* 以key、value置入AppContext
 */
func (ctx *CMap) Set(key string, value interface{}) error {
	ctx.Lock()
	ctx.contextMap[key] = value
	ctx.Unlock()
	return nil
}

/*
* 读取指定key在AppContext中的内容
 */
func (ctx *CMap) Get(key string) (value interface{}, exists bool) {
	ctx.RLock()
	value, exists = ctx.contextMap[key]
	ctx.RUnlock()
	return value, exists
}

//remove item by gived key
//if not exists key, do nothing...
func (ctx *CMap) Remove(key string) {
	ctx.Lock()
	delete(ctx.contextMap, key)
	ctx.Unlock()
}

//get item by gived key, and remove it
//only can be read once, it will be locked
func (ctx *CMap) Once(key string) (value interface{}, exists bool) {
	ctx.Lock()
	defer ctx.Unlock()
	value, exists = ctx.contextMap[key]
	if exists {
		delete(ctx.contextMap, key)
	}
	return value, exists
}

/*
* 读取指定key在AppContext中的内容，以string格式输出
 */
func (ctx *CMap) GetString(key string) string {
	value, exists := ctx.Get(key)
	if !exists {
		return ""
	}
	return fmt.Sprint(value)
}

/*
* 读取指定key在AppContext中的内容，以int格式输出
 */
func (ctx *CMap) GetInt(key string) int {
	value, exists := ctx.Get(key)
	if !exists {
		return 0
	}
	return value.(int)
}

//check exists key
func (ctx *CMap) Exists(key string) bool {
	_, exists := ctx.contextMap[key]
	return exists
}

//get current map, returns map[string]interface{}
func (ctx *CMap) GetCurrentMap() map[string]interface{} {
	return ctx.contextMap
}

//get context length
func (ctx *CMap) Len() int {
	return len(ctx.contextMap)
}
