<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <title>Observations</title>
    </head>
    <body>
        <h1>Observations</h1>

        <div>
            {{range .Observations}}
            <p>{{$time := .Timestamp.Format "2006-01-02 15:04:05"}}{{$time}} {{.Point.Name}}: {{.Celsius}} Celsius</p>
            {{end}}

            {{if .NotFirstPage}}
            <a href="/app/observation?pageNumber={{.LastPage}}">Prev</a>
            {{end}}

            {{if .NotLastPage}}
            <a href="/app/observation?pageNumber={{.NextPage}}">Next</a>
            {{end}}
        </div>

        <h2>Add new observation</h2>            
        <div>
            <form action="/app/observation" method="POST" >
                <p>
                    <select id="observationPointId" name="observationPointId">
                    {{range $i, $op := .ObservationPoints}}
                    <option value="{{$i}}">{{$op.Name}}</option>
                    {{end}}
                    <label for="observationPointId">Location</label> 
                </p>  
                <p>  
                    <input id="Celsius" name="Celsius" type="text" value=""/>  
                    <label for="Celsius">Celsius</label> 
                </p>
                <p class="submit">  
                    <input type="submit" value="Add" />  
                </p>
            </form>
        </div>

        
        <div>
            <p><a href="/app/observation">Observations</a></p>
        </div>

        <div>
            <p><a href="/app/observationpoint">Observation points</a></p>
        </div>
    </body>
</html>
