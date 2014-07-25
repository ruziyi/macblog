package index

import(
	"github.com/Unknwon/macaron"
    "github.com/macaron-contrib/session"
)

func Index(ctx *macaron.Context, sess session.SessionStore) {

    ctx.Data["hi"] = "hello world!"
    ctx.Data["UserName"] = sess.Get("Uname")

    ctx.HTML(200, "index/index", ctx.Data)
}
