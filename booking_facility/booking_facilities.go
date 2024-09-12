// Based on given booking inputs (bookingInputs) try booking the slots
package main

import (
	"fmt"
	"time"
)

var bookingInputs = []struct {
	facility string
	date     string
	start    string
	end      string
}{
	{Clubhouse, "26-10-2020", "16:00", "22:00"},
	{TennisCourt, "26-10-2020", "16:00", "20:00"},
	{Clubhouse, "26-10-2020", "16:00", "22:00"},
	{TennisCourt, "26-10-2020", "17:00", "21:00"},
	{BasketBallCourt, "27-10-2020", "17:00", "21:00"},
	{BasketBallCourt, "27-10-2020", "20:00", "21:00"},
}

const (
	Clubhouse       = "Clubhouse"
	TennisCourt     = "Tennis Court"
	BasketBallCourt = "Basket Ball Court"
)

// Define Pricing configgurations
const (
	clubhouseDayRate     = 100
	clubhouseEveningRate = 500
	tennisCourtRate      = 50
	basketBallCourtRate  = 500
)

// Define booking structure
type Booking struct {
	StartTime time.Time
	EndTime   time.Time
}

// Store bookings in memory
var bookings = map[string][]Booking{}

// Calculate cost for Clubhouse
func calculateClubhouseCost(start, end time.Time) int {
	totalCost := 0
	for t := start; t.Before(end); t = t.Add(time.Hour) {
		if t.Hour() >= 10 && t.Hour() < 16 {
			totalCost += clubhouseDayRate
		} else if t.Hour() >= 16 && t.Hour() < 22 {
			totalCost += clubhouseEveningRate
		}
	}
	return totalCost
}

// Calculate cost for Tennis Court
func calculateTennisCourtCost(start, end time.Time) int {
	duration := end.Sub(start).Hours()
	return int(duration) * tennisCourtRate
}

// Calculate cost for Basket Ball Court
func calculateBasketballCourtCost(start, end time.Time) int {
	duration := end.Sub(start).Hours()
	return int(duration) * basketBallCourtRate
}

// Check if a time slot is available
func isAvailable(facility string, start, end time.Time) bool {
	for _, booking := range bookings[facility] {
		if start.Before(booking.EndTime) && end.After(booking.StartTime) {
			return false
		}
	}
	return true
}

// Book a facility
func bookFacility(facility string, start, end time.Time) (string, int) {
	if !isAvailable(facility, start, end) {
		return "Booking Failed, Already Booked", 0
	}

	var cost int
	switch facility {
	case Clubhouse:
		cost = calculateClubhouseCost(start, end)
	case TennisCourt:
		cost = calculateTennisCourtCost(start, end)
	case BasketBallCourt:
		cost = calculateBasketballCourtCost(start, end)
	}

	bookings[facility] = append(bookings[facility], Booking{StartTime: start, EndTime: end})
	return "Booked", cost
}

func main() {
	// Example inputs

	for _, test := range bookingInputs {
		dateTimeLayout := "02-01-2006 15:04"
		startTime, _ := time.Parse(dateTimeLayout, fmt.Sprintf("%s %s", test.date, test.start))
		endTime, _ := time.Parse(dateTimeLayout, fmt.Sprintf("%s %s", test.date, test.end))

		status, cost := bookFacility(test.facility, startTime, endTime)
		if status == "Booked" {
			fmt.Printf("%s, %s, %s - %s, %s Rs. %d\n", test.facility, test.date, test.start, test.end, status, cost)
		} else {
			fmt.Printf("%s, %s, %s - %s\n", test.facility, test.date, test.start, status)
		}
	}
}
