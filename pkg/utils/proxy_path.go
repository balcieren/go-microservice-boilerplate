package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type IncomingPath struct {
	Full   string `json:"full"`
	Parent string `json:"parent"`
}

var ErrInvalidPath = errors.New("invalid path")

func SplitProxyPath(c *fiber.Ctx, sep string) (IncomingPath, error) {
	var fullPath, parentPath string = "", ""

	if len(strings.Split(string(c.Request().URI().RequestURI()), sep)) >= 2 {
		fullPath = strings.Split(string(c.Request().URI().RequestURI()), sep)[1]
		fullPath = strings.Split(fullPath, "?")[0]
	} else {
		return IncomingPath{}, ErrInvalidPath
	}

	if len(strings.Split(fullPath, "/")) >= 2 {
		parentPath = strings.Split(fullPath, "/")[1]
	} else {
		return IncomingPath{}, ErrInvalidPath
	}

	return IncomingPath{
		Full:   fullPath,
		Parent: parentPath,
	}, nil
}

type OutgoingPath struct {
	Version   string `json:"version"`
	HostSplit string `json:"host_split"`
	Port      string `json:"port"`
}

func GenerateURL(incoming IncomingPath, outgoing OutgoingPath) string {
	if outgoing.HostSplit == "" {
		outgoing.HostSplit = "-api"
	}

	return fmt.Sprintf("http://%s:%s/%s%s", incoming.Parent+outgoing.HostSplit, outgoing.Port, outgoing.Version, incoming.Full)
}
