package main

import (
  "encoding/json"
  "fmt"
)

// Profile holds the profile variables of Arturo build
type Profile struct {
  PipelineStage string `json:"pipelineStage"`
  PipelineBuildPrefix string `json:"pipelineBuildPrefix"`
}

// Message holds the Webhook payload from Arturo
type Message struct {
  Event string `json:"event"`
  Profile *Profile `json:"profile"`
}

func main() {
  in :=
  `{
     "Event":"teardown",
     "History": "build_history",
     "Profile": {
       "pipelineStage": "COMPLETE",
       "pipelineBuildPrefix": "builds/degitaliam/isamlifecyclemanagement/feature-IDMAN-1957/43",
       "end": "end"
     }
   }`
  bytes := []byte(in)

  var message Message
  if err := json.Unmarshal(bytes, &message); err != nil {
    panic(err)
  }

  messageJSONString, _ := json.Marshal(message)
  fmt.Println(string(messageJSONString))
}
