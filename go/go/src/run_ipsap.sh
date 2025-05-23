#echo "*** 기존 실행되던 서버 종료 ***"
#kill -9 `ps -ef | grep 'ipsap_main' | awk '{print $2}'`

/root/go/swag init -d ./ipsap/ -o ./ipsap/docs -g ipsap.go
go run ./ipsap_main.go '/root/go/config/hks.toml'
