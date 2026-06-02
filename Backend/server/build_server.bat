@echo off

set image_name="image-aetova-serv"
set container_name="aetova-serv"
set machine_port="51418"
set container_port="51418"

REM Stop the container
docker stop %container_name%
docker container rm %container_name%

REM Rebuild image
docker image rm %image_name%

docker build --build-arg PORT=%container_port% -t %image_name% .
if %errorlevel% neq 0 (
	echo Error Build
	pause
	exit /b 1
)

REM Relaunch the container
docker run -d --name %container_name% -p %machine_port%:%container_port% %image_name%
if %errorlevel% neq 0 (
	echo Eroor run
	pause
	exit /b 1
)