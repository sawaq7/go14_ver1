package trans3

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go14_ver1/general/type5"

	    "cloud.google.com/go/datastore"
        "context"
        "os"

                                                )

///                           ///
///    get copy list data
///                          ///

func Ds_copy_list( w http.ResponseWriter, r *http.Request )  ([]type5.Ds_Copy_List ) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//      OUT             : slice of Ds_Copy_List (struct)

//    fmt.Fprintf( w, "trans.Ds_copy_list start \n" )

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "trans.Ds_copy_list :  projectID unset \n"  )

      project_name = "sample-7777"

	}

    ctx := context.Background()

	query := datastore.NewQuery("Ds_Copy_List")

	client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return  nil
    }

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	ds_copy_list      := make([]type5.Ds_Copy_List, 0, count)

	ds_copy_list_view := make([]type5.Ds_Copy_List, 0)

    keys, err := client.GetAll(ctx, query , &ds_copy_list)

    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)
       return	nil
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, ds_copy_listw := range ds_copy_list {


         ds_copy_list_view = append(ds_copy_list_view, type5.Ds_Copy_List { keys_wk[pos] ,
                                                                    ds_copy_listw.Basic_Name  ,
                                                                    ds_copy_listw.Copy_Name   ,
                                                                    ds_copy_listw.New_Name    ,

                                                                                                 })

	}

    return	ds_copy_list_view
}
