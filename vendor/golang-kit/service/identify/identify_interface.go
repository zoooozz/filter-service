package identify

import (
	"golang-kit/net/context"
)

type Indentify interface {
	Login(c context.Context) (err error)
	Verify(c context.Context) (err error)
	IsLogin(c context.Context) (err error)
	Logout(c context.Context) (err error)
}
