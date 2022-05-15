package web

type Application interface {
	RouterDraw(*Router)
	HomePath() string
}
