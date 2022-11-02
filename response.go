package chagin

import (
	"time"
	"runtime"
	"log"
	
	"github.com/gin-gonic/gin"
	"github.com/UedaTakeyuki/erapse"
)

func RespondErr(c *gin.Context, err error) {
	defer erapse.ShowErapsedTIme(time.Now())

	// get line of error occuring
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)

	log.Println(err, f.Name(), file, line)
	c.JSON(500, gin.H{
		"err": err.Error(),
	})
}
