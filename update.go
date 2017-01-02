package main

import (
	"regexp"
	"strings"
	"errors"
)

func updateCopyrightYear(content string, copyrightPattern string, newYear string) (string, error) {
	re, err := regexp.Compile(copyrightPattern)
	if err != nil {
		return "", err
	}

	copyright := re.FindStringSubmatch(content)
	if len(copyright) == 0 {
		return "", errors.New(labels.CopyrightPatternNotFound)
	}

	lastIndex := len(copyright) - 1
	if copyright[lastIndex] == "" {
		copyright = copyright[:lastIndex]
		lastIndex--
	}
	oldCopyright := copyright[0]
	oldYear := copyright[lastIndex]

	if newYear == oldYear {
		return "", errors.New(labels.YearHasNotChanged)
	}

	newCopyright := strings.Replace(oldCopyright, oldYear, newYear, 1)
	content = strings.Replace(content, oldCopyright, newCopyright, 1)

	return content, nil
}
