package main

import (
  "encoding/json"
  "net/http"
  "strings"

  "github.com/gin-gonic/gin"
)

// Profile holds the profile variables of Arturo build
type Profile struct {
  Branch string `json:"pipelineBranch"`
  BuildNumber string `json:"pipelineBuildNumber"`
  ApplicationName string `json:"pipelineApplication"`
  ServiceName string `json:"pipelineService"`
  Domain string `json:"pipelineDomain"`
}

// Message holds the Webhook payload from Arturo
type Message struct {
  Event string `json:"event"`
  Profile *Profile `json:"profile"`
  Webseal [2]string `json:"webseal"`
}

func ConstructWebsealNames(strs ...string) [2]string {
  var ret [2]string
  var sb strings.Builder

  for _, str := range strs {
    sb.WriteString(str)
  }
  ret[0] = "AWS-1-1.webseald" + sb.String()
  ret[1] = "AWS-2-1.webseald" + sb.String()

  return ret
}

func main() {
  router := gin.Default()

  router.POST("/event", func(c *gin.Context) {
    var message Message

    if err := c.ShouldBindJSON(&message); err == nil {

      message.Webseal = ConstructWebsealNames(".",
        string(message.Profile.Branch),"-",
	string(message.Profile.BuildNumber),".",
	string(message.Profile.ApplicationName),".",
	string(message.Profile.ServiceName),".",
	string(message.Profile.Domain),
      )
      res, _ := json.Marshal(&message)
      c.JSON(http.StatusOK, gin.H{"status":string(res)})

    } else {

      c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})

    }
  })

  router.Run()
}
