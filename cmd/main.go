package main

//0 */2 * * * date >> /home/qw/weather
import (
	"github.com/zhanglt/weather/internal/model"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(model.ProvideConfig),
		fx.Provide(model.ProvideLog),
		fx.Provide(model.ProvideDbClient),
		fx.Invoke(bootInvoke),
	)
	/*
		scm := di.ServiceConstructorMap{
			"config": func(get di.Get) interface{} {
				return model.ProvideConfig()
			},
			"log": func(get di.Get) interface{} {
				return model.ProvideLog(get("config").(*model.Config))
			},
		}

		container := di.NewContainer(scm)
		config := container.Get("config").(*model.Config)
		log := container.Get("log").(*log.Logger)
		log.Println(config.Area)
	*/
}
