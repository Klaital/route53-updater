package ipsource

type IPSource interface {
	GetPublicIP() (string, error)
}