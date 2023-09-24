// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongtable/go")
	{ // insertion point for registrations
		v1.GET("/v1/cells", GetController().GetCells)
		v1.GET("/v1/cells/:id", GetController().GetCell)
		v1.POST("/v1/cells", GetController().PostCell)
		v1.PATCH("/v1/cells/:id", GetController().UpdateCell)
		v1.PUT("/v1/cells/:id", GetController().UpdateCell)
		v1.DELETE("/v1/cells/:id", GetController().DeleteCell)

		v1.GET("/v1/cellbooleans", GetController().GetCellBooleans)
		v1.GET("/v1/cellbooleans/:id", GetController().GetCellBoolean)
		v1.POST("/v1/cellbooleans", GetController().PostCellBoolean)
		v1.PATCH("/v1/cellbooleans/:id", GetController().UpdateCellBoolean)
		v1.PUT("/v1/cellbooleans/:id", GetController().UpdateCellBoolean)
		v1.DELETE("/v1/cellbooleans/:id", GetController().DeleteCellBoolean)

		v1.GET("/v1/cellfloat64s", GetController().GetCellFloat64s)
		v1.GET("/v1/cellfloat64s/:id", GetController().GetCellFloat64)
		v1.POST("/v1/cellfloat64s", GetController().PostCellFloat64)
		v1.PATCH("/v1/cellfloat64s/:id", GetController().UpdateCellFloat64)
		v1.PUT("/v1/cellfloat64s/:id", GetController().UpdateCellFloat64)
		v1.DELETE("/v1/cellfloat64s/:id", GetController().DeleteCellFloat64)

		v1.GET("/v1/cellicons", GetController().GetCellIcons)
		v1.GET("/v1/cellicons/:id", GetController().GetCellIcon)
		v1.POST("/v1/cellicons", GetController().PostCellIcon)
		v1.PATCH("/v1/cellicons/:id", GetController().UpdateCellIcon)
		v1.PUT("/v1/cellicons/:id", GetController().UpdateCellIcon)
		v1.DELETE("/v1/cellicons/:id", GetController().DeleteCellIcon)

		v1.GET("/v1/cellints", GetController().GetCellInts)
		v1.GET("/v1/cellints/:id", GetController().GetCellInt)
		v1.POST("/v1/cellints", GetController().PostCellInt)
		v1.PATCH("/v1/cellints/:id", GetController().UpdateCellInt)
		v1.PUT("/v1/cellints/:id", GetController().UpdateCellInt)
		v1.DELETE("/v1/cellints/:id", GetController().DeleteCellInt)

		v1.GET("/v1/cellstrings", GetController().GetCellStrings)
		v1.GET("/v1/cellstrings/:id", GetController().GetCellString)
		v1.POST("/v1/cellstrings", GetController().PostCellString)
		v1.PATCH("/v1/cellstrings/:id", GetController().UpdateCellString)
		v1.PUT("/v1/cellstrings/:id", GetController().UpdateCellString)
		v1.DELETE("/v1/cellstrings/:id", GetController().DeleteCellString)

		v1.GET("/v1/checkboxs", GetController().GetCheckBoxs)
		v1.GET("/v1/checkboxs/:id", GetController().GetCheckBox)
		v1.POST("/v1/checkboxs", GetController().PostCheckBox)
		v1.PATCH("/v1/checkboxs/:id", GetController().UpdateCheckBox)
		v1.PUT("/v1/checkboxs/:id", GetController().UpdateCheckBox)
		v1.DELETE("/v1/checkboxs/:id", GetController().DeleteCheckBox)

		v1.GET("/v1/displayedcolumns", GetController().GetDisplayedColumns)
		v1.GET("/v1/displayedcolumns/:id", GetController().GetDisplayedColumn)
		v1.POST("/v1/displayedcolumns", GetController().PostDisplayedColumn)
		v1.PATCH("/v1/displayedcolumns/:id", GetController().UpdateDisplayedColumn)
		v1.PUT("/v1/displayedcolumns/:id", GetController().UpdateDisplayedColumn)
		v1.DELETE("/v1/displayedcolumns/:id", GetController().DeleteDisplayedColumn)

		v1.GET("/v1/formdivs", GetController().GetFormDivs)
		v1.GET("/v1/formdivs/:id", GetController().GetFormDiv)
		v1.POST("/v1/formdivs", GetController().PostFormDiv)
		v1.PATCH("/v1/formdivs/:id", GetController().UpdateFormDiv)
		v1.PUT("/v1/formdivs/:id", GetController().UpdateFormDiv)
		v1.DELETE("/v1/formdivs/:id", GetController().DeleteFormDiv)

		v1.GET("/v1/formeditassocbuttons", GetController().GetFormEditAssocButtons)
		v1.GET("/v1/formeditassocbuttons/:id", GetController().GetFormEditAssocButton)
		v1.POST("/v1/formeditassocbuttons", GetController().PostFormEditAssocButton)
		v1.PATCH("/v1/formeditassocbuttons/:id", GetController().UpdateFormEditAssocButton)
		v1.PUT("/v1/formeditassocbuttons/:id", GetController().UpdateFormEditAssocButton)
		v1.DELETE("/v1/formeditassocbuttons/:id", GetController().DeleteFormEditAssocButton)

		v1.GET("/v1/formfields", GetController().GetFormFields)
		v1.GET("/v1/formfields/:id", GetController().GetFormField)
		v1.POST("/v1/formfields", GetController().PostFormField)
		v1.PATCH("/v1/formfields/:id", GetController().UpdateFormField)
		v1.PUT("/v1/formfields/:id", GetController().UpdateFormField)
		v1.DELETE("/v1/formfields/:id", GetController().DeleteFormField)

		v1.GET("/v1/formfielddates", GetController().GetFormFieldDates)
		v1.GET("/v1/formfielddates/:id", GetController().GetFormFieldDate)
		v1.POST("/v1/formfielddates", GetController().PostFormFieldDate)
		v1.PATCH("/v1/formfielddates/:id", GetController().UpdateFormFieldDate)
		v1.PUT("/v1/formfielddates/:id", GetController().UpdateFormFieldDate)
		v1.DELETE("/v1/formfielddates/:id", GetController().DeleteFormFieldDate)

		v1.GET("/v1/formfielddatetimes", GetController().GetFormFieldDateTimes)
		v1.GET("/v1/formfielddatetimes/:id", GetController().GetFormFieldDateTime)
		v1.POST("/v1/formfielddatetimes", GetController().PostFormFieldDateTime)
		v1.PATCH("/v1/formfielddatetimes/:id", GetController().UpdateFormFieldDateTime)
		v1.PUT("/v1/formfielddatetimes/:id", GetController().UpdateFormFieldDateTime)
		v1.DELETE("/v1/formfielddatetimes/:id", GetController().DeleteFormFieldDateTime)

		v1.GET("/v1/formfieldfloat64s", GetController().GetFormFieldFloat64s)
		v1.GET("/v1/formfieldfloat64s/:id", GetController().GetFormFieldFloat64)
		v1.POST("/v1/formfieldfloat64s", GetController().PostFormFieldFloat64)
		v1.PATCH("/v1/formfieldfloat64s/:id", GetController().UpdateFormFieldFloat64)
		v1.PUT("/v1/formfieldfloat64s/:id", GetController().UpdateFormFieldFloat64)
		v1.DELETE("/v1/formfieldfloat64s/:id", GetController().DeleteFormFieldFloat64)

		v1.GET("/v1/formfieldints", GetController().GetFormFieldInts)
		v1.GET("/v1/formfieldints/:id", GetController().GetFormFieldInt)
		v1.POST("/v1/formfieldints", GetController().PostFormFieldInt)
		v1.PATCH("/v1/formfieldints/:id", GetController().UpdateFormFieldInt)
		v1.PUT("/v1/formfieldints/:id", GetController().UpdateFormFieldInt)
		v1.DELETE("/v1/formfieldints/:id", GetController().DeleteFormFieldInt)

		v1.GET("/v1/formfieldselects", GetController().GetFormFieldSelects)
		v1.GET("/v1/formfieldselects/:id", GetController().GetFormFieldSelect)
		v1.POST("/v1/formfieldselects", GetController().PostFormFieldSelect)
		v1.PATCH("/v1/formfieldselects/:id", GetController().UpdateFormFieldSelect)
		v1.PUT("/v1/formfieldselects/:id", GetController().UpdateFormFieldSelect)
		v1.DELETE("/v1/formfieldselects/:id", GetController().DeleteFormFieldSelect)

		v1.GET("/v1/formfieldstrings", GetController().GetFormFieldStrings)
		v1.GET("/v1/formfieldstrings/:id", GetController().GetFormFieldString)
		v1.POST("/v1/formfieldstrings", GetController().PostFormFieldString)
		v1.PATCH("/v1/formfieldstrings/:id", GetController().UpdateFormFieldString)
		v1.PUT("/v1/formfieldstrings/:id", GetController().UpdateFormFieldString)
		v1.DELETE("/v1/formfieldstrings/:id", GetController().DeleteFormFieldString)

		v1.GET("/v1/formfieldtimes", GetController().GetFormFieldTimes)
		v1.GET("/v1/formfieldtimes/:id", GetController().GetFormFieldTime)
		v1.POST("/v1/formfieldtimes", GetController().PostFormFieldTime)
		v1.PATCH("/v1/formfieldtimes/:id", GetController().UpdateFormFieldTime)
		v1.PUT("/v1/formfieldtimes/:id", GetController().UpdateFormFieldTime)
		v1.DELETE("/v1/formfieldtimes/:id", GetController().DeleteFormFieldTime)

		v1.GET("/v1/formgroups", GetController().GetFormGroups)
		v1.GET("/v1/formgroups/:id", GetController().GetFormGroup)
		v1.POST("/v1/formgroups", GetController().PostFormGroup)
		v1.PATCH("/v1/formgroups/:id", GetController().UpdateFormGroup)
		v1.PUT("/v1/formgroups/:id", GetController().UpdateFormGroup)
		v1.DELETE("/v1/formgroups/:id", GetController().DeleteFormGroup)

		v1.GET("/v1/formsortassocbuttons", GetController().GetFormSortAssocButtons)
		v1.GET("/v1/formsortassocbuttons/:id", GetController().GetFormSortAssocButton)
		v1.POST("/v1/formsortassocbuttons", GetController().PostFormSortAssocButton)
		v1.PATCH("/v1/formsortassocbuttons/:id", GetController().UpdateFormSortAssocButton)
		v1.PUT("/v1/formsortassocbuttons/:id", GetController().UpdateFormSortAssocButton)
		v1.DELETE("/v1/formsortassocbuttons/:id", GetController().DeleteFormSortAssocButton)

		v1.GET("/v1/options", GetController().GetOptions)
		v1.GET("/v1/options/:id", GetController().GetOption)
		v1.POST("/v1/options", GetController().PostOption)
		v1.PATCH("/v1/options/:id", GetController().UpdateOption)
		v1.PUT("/v1/options/:id", GetController().UpdateOption)
		v1.DELETE("/v1/options/:id", GetController().DeleteOption)

		v1.GET("/v1/rows", GetController().GetRows)
		v1.GET("/v1/rows/:id", GetController().GetRow)
		v1.POST("/v1/rows", GetController().PostRow)
		v1.PATCH("/v1/rows/:id", GetController().UpdateRow)
		v1.PUT("/v1/rows/:id", GetController().UpdateRow)
		v1.DELETE("/v1/rows/:id", GetController().DeleteRow)

		v1.GET("/v1/tables", GetController().GetTables)
		v1.GET("/v1/tables/:id", GetController().GetTable)
		v1.POST("/v1/tables", GetController().PostTable)
		v1.PATCH("/v1/tables/:id", GetController().UpdateTable)
		v1.PUT("/v1/tables/:id", GetController().UpdateTable)
		v1.DELETE("/v1/tables/:id", GetController().DeleteTable)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtable/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
