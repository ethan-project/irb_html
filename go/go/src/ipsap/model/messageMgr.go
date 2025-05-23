package model

import (
  "ipsap/common"
  "time"
  // "log"
  // "fmt"
)

type MessageMgr struct {
  Application_seq     uint
  Institution_seq     uint
  Msg_ID              int
  User_info           map[string]interface{}
  Mebership_info      map[string]string
}

func (ins *MessageMgr)SendMsgExpert() {
  if ins.Msg_ID == 0  {
    return
  }

  application := Application{
    Application_seq : ins.Application_seq,
  }

  judge_alarm_type := 0
  switch ins.Msg_ID {
    case DEF_MSG_EXPER_JUDGE_START: judge_alarm_type = DEF_JUDGE_ALARM_EXPER_JUDGE_START
    case DEF_MSG_BEFORE_24_HOURS: judge_alarm_type = DEF_JUDGE_ALARM_BEFORE_24_HOURS
    case DEF_MSG_BEFORE_4_HOURS: judge_alarm_type = DEF_JUDGE_ALARM_BEFORE_4_HOURS
    case DEF_MSG_JUDGE_DELAYED:  judge_alarm_type = DEF_JUDGE_ALARM_JUDGE_DELAYED
  }

  if !application.CheckSendMsg(judge_alarm_type) {
    return
  }

  app_member := AppMember{ Application : &application }
  if !app_member.Load("expert_member", nil)  {   //  전문위원 아이템 이름
    return;
  }

  nowUnixTime := time.Now().Unix()
  expert_deadline := application.GetUnixtimeForExpertDeadline()
  if expert_deadline == 0  {  return  }

  // log.Println("======== Msg_ID : ", ins.Msg_ID)
  // log.Println(app_member.Datas)

  for _, user := range app_member.Datas {
    //  MsgID 별로 bit 단위로 flag가 설정 되어 있음.
    if ((common.ToInt(user["chk_flag"]) & (1 << ins.Msg_ID)) > 0) {
      continue //  이미 메세지를 전송한 경우는 Skip
    }
    ins.User_info = user["info"].(map[string]interface{})
    if (ins.SendMessage())  {
      app_member.SetChkFlag("expert_member", ins.User_info["user_seq"], judge_alarm_type)
      if (ins.Msg_ID == DEF_MSG_EXPER_JUDGE_START)  {
        before_24h := expert_deadline - 24 * 60 * 60
        if nowUnixTime >= before_24h  {   //  종료까지 24시간이 안남았으면
          app_member.SetChkFlag("expert_member", ins.User_info["user_seq"], DEF_JUDGE_ALARM_BEFORE_24_HOURS)
        }
        before_4h := expert_deadline - 4 * 60 * 60
        if nowUnixTime >= before_4h  {   //  종료까지 4시간이 안남았으면
          app_member.SetChkFlag("expert_member", ins.User_info["user_seq"], DEF_JUDGE_ALARM_BEFORE_4_HOURS)
        }
      }
    }
  }
  return
}

func (msg *MessageMgr)SendMsgNormal() {
  if msg.Msg_ID == 0  {
    return
  }

  application := Application{
    Application_seq : msg.Application_seq,
  }

  judge_alarm_type := 0
  switch msg.Msg_ID {
    case DEF_MSG_NORMAL_JUDGE_START: judge_alarm_type = DEF_JUDGE_ALARM_NORMAL_JUDGE_START
    case DEF_MSG_BEFORE_24_HOURS: judge_alarm_type = DEF_JUDGE_ALARM_BEFORE_24_HOURS
    case DEF_MSG_BEFORE_4_HOURS: judge_alarm_type = DEF_JUDGE_ALARM_BEFORE_4_HOURS
    case DEF_MSG_JUDGE_DELAYED:  judge_alarm_type = DEF_JUDGE_ALARM_JUDGE_DELAYED
  }

  if !application.CheckSendMsg(judge_alarm_type) {
    return
  }

  app_normal_member := AppMember{ Application : &application }

  // 21-09-06 : 일반심사 완료한 인원 메세지 보내지 않기
  if !app_normal_member.GetToSendNormalMembers()  {
    return;
  }

  // app_member_in := AppMember{ Application : &application }
  // app_member_ex := AppMember{ Application : &application }
  // if !app_member_in.Load("committee_in_member", nil)  {   //  내부심사위원 아이템 이름
  //   return;
  // }
  //
  // if !app_member_ex.Load("committee_ex_member", nil)  {   //  외부심사위원 아이템 이름
  //   return;
  // }

  nowUnixTime := time.Now().Unix()
  normal_deadline := application.GetUnixtimeForNormalDeadline()
  if normal_deadline == 0  {  return  }
  for _, user := range app_normal_member.Datas {
    //  MsgID 별로 bit 단위로 flag가 설정 되어 있음.
    if ((common.ToInt(user["chk_flag"]) & (1 << msg.Msg_ID)) > 0) {
      continue //  이미 메세지를 전송한 경우는 Skip
    }

    msg.User_info = user["info"].(map[string]interface{})
    if (msg.SendMessage()) {
      app_normal_member.SetChkFlag("committee_in_member", msg.User_info["user_seq"], judge_alarm_type)
      if (msg.Msg_ID == DEF_MSG_NORMAL_JUDGE_START)  {
        before_24h := normal_deadline - 24 * 60 * 60
        if nowUnixTime >= before_24h  {   //  종료까지 24시간이 안남았으면
          app_normal_member.SetChkFlag("committee_in_member", msg.User_info["user_seq"], DEF_JUDGE_ALARM_BEFORE_24_HOURS)
        }
        before_4h := normal_deadline - 4 * 60 * 60
        if nowUnixTime >= before_4h  {   //  종료까지 4시간이 안남았으면
          app_normal_member.SetChkFlag("committee_in_member", msg.User_info["user_seq"], DEF_JUDGE_ALARM_BEFORE_4_HOURS)
        }
      }
    }
  }

//   for _, user := range app_member_in.Datas {
//     //  MsgID 별로 bit 단위로 flag가 설정 되어 있음.
//     if ((common.ToInt(user["chk_flag"]) & (1 << msg.Msg_ID)) > 0) {
//       continue //  이미 메세지를 전송한 경우는 Skip
//     }
//     msg.User_info = user["info"].(map[string]interface{})
//     if (msg.SendMessage()) {
//       app_member_in.SetChkFlag("committee_in_member", msg.User_info["user_seq"], judge_alarm_type)
//       if (msg.Msg_ID == DEF_MSG_NORMAL_JUDGE_START)  {
//         before_24h := normal_deadline - 24 * 60 * 60
//         if nowUnixTime >= before_24h  {   //  종료까지 24시간이 안남았으면
//           app_member_in.SetChkFlag("committee_in_member", msg.User_info["user_seq"], DEF_JUDGE_ALARM_BEFORE_24_HOURS)
//         }
//         before_4h := normal_deadline - 4 * 60 * 60
//         if nowUnixTime >= before_4h  {   //  종료까지 4시간이 안남았으면
//           app_member_in.SetChkFlag("committee_in_member", msg.User_info["user_seq"], DEF_JUDGE_ALARM_BEFORE_4_HOURS)
//         }
//       }
//     }
//   }
//
//   for _, user := range app_member_ex.Datas {
//     //  MsgID 별로 bit 단위로 flag가 설정 되어 있음.
//     if ((common.ToInt(user["chk_flag"]) & (1 << msg.Msg_ID)) > 0) {
//       continue //  이미 메세지를 전송한 경우는 Skip
//     }
//     msg.User_info = user["info"].(map[string]interface{})
//     if (msg.SendMessage()) {
//       app_member_ex.SetChkFlag("committee_ex_member", msg.User_info["user_seq"], judge_alarm_type)
//       if (msg.Msg_ID == DEF_MSG_NORMAL_JUDGE_START)  {
//         before_24h := normal_deadline - 24 * 60 * 60
//         if nowUnixTime >= before_24h  {   //  종료까지 24시간이 안남았으면
//           app_member_ex.SetChkFlag("committee_ex_member", msg.User_info["user_seq"], DEF_JUDGE_ALARM_BEFORE_24_HOURS)
//         }
//         before_4h := normal_deadline - 4 * 60 * 60
//         if nowUnixTime >= before_4h  {   //  종료까지 4시간이 안남았으면
//           app_member_ex.SetChkFlag("committee_ex_member", msg.User_info["user_seq"], DEF_JUDGE_ALARM_BEFORE_4_HOURS)
//         }
//       }
//     }
//   }
}
