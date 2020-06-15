package payment_delete

import (

	"net/http"
	"strconv"
//	"fmt"

	"github.com/sawaq7/go14_ver1/client/reserve/process4"

    "github.com/sawaq7/go14_ver1/client/reserve/datastore6/check5"

    "cloud.google.com/go/datastore"
	"context"
	"os"

                                            )

///                         　　　　
///     delete payment inf. in d.s.
///

func Payment_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "payment_delete start \n" )

    id := r.FormValue("id")
//    fmt.Fprintf( w, "payment_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "payment_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "payment_delete : delid %v\n", delid )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    key := datastore.IDKey("Guest_Payment", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    general_work := check5.Guest_temp (w , r )

    guest_no   := general_work[0].Int64_Work

///
///      show payment inf. on web
///

    process4.Payment_register( w , r ,guest_no )

}
