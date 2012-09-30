<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <title>Order information</title>
    </head>
    <body>
        <h1>Hello {{index .Values "Name"}}  {{range .Flashes}}{{.}}{{end}}</h1>

        <p>Your order will be posted to {{index .Values "Address"}}</p>

        <p>Ordered items:<br/>

        <div>
            <ol>
                {{range index .Values "Items"}}
                <li>{{.}}</li>
                {{end}}
            </ol>
        </div> 
   </p>
</body>
</html>
