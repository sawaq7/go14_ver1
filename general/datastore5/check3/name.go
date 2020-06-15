package check3

import (

	    "net/http"
	    "fmt"
	    "errors"

                     )

///
///    d.s. name check
///

func Name( w http.ResponseWriter, basic_name string  ) (err error ){

//     IN    w        : response-writer
//     IN  basic_Name :

    fmt.Fprintf( w, "check3.name start \n" )
    fmt.Fprintf( w, "check3.name basic_name %v\n" ,basic_name)

///
///     set error message
///

var (

	  Err1 = errors.New("can't find this datastore's file (check3.name)")

	                                                                        )
///
///       d.s. name check
///

    ok_flag := 0

    switch basic_name {

      case "Deliver" :

        ok_flag = 1


      break;

    }

    if ok_flag == 0 {

		return Err1
	}

    fmt.Fprintf( w, "check3.name ok_flag %v\n" ,ok_flag)
	fmt.Fprintf( w, "check3.name normal end \n" )

    return nil
}
