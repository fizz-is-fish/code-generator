package {{.packageName}}

import (
	"git.zuoyebang.cc/pkg/golib/v2/base"
	"github.com/gin-gonic/gin"
	"hxx-aiclass-course/components"
	"hxx-aiclass-course/components/dto/dtobigdataschool"
	"hxx-aiclass-course/service/bigdata"
)

func {{.methodName}}(ctx *gin.Context) {
	var param dtobigdataschool.SchoolCommonReq
	if err := ctx.ShouldBind(&param); err != nil {
		base.RenderJsonFail(ctx, components.ErrorParamInvalid)
		return
	}
	res, err := bigdata.NewBigDataSchoolSvc().LearningAnalysis(ctx, &param)
	if err != nil {
		base.RenderJsonFail(ctx, err)
		return
	}
	base.RenderJsonSucc(ctx, res)
}
