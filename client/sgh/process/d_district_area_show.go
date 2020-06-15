package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go14_ver1/client/sgh/html2"
	    "github.com/sawaq7/go14_ver1/client/sgh/datastore2"
	    "github.com/sawaq7/go14_ver1/client/sgh/type2"
                                                )

///
///     show  area inf. in d.s.
///

func D_district_area_show(w http.ResponseWriter, r *http.Request ,district_no int64) {

//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter
//     IN 　district_no : district no

//    fmt.Fprintf( w, "d_district_area_show start \n" )

//     set template

//     monitor := template.Must(template.New("html").Parse(html2.D_district_area_show))
     monitor := template.Must(template.New("html").Parse(html2.D_district_area))

//    get area inf in d.s.

     d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

//   get value from interface data

    value, _ := d_area_view.([]type2.D_Area)

///
///     show area inf. on web
///

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

