package csv_make

import (
	    "github.com/sawaq7/go14_ver1/storage2"
//	    "fmt"
	    "net/http"

	    "github.com/sawaq7/go14_ver1/general/type5"
//	    "github.com/sawaq7/go14_ver1/general/process3"
	    "github.com/sawaq7/go14_ver1/general/datastore5/trans3"
//	    "strings"
        "github.com/sawaq7/go14_ver1/general/html5"
        "html/template"

        "os"

        "cloud.google.com/go/datastore"
        "context"

                                                  )

func Csv_make(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_make start \n" )

    var bucket ,filename  ,project_name string

    var csv_inf_first type5.Csv_Inf

///
/// 入力データをGET 　
///

    filename = r.FormValue("file_name")  //  get new file name

///
///    get project name
///

    project_name = os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

	query := datastore.NewQuery("Storage_B_O_Temp")

	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err := client.GetAll(ctx, query , &storage_b_o_temp );  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

           project_name       = storage_b_o_tempw.Project_Name    //  set project name
           bucket   = storage_b_o_tempw.Bucket_Name    // set bucket name

        }

	  }

	}

///
/// 　　　make csv file　
///

     writer := storage2.File_Create2( w ,r ,bucket ,filename ,"text/plain" )
//     writer := storage2.File_Create2( w ,r ,bucket ,filename ,"text/csv" )

	defer writer.Close()

	csv_inf := trans3.Csv_inf ( w ,r )  ///      get csv inf.

	colum_num := int ( csv_inf[0].Column_Num )  // set column number
	filename2 := csv_inf[0].File_Name            //  set file name
	first_id  :=  csv_inf[0].Id


//	fmt.Fprintf( w, "csv_make : colum_num %v\n", colum_num )
//	fmt.Fprintf( w, "csv_make : record_num %v\n", len(csv_inf) )

	record := make ( []string ,colum_num )   //　allocate work area for records

///
///    make csv file
///

    for _ , csv_infw := range csv_inf {

      for ii := 0 ; ii < colum_num ; ii++ {  //　set record data

        switch ii {

          case 0 :

            record[ii] = csv_infw.Column1

          break;

          case 1 :

            record[ii] = csv_infw.Column2

          break;

          case 2 :

            record[ii] = csv_infw.Column3

          break;

          case 3 :

            record[ii] = csv_infw.Column4

          break;

          case 4 :

            record[ii] = csv_infw.Column5

          break;

          case 5 :

            record[ii] = csv_infw.Column6

          break;

          case 6 :

            record[ii] = csv_infw.Column7

          break;

          case 7 :

            record[ii] = csv_infw.Column8

          break;

          case 8 :

            record[ii] = csv_infw.Column9

          break;

          case 9 :

            record[ii] = csv_infw.Column10

          break;

        }
      }

//      fmt.Fprintf( w, "csv_make : record %v\n", record )

      storage2.File_Write_Csv2 ( w  ,writer ,record )  //  write record with csv format
//      storage2.File_Write_Csv ( w  ,writer ,record )

	}

///
///　　modify  file name
///

    if filename != filename2 {

       key := datastore.IDKey("Csv_Inf", first_id, nil)

       if err := client.Get(ctx, key , &csv_inf_first ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	   }

       csv_inf_first.File_Name = filename

       if _, err := client.Put(ctx, key, &csv_inf_first ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	   }
    }

///
///　　  show csv inf. on web
///

     csv_inf_view := trans3.Csv_inf ( w ,r )  ///   set csv inf.

     monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // set template

     err = monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_make normal end \n" )

}
