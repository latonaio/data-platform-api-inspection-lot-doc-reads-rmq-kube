package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToHeaderDoc(rows *sql.Rows) (*[]HeaderDoc, error) {
	defer rows.Close()
	headerDoc := make([]HeaderDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &HeaderDoc{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &headerDoc, err
		}

		data := pm
		headerDoc = append(headerDoc, HeaderDoc{
			InspectionLot:            data.InspectionLot,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &headerDoc, nil
	}

	return &headerDoc, nil
}

func ConvertToOperationDoc(rows *sql.Rows) (*[]OperationDoc, error) {
	defer rows.Close()
	operationDoc := make([]OperationDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &OperationDoc{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &operationDoc, err
		}

		data := pm
		operationDoc = append(operationDoc, OperationDoc{
			InspectionLot:            data.InspectionLot,
			Operations:      		  data.Operations,
			OperationsItem:      	  data.OperationsItem,
			OperationID:      	 	  data.OperationID,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &operationDoc, nil
	}

	return &operationDoc, nil
}
