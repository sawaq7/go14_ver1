package trans

import (

	    "net/http"
//	    "fmt"
//	    "html/template"

	    "github.com/sawaq7/go14_ver1/client/sgh/type2"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///     get car inf which run on district
///

func Car_district( district_no int64  ,w http.ResponseWriter, r *http.Request )  ( car_district_view []type2.Car ) {

//     IN  district_no  : 地域No
//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

//     OUT car_district_view  :   slice of struct ( Car )

//    fmt.Fprintf( w, "trans.car_district start \n" )

     project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    query := datastore.NewQuery("Car").Order("Car_No")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	nil
	}

	car_district      := make([]type2.Car, 0, count)
	car_district_view = make([]type2.Car, 0)

    keys, err := client.GetAll(ctx, query , &car_district)

    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)

//		fmt.Fprintf( w, "d_district_area_show err \n" ,err)
		return	nil
	}

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, car_districtw := range car_district {

//	  fmt.Fprintf( w, "trans.car_district car_districtw %v\n" ,car_districtw)


      if district_no == car_districtw.District_No {

         car_district_view = append(car_district_view, type2.Car { keys_wk[pos]   ,
                                                          car_districtw.District_No    ,
                                                          car_districtw.District_Name  ,
                                                          car_districtw.Car_No         ,
                                                          car_districtw.Car_Name       ,
                                                          car_districtw.Car_Explain    ,
                                                          car_districtw.Number_Total   ,
                                                          car_districtw.Time_Total     ,
                                                          car_districtw.Productivity      })

      }
	}

    return	car_district_view
}

