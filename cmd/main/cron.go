package main

import (
	"fmt"
	"log"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/captainkie/websync-api/config"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/internal/app/service"
	"github.com/go-co-op/gocron/v2"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func imageScheduler() {
	// default timezone
	os.Setenv("TZ", "Asia/Bangkok")
	// create a scheduler
	s, _ := gocron.NewScheduler()

	// add a job to the scheduler (every 10 minutes = 600*time.Second)
	j, err := s.NewJob(
		gocron.DurationJob(
			600*time.Second,
		),
		gocron.NewTask(
			func(a string, b int) {
				timeZone, offset := time.Now().Zone()
				log.Printf("🚗 Cronjob imageScheduler is processing (time zone: %s, %d)\n", timeZone, offset)

				db := config.ConnectDatabase()
				validate := validator.New()

				queueRepository := repository.NewQueueRepositoryImpl(db)
				queueService := service.NewQueueServiceImpl(queueRepository, validate)

				queueService.CreateImageQueue()
			},
			"image_cron_process",
			1,
		),
	)
	if err != nil {
		fmt.Println("Error creating job", err)
	}
	// each job has a unique id
	fmt.Println("🚀 Cronjob Start imageScheduler ID: ", j.ID())
	// start the scheduler
	s.Start()
}

func syncProductScheduler() {
	// default timezone
	os.Setenv("TZ", "Asia/Bangkok")

	// create a scheduler
	s, _ := gocron.NewScheduler()

	// add a job to the scheduler (every 00.01 PM)
	j, err := s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(00, 01, 0),
			),
		),
		gocron.NewTask(
			func(a string, b int) {
				timeZone, offset := time.Now().Zone()
				log.Printf("🚗 Cronjob syncProductScheduler is processing (time zone: %s, %d)\n", timeZone, offset)

				db := config.ConnectDatabase()
				validate := validator.New()

				queueRepository := repository.NewQueueRepositoryImpl(db)
				imageRepository := repository.NewImageRepositoryImpl(db)

				queueService := service.NewQueueServiceImpl(queueRepository, validate)
				imageService := service.NewImageServiceImpl(imageRepository, validate)

				requestID := uuid.New().String()
				queueService.CreateConnectionQueue(requestID)

				// Use the imageService to delete images older than targetDate
				targetDate := time.Now().AddDate(0, 0, -1)
				imageService.DeleteImage(targetDate)

				// Delete old image folder
				targetFolder := time.Now().AddDate(0, 0, -2).Format("20060102")
				baseDirectory := "public/uploads"
				directoryPath := fmt.Sprintf("%s/%s", baseDirectory, targetFolder)
				if _, err := os.Stat(directoryPath); !os.IsNotExist(err) {
					os.RemoveAll(directoryPath)
				}
			},
			"product_cron_process",
			1,
		),
	)
	if err != nil {
		fmt.Println("Error creating job", err)
	}
	// each job has a unique id
	fmt.Println("🚀 Cronjob Start syncProductScheduler ID: ", j.ID())

	// start the scheduler
	s.Start()
}
