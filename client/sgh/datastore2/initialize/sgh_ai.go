package initialize

import (

	      "net/http"
//	      "fmt"

	      "github.com/sawaq7/go14_ver1/client/sgh/type2"

	      "cloud.google.com/go/datastore"
	      "context"
	      "os"

                                                )

///
///　　　コースNoの計算式を削除
///

func Sgh_ai( course_no int64 ,w http.ResponseWriter, r *http.Request )   {

//     IN  course_no  　: コースNo
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//   fmt.Fprintf( w, "initialize.sgh_ai start \n" )

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    query := datastore.NewQuery("Sgh_Ai").Order("Course_No")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "initialize.sgh_ai count \n" ,count )

	sgh_ai     := make([]type2.Sgh_Ai, 0, count)

	keys, err := client.GetAll(ctx, query , &sgh_ai)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

    keys_wk := make([]int64, count)

    for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }
///
///        delete expression of course-no
///

	for pos2, sgh_aiw := range sgh_ai {

      if  course_no == sgh_aiw.Course_No {

        key := datastore.IDKey("Sgh_Ai", keys_wk[pos2], nil)

        if err := client.Delete(ctx, key ); err != nil {

		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	    }
	  }
//      fmt.Fprintf( w, "initialize.sgh_ai pos2 %v   \n" , pos2  )

    }

	return
}
