package process3

import (

	    "net/http"
//	    "fmt"
//	    "html/template"
	    "github.com/sawaq7/go14_ver1/general/datastore5/reformat"

//	    "github.com/sawaq7/go14_ver1/general/html5"
	    "github.com/sawaq7/go14_ver1/general/type5"
	    "github.com/sawaq7/go14_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go14_ver1/general/datastore5/set1"
	    "github.com/sawaq7/go14_ver1/storage2"
	    "io"
	    "strings"
	    "bufio"
	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///     add one-column which is selected
///

func Csv_column_join(w http.ResponseWriter, r *http.Request , filename string ,column_no int ) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter
//     IN  filename 　  : file-name
//     IN  column_no　  : column-no which is added

//    fmt.Fprintf( w, "process3.csv_column_join start \n" )

    var index   int64

    var column   int

    var str_work[10] string

    var bucket string

    csv_inf_join := make( []string , 0)

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      project_name = "sample-7777"

	}
    ctx := context.Background()

	query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, project_name)
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

    if _, err := client.GetAll(ctx, query , &storage_b_o_temp); err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

//           project_name       = storage_b_o_tempw.Project_Name
           bucket   = storage_b_o_tempw.Bucket_Name      //   set bucket-name
//           filename = storage_b_o_tempw.Object_Name    //   set object-name

        }
	  }
	}

///
/// 　　　modify csv-inf.　
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///      get csv-inf.

    csv_inf2 := reformat.Csv_inf ( 1,csv_inf[0].Column_Num+1 ,csv_inf ,w ,r )
                                                           /// expand one-column-area
///
///      get csv-inf. for adding
///

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket ,filename , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

    defer reader.Close()

    csv_reader := bufio.NewReaderSize(reader, 4096)

    index = 0

    for {

        index ++
	    record ,err  := csv_reader.ReadString('\n')

//	    fmt.Fprintf( w, "csv_show : record %v\n", record )

	    if err == io.EOF {

	      break

		}
		if err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
	      return

		}
		if index == 1 {

		  column = strings.Count( record ,",") + 1   //  cal. column-number

//		  fmt.Fprintf( w, "csv_show : column %v\n", column )

		}

		str := strings.Split ( record, ","  )

//		fmt.Fprintf( w, "csv_show : str %v\n", str )

		for ii := 0 ; ii < column ; ii++ {

           str_work[ii] = str[ii]

        }

        csv_inf_join = append( csv_inf_join ,str_work[column_no-1] )
                                                      ///    make csv-inf. for adding

    }

    csv_inf_new := set1.Csv_inf (  csv_inf2 ,csv_inf_join ,int(csv_inf2[0].Column_Num) , w ,r )
                                                        ///   add one-column-inf. which was maked in existing-csv-inf.

///
/// 　　　set new-csv-inf. in d.s.
///

    for _, csv_inf_neww := range csv_inf_new {

//   	  fmt.Fprintf( w, "process3.csv_column_join csv_inf_neww %v\n", csv_inf_neww )

      key := datastore.IDKey("Csv_Inf", csv_inf_neww.Id, nil)

      if _, err := client.Put(ctx, key, &csv_inf_neww ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

//	fmt.Fprintf( w, "process3.csv_column_join normal end \n" )


}
