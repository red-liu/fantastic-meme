go1.11后，go使用modules管理依赖，项目依赖构建时不需要再依赖GOPATH环境变量。
    1.激活modules
        .首先要升级go到1.11版本
        .设置环境变量GO111MODULE=on或者auto
    
    2.在任意目录下构建项目，项目名建议使用域名命名
        例如，
            #mkdir -p my.com

    3.创建模块
        例子中三个模块：app，firstmod，secondmod
            #cd my.com
            #mkdir app firstmod secondmod
            #cd app;go mod init my.com/app
            #cd ../firstmod;go mod init my.com/firstmod
            #cd ../secondmod;go mod init my.com/secondmod

    4.为各个module增加代码
        详细参照app，firstmod，secondmod下的代码

    5.建立依赖关系
        .在firstmod，secondmod下运行go build，自动更新go.mod内容
        .如果你希望下载特定版本的依赖第三方包，使用命令go get
         
        例如：
            #cd secondmod
            #go list -m -versions rsc.io/quote/v3
                rsc.io/quote/v3 v3.0.0 v3.1.0
        下载特定版本,并且自动变更go.mod文件:
            #go get rsc.io/quote/v3@v3.1.0

        下载的版本在$GOPATH/pkg/mod/cache/download,如果完全不使用$GOPATH环境变量，可以使用go mod vendor 命令。
            #cd secondmod
            #go mod vendor

        .如果依赖项目下的模块，使用模块名
            import "my.com/firstmod"
            import "my.com/secondmod"

            修改go.mod:
            require my.com/firstmod v0.0.0
            require my.com/secondmod v0.0.0
            replace my.com/firstmod v0.0.0 => ../firstmod/v2
            replace my.com/secondmod v0.0.0 => ../secondmod

            甚至你能依赖模块的某个版本例如
                replace my.com/firstmod v0.0.0 => ../firstmod/v2
                是依赖firstmod下的v2版本
        
        .显示依赖关系图：在app下运行命令：
            #go mod graph
                my.com/app my.com/firstmod@v0.0.0
                my.com/app my.com/secondmod@v0.0.0
                my.com/secondmod@v0.0.0 rsc.io/quote/v3@v3.1.0
                rsc.io/quote/v3@v3.1.0 rsc.io/sampler@v1.3.0
                rsc.io/sampler@v1.3.0 golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c
    
        6.编译项目，在app下go build即可。


        

