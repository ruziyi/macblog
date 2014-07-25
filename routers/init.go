package routers

import(
    "github.com/weisd/macblog/modules/setting"
    "github.com/weisd/macblog/modules/log"
    "github.com/weisd/macblog/models"
)

func GlobalInti(){
    setting.NewConfigContext()
    setting.NewLogService()
    models.LoadModelsConfig()

    if err := models.NewEngine(); err != nil {
        log.Fatal("Fail to initialize ORM engine: %v", err)
    }
}
