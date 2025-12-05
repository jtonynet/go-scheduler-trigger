package ginHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jtonynet/go-scheduler-trigger/api/bootstrap"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/core/dto"
)

// @Summary Payment Execute Transaction
// @Description Schedule trigger to send email at a pre-determined UTC date/time.
// @Tags SchedulerTrigger
// @Accept json
// @Produce json
// @Param request body dto.SchedulerTriggerReq true "Request body for Create Scheduler Trigger"
// @Router /schedules [post]
// @Success 200 {object} dto.SchedulerTriggerResp
func SchedulerTriggerCreate(ctx *gin.Context) {
	app := ctx.MustGet("app").(bootstrap.REST)

	var scheduleReq dto.SchedulerTriggerReq
	if err := ctx.ShouldBindBodyWith(&scheduleReq, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.SchedulerTriggerResp{
			Message: "Invalid Request",
		})

		return
	}

	uid, err := app.SchedulerTriggerCreate.Execute(scheduleReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.SchedulerTriggerResp{
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusAccepted, &dto.SchedulerTriggerResp{
		Message: "Request Accepted",
		UID:     *uid,
	})
}
