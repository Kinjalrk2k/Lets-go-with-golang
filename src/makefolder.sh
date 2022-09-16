echo $1

folder=$1
len=${#folder}
modname=${folder:2:$len}

mkdir $1
cd $1
go mod init $modname
touch main.go
code main.go