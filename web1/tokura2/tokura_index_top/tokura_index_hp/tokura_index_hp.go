package tokura_index_hp

import (

	    "net/http"
	    "github.com/sawaq7/go14_ver1/client/tokura/html4"
	    "html/template"
                                                  )

///
/// show the menu of tokuras' applications
///



func Tokura_index_hp(w http.ResponseWriter, r *http.Request) {

   var cdmy string

///     set template

    monitor := template.Must(template.New("html").Parse(html4.Tokura_index_hp))

///           show tokura's hp menu on web

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

