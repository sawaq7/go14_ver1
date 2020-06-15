package initialize

import (

        "cloud.google.com/go/datastore"

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go14_ver1/general/type5"
	    "context"
	    "os"

                                                )

///                                   ///
/// delete all data of temporary file ///
///                                   ///

func Storage_b_o_temp(w http.ResponseWriter, r *http.Request )   {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//    fmt.Fprintf( w, "init/storage_b_o_temp start \n" )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	query := datastore.NewQuery("Storage_B_O_Temp")

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "init/storage_b_o_temp count \n" ,count )

	storage_b_o_temp     := make([]type5.Storage_B_O_Temp, 0, count)

	keys, err := client.GetAll(ctx, query , &storage_b_o_temp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

	for _, keysw := range keys {

      if err := client.Delete(ctx, datastore.IDKey("Storage_B_O_Temp", keysw.ID, nil)); err != nil {

		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	  }

    }

	return

}
