package process4

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go14_ver1/client/reserve/html6"
//	    "github.com/sawaq7/go14_ver1/client/reserve/type6"
	    "github.com/sawaq7/go14_ver1/client/reserve/datastore6/trans5"
                                                )


func Payment_register(w http.ResponseWriter, r *http.Request ,guest_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  guest_no     : guest no

//    fmt.Fprintf( w, "payment_register start \n" )

///
///     set template
///

    monitor := template.Must(template.New("html").Parse(html6.Payment_register))

//    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

///
///      get Guest_payment in d.s.
///

    guest_payment_slice := trans5.Guest_payment ( guest_no , w , r  )

//    value, _ := d_area_view.([]type2.D_Area)

///
///     show payment inf. on web
///

//   fmt.Fprintf( w, "payment_register d_area_view %v\n" ,d_area_view)

//	err := monitor.Execute(w, value)

	err := monitor.Execute(w, guest_payment_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

