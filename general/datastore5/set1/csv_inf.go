package set1

import (

	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go14_ver1/general/type5"

//	    "time"
                                                )

///
///    set csv file for column
///

func Csv_inf ( csv_inf []type5.Csv_Inf ,csv_inf_join []string ,column_no int ,w http.ResponseWriter, r *http.Request )  ([]type5.Csv_Inf ) {

//     IN  csv_inf  　　: basic csv-inf
//     IN  csv_inf_join : csv-inf for adding　　　
//     IN  column_no  　　: 対象の衁E
//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//     OUT csv_inf_new  : csvinf after change

//    fmt.Fprintf( w, "set1.csv_inf start %v\n" )
//    fmt.Fprintf( w, "set1.csv_inf csv_inf_join %v\n" ,csv_inf_join )
//    fmt.Fprintf( w, "set1.csv_inf column_no %v\n" ,column_no)

    str_work := make([]string ,10 )
	csv_inf2 := make([]type5.Csv_Inf, 0)

	index := 0

	for _, csv_infw := range csv_inf {

//	  fmt.Fprintf( w, "trans.csv_inf csv_infw %v\n" ,csv_infw)

      str_work[0]  = csv_infw.Column1
	  str_work[1]  = csv_infw.Column2
	  str_work[2]  = csv_infw.Column3
	  str_work[3]  = csv_infw.Column4
	  str_work[4]  = csv_infw.Column5
	  str_work[5]  = csv_infw.Column6
	  str_work[6]  = csv_infw.Column7
	  str_work[7]  = csv_infw.Column8
	  str_work[8]  = csv_infw.Column9
	  str_work[9]  = csv_infw.Column10

	  str_work[column_no-1]  = csv_inf_join[index]

	  index ++

      csv_inf2 = append(csv_inf2, type5.Csv_Inf { csv_infw.Id           ,
                                                     csv_infw.Line_No      ,
                                                     csv_infw.File_Name    ,
                                                     csv_infw.Column_Num   ,
                                                     str_work[0]           ,
                                                     str_work[1]           ,
                                                     str_work[2]           ,
                                                     str_work[3]           ,
                                                     str_work[4]           ,
                                                     str_work[5]           ,
                                                     str_work[6]           ,
                                                     str_work[7]           ,
                                                     str_work[8]           ,
                                                     str_work[9]           })




	}

//    fmt.Fprintf( w, "set1.csv_inf csv_inf2 %v\n" ,csv_inf2)

    return	csv_inf2
}
