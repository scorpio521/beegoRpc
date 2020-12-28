package dto

import (
	"errors"
	"github.com/astaxie/beego/validation"
)

type DocumentInput struct {
	DocumentId         int                              `json:"document_id" form:"document_id" `

}

//通过自定义参数验证
func (o *DocumentInput) BindingValidParams() error {

	valid := validation.Validation{}
	valid.Required(o.DocumentId, "period").Message("文档id不能为空!")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			return errors.New(err.Key + "：" + err.Message)
		}
	}

	return nil
}

