package main

import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("i want to learn consepts of time in golang:")
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Printf("Type of currentTime: %T\n", currentTime)

	// Formatting time
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)


	// Parsing time
	timeStr := "2023-10-15 14:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed Time:", parsedTime)
	}


	// Time arithmetic
	futureTime := currentTime.Add(2 * time.Hour)
	fmt.Println("Time after 2 hours:", futureTime)
	duration := futureTime.Sub(currentTime)
	fmt.Println("Duration between current time and future time:", duration)


	// Extracting components
	year, month, day := currentTime.Date()
	hour, min, sec := currentTime.Clock()
	fmt.Printf("Year: %d, Month: %d, Day: %d\n", year, month, day)
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", hour, min, sec)
	weekday := currentTime.Weekday()
	fmt.Println("Weekday:", weekday)

	// Time zones
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading location:", err)
	} else {
		nyTime := currentTime.In(location)
		fmt.Println("Current Time in New York:", nyTime)
	}

	timestamp := currentTime.Unix()
	fmt.Println("Unix Timestamp:", timestamp)
	sleepDuration := 2 * time.Second
	fmt.Printf("Sleeping for %v...\n", sleepDuration)
	time.Sleep(sleepDuration)
	fmt.Println("Awake now!")

	formattedUTC := currentTime.UTC().Format(time.RFC1123)
	fmt.Println("Current Time in UTC (RFC1123):", formattedUTC)

	customFormat := currentTime.Format("02-Jan-2006 03:04:05 PM")
	fmt.Println("Custom Formatted Time:", customFormat)

	durationExample := time.Duration(5 * time.Minute)
	fmt.Println("Duration Example (5 minutes):", durationExample)

	ticker := time.NewTicker(1 * time.Second)
	fmt.Println("Starting ticker for 3 seconds:")

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	fmt.Println("Time manipulation and formatting in Go completed.")
    
	




}