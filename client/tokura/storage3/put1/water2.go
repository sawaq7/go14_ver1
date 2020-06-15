package put1

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go14_ver1/storage2"
	    "bufio"

	    "io"

	    "github.com/sawaq7/go14_ver1/client/tokura/suiri/type4"
	    "cloud.google.com/go/storage"

                                                )

///                           　　　　　　　　　　　
///      put water-inf. which was selected in storage
///                          　　　　　　　　　　　

func Water2( w http.ResponseWriter, r *http.Request ,water_inf type4.Water2 ) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter
//     IN   water_inf   : slice of Water2 (struct)

//    fmt.Fprintf( w, "put1.water2 start \n" )

    var lf_flag int64

    bucket := "sample-7777"
    filename1 := "Water2.txt"
    filename2 := "Water2_2.txt"

///
/// 　　　rename file-name
///

    storage2.File_Rename ( w ,r ,bucket ,filename1 ,filename2 )

///
///      open water-file which was renamed
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename2 , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

//    reader := storage2.File_Open(w ,r ,bucket ,filename2)

    sreader := bufio.NewReaderSize(reader, 4096)

///
///     create new-water-file
///

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,filename1 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

//    writer := storage2.File_Create( w ,r ,bucket ,filename1 )

    defer writer.Close()

    index := 0

    lf_flag = 0

///
///      read Water2-file
///

    for {

//      fmt.Fprintf(w, "put1.water2 : lndex %v\n", index )

      line ,_  := sreader.ReadString('\n') // read one-record in Water2-file

      num := len(line)

//      fmt.Fprintf(w, "put1.water2 : num %v\n", num )

      if num > 1 {

         index ++
//         fmt.Fprintf(w, "put1.water2 : line %s\n", line )
         storage2.File_Write_Struct ( w ,writer ,lf_flag ,line )

      } else if num == 0 {

          index ++
          lf_flag = 1

          water_inf.Id = int64(index)

          storage2.File_Write_Struct ( w ,writer ,lf_flag ,water_inf )

//         io.WriteString(w, "\n put1.water2 : data end \n")

         break

      }
   }

///
///      delete temp.-file
///

   storage2.File_Delete ( w , r ,bucket ,filename2 )

  return

}
