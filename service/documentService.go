package service

import (
	"doc-lixiang/models"
	"github.com/pkg/errors"
)

func  DocumentDataById(documentId int) (data []models.LbpDocumentParameters, err error){

	//业务逻辑处理

	documentDao := &models.LbpDocumentParameters{}
	dataDoc, err := documentDao.GetLbpDocumentParametersAllById(documentId)
	if err !=nil {
		//加log记录信息
		return data, errors.New("未查询到数据")
	}


	return dataDoc, nil
}
