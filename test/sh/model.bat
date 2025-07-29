@REM 项目根目录  运行例子： .\test\sh\model.bat sys sys_users
@REM  第一个参数目录名  第二个参数是表名
set name=%1
set tableName=%2
set mysql=root:123456@tcp(127.0.0.1:3306)/go-zero-fast


goctl model mysql datasource --dir service/%name%/model --table %tableName% --cache true --url=%mysql% -style=go_zero -cache=true
