package main

import (
	"encoding/json"
	"flag"
	"time"

	"github.com/apex/go-apex"
	"github.com/golang/glog"
	"github.com/taskd19/tegami"
)

type dummyResult struct {
	count int
}

func main() {
	//To avoid "Unexpected Token" in AWS Lambda
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()

	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		timeframe, consumerKey, consumerSecret, accessToken, accessTokenSecret := tegami.GetEnvValues()
		client := tegami.NewTwitterClient(consumerKey, consumerSecret, accessToken, accessTokenSecret)

		currentTime := time.Now()
		glog.Infof("currentTime: %s", currentTime)
		glog.Infof("timeframe: %d", timeframe)
		result, err := tegami.CheckEvent(currentTime, timeframe, client)

		if err != nil {
			glog.Fatalf("Tegami got an error while checking event. err: %s", err)
		}
		glog.Infof("Tegami notified %d events at %s", result, currentTime)

		return &dummyResult{count: result}, err
	})
}
