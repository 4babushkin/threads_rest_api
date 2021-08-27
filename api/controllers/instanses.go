package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Get(context *gin.Context) {
	name := context.Param("instans")
	config, err := ReadConfig(os.Getenv("APP_THREAD_LIC_DIR") + "/" + name + ".txt")
	if err != nil {
		context.JSON(404, gin.H{"error": err.Error()})
		return
	}
	type Instance struct {
		Customer        string `json:"lic.customer" gorm:"primary_key"`
		Expiration      string `json:"lic.expiration-date"`
		Features        string `json:"lic.features"`
		Max_active_user int    `json:"lic.max-active-users"`
	}

	max_active_users, err := strconv.Atoi(config["lic.max-active-users"])
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": "cannot convert lic.max-active-users to int"})
		return
	}
	var instance = Instance{name, config["lic.expiration"], config["lic.features"], max_active_users}

	context.JSON(http.StatusOK, gin.H{"message": instance})
}

///////////////////////////////////
func GetAllInstanses(context *gin.Context) {

	type InstItem struct {
		Name string
	}

	type Inst struct {
		Items []InstItem
	}

	file, err := os.Open(os.Getenv("APP_THREAD_LIC_DIR"))
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders

	onlyFull := []string{}
	// onlyFull := map[string]interface{}
	for _, v := range list {
		if strings.HasSuffix(v, ".txt") {
			item := strings.TrimSuffix(v, ".txt")
			// item = "{Name:Instance: " + item + "},"
			onlyFull = append(onlyFull, item)
		}
	}

	context.JSON(http.StatusOK, onlyFull)

}
