package sample_index

import (

	    "net/http"
	    "github.com/sawaq7/go14_ver1/temp/html1000"
	    "html/template"
                                                  )

///
/// show the menu of samples' applications
///

func Sample_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// ãE³ãã¬ã¼ããEãããã¼ãGET

    monitor := template.Must(template.New("html").Parse(html1000.Sample_index))

// ã¢ãã¿ã¼ã«è¡¨ç¤º

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

