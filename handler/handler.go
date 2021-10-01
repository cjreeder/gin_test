package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cjreeder/gin_test/flightdeck"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Deployer flightdeck.Deployer
}

func AsyncHandler() {
	time.Sleep(8 * time.Second)
	fmt.Println("I wait 8 seconds for a response")
}

func (h *Handlers) RefloatByDeviceID(c *gin.Context) {

	//c.Status(http.StatusOK)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Minute)
	defer cancel()

	//var wg sync.WaitGroup
	//wg.Add(1)
	/*
	   go func() {
	           //      defer wg.Done()
	           if err := h.Deployer.Refloat(ctx, c.Param("deviceID")); err != nil {
	                   c.String(http.StatusInternalServerError, err.Error())
	                   return
	           }

	           c.JSON(http.StatusOK, map[string]string{"Success": "Command Sent"})
	   }()
	*/
	//defer wg.Wait()

	if err := h.Deployer.Refloat(ctx, c.Param("deviceID")); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//c.Status(http.StatusOK)
}
