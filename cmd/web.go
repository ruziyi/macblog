package cmd

import(
    "fmt"

	"github.com/codegangsta/cli"

    "github.com/Unknwon/macaron"
    
    "github.com/weisd/macblog/routers/user"
    "github.com/macaron-contrib/session"

    "github.com/macaron-contrib/pongo2"
)

var CmdWeb = cli.Command{
    Name: "web",
    Usage:  "start macblog web server",
    Description:  "macblog for weisd",
    Action: runWeb,
    Flags:  []cli.Flag{},
}

func runWeb(*cli.Context){
    fmt.Println("run web server")

    m := macaron.Classic()

    m.Use(session.Sessioner())
    m.Use(pongo2.Pongoer())

    m.Get("/",func() string {
        return "hello world!"
    })

    m.Get("/login", user.Login)

    m.Run()
}

