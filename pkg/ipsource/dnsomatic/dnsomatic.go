package dnsomatic

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type DnsomaticIPSource struct {
	Hostname string
	Path     string
}

func New() *DnsomaticIPSource {
	return &DnsomaticIPSource{
		Hostname: "http://myip.dnsomatic.com",
		Path:     "/mypublicip.txt",
	}
}

func (src *DnsomaticIPSource) GetPublicIP() (string, error) {
	url := fmt.Sprintf("%s%s", src.Hostname, src.Path)
	resp, err := http.Get(url)
	if err != nil {
		log.WithFields(log.Fields{
			"hostname": src.Hostname,
			"path":     src.Path,
			"url":      url,
		}).WithError(err).Error("Error making GET request")
		return "", err
	}
	// TODO: if statuscode == 429, backoff and retry
	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"hostname": src.Hostname,
			"path":     src.Path,
			"url":      url,
			"status":   resp.Status,
		}).Error("HTTP Error from dnsomatic")
		return "", fmt.Errorf("dnsomatic http response %d indicates failure", resp.StatusCode)
	}

	defer resp.Body.Close()
	ipbuffer := make([]byte, 17)
	_, err = resp.Body.Read(ipbuffer)
	if err == io.EOF {
		return string(ipbuffer), nil
	}
	if err != nil {
		log.WithFields(log.Fields{
			"hostname": src.Hostname,
			"path":     src.Path,
			"url":      url,
			"status":   resp.Status,
		}).WithError(err).Error("Failed to read dnsomatic response")
		return "", err
	}

	return string(ipbuffer), nil
}
