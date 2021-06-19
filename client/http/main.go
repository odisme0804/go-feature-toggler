package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jessevdk/go-flags"
)

type Args struct {
	GRPCAddr string `long:"grpc.addr" env:"GRPC_ADDR" default:":8081"`
}

func main() {

	args := Args{}
	_, err := flags.NewParser(&args, flags.Default).Parse()
	if err != nil {
		panic(err)
	}

	logger := log.NewLogfmtLogger(os.Stderr)
	logger.Log("args", args)

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/simpleBooleanFlag", nil)
	if err != nil {
		level.Error(logger).Log("http.NewRequest", err)
		return
	}

	client := &http.Client{}

	timer := time.NewTicker(3 * time.Second)
	defer timer.Stop()

	users := []string{"Alice", "Bob", "Cindy", "Derrick", "Eric"}
	logger.Log("Name", fmt.Sprintf("| %20s | %20s | %20s | %20s | %20s", "Alice", "Bob", "Cindy", "Derrick", "Eric"))
	for {
		select {
		case <-timer.C:
			ss := ""
			for _, name := range users {
				req.Header.Set("X-Member-ID", name)
				resp, err := client.Do(req)
				if err != nil {
					level.Error(logger).Log("client.Do", err)
				}
				defer resp.Body.Close()

				body := map[string]interface{}{}
				if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
					level.Error(logger).Log("Decode", err)
				}

				ss += fmt.Sprintf("| %+v ", body)
			}
			logger.Log("@Got", fmt.Sprintf("%s ", ss))
		}
	}
}
