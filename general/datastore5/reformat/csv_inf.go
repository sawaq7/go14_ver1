package reformat

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go14_ver1/general/type5"
	    "github.com/sawaq7/go14_ver1/basic/array"
//	    "time"
                                                )

///
/// csvファイルのフォーマットを変更する  change format of csv-file
///

func Csv_inf(funct int64 ,column_no int64 ,csv_inf []type5.Csv_Inf ,w http.ResponseWriter, r *http.Request )  ([]type5.Csv_Inf ) {

//     IN  funct  　　　: function
//     　　　　　delete
//     　　　　　insert
//     IN  column_no  　　
//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//     OUT csv_inf_new  : after reformating csv inf

//    fmt.Fprintf( w, "reformat.csv_inf start \n" )
//    fmt.Fprintf( w, "reformat.csv_inf funct \n" ,funct )
//    fmt.Fprintf( w, "reformat.csv_inf column_no \n" ,column_no)

    var column_wk int64

    str_work := make([]string ,10 )   // allocate work-area
	csv_inf2 := make([]type5.Csv_Inf, 0)

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

      str_work_new := array.Array_string ( funct ,column_no ,str_work )

      if funct == 0 {

		column_wk = csv_infw.Column_Num - 1

	  } else {

	  	column_wk = csv_infw.Column_Num + 1

	  }

      csv_inf2 = append(csv_inf2, type5.Csv_Inf { csv_infw.Id           ,
                                                     csv_infw.Line_No      ,
                                                     csv_infw.File_Name    ,
                                                     column_wk             ,
                                                     str_work_new[0]       ,
                                                     str_work_new[1]       ,
                                                     str_work_new[2]       ,
                                                     str_work_new[3]       ,
                                                     str_work_new[4]       ,
                                                     str_work_new[5]       ,
                                                     str_work_new[6]       ,
                                                     str_work_new[7]       ,
                                                     str_work_new[8]       ,
                                                     str_work_new[9]           })
	  }

    return	csv_inf2
}
