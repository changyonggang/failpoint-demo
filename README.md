# failpoint-demo

参考文档: https://www.luozhiyun.com/archives/595

bin/failpoint-ctl enable .
GO_FAILPOINTS='main/testValue=10*return("abc")' go run main.go binding__failpoint_binding__.go
bin/failpoint-ctl disable . 