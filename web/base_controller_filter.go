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
OUTER:
	for _, actionFilter := range actionFilters {
		if actionFilter.only != nil {
			for _, actionName := range actionFilter.only {
				if actionName == c.actionName {
					actionFilter.filter()
					continue OUTER
				}
			}

		}
		if actionFilter.except != nil {
			for _, actionName := range actionFilter.except {
				if actionName == c.actionName {
					continue OUTER
				}
			}
			actionFilter.filter()
			continue OUTER
		}

		if actionFilter.only == nil && actionFilter.except != nil {
			actionFilter.filter()
		}

	}
}

func (c *BaseController) BeforeAction(filter func(), options ...AFH) {
	beforeAction := createActionFilter(filter, options...)
	if beforeAction != nil {
		c.beforeActions = append(c.beforeActions, beforeAction)
	}
}

func (c *BaseController) AfterAction(filter func(), options ...AFH) {
	afterAction := createActionFilter(filter, options...)
	if afterAction != nil {
		c.beforeActions = append(c.beforeActions, afterAction)
	}
}
