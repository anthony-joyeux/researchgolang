package arangoapi

import "gopkg.in/kataras/iris.v6"
import "time"




type TimeElapsed struct {
	Startime int64 `json:"startime"`
	Endtime  int64 `json:"endtime"`
	Duration int64 `json:"duration"`
}

func Read(ctx *iris.Context) {

	var timer TimeElapsed
	timer.Startime =time.Now().UnixNano()
	CreateRecord()
	timer.Endtime =time.Now().UnixNano()
	timer.Duration = timer.Endtime-timer.Startime
	ctx.JSON(iris.StatusOK, timer)
}


func ReadAll(ctx *iris.Context) {

	var timer TimeElapsed


	ctx.JSON(iris.StatusOK, timer)
}

func Create(ctx *iris.Context) {

	var timer TimeElapsed
	timer.Startime =time.Now().UnixNano()
	CreateRecord()
	timer.Endtime =time.Now().UnixNano()
	timer.Duration = timer.Endtime-timer.Startime


	ctx.JSON(iris.StatusOK, timer)
}

func CreateInNumbers(ctx *iris.Context) {



	var timer TimeElapsed
	timer.Startime =time.Now().UnixNano()

	CreateRecord()

	timer.Endtime =time.Now().UnixNano()
	timer.Duration = timer.Endtime-timer.Startime


	ctx.JSON(iris.StatusOK, timer)
}



func Delete(ctx *iris.Context) {

	var timer TimeElapsed


	ctx.JSON(iris.StatusOK, timer)
}

