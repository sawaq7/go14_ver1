package html4

const Pipe_line1_show = `
   <!DOCTYPE html>
   <html>

     <head>
       <meta charset="UTF-8">
       <title>pipe_line1_show</title>
       <link rel="stylesheet" href="css/tokura.css" type="text/css">
     </head>

     <body>
       <section>
         <table border="2" cellpadding="8" align="center" bgcolor="#cd5c5c">

         <h2 align="center">List Of Deliverly Situation</h2>

         <tr> <th>water-name</th> <th>water-high</th> <th>R.F.</th> <th>change</th> <th>delete</th> <th>excute</th> </tr>
         {{range .}}
           <tr>

           <form method="GET" action="/pipe_line1_change" >

             <td>
                  <input type="text" name="water_name"  value="{{.Name|html}}"/>
                  <input type="hidden" name="water_id"  value="{{.No|html}}"/>
             </td>

             <td>
                  <input type="text" name="water_high" value="{{.High|html}}"/>
             </td>

             <td>
                  <input type="text" name="water_factor" value="{{.Roughness_factor|html}}"/>
             </td>

             <td>
                  <input type="submit" value="ON"  />
             </td>
           </form>

           <form method="GET" action="/pipe_line1_delete" >
             <td>
                  <input type="hidden" name="water_id"  value="{{.No|html}}"/>
                  <input type="submit" value="ON"  />
             </td>
           </form>

           <form method="GET" action="/pipe_line1_excute_single" >
             <td>
                  <input type="hidden" name="water_id"  value="{{.No|html}}"/>
                  <input type="submit" value="ON"  />
             </td>
           </form>

           </tr>
         {{end}}
         </table>
       </section>

       <section>
          <table border="2" cellpadding="8" align="center" bgcolor="#cd5c5c">

          <h2 align="center"> excute </h2>
          <tr>  <th>All</th> </tr>
          {{range .}}
          <tr>
          <form method="GET" action="/pipe_line1_excute_all" >

            <td>
               <input type="submit" value="all"  />
            </td>

          </form>
          </tr>
          {{end}}
          </table>
       </section>
     </body>
   </html>
`
