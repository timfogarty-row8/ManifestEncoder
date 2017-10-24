package manifest-encoder

import (
	"net/http"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/subosito/gotenv"
)


var Redis *pool.Pool

func init() {
	var err error

	gotenv.Load()
	
	server := os.Getenv("REDIS_SERVER")
	
	num_connections, err := strconv.Atoi(os.Getenv("REDIS_POOL", "10"))
    if err != nil {
        log.Fatal(err)
    }
    
    var err error
    	Redis, err := pool.New("tcp", server, num_connections)
    if err != nil {
        log.Fatal(err)
    }
}


func main() {
	init()
	
	gin.DisableConsoleColor()

	r := gin.Default()
	r.Use(cors.Default())
	
	r.GET("/version", func(c *gin.Context) {
		c.String(200,"Manifest Encoder v0.1.0")
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.Use(gin.Logger())
		v1.GET("/:format/:mediaGuid/:token", encodeManifest)

		v1.Use( authMiddleware.MiddlewareFunc() ) {
			v1.GET("/list", listStats)
			v1.GET("/refresh", refreshCache)
		}
	}

	// Listen and Server in 0.0.0.0:9000
	r.Run(":9000")
}
