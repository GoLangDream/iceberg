package web

import "github.com/GoLangDream/rgo/rstring"

func (c *BaseController) beforeActionFilter() {
	c.callBeforeAction()
}

func (c *BaseController) afterActionFilter() {
	c.render()
}

func (c *BaseController) callBeforeAction() {
	filters, ok := c.beforeActions[c.actionName]
	if ok {
		for _, filter := range filters {
			filter()
		}
	}
}

func (c *BaseController) callAfterAction() {
	filters, ok := c.afterActions[c.actionName]
	if ok {
		for _, filter := range filters {
			filter()
		}
	}
}

func (c *BaseController) BeforeActon(actionName string, filterFunc func()) {
	name := rstring.Underscore(actionName)
	filters, ok := c.beforeActions[name]
	if !ok {
		filters = []func(){}
	}
	c.beforeActions[name] = append(filters, filterFunc)
}

func (c *BaseController) AfterActon(actionName string, filterFunc func()) {
	name := rstring.Underscore(actionName)
	filters, ok := c.afterActions[name]
	if !ok {
		filters = []func(){}
	}
	c.afterActions[name] = append(filters, filterFunc)
}
