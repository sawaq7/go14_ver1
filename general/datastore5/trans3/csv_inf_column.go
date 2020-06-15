package trans3

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go14_ver1/general/type5"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///    get csv data in column
///

func Csv_inf_column ( w http.ResponseWriter, r *http.Request ,column_no int )  ( []string ) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//     IN  column_no　      :

//     OUT csv_inf_column  : csv-inf-column
//    fmt.Fprintf( w, "trans.csv_inf_column start \n" )

    var string_wk string

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return nil
    }

	query := datastore.NewQuery("Csv_Inf").Order("Line_No")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	csv_inf      := make([]type5.Csv_Inf, 0, count)
	csv_inf_column := make([]string, 0)

    _, err = client.GetAll(ctx, query , &csv_inf)

    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	for _, csv_infw := range csv_inf {

///    set column-inf.

     switch column_no {

          case 1 :

            string_wk = csv_infw.Column1

          break;

          case 2 :

            string_wk = csv_infw.Column2

          break;

          case 3 :

            string_wk = csv_infw.Column3

          break;

          case 4 :

            string_wk = csv_infw.Column4

          break;

          case 5 :

            string_wk = csv_infw.Column5

          break;

          case 6 :

            string_wk = csv_infw.Column6

          break;

          case 7 :

            string_wk = csv_infw.Column7

          break;

          case 8 :

            string_wk = csv_infw.Column8

          break;

          case 9 :

            string_wk = csv_infw.Column9

          break;

          case 10 :

            string_wk = csv_infw.Column10

          break;

      }

      csv_inf_column = append( csv_inf_column, string_wk )

	}

//    fmt.Fprintf( w, "trans.csv_inf_column normal end )

    return	csv_inf_column
}
