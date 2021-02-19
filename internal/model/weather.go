package model

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

// 从天气预报API获取数据并发序列化到 stuct对象
func getWeatherNew(area int, conf *Config, logger *log.Logger) (Weather, error) {

	var wn Weather
	var err error
	res, err := http.Get(conf.Service.APIHost + strconv.Itoa(area))
	if err != nil {

		logger.Println("获取天气信息错误，区域编码：", area, "错误信息：", err)
		return wn, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Println("读取天气信息错误，区域编码：", area, "错误信息：", err)
		return wn, err
	}
	err = json.Unmarshal(body, &wn)
	if err != nil {
		logger.Println("读取天气信息错误，反序列化错误：", area, "错误信息：", err)
		return wn, err
	}
	return wn, err
}

// getWeatherF 从mongo读取对应area的数据条目，并发序列化到struc对象

func getWeatherF(ctx context.Context, client *qmgo.QmgoClient, area int, logger *log.Logger) (Weather_f, error) {
	var wf Weather_f
	var err error
	filter := bson.M{"areaid": area}        //查询条件
	err = client.Find(ctx, filter).One(&wf) //从mongodb数据库中查询区域对应的天气信息并反序列化
	if err != nil {
		logger.Println("mongoDB查询天气信息，反序列化错误，区域编码：", area, "错误信息：", err)
		return wf, err
	}

	if len(wf.Content.F.F1) < 2 {
		logger.Println("wf 反序列化原始数据不完整：", area)
		return wf, errors.New("原始数据错误不完整")
	}
	return wf, err
}

// UpdateWeather 利用天气api获取的数据，更新mongdb数据对象，并更新写入mongo数据库
func UpdateWeather(ctx context.Context, client *qmgo.QmgoClient, conf *Config, area int, logger *log.Logger) (Weather_f, error) {

	wn, er := getWeatherNew(area, conf, logger)
	wf, err := getWeatherF(ctx, client, area, logger)
	if er != nil || err != nil {
		logger.Println("wn或者wf获取错误:", er, err, area)
		return wf, errors.New("wn或者wf获取错误")
	} else {
		wf.Date = time.Now()
		for i := 0; i < 3; i++ { //转换三天的数据
			//更新信息
			wf.Date = time.Now()
			//wf.Content.F.F0 = "202101271907"
			windPower := wn.Data.Forecast[i].Fengli
			windPower = windPower[9 : len(windPower)-1]
			windPower = windPower[:len(windPower)-5]
			tempertureH := wn.Data.Forecast[i].High
			tempertureH = tempertureH[6 : len(tempertureH)-1]
			tempertureH = tempertureH[:len(tempertureH)-2]
			tempertureL := wn.Data.Forecast[i].Low
			tempertureL = tempertureL[6 : len(tempertureL)-1]
			tempertureL = tempertureL[:len(tempertureL)-2]
			wf.Content.F.F1[i].Fa = ConvertWeather(wn.Data.Forecast[i].Type)
			wf.Content.F.F1[i].Fb = ConvertWeather(wn.Data.Forecast[i].Type)
			wf.Content.F.F1[i].Fc = tempertureH
			wf.Content.F.F1[i].Fd = tempertureL
			wf.Content.F.F1[i].Fe = ConvertWind(wn.Data.Forecast[i].Fengxiang)
			wf.Content.F.F1[i].Ff = ConvertWind(wn.Data.Forecast[i].Fengxiang)
			wf.Content.F.F1[i].Fg = windPower
			wf.Content.F.F1[i].Fh = windPower
		}

	}
	return wf, err
}
