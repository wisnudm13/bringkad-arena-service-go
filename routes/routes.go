package routes

import (
	// import v1 routes
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type VersionRouter interface {
	RegisterRoutes(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine) {
	apiVersions := FindAPIVersion()

	for _, version := range apiVersions {
		log.Printf(version)
	}
}

func FindAPIVersion() []string {
	/*
		Read all directories on routes dir, return all directories with prefix v
	*/

	var versions []string
	entries, err := os.ReadDir("routes")

	if err != nil {
		log.Fatalf("Failed to read directory routes: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "v") {
			versions = append(versions, entry.Name())
		}
	}

	return versions

}
