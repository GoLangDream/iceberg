package web

// func Route(name string, controllerHandler func(Controller)) {
// 	controller := createController(name)
// 	controllerHandler(controller)
// }

func Routes(block func(Router)) {
	router := Router{}
	block(router)
}
