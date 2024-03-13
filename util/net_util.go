package util

import (
	"net"
)

// getHostname obtiene el hostname a través de llamada al sistema operativo.
func getHostname(ip string) string {
	names, err := net.LookupAddr(ip)
	if err != nil {
		return ip
	}
	return names[0]
}

// IPCache es un tipo para implementar un caché de indagaciones a DNS
type IPCache map[string]string

// updCacheIP agrega un nuevo registro al cache cuando se detecta que la ip
// no se encuentra en él.
func (c IPCache) updCacheIP(ip string) {
	if c[ip] == "" {
		c[ip] = getHostname(ip)
	}
}

// NewCacheIP crea y entrega un caché vació
func NewCacheIP() IPCache {
	return make(IPCache)
}

// Hostname devueelve el hostname que corresponde a la IP que se recibe de
// parámetro.  Para reducir las llamadas al sistema, se indaga contra un caché
// implemetado por la aplicación.
// (net_util.go)
func Hostname(ip string, cache IPCache) string {
	cache.updCacheIP(ip)
	return cache[ip]
}
