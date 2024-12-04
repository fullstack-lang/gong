package models

import (
	"fmt"
	"regexp"
	"strconv"
)

func extractWidthNumber(s string) (int, error) {
	re := regexp.MustCompile(`gong:width\s*([0-9]+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) < 2 {
		return 0, fmt.Errorf("No width number found")
	}
	widthStr := matches[1]
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return 0, err
	}
	return width, nil
}

func extractHeightNumber(s string) (int, error) {
	re := regexp.MustCompile(`gong:height\s*([0-9]+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) < 2 {
		return 0, fmt.Errorf("No height number found")
	}
	heightStr := matches[1]
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return 0, err
	}
	return height, nil
}

func extractTimeFormat(s string) (string, error) {
	re := regexp.MustCompile(`gong:bespoketimeserializeformat "([^"]*)"`)
	matches := re.FindStringSubmatch(s)
	if len(matches) < 2 {
		return "", fmt.Errorf("No time format found")
	}
	format := matches[1]
	return format, nil
}
