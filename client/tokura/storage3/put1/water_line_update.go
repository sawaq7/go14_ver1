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

                                                )

///                           　　　　　　　　　　　
///      change water-line-inf. which was selected in storage
///                          　　　　　　　　　　　

func Water_line_update( w http.ResponseWriter, r *http.Request ,updid int64 ,wname string ,water_inf type4.Water_Line ) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter
//     IN     updid     : update-id
//     IN    wname      : water-name
//     IN   water_inf   : slice of Water_Line (struct)

//   fmt.Fprintf( w, "put1.water_line_update start \n" )

    var lf_flag int64

    bucket := "sample-7777"
    filename1 := "Water_Line.txt"
    filename2 := "Water_Line_2.txt"

    lf_flag   = 1

///
/// 　　　rename file-name
///

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

///
///        open water-file which was renamed
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    sreader := bufio.NewReaderSize(reader, 4096)

///
///        create new-water-file
///

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()

    index := 0

    for {

//      fmt.Fprintf(w, "put1.water_line_update : lndex %v\n", index )

      line ,_  := sreader.ReadString('\n')    // read one-record in Water_Line-file

      num := len(line)

//     fmt.Fprintf(w, "put1.water_line_update : num %v\n", num )

      if num > 1 {

         index ++
//         fmt.Fprintf(w, "put1.water_line_update : line %s\n", line )

         water_line_struct := struct_set.Water_line( w , line )

       if  water_line_struct.Id   == updid &&
           water_line_struct.Name == wname      {

         storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

       } else {

         storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_line_struct )

       }

      } else if num == 0 {

//          io.WriteString(w, "\n put1.water_line_update : data end \n")

         break

      }
   }

///
///       delete temp.-file
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}
