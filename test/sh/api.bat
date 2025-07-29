@REM 项目根目录  运行例子1： .\pkg\test\api.bat applet
@REM 项目根目录  运行例子2： .\pkg\test\api.bat applet api

set dir=%1
set api=%2

if "%api%" == "api" (
   goctl api go -api application\%dir%\api\%dir%.api -dir application\%dir%\api\
) else (
   goctl api go -api application\%dir%\%dir%.api -dir application\%dir%\
)



@REM goctl api go -api application\%dir%\%api% -dir %dir% -home ./dev/goctl
@REM goctl api go --dir=./ --api applet.api
@REM goctl api go -api service/sys/desc/all.api -dir service/sys/ -home ./template/185/ -style=go_zero