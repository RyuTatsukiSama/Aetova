set SERVER_URL=http://localhost:51418
set API_KEY=c7e642cc-9928-4248-bd3f-c9588490bb60
set PORT=51419

go build

robocopy "./" "../../Frontend/clientGo/" client.exe
if %errorlevel% gtr 7 (
    pause
    exit
)

pause