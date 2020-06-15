package csv_show

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go14_ver1/storage2/vaccine1"
	    "github.com/sawaq7/go14_ver1/storage2"
	    "strconv"
	    "strings"
	    "io"
	    "bufio"
	    "html/template"

	    "github.com/sawaq7/go14_ver1/general/type5"
	    "github.com/sawaq7/go14_ver1/general/html5"
	    "github.com/sawaq7/go14_ver1/general/datastore5/initialize"
	    "github.com/sawaq7/go14_ver1/general/datastore5/trans3"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func Csv_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_show start \n" )

    var bucket ,filename string

    var index   int64

    var column  int

    var str_work[10] string

    var csv_inf type5.Csv_Inf

    var storage_b_o_temp2 type5.Storage_B_O_Temp

    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "csv_show : line_no %v\n", line_no )

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "csv_show : select_id %v\n", select_id )

///
///   バケチE�E��E�名をゲチE�E��E�
///

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
///       initialize csv-file
///

    initialize.Csv_inf (w , r )

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0 // レコードカウンターをイニシャライズ

    warn_flag := 0  // ワーニングフラグのイニシャライズ

    for {

	    record ,err  := csv_reader.ReadString('\n')

//	    fmt.Fprintf( w, "csv_show : record %v\n", record )

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}

        record = strings.Replace( record, ",", " ", -1)

        column = strings.Count( record ," ") + 1

//        fmt.Fprintf( w, "csv_show : column %v\n", strings.Count( record ," ") + 1 )


        if  column > 1 {
          index ++
          str := strings.Fields(record)

//		fmt.Fprintf( w, "csv_show : str %v\n", str )

		  for ii := 0 ; ii < column ; ii++ {

             str_work[ii] = str[ii]

          }

///    ワークエリア(チE�E�Eタストア�E�E�E�にcsv惁E�E��E�をセチE�E��E�

          csv_inf.Line_No    = index
          csv_inf.File_Name  = filename
          csv_inf.Column_Num = int64(column)
	      csv_inf.Column1    = str_work[0]
	      csv_inf.Column2    = str_work[1]
	      csv_inf.Column3    = str_work[2]
	      csv_inf.Column4    = str_work[3]
	      csv_inf.Column5    = str_work[4]
	      csv_inf.Column6    = str_work[5]
	      csv_inf.Column7    = str_work[6]
	      csv_inf.Column8    = str_work[7]
	      csv_inf.Column9    = str_work[8]
	      csv_inf.Column10   = str_work[9]

          new_key := datastore.IncompleteKey("Csv_Inf", nil)

//    fmt.Fprintf(w, "storage_object_copy_excute: new_key %v\n", new_key )

         _, err = client.Put(ctx, new_key, &csv_inf )
	     if err != nil {
		   http.Error(w,err.Error(), http.StatusInternalServerError)
		   return
	     }

       } else {

      	 warn_flag = 1

       }

	}

    initialize.Storage_b_o_temp (w , r ) // ini. storage-temp.-file

    storage_b_o_temp2.Line_No =  1
    storage_b_o_temp2.Project_Name = projectID
    storage_b_o_temp2.Bucket_Name = bucket
    storage_b_o_temp2.Object_Name = filename

    new_key := datastore.IncompleteKey("Storage_B_O_Temp", nil)

    if _, err := client.Put(ctx, new_key, &storage_b_o_temp2 ); err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError )
    }

///
///　　 revise file-inf.
///

    if warn_flag == 1 {

       vaccine1.File_Pack ( w , r ,bucket  ,filename )

	}

///
///　　show csv-inf. on web
///

     csv_inf_view := trans3.Csv_inf ( w ,r )

     monitor := template.Must( template.New("html").Parse( html5.Csv_show ))

     err = monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_show : normal end \n" )

}
