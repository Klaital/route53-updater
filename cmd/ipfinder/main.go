package main

import (
	"fmt"
	"github.com/klaital/route53-updater/pkg/ipsource"
	"github.com/klaital/route53-updater/pkg/ipsource/dnsomatic"
	log "github.com/sirupsen/logrus"
)

func main() {
	var src ipsource.IPSource
	src = dnsomatic.New()
	ip, err := src.GetPublicIP()

	if err != nil {
		log.WithError(err).Fatal("Failed to get public IP")
	}

	fmt.Println(ip)
}
