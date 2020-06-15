package csv_show_test1

import (

	    "net/http"
	    "fmt"


//	    "github.com/sawaq7/go14_ver1/storage2/get"
	    "github.com/sawaq7/go14_ver1/storage2"
	    "strconv"
	    "strings"
	    "io"
	    "bufio"


	    "github.com/sawaq7/go14_ver1/general/type5"

	    "github.com/sawaq7/go14_ver1/general/datastore5/initialize"


        "cloud.google.com/go/datastore"
        "context"
        "os"


                                                  )


func Csv_show_test1(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

    fmt.Fprintf( w, "csv_show_test start \n" )

    var bucket ,filename string

    var index   int64

    var column int


    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "csv_show : line_no %v\n", line_no )

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "csv_show : select_id %v\n", select_id )

     projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      projectID = "sample-7777"

	}

    ctx := context.Background()

    query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {

         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
    }

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp ) ;  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {
           bucket    = storage_b_o_tempw.Bucket_Name

        }
	  }
	}

//    fmt.Fprintf( w, "csv_show : bucket2: %v\n", bucket2 )

///
///     get file-name
///

	objects :=  storage2.Object_List ( w  ,r , bucket )

//    fmt.Fprintf( w, "csv_show : objects: %v\n", objects )

    for pos2 , objectsw := range objects {

      if pos2 == select_id-1 {

        filename = objectsw


      }

    }

///
///      initialize csv-inf.
///

    initialize.Csv_inf (w , r )

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0

    for {

        index ++
	    record ,err  := csv_reader.ReadString('\n')

	    fmt.Fprintf( w, "csv_show : record %v\n", record )

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}

		record = strings.Replace( record, ",", " ", -1)     ///   change　delimiter

		if index == 1 {   //   get column-number

		  column = strings.Count( record ," ") + 1

		  fmt.Fprintf( w, "csv_show : column %v\n", column )

		}

	}

//	fmt.Fprintf( w, "csv_show : normal end \n" )

}
