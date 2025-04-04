/*
 * Catering service
 *
 * Auth service.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"

	loggerconfig "github.com/KusakinDev/Catering-Auth-Service/internal/config/logger"
	routespkg "github.com/KusakinDev/Catering-Auth-Service/internal/routes"
)

func main() {
	loggerconfig.Init()
	routes := routespkg.ApiHandleFunctions{}

	log.Printf("Server started")

	router := routespkg.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}
