package test1

import (

	    "net/http"
	    "github.com/sawaq7/go14_ver1/client/tokura/suiri/cal4"
//	    "github.com/sawaq7/go14_ver1/client/tokura/suiri/type4"

//	    "github.com/sawaq7/go14_ver1/client/tokura/datastore4/trans2"
//	    "github.com/sawaq7/go14_ver1/basic/type3"
//	    "strconv"
	    "fmt"

//        "cloud.google.com/go/datastore"
//        "context"
//        "os"

                                                  )

func Test1(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_ds_cal start \n"  )

///
///
///

//      p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal4.Pipe_line1( w  ,r  )
 //     p_number  := cal4.Pipe_line1( w  ,r )
         p_number  := cal4.Pipe_line2( )
//     p_number  := 1
    fmt.Fprintf( w, "sky/pipe_line_ds_cal : p_number %v\n", p_number )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup %v\n", ad_eneup )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_enedown %v\n", ad_enedown )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glineup %v\n", ad_glineup )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_glinedown %v\n", ad_glinedown )
//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : ad_eneup len %v\n", len(ad_eneup) )

/// グラフ�E作�E ///

//    p_number := 1
//    ad_eneup := make([]type3.Point ,20 ,50)
//   ad_enedown := make([]type3.Point ,20 ,50)
//   ad_glineup := make([]type3.Point ,20 ,50)
//   ad_glinedown := make([]type3.Point ,20 ,50)

//   ad_eneup[0].X = 1.
//      ad_enedown[0].Y = 1.
//      ad_glineup[0].X = 1.
//      ad_glinedown[0].Y = 1.

//    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_ds_cal : f_name %v\n", f_name )

///
    f_name := "1"
    cal4.Pipe_line1_show_graf( w ,r ,f_name )

}
