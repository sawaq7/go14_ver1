    ///////////////////////////////////////////////////////
   ///                                                 ///
  ///     template　for ( pipe_line1_show )            ///
 ///                                                  ///
////////////////////////////////////////////////////////

package html4

const Pipe_line1_show2 = `
   <!DOCTYPE html>
   <html>

     <head>
       <meta charset="UTF-8">
       <title> constructor "tokura"</title>
       <link rel="stylesheet" href="css/css/tokura.css" type="text/css">
     </head>

     <body>

       <section>
          <table border="2" cellpadding="8" align="center" bgcolor="#cd5c5c">

          <h2 align="center"> excute all </h2>
          <tr> <th>water-name</th> <th>excute</th> </tr>

          <tr>

          {{.Name|html}}
          {{.No|html}}
              <
          </tr>

          </table>
       </section>
     </body>
   </html>
`
