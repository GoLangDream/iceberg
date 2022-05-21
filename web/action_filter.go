package web

type actionFilter struct {
	filter func()
	only   []string
	except []string
}

func createActionFilter(filter func(), options ...map[string][]string) *actionFilter {
	actionFilter := &actionFilter{
		filter: filter,
		only:   nil,
		except: nil,
	}
	switch len(options) {
	case 0:
		return actionFilter
	case 1:
		if only, ok := options[0]["only"]; ok {
			actionFilter.only = only
		}
		if except, ok := options[0]["except"]; ok {
			actionFilter.except = except
		}
	}
	return nil
}
