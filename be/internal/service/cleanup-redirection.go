package service

import (
	"fmt"
	"github.com/novando/byfood/be/internal/dto"
	"strings"
)

const (
	REDIRECTION = "redirection"
	CANONICAL   = "canonical"
	ALL         = "all"
)

func processCanonical(url string) string {
	urlArr := strings.Split(url, "?")
	return strings.TrimSuffix(urlArr[0], "/")
}

func processRedirection(url string) string {
	urlArr := strings.Split(url, "://")
	pathArr := strings.Split(urlArr[1], "/")
	pathArr[0] = "www.byfood.com"
	newUrl := fmt.Sprintf("%v://%v", urlArr[0], strings.Join(pathArr, "/"))
	return strings.ToLower(newUrl)
}

func ProcessUrl(params dto.CleanupRequest) string {
	switch params.Operation {
	case CANONICAL:
		return processCanonical(params.Url)
	case REDIRECTION:
		return processRedirection(params.Url)
	default:
		urlRedirection := processRedirection(params.Url)
		return processCanonical(urlRedirection)
	}
}
