$jsonObj = @'
{"value":"Say smth"}
'@
Invoke-RestMethod -Method Post -Uri "http://localhost:8081/v1/example/echo" -body $jsonObj -ContentType "application/json"
