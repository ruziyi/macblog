package user

import(
	"github.com/Unknwon/macaron"
    "github.com/macaron-contrib/session"
//    "github.com/macaron-contrib/binding"

    "fmt"
)

// 登陆页
func Login(ctx *macaron.Context){
    ctx.Data["name"] = "dada"
    ctx.HTML(200, "user/login", ctx.Data)
}

// 登陆
func LoginPost(ctx *macaron.Context, sess session.SessionStore, f *session.Flash){

    ctx.Req.ParseForm()
    userName := ctx.Query("username")
    password := ctx.Query("password")
    pasord := ctx.Query("pasord")

    if userName == "weisd" && password == "sdfsdf" {
        sess.Set("Uname", userName)
        sess.Set("Uid", 1)
        ctx.Redirect("/")
    }

    f.Error("用户名，密码错误")

    ctx.Redirect("/login")


    fmt.Println("auth", pasord)

    fmt.Println(ctx.Req.Form)

}

// 退出登陆
func Logout(ctx *macaron.Context, sess session.SessionStore){
    sess.Delete("Uname")
    sess.Delete("Uid")

    ctx.Redirect("/")
}


// 判断是否已经登陆
func CheckLogin(ctx *macaron.Context, sess session.SessionStore){

    data := sess.Get("uid")
    if uid, ok := data.(int64); !ok || uid == 0{
        ctx.Redirect("/login", 302)
    }
}


/*
// 用户登陆表单

type Form interface {
    Name(field string) string
}

type UserLoginForm struct {
    UserName string `form:"username" binding:"Required;MaxSize(35)"`
    Password string `form:"password" binding:"Required;MinSize(6);MaxSize(30)"`
    Remember bool   `form:"remember"`
}

func (f *UserLoginForm) Name(field string) string {
    names := map[string]string{
        "UserName": "UserName",
        "Password": "Password",
    }

    return names[field]
}

// 用户登陆表单 验证
func (f *UserLoginForm) Validate(errs *binding.Errors, req *http.Request, ctx macaron.Context){
    

}
*/
