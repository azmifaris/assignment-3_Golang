package controllers

import (
	"assignment-3_AzmiFarisM/helpers"
	"assignment-3_AzmiFarisM/models"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStatus(ctx *gin.Context) {
	time.Sleep(15 * time.Second)
	helpers.UpdateDataJSON()
	jsonData, err := os.Open("data.json")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	tpl, errTemp := template.ParseFiles("templates/index.html")
	var w http.ResponseWriter
	if errTemp != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	bytes, err := io.ReadAll(jsonData)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var data models.Status
	json.Unmarshal(bytes, &data)
	waterStatus, windStatus := helpers.GetWaterStatus(data.Status.Water), helpers.GetWindStatus(data.Status.Wind)
	waterClassified, windClassified := helpers.GetWaterClass(data.Status.Water), helpers.GetWindClass(data.Status.Wind)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"water":       data.Status.Water,
		"waterStatus": waterStatus,
		"waterClass":  waterClassified,
		"wind":        data.Status.Wind,
		"windStatus":  windStatus,
		"windClass":   windClassified,
	})
	tpl.Execute(w, data)
}
