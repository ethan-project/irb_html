package model

import (
  "ipsap/common"
  "time"
  "log"
  "fmt"
)

type BackgroundMonitor struct {
}

func (bkgm *BackgroundMonitor)Start() {
  go bkgm.realStart()
}

func (bkgm *BackgroundMonitor)realStart() {

  beforeHourMap := map[uint]interface{} {}   //  app_seq -> [24, 4] : 신청서당 처리한 시간대

  for { //  loop forever
    // 1. '심사중'인 신청서 리스트 추출
    app_list := bkgm.loadAppListInJudgeIng()

    for idx, app := range app_list {
      application := Application{
        Application_seq : common.ToUint(app["application_seq"]),
      }
      beforeMap, exists := beforeHourMap[application.Application_seq]
      if !exists {
        beforeHourMap[application.Application_seq] = make(map[string]bool)
        beforeMap, exists = beforeHourMap[application.Application_seq]
      }

      nowUnixTime := time.Now().Unix()
      expert_deadline := application.GetUnixtimeForExpertDeadline()
      if expert_deadline == 0  {  continue  }
      normal_deadline := application.GetUnixtimeForNormalDeadline()
      if normal_deadline == 0  {  continue  }

//      log.Println("nowUnixTime:", nowUnixTime, "app_seq:", application.Application_seq, "e_deadline:", expert_deadline, normal_deadline)
      application_step := common.ToInt(app["application_step"])
      application_result := common.ToInt(app["application_result"])

      switch(application_step) {
      case DEF_APP_STEP_PRO  :      //  전문심사 단계
        if nowUnixTime >= expert_deadline {
          if (application_result == DEF_APP_RESULT_JUDGE_ING ||   // 심사중, 심사중 2
              application_result == DEF_APP_RESULT_JUDGE_ING_2) { //  지연 상태로 변경
            if application.UpdateStepAndResult(DEF_APP_STEP_PRO, DEF_APP_RESULT_JUDGE_DELAY) {
              log.Println("전문심사 (심사중 -> 지연중) : 성공", app);
              msgMgr := MessageMgr	{
                Application_seq : application.Application_seq,
                Msg_ID : DEF_MSG_JUDGE_DELAYED,
              }
              msgMgr.SendMsgExpert();
            } else {
              log.Println("전문심사 (심사중 -> 지연중) : Error!!!!!", app);
            }
          }
        } else {
          before_24h := expert_deadline - 24 * 60 * 60
          if nowUnixTime >= before_24h  {   //  종료까지 24시간이 안남았으면
            _, has := beforeMap.(map[string]bool)["e_24"]
            if !has {
              beforeMap.(map[string]bool)["e_24"] = true
              msgMgr := MessageMgr	{
                Application_seq : application.Application_seq,
                Msg_ID : DEF_MSG_BEFORE_24_HOURS,
              }
              msgMgr.SendMsgExpert();
            }
          }
          before_4h := expert_deadline - 4 * 60 * 60
          if nowUnixTime >= before_4h  {   //  종료까지 4시간이 안남았으면
            _, has := beforeMap.(map[string]bool)["e_04"]
            if !has {
              beforeMap.(map[string]bool)["e_04"] = true
              msgMgr := MessageMgr	{
                Application_seq : application.Application_seq,
                Msg_ID : DEF_MSG_BEFORE_4_HOURS,
              }
              msgMgr.SendMsgExpert();
            }
          }
        }
      case DEF_APP_STEP_NORMAL  :   //  일반심사 단계
        if nowUnixTime >= normal_deadline {
          if (application_result == DEF_APP_RESULT_JUDGE_ING ||   // 심사중, 심사중 2
              application_result == DEF_APP_RESULT_JUDGE_ING_2) { //  지연 상태로 변경
            if application.UpdateStepAndResult(DEF_APP_STEP_NORMAL, DEF_APP_RESULT_JUDGE_DELAY) {
              log.Println("일반심사 (심사중 -> 지연중) : 성공", app);
              msgMgr := MessageMgr	{
                Application_seq : application.Application_seq,
                Msg_ID : DEF_MSG_JUDGE_DELAYED,
              }
              msgMgr.SendMsgNormal();
            } else {
              log.Println("일반심사 (심사중 -> 지연중) : Error!!!!!", app);
            }
          }
        } else {
          before_24h := normal_deadline - 24 * 60 * 60
          if nowUnixTime >= before_24h  {   //  종료까지 24시간이 안남았으면
            _, has := beforeMap.(map[string]bool)["n_24"]
            if !has {
              beforeMap.(map[string]bool)["n_24"] = true
              msgMgr := MessageMgr	{
                Application_seq : application.Application_seq,
                Msg_ID : DEF_MSG_BEFORE_24_HOURS,
              }
              msgMgr.SendMsgNormal();
            }
          }
          before_4h := normal_deadline - 4 * 60 * 60
          if nowUnixTime >= before_4h  {   //  종료까지 4시간이 안남았으면
            _, has := beforeMap.(map[string]bool)["n_04"]
            if !has {
              beforeMap.(map[string]bool)["n_04"] = true
              msgMgr := MessageMgr	{
                Application_seq : application.Application_seq,
                Msg_ID : DEF_MSG_BEFORE_4_HOURS,
              }
              msgMgr.SendMsgNormal();
            }
          }
        }
      default:
        log.Println("Error!!!!!", idx, app);
      }
    }

    time.Sleep(time.Second * 60)   //  1분단위로 작업
  }
}


func (bkgm *BackgroundMonitor)loadAppListInJudgeIng() (ret []map[string]interface{}) {
  application := Application{ }

  //  심사중1, 심사중2, 지연중
  moreCondition := fmt.Sprintf(`
    AND app.application_result in (%v, %v, %v)
  `, DEF_APP_RESULT_JUDGE_ING, DEF_APP_RESULT_JUDGE_ING_2, DEF_APP_RESULT_JUDGE_DELAY);

  return application.LoadList("", "", moreCondition, "").([]map[string]interface{})
}
