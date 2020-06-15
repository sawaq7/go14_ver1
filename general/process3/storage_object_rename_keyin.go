package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go14_ver1/general/html5"

	    "github.com/sawaq7/go14_ver1/general/type5"

	    "github.com/sawaq7/go14_ver1/storage2"
	    "strconv"

	    "os"
	    "log"
        "cloud.google.com/go/datastore"
        "context"
                                                )

func Storage_object_rename_keyin(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//    fmt.Fprintf( w, "process3.storage_object_rename_keyin start \n" )

    var sdmy string

    var storage_b_o_temp2 type5.Storage_B_O_Temp

    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "storage_object_show : line_no %v\n", line_no )

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "storage_object_show : select_id %v\n", select_id )

     projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      projectID = "sample-7777"

	}
    ctx := context.Background()

    query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       log.Fatalf("Failed to create client: %v", err)
    }

	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp) ; err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

           storage_b_o_temp2.Line_No = 1
           storage_b_o_temp2.Project_Name = storage_b_o_tempw.Project_Name
           storage_b_o_temp2.Bucket_Name = storage_b_o_tempw.Bucket_Name

           updid := datastore.IDKey( "Storage_B_O_Temp",int64(pos+1) , nil)

           objects :=  storage2.Object_List ( w  ,r , storage_b_o_tempw.Bucket_Name )

           for pos2 , objectsw := range objects {

             if pos2+1 == select_id {

                storage_b_o_temp2.Object_Name = objectsw

             }

           }

           if _, err := client.Put(ctx, updid, &storage_b_o_temp2 ); err != nil {

             log.Fatalf("Failed to save task: %v", err)
             return
           }

        }
	  }
	}

///
///    set template-inf.
///

     monitor := template.Must(template.New("html").Parse(html5.Storage_object_rename_keyin))

///
///       show storage-inf. on web
///

	err = monitor.Execute(w, sdmy)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}
