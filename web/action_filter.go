package web

import (
	"github.com/GoLangDream/rgo/rstring"
	"github.com/thoas/go-funk"
)

type AFH = map[string][]string

type actionFilter struct {
	filter func()
	only   []string
	except []string
}

func createActionFilter(filter func(), options ...AFH) *actionFilter {
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
			actionFilter.only = funk.Map(only, func(item string) string {
				return rstring.Underscore(item)
			}).([]string)
		}
		if except, ok := options[0]["except"]; ok {
			actionFilter.except = funk.Map(except, func(item string) string {
				return rstring.Underscore(item)
			}).([]string)
		}
		return actionFilter
	}
	return nil
}
