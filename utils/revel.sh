#!/bin/bash
#Program:
#       golang revel 自動 構建 腳本
#History:
#       2018-03-29 king first release
#Email:
#       zuiwuchang@gmail.com

# 定義的 各種 輔助 函數
MkDir(){
	mkdir -p "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}
NewFile(){
	echo "$2" > "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}
WriteFile(){
	echo "$2" >> "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}


CreateGoVersion(){
	# 返回 git 信息 時間
	tag=`git describe`
	if [ "$tag" == '' ];then
		tag="[unknown tag] "
	else
		tag="$tag "
	fi

	commit=`git rev-parse HEAD`
	if [ "$commit" == '' ];then
		commit="[unknow commit]"
	fi
	
	date=`date +'%Y-%m-%d %H:%M:%S'`

	# 打印 信息
	echo ${tag} $commit
	echo $date


	# 自動 創建 go 代碼
	NewFile $1	"package $2"
	WriteFile $1	''
	WriteFile $1	'// Version auto create build version info'
	WriteFile $1	"const Version = \`$tag $commit"
	WriteFile $1	"build at $date\`"
}

createRevelVersion(){
    cd $1
    if [ "$?" != 0 ] ;then
		exit 1
	fi
    CreateGoVersion version.go app
}
CreateRevelVersion(){
    base=`pwd`
    IFS=':' read -r -a array <<< "$GOPATH"
    for element in "${array[@]}"
    do
        if test -d "$element/src/$1/app";then
            createRevelVersion "$element/src/$1/app"
            cd $base
            break
        fi
    done
}

if [ "$1" == "run" ];then
    CreateRevelVersion $2
    revel $@
elif [ "$1" == "package" ];then
    CreateRevelVersion $2
    revel $@
else
    revel $@
fi