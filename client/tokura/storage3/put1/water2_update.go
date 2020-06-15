package put1

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go14_ver1/storage2"
	    "bufio"

	    "io"
        "github.com/sawaq7/go14_ver1/client/tokura/storage3/struct_set"
	    "github.com/sawaq7/go14_ver1/client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"
	    "github.com/sawaq7/go14_ver1/general/type5"

                                                )

///                           　　　　　　　　　　　
///     update water-inf. which was selected in storage
///                          　　　　　　　　　　　

func Water2_update( w http.ResponseWriter, r *http.Request ,updid int64 ,water_inf type4.Water2 ) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter
//     IN     updid     : update-id
//     IN   water_inf   : slice of Water2 (struct)

//    fmt.Fprintf( w, "put1.water2_update start \n" )

    var lf_flag int64

    bucket := "sample-7777"
    filename1 := "Water2.txt"
    filename2 := "Water2_2.txt"

    lf_flag   = 1

///
/// 　　　rename file-name
///

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

///
///       open water-file which was renamed
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    sreader := bufio.NewReaderSize(reader, 4096)

///
///       create new-water-file
///

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()

    index := 0

    for {

//      fmt.Fprintf(w, "put1.water2_update : lndex %v\n", index )

      line ,_  := sreader.ReadString('\n')     // read one-record in Water2-file

      num := len(line)

//      fmt.Fprintf(w, "put1.water2_update : num %v\n", num )

      if num > 1 {

         index ++
//         fmt.Fprintf(w, "put1.water2_update : line %s\n", line )

         water2_struct := struct_set.Water2( w , line )

         general_work := make([]type5.General_Work, 1 )   // allocate work area for records
         general_work[0].Sw_Work    = writer     //　   set storage-writer
         general_work[0].Int64_Work = lf_flag    //　  set line-feed-flag

         if  water2_struct.Id == updid {

           storage2.Storage_basic( "write2" ,general_work ,water_inf , w , r  )

//         storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

         } else {

           storage2.Storage_basic( "write2" ,general_work ,water2_struct , w , r  )

//         storage2.File_Write_Struct ( w ,writer ,lf_flag ,water2_struct )

         }

      } else if num == 0 {

//          io.WriteString(w, "\n put1.water2_update : data end \n")

         break

      }
   }

///
///       delete temp.-file
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}
