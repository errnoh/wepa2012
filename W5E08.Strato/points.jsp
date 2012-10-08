<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <title>Observation points</title>
    </head>
    <body>

        <h1>Observation points</h1>

        <div>
            {{range .ObservationPoints}}
            <p>{{.Name}} (lat: {{.Latitude}}, lon: {{.Longitude}}) (<a href="http://maps.google.com/?q={{.Latitude}},{{.Longitude}}" target="_blank">map</a>)</p>
            {{end}}
        </div>

        <h2>Add new</h2>            
        <div>
            <form id="observationPoint" action="/app/observationpoint" method="POST">
                <p>  
                    <input id="Name" name="Name" type="text" value=""/>  
                    <label for="Name">Name</label>  
                </p>  
                <p>  
                    <input id="Latitude" name="Latitude" type="text" value=""/>  
                    <label for="Latitude">Latitude (Attn! Use dot (.) and not comma (,) as separator)</label>
                </p>  
                <p>  
                    <input id="Longitude" name="Longitude" type="text" value=""/>
                    <label for="Longitude">Longitude (Attn! Use dot (.) and not comma (,) as separator )</label>
                </p>
                <p>  
                    <input type="submit" value="Add" />  
                </p>
            </form>
        </div>


        <div>
            <p><a href="observation">Observations</a></p>
        </div>

        <div>
            <p><a href="observationpoint">Observation points</a></p>
        </div>
    </body>
</html>
