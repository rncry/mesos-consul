package mesos

import (
	"net"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func cleanName(name string, separator string) string {
	reg, err := regexp.Compile("[^\\w-]")
	if err != nil {
		log.Warn(err)
		return name
	}

	s := reg.ReplaceAllString(name, "-")

	return strings.ToLower(strings.Replace(s, "_", separator, -1))
}

func leaderIP(leader string) string {
	host := strings.Split(leader, "@")[1]
	host = strings.Split(host, ":")[0]

	return toIP(host)
}

func toIP(host string) string {
	// Check if host string is already an IP address
	ip := net.ParseIP(host)
	if ip != nil {
		return host
	}

	// Try to resolve host
	i, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		// Return the hostname if unable to resolve
		return host
	}

	return i.String()
}

func toPort(p string) int {
	ps, err := strconv.Atoi(p)
	if err != nil {
		log.Warnf("Invalid port number: %s", p)
	}

	return ps
}
