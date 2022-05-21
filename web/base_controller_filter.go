package web

func (c *BaseController) beforeActionFilter() {
	c.callBeforeAction()
}

func (c *BaseController) afterActionFilter() {
	c.render()
	c.callAfterAction()
}

func (c *BaseController) callBeforeAction() {
	c.callActionFilter(c.beforeActions)
}

func (c *BaseController) callAfterAction() {
	c.callActionFilter(c.afterActions)
}

func (c *BaseController) callActionFilter(actionFilters []*actionFilter) {
	for _, actionFilter := range actionFilters {
		if actionFilter.only != nil {
			for _, actionName := range actionFilter.only {
				if actionName == c.actionName {
					actionFilter.filter()
				}
			}
			continue
		}
		if actionFilter.except != nil {
			for _, actionName := range actionFilter.except {
				if actionName == c.actionName {
					continue
				}
			}
		}
		actionFilter.filter()
	}
}

func (c *BaseController) BeforeAction(filter func(), options ...map[string][]string) {
	beforeActions := createActionFilter(filter, options...)
	c.beforeActions = append(c.beforeActions, beforeActions)
}

func (c *BaseController) AfterAction(filter func(), options ...map[string][]string) {
	afterActions := createActionFilter(filter, options...)
	c.beforeActions = append(c.beforeActions, afterActions)
}
