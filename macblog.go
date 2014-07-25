package main

import(
    "runtime"
//    "os"

    /*
	"github.com/codegangsta/cli"
    "github.com/weisd/macblog/cmd"
    */

	"github.com/Unknwon/macaron"
    "github.com/macaron-contrib/session"
    "github.com/macaron-contrib/pongo2"

    "github.com/weisd/macblog/routers"
    "github.com/weisd/macblog/routers/user"
    "github.com/weisd/macblog/routers/index"
)

const APP_VER = "0.0.1"

func init(){
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){

    /*
    app := cli.NewApp()
    app.Name = "macblog"
    app.Usage = "macblog for weisd"
    app.Version = APP_VER
    app.Commands = []cli.Command{
        cmd.CmdWeb,
    }

    app.Flags = append(app.Flags, []cli.Flag{}...)
    app.Run(os.Args)
    */

    routers.GlobalInti()

    m := macaron.Classic()

    m.Use(session.Sessioner())
    m.Use(pongo2.Pongoer())

    m.Get("/",index.Index)
    m.Get("/login", user.Login)
    m.Post("/login", user.LoginPost)
    m.Get("/logout", user.Logout)

    m.NotFound(func(ctx *macaron.Context){
//        ctx.Redirect("/")
          ctx.HTML(200, "404", ctx.Data)
    })

    /*
    m.Group("/user", func(r *macaron.Router){
        r.Get("/login", user.Login)
        r.Post("/login", user.LoginPost)
    })
    */

    m.Run()

}
