@REM 项目根目录  运行例子1： .\test\sh\api.bat sys all

@REM 路径
set dir=%1
@REM 文件名
set apiName=%2

goctl api go -api service/%dir%/api/desc/%apiName%.api -dir service/%dir%/api/ -home ./template/185/ -style=go_zero

@REM if "%api%" == "api" (
@REM    goctl api go -api application\%dir%\api\%dir%.api -dir application\%dir%\api\
@REM ) else (
@REM    goctl api go -api application\%dir%\%dir%.api -dir application\%dir%\
@REM )



@REM goctl api go -api application\%dir%\%api% -dir %dir% -home ./dev/goctl
@REM goctl api go --dir=./ --api applet.api
@REM goctl api go -api service/sys/api/desc/all.api -dir service/sys/api/ -home ./template/185/ -style=go_zero