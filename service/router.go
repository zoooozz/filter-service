package service

import (
	"golang-kit/net/router"
)

func initInner(r *router.Router) {
	r.GuestPost("/x/internal/filter/keyword/add", internalAddkeyword)
	r.GuestPost("/x/internal/filter/keyword/state", internalUpdateStatekeyword)
	r.GuestPost("/x/internal/filter/keyword/edit", internalUpdateInfokeyword)
	r.GuestGet("/x/internal/filter/keyword/list", internalList)
	r.GuestGet("/x/internal/filter/business/list", internalBusinessList)

}

func initOutter(r *router.Router) {
	r.GuestPost("/x/outter/filter/list", list)
}
