package put1

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go14_ver1/storage2"

	    "github.com/sawaq7/go14_ver1/client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"

                                                )

///                           　　　　　　　　　　　                        　　　　　　　　　　　
///      if Water2-file don't exist in storage , create new-Water2-file
///                          　　　　　　　　　

func Water2_new( w http.ResponseWriter, r *http.Request ,water_inf type4.Water2 ) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter
//     IN   water_inf   : slice of Water2 (struct)

//    fmt.Fprintf( w, "put1.water2_new start \n" )

    bucket := "sample-7777"
    filename1 := "Water2.txt"

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()

    lf_flag := int64( 0 )

    water_inf.Id = int64( 1 )

    storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

    return

}
