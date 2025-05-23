
# echo *** 기존 실행되던 서버 종료 ***

#:: 기존 프로세스 종료
# tasklist | findstr "win_ipsap_main" > nul && taskkill /F /IM win_ipsap_main.exe

# :: Go 프로그램 실행
#cd C:\root\go\  :: Go 작업 디렉터리로 이동
#go run .\win_ipsap_main.go "C:\ipsap\go\go\config\win_ipsap.toml"
./win_ipsap_main.exe ../config/win_ipsap.toml
