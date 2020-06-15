package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go14_ver1/client/sgh/html2"
	    "github.com/sawaq7/go14_ver1/client/sgh/type2"
	    "github.com/sawaq7/go14_ver1/general/type5"
//	    "github.com/sawaq7/go14_ver1/client/sgh/datastore2/trans"

        "github.com/sawaq7/go14_ver1/client/sgh/datastore2"
//	    "time"
                                                )

///
///     show deliver inf. in d.s.
///

func Deliver_showall2( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　　: response-writer
//     IN    r      　　: request-parameter

//   fmt.Fprintf( w, "process.deliver_show_all2 start \n" )

///
///     set template
///

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showall1))

///
///    get deliver inf in d.s.
///


     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          //  district no
     general_work[1].Int64_Work = course_no  //　cource no

     deliver_view := datastore2.Datastore_sgh( "Deliver"  ,"trans"  ,general_work , w , r  )

   //  get value from interface data
     value, _ := deliver_view.([]type2.Deliver)

///
///       show deliver inf. on web
///

    err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

