@echo off

set image_name="image-aetova-serv"
set container_name="aetova-serv"

REM Get the .env port
for /f "tokens=1,2 delims==" %%a in (.env) do (
    if "%%a"=="PORT" (
        set machine_port=%%b
        set container_port=%%b
    )
)

REM Stop the container
docker stop %container_name%
docker container rm %container_name%

REM Rebuild image
docker image rm %image_name%
docker build --build-arg PORT=%container_port% -t %image_name% .

REM Relaunch the container
REM If a container already run
docker stop %container_name%  
docker rm %container_name%
docker run -d --name %container_name% --env-file .env -p %machine_port%:%container_port% %image_name%
docker up %container_name%