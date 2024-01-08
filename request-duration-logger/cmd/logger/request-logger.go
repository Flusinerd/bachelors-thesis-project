package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"request-logger/requests"
	"sort"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "request-duration-logger",
		Action: func(ctx *cli.Context) error {
			url := ctx.Args().Get(0)
			totalCount := ctx.Args().Get(1)
			parallelCount := ctx.Args().Get(2)

			if url == "" {
				panic("url is required")
			}

			if totalCount == "" {
				totalCount = "1"
			}

			if parallelCount == "" {
				parallelCount = "1"
			}

			totalCountInt, err := strconv.Atoi(totalCount)
			if err != nil {
				panic(err)
			}

			parallelCountInt, err := strconv.Atoi(parallelCount)
			if err != nil {
				panic(err)
			}

			if totalCountInt < 1 {
				panic("totalCount must be greater than 0")
			}

			if parallelCountInt < 1 {
				panic("parallelCount must be greater than 0")
			}

			fmt.Printf("url: %s, totalCount: %s, parallelCount: %s\n", url, totalCount, parallelCount)

			client := requests.NewRequestClient(requests.RequestConfig{
				TotalRequestsCount: uint(totalCountInt),
				ParallelCount:      uint(parallelCountInt),
				Url:                url,
				Method:             "GET",
				Body:               "",
			})

			results := client.Run()

			timeStamp := time.Now().Unix()
			fileName := fmt.Sprintf("results-%d.csv", timeStamp)

			// file relative to current cwd
			currentDir, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			fileName = path.Join(currentDir, fileName)

			file, err := os.Create(fileName)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			_, err = file.WriteString("Index,Time,TotalTime,StatusCode\n")
			if err != nil {
				panic(err)
			}

			// Sort the results by startTime ASC
			sort.Slice(results, func(i, j int) bool {
				return results[i].StartedAt.Before(results[j].StartedAt)
			})

			var startTime time.Time
			for index, result := range results {
				if index == 0 {
					startTime = result.StartedAt
				}

				totalTime := result.FinishedAt.Sub(startTime)

				// TODO: runden

				_, err := file.WriteString(fmt.Sprintf("%d,%d,%d,%d\n", index, result.TotalTime.Milliseconds(), totalTime.Milliseconds(), result.StatusCode))
				if err != nil {
					panic(err)
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
