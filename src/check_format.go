package main

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

func checkDate(d string) bool {

	correct := true

	s := strings.Split(d, ("."))

	if !isInt(s[0]) {
		correct = false
	}

	if !isInt(s[1]) {
		correct = false
	}

	if !isInt(s[2]) {
		correct = false
	}

	day, _ := strconv.Atoi(s[0])
	month, _ := strconv.Atoi(s[1])
	year, _ := strconv.Atoi(s[2])
	if month > 12 {
		correct = false
	}

	if year < 2023 {
		correct = false
	}

	if day > daysMonth[month-1] {
		correct = false
	}

	return correct
}

func checkTime(d string) bool {
	correct := true

	s := strings.Split(d, (":"))

	if !isInt(s[0]) {
		correct = false
	}

	if !isInt(s[1]) {
		correct = false
	}

	hour, _ := strconv.Atoi(s[0])
	minutes, _ := strconv.Atoi(s[1])

	if hour > 24 {
		correct = false
	}

	if minutes > 60 {
		correct = false
	}

	return correct
}

func isInt(st string) bool {
	for _, c := range st {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
func parseDate(date string, rawtime string) time.Time {
	s := strings.Split(date, ("."))
	day := s[0]
	month := s[1]
	year := s[2]
	t := strings.Split(rawtime, (":"))
	hour := t[0]
	minute := t[1]
	frmtd := year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + "00"
	vl, _ := time.Parse("2006-01-02 15:04:05", frmtd)
	return vl
}
