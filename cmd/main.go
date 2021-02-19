package main

//0 */2 * * * date >> /home/qw/weather
import (
	"context"
	"fmt"
	"log"

	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/qiniu/qmgo"
	"github.com/zhanglt/report/internal/model"
	"go.uber.org/fx"
	"gopkg.in/mgo.v2/bson"
)

func createReport(ctx context.Context, client *qmgo.QmgoClient, conf *model.Config, logger *log.Logger) {
	var air = []model.Quality{}
	var err error
	filter := bson.M{} //查询条件
	//err = client.Find(ctx, filter).One(&air) //从mongodb数据库中查询区域对应的天气信息并反序列化
	err = client.Find(ctx, filter).All(&air)
	if err != nil {
		logger.Println("mongoDB查询天气信息，反序列化错误，区域编码： 错误信息：", err)
	}
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("air")
	for i := 1; i < len(air); i++ {
		//f.SetCellValue("Sheet1", "B2", 100)
		if air[i].AreaCode > 5 {
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), air[i].AreaCode)
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), air[i].Date.Format("2006-01-02 15:04:05"))
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(i), air[i].CityCode)
			f.SetCellValue("Sheet1", "D"+strconv.Itoa(i), air[i].AreaName)
			if len(air[i].ApplyContent.Detail) > 0 {
				f.SetCellValue("Sheet1", "E"+strconv.Itoa(i), air[i].ApplyContent.Detail[0].PrimaryPollutant)
				f.SetCellValue("Sheet1", "F"+strconv.Itoa(i), air[i].ApplyContent.Detail[0].AirIndexLevel)
				f.SetCellValue("Sheet1", "G"+strconv.Itoa(i), air[i].ApplyContent.Detail[0].AirQualityIndex)
			}
			if len(air[i].ApplyContent.Detail) > 1 {
				f.SetCellValue("Sheet1", "H"+strconv.Itoa(i), air[i].ApplyContent.Detail[1].PrimaryPollutant)
				f.SetCellValue("Sheet1", "I"+strconv.Itoa(i), air[i].ApplyContent.Detail[1].AirIndexLevel)
				f.SetCellValue("Sheet1", "J"+strconv.Itoa(i), air[i].ApplyContent.Detail[1].AirQualityIndex)
			}
			if len(air[i].ApplyContent.Detail) > 2 {
				f.SetCellValue("Sheet1", "K"+strconv.Itoa(i), air[i].ApplyContent.Detail[2].PrimaryPollutant)
				f.SetCellValue("Sheet1", "L"+strconv.Itoa(i), air[i].ApplyContent.Detail[2].AirIndexLevel)
				f.SetCellValue("Sheet1", "M"+strconv.Itoa(i), air[i].ApplyContent.Detail[2].AirQualityIndex)
			}
		}
		// Set active sheet of the workbook.

		f.SetActiveSheet(index)
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(air))

}

/*
func bootInvoke(ctx context.Context, client *qmgo.QmgoClient, conf *model.Config, logger *log.Logger) {
	t1 := time.Now()
	count := 0
	for name, area := range conf.Area.Area {
		filter := bson.M{"areaid": area} //查询条件
		up, ok := model.UpdateWeather(ctx, client, conf, area, logger)
		if ok != nil {
			logger.Println("更新错误：", name, area, ok)
		} else {
			result, err := client.Upsert(ctx, filter, up)
			if err != nil {
				logger.Println("更新失败错误信息：", area, "|", err)
			}
			count++
			if conf.Writable.LogLevel == "DEBUG" {
				logger.Println("更新信息：", area, result.MatchedCount, ":", result.ModifiedCount, ":", result.UpsertedCount)
			}
		}

	}
	t2 := time.Now()
	logger.Println(time.Now().Format("2006/1/2 15:04:05"), "共同步：", count, "条数据,用时：", t2.Sub(t1))
}
*/
func main() {
	//excel()
	fx.New(
		fx.Provide(model.ProvideConfig),
		fx.Provide(model.ProvideLog),
		fx.Provide(model.ProvideDbClient),
		fx.Invoke(createReport),
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
