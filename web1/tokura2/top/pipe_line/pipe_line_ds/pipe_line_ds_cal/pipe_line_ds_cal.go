package pipe_line_ds_cal

import (

	    "net/http"
//	    "github.com/sawaq7/go14_ver1/basic/graph1"
	    "github.com/sawaq7/go14_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go14_ver1/client/tokura/datastore4"
//	    "strconv"
//	    "fmt"
        "github.com/sawaq7/go14_ver1/client/tokura/suiri/cal"

//        "github.com/sawaq7/go14_ver1/storage2"

        "cloud.google.com/go/datastore"
        "context"
        "os"

//        "github.com/sawaq7/go14_ver1/basic/type3"

                                                  )

func Pipe_line_ds_cal(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_ds_cal start \n"  )



///    get key-in data

   var water type4.Water2

   project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)

    query := datastore.NewQuery("Water2_Temp").Order("Name")

    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal  \n" ,count )

    water_temp      := make([]type4.Water2, 0, count)

    keys, err := client.GetAll(ctx, query , &water_temp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : len(water) %v\n", len(water) )

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos2, waterw := range water_temp {

       water.Id   = keys_wk[pos2]

	   water.Name = waterw.Name
	   water.High =  waterw.High
	   water.Roughness_Factor = waterw.Roughness_Factor

    }

///
///      get water-line inf.
///

//    water_line := trans2.Water_line (1  ,water.Name , w ,r )

     water_line := datastore4.Datastore_tokura( "Water_Line"  ,"trans"  ,water.Name , w , r  )

     value, _ := water_line.([]type4.Water_Line)

//     _= value

///
///        cal. water-slope-line
///

    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal.Pipe_line1( water  ,value  )

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : p_number %v\n", p_number )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup %v\n", ad_eneup )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_enedown %v\n", ad_enedown )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glineup %v\n", ad_glineup )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glinedown %v\n", ad_glinedown )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup len %v\n", len(ad_eneup) )

///
///        make water-slope-line otherts,various-data
///

   f_name :=  cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,
                                                  ad_enedown  ,ad_glineup ,ad_glinedown )

//    bucket := "sample-7777"

//    storage2.Storage_basic( "show1" ,bucket ,f_name , w , r  )

    cal.Pipe_line1_show_graf( w ,r ,f_name )

}
