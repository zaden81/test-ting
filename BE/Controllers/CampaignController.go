package Controllers

import (
	"net/http"
	"strconv"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	var guid = uuid.New().String()

	pageIndexStr := context.DefaultQuery("pageIndex", "1")
	pageSizeStr := context.DefaultQuery("pageSize", "10")

	pageIndex, err1 := strconv.ParseInt(pageIndexStr, 10, 64)
	pageSize, err2 := strconv.ParseInt(pageSizeStr, 10, 64)

	if err1 != nil || err2 != nil || pageIndex <= 0 || pageSize <= 0 {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = "Invalid pageIndex/pageSize"
		KolsVM.PageIndex = 1
		KolsVM.PageSize = 10
		KolsVM.Guid = guid
		KolsVM.TotalCount = 0
		KolsVM.KOL = make([]*DTO.KolDTO, 0)
		context.JSON(http.StatusBadRequest, KolsVM)
		return
	}

	kols, total, err := Logic.GetKolLogic(pageIndex, pageSize)
	if err != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = err.Error()
		KolsVM.PageIndex = pageIndex
		KolsVM.PageSize = pageSize
		KolsVM.Guid = guid
		KolsVM.TotalCount = 0
		KolsVM.KOL = make([]*DTO.KolDTO, 0)
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = pageIndex
	KolsVM.PageSize = pageSize
	KolsVM.Guid = guid
	KolsVM.TotalCount = total
	KolsVM.KOL = kols

	context.JSON(http.StatusOK, KolsVM)
}

