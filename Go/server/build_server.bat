@echo off

set image_name="image-aetova-serv"
set container_name="aetova-serv"
set app_port="8080" # Port that the golang code listen
set container_port="8080" # Port to access the code

REM Rebuild image
docker image rm %image_name%
docker build -t %image_name% .

REM Relaunch the container
REM If a container already run
docker stop %container_name%  
docker rm %container_name%
docker run -d --name %container_name% --env-file .env -p 8080:8080 %image_name%

pause