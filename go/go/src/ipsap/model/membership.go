package model

import (
	"github.com/nleeper/goment"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"ipsap/common"
	"strings"
	"math"
	"fmt"
	"log"
)

type FreeMembership struct {
	Free_seq				uint		`json:"-"`
	Institution_seq	uint		`json:"institution_seq" example:"1"`
	Usage_limit			uint		`json:"usage_limit" example:"1"`
	Free_period 		string	`json:"free_period" example:"1"`
	Reason 					string	`json:"reason" example:"가입 기념 무료 지급"`
	StartIdx 				string	`json:"-"`
	RowCnt 	 				string	`json:"-"`
	UseStatus 			string	`json:"-"`
	SearchWord 			string	`json:"-"`
	PurchasedDate   string  `json:"-"`
}

type Membership struct {
	Institution_seq uint
}

type PaymentCancelReq struct {
	TID								string
	MID								string
	Moid							string
	CancelAmt					string
	CancelMsg					string
	EdiDate						string
	SignData					string
	MallReserved			string
	PartialCancelCode	string
}

type PaymentCancelResp struct {
	ResultCode 				string	`json:"ResultCode"`
  ResultMsg  				string	`json:"ResultMsg"`
  ErrorCD		 				string	`json:"ErrorCD"`
  ErrorMsg	 				string	`json:"ErrorMsg"`
  CancelAmt					string	`json:"CancelAmt"`
  MID								string	`json:"MID"`
  Moid							string	`json:"Moid"`
  PayMethod					string	`json:"PayMethod"`
  TID								string	`json:"TID"`
  CancelDate				string	`json:"CancelDate"`
  CancelTime				string	`json:"CancelTime"`
  CancelNum 				string	`json:"CancelNum"`
  RemainAmt					string	`json:"RemainAmt"`
  Signature					string	`json:"Signature"`
  MallReserved			string	`json:"MallReserved"`
}

type MembershipCancel struct {
	Institution_seq			 uint
  MembershipAmt				 uint
	PlanAmt							 uint
	CancelUserSeq				 uint
	OrderSeqArr					 []uint
	CancelAmtArr				 []uint
	PartialCancelCodeArr []string
	Pname 							 string
	MebershipJoinDate 	 string
	AdminUserSeqArr		 	 string
}

func (free *FreeMembership) InsertFreeMembership() (succ bool) {
	succ = false
	free.Apply(false)

	sql := fmt.Sprintf(`
					INSERT INTO t_institution_free_period( institution_seq, usage_limit, free_period,
																								 reason, purchased_date, reg_date)
					VALUES(%v,%v,%v,
								 '%v',%v,UNIX_TIMESTAMP())`,
							 	 free.Institution_seq, free.Usage_limit, free.Free_period,
								 free.Reason, free.PurchasedDate)

  _, err := common.DBconn().Exec(sql)
  if nil != err {
    log.Println(err)
    return
  }

  succ = true
	return
}

func (free *FreeMembership) Apply(batchFlag bool) {
	instt := Institution{
		Institution_seq : free.Institution_seq,
	}

	instt.Load()

	now, _ := goment.New()
	today := common.ToUint(now.Format("YYYYMMDD"))
	t1, _ := goment.New()
	t1.EndOf("month")
	t1.Add(-7, "days")
	date1 := common.ToUint(t1.Format("YYYYMMDD"))
	expirationDate, _ := goment.New(common.ToStr(instt.Data["expiration_date"]))
	date2 := common.ToUint(expirationDate.Format("YYYYMMDD"))
	expirationDate.Add(-7, "days")
	date3 := common.ToUint(expirationDate.Format("YYYYMMDD"))

	free.PurchasedDate = "0"
	// 오늘이 이번달말일 -7일 이후일때
  if today >= date1 {
		if date1 <= date3 { // 이번달 -7일이 기관 만료일 -7일 보다 작거나 같을때
			go free.InstitutionApply(false)
			free.PurchasedDate = "UNIX_TIMESTAMP()"
		}
	} else {
		// 기관 만료일이 지났을때
		if today >= date2 {
			go free.InstitutionApply(true)
			free.PurchasedDate = "UNIX_TIMESTAMP()"
		}
	}

	// 무료이용권 지급일 UPDATE
	if batchFlag {
		if free.Free_seq == 0 {
			return
		}
		sql := `UPDATE t_institution_free_period
							 SET purchased_date = UNIX_TIMESTAMP()
						 WHERE free_seq = ?`
		_, err := common.DBconn().Exec(sql, free.Free_seq)
		if nil != err {
			log.Println(err)
		}
	}

	return
}

func (free *FreeMembership) InstitutionApply(pastExpirationDate bool) {
	// 기관 만료일이 지났을때 무료 시작일은 오늘  무료 시작일 이번달 마지막일 + 1
	t1, _ := goment.New()
	freePeriod := common.ToInt(free.Free_period)
	if pastExpirationDate {
		freePeriod--;
	} else {
		t1.EndOf("month")
		t1.Add(1, "days")
	}

	freeStartDate := t1.Format("YYYYMMDD")
	freeStartDateStr := common.GetDateStr(t1)
	t1.Add(freePeriod, "month")
	t1.EndOf("month")
	freeEndDate := t1.Format("YYYYMMDD")
	freeEndDateStr := common.GetDateStr(t1)

	sql := `UPDATE t_institution
						 SET service_status = 1,
						 		 free_start_date = ?,
								 free_end_date = ?,
						 		 expiration_date = ?,
								 usage_limit = ?
					 WHERE institution_seq = ?`
	 _, err := common.DBconn().Exec(sql, freeStartDate, freeEndDate,
		 															freeEndDate, free.Usage_limit, free.Institution_seq)
	if nil != err {
		log.Println(err)
	}

	sql2 := fmt.Sprintf(`SELECT (SELECT group_concat(user.user_seq)
												 					FROM t_user user
															 	 WHERE user.institution_seq = instt.institution_seq
																 	 AND user.user_type LIKE '%%%v%%'
																	 AND user.user_status = 2) AS user_arr,
										 					instt.membership_payment_date
												 FROM t_institution instt
						 					  WHERE instt.institution_seq = %v`,DEF_USER_TYPE_ADMIN_SECRETARY, free.Institution_seq)
	row := common.DB_fetch_one(sql2, nil)
	membershipInfo := make(map[string]string)
	t2, _ := goment.Unix(common.ToInt64(row["membership_payment_date"]))
	membershipJoinDate, _ := goment.New(t2)
	membershipInfo[DEF_MEMBERSHIP_JOIN_DATE] = common.GetDateStr(membershipJoinDate)
	membershipInfo[DEF_MEMBERSHIP_EXPIRATION_DATE] = common.GetMembershipExpirationDate(membershipJoinDate)
	membershipInfo[DEF_FREE_START_DATE] = freeStartDateStr
	membershipInfo[DEF_FREE_END_DATE] = freeEndDateStr
	userArr := strings.Split(common.ToStr(row["user_arr"]), ",")
	msg := MessageMgr{}
	for _, userSeq := range userArr {
		user := User{
			User_seq : common.ToUint(userSeq),
		}
		if !user.Load(){
			continue;
		}
		msg.Msg_ID					= DEF_MSG_MEMBERSHIP_FREE
		msg.User_info				= user.Data
		msg.Mebership_info	= membershipInfo
		msg.Institution_seq	= free.Institution_seq
		msg.SendMessage()
	}

	return
}

func (free *FreeMembership) LoadList() (ret gin.H) {

	moreCondition := ""

	if "" != free.UseStatus {
		if "1" == free.UseStatus {
			moreCondition += fmt.Sprintf(` AND free.purchased_date > 0`)
		} else {
			moreCondition += fmt.Sprintf(` AND free.purchased_date = 0`)
		}
	}

	if "" != free.SearchWord {
		moreCondition += fmt.Sprintf(`  AND (instt.name_ko LIKE '%%%v%%'
																		  	OR free.reason LIKE '%%%v%%')`, free.SearchWord, free.SearchWord)
	}

	sql := fmt.Sprintf(`
		SELECT free.free_seq, free.institution_seq, free.purchased_date,
		 			 free.usage_limit, free.free_period, free.reason, instt.name_ko,
					 free.reg_date
			FROM t_institution_free_period free, t_institution instt
		 WHERE free.institution_seq = instt.institution_seq
		 	 AND deleted_flag = 0
			 	%v #moreCondition
		 ORDER BY purchased_date ASC
		 LIMIT %v, %v
		    `, moreCondition, free.StartIdx, free.RowCnt)
	list := common.DB_fetch_all(sql, nil)

	sql2 := fmt.Sprintf(`
		SELECT COUNT(free.free_seq) as totalCnt
			FROM t_institution_free_period free, t_institution instt
		 WHERE free.institution_seq = instt.institution_seq
		 	 AND deleted_flag = 0
			 	%v #moreCondition`, moreCondition)
	row := common.DB_fetch_one(sql2, nil)
	totalCnt := common.ToUint(row["totalCnt"])

	pageInfo := gin.H {
		"rowCnt"  : len(list),
		"startIdx": free.StartIdx,
		"totalCnt": totalCnt,
	}

	ret = gin.H {
		"pageInfo": pageInfo,
		"list" : list,
	}

	return
}

func (free *FreeMembership) DeleteFreeMembership() (succ bool) {
	succ = false
	sql := `UPDATE t_institution_free_period
						 SET deleted_flag = 1
					 WHERE free_seq = ?`
	_, err := common.DBconn().Exec(sql, free.Free_seq)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (membership *Membership) GetInstitutionMembershipInfo() (result gin.H) {
	sql := `SELECT instt.membership_fee_status, instt.product_seq,
								 instt.payment_setting, instt.bid,
								 (SELECT plan.name
								 	  FROM t_products prod, t_membership_plan plan
								   WHERE prod.plan_seq = plan.plan_seq
								 		 AND prod.product_seq = instt.product_seq) as plan_name,
								  IFNULL(
								 		IF(
								 			instt.usage_limit > 0,
								 			instt.usage_limit,
								 			(
								 				SELECT
								 					usage_limit
								 				FROM
								 					t_products prod,
								 					t_membership_plan plan
								 				WHERE
								 					prod.plan_seq = plan.plan_seq
								 					AND prod.product_seq = instt.product_seq
								 			)
								 		),
								 		0
								 	) as usage_limit,
								 instt.usage_limit as free_usage_limit,
 								 instt.membership_payment_date, # 회원가입 결제일
								 instt.expiration_date, # 만료일
							   date_format(LAST_DAY(FROM_UNIXTIME(instt.membership_payment_date)), '%Y-%m-%d') AS expiration_date2, # 만료일 YYYY-MM-DD
								 DATE_ADD(instt.expiration_date, INTERVAL 1 DAY) AS payment_due_date, # 다음 결제일
								 date_format(instt.free_start_date, '%Y-%m-%d') AS free_start_date,
								 date_format(instt.free_end_date, '%Y-%m-%d') AS free_end_date
						FROM t_institution instt
					 WHERE instt.institution_seq = ?`
	row := common.DB_fetch_one(sql, nil, membership.Institution_seq)
	mStatus := common.ToUint(row["membership_fee_status"])
	paymentSetting := common.ToUint(row["payment_setting"])
	useMembershipName := common.ToStr(row["plan_name"])
	paymentDateInt := cast.ToInt64(row["membership_payment_date"])
	billKey := common.ToStr(row["bid"])
	usageLimit := common.ToUint(row["usage_limit"])
	paymentDueDate  := common.ToStr(row["payment_due_date"])
	freeUsageLimit := common.ToUint(row["free_usage_limit"])
	defaultFreeStart := ""
	defaultFreeEnd := ""
	addFreeStart := common.ToStr(row["free_start_date"])
	addFreeEnd := common.ToStr(row["free_end_date"])

	sql = `SELECT plan_seq, name, desc_text, usage_limit,
								IF((SELECT plan_seq
											FROM t_institution instt, t_products prod
										 WHERE prod.product_seq = instt.product_seq
										 	 AND instt.institution_seq = ?) = plan_seq, true , false) as current_use
					 FROM t_membership_plan
					WHERE plan_available = 1
						AND deleted_flag = 0`
	filter := func (row map[string]interface{}) {
		sql2 := `SELECT product_seq, plan_seq,	category, price,
									  discount_rate, discounted_amount
							 FROM t_products
							WHERE plan_seq = ?`
	  rows2 := common.DB_fetch_all(sql2, nil, row["plan_seq"])
		row["membership_plan"] = rows2
	}

	plans := common.DB_fetch_all(sql, filter, membership.Institution_seq)
	sql = `SELECT pname, amount, order_type, auth_date, pay_method
				   FROM t_orders
					WHERE institution_seq = ?
					  AND order_status = 1
						AND order_type IN (1, 2)
					ORDER BY auth_date DESC
					LIMIT 0, 1`
	paymentInfo := common.DB_fetch_one(sql, nil, membership.Institution_seq)

	// 회원 가입이 되어 있을때
	if mStatus > 0 && paymentDateInt > 0 {

		// 이용기간
		paymentDateToUnix, _ := goment.Unix(paymentDateInt)
		paymentDate, err := goment.New(paymentDateToUnix)
		if nil != err {
			log.Println(err)
		}

		now, _ := goment.New()

		// 회원 가입한 달과 현재 달이 같은 경우
		if paymentDate.Format("YYYYMM") == now.Format("YYYYMM") {
			defaultFreeStart = paymentDate.Format("YYYY-MM-DD")
			defaultFreeEnd = common.ToStr(row ["expiration_date2"])
		} else {
			if freeUsageLimit > 0 {
				useMembershipName = "IPSAP 무료 멤버십"
				// 무료 이용 기간
			}
		}
	}

	// Free period of Use
	// 결제 관련
	billKeyExist := false
	if "" != billKey {
		billKeyExist = true
	}

	instt := Institution {
		Institution_seq : membership.Institution_seq,
	}
	result = gin.H {
		"default_free_start" : defaultFreeStart,
		"default_free_end" : defaultFreeEnd,
		"add_free_start" : addFreeStart,
		"add_free_end" : addFreeEnd,
		"membership_payment_date" : paymentDateInt,
		"bill_key_exist" : billKeyExist,
		"use_membership_name" : useMembershipName,
		"membership_fee_status" : mStatus,
		"plans" : plans,
		"payment_info" : paymentInfo,
		"payment_setting" : paymentSetting,
		"this_month_service_count" : instt.GetInstitutionApplicationCnt(),
		"usage_limit" : usageLimit,
		"payment_due_date" : paymentDueDate,
	}

	return
}

func (mCancel *MembershipCancel) getMembershipCancelQuery(moreCondition string) (sql string) {
	sql = fmt.Sprintf(`
					SELECT instt.membership_fee_status, instt.membership_payment_date,
								 ord.amount, ord.order_seq
						FROM t_institution instt
						LEFT OUTER JOIN t_orders ord ON (ord.institution_seq = instt.institution_seq)
						LEFT OUTER JOIN t_products prod ON (ord.product_seq = prod.product_seq)
					 WHERE instt.institution_seq = %v
						 AND ord.order_status = %v
						 AND ord.order_type IN(%v, %v)
						 AND (SELECT COUNT(ord2.order_seq)
										FROM t_orders ord2
									 WHERE ord2.tid = ord.tid
										 AND order_type IN (%v, %v)
										 AND ord.order_status = %v)= 0
 						  %v #moreCondition
						`,mCancel.Institution_seq, DEF_ORDER_STATUS_COMPLETED,
							DEF_ORDER_TYPE_NORMAL_PAYMENT, DEF_ORDER_TYPE_REGULAR_PAYMENT,
						  DEF_ORDER_TYPE_ALL_CANCEL, DEF_ORDER_TYPE_PARTIAL_CANCEL,
							DEF_ORDER_STATUS_COMPLETED, moreCondition)
	return
}

func (mCancel *MembershipCancel) GetMembershipCancelInfo() {
	now, _ := goment.New()
	moreCondition := fmt.Sprintf(`AND prod.category = 'membership'`)
  sql := mCancel.getMembershipCancelQuery(moreCondition)
	row := common.DB_fetch_one(sql, nil)
	// 회원 가입비 계산
	if nil != row {
	  membershipFeeStatus := common.ToUint(row["membership_fee_status"])
		if membershipFeeStatus == 0 {
			return
		}

		// 환불금액: 가입비 - {(가입비 ÷ 365일)*해당년도 이용일수}
		// 당일 사용 하여도 하루치를 빼고 환불한다.
		paymentDateInt := cast.ToInt64(row["membership_payment_date"])
		paymentDateToUnix, _ := goment.Unix(paymentDateInt)
		paymentDate, _ := goment.New(paymentDateToUnix)
		usedDay := now.Diff(paymentDate, "days") + 1

		if usedDay < 365 {
			// 회원 가입비 환불액 0원
			membershipFee := cast.ToFloat64(row["amount"])
			amount := membershipFee - (float64((membershipFee/365)) * float64(usedDay))
			amountStr := common.ToStr(math.Floor(amount))
			tmp := amountStr[:len(amountStr) -1] + "0"
			mCancel.MembershipAmt = common.ToUint(tmp)
			if mCancel.MembershipAmt > 0 {
				mCancel.OrderSeqArr = append(mCancel.OrderSeqArr, common.ToUint(row["order_seq"]))
				mCancel.CancelAmtArr = append(mCancel.CancelAmtArr,mCancel.MembershipAmt)
				mCancel.PartialCancelCodeArr= append(mCancel.PartialCancelCodeArr, "1") // 부분취소
			}
		}
	}

	// 이번달 월정액 결제 내역이 있을시
	moreCondition = fmt.Sprintf(`AND ord.the_date = date_format(NOW(), '%%Y%%m')
															 AND prod.category = 'plan'` )
  sql = mCancel.getMembershipCancelQuery(moreCondition)
	row2 := common.DB_fetch_one(sql, nil)
	if nil != row2 {
		daysInMonth := now.DaysInMonth()
	  toDay	 := now.Date()
		if (daysInMonth - toDay) != 0 {
			planFee := cast.ToFloat64(row2["amount"])
			amount := planFee - (float64((planFee/float64(daysInMonth))) * float64(toDay))
			amountStr := common.ToStr(math.Floor(amount))
			tmp := amountStr[:len(amountStr) -1] + "0"
			mCancel.PlanAmt = common.ToUint(tmp)
			if mCancel.PlanAmt > 0 {
				mCancel.OrderSeqArr = append(mCancel.OrderSeqArr, common.ToUint(row2["order_seq"]))
				mCancel.CancelAmtArr = append(mCancel.CancelAmtArr,mCancel.PlanAmt)
				mCancel.PartialCancelCodeArr= append(mCancel.PartialCancelCodeArr, "1")
			}
		}
	}

	// 미리 결제한 월정액 결제 내역이 있을시
	moreCondition = fmt.Sprintf(`AND ord.the_date > date_format(NOW(), '%%Y%%m')
															 AND prod.category = 'plan'`)
  sql = mCancel.getMembershipCancelQuery(moreCondition)
	row3 := common.DB_fetch_one(sql, nil)
	if nil != row3 {
		mCancel.PlanAmt += common.ToUint(row3["amount"])
		if common.ToUint(row3["amount"]) > 0 {
			mCancel.OrderSeqArr = append(mCancel.OrderSeqArr, common.ToUint(row3["order_seq"]))
			mCancel.CancelAmtArr = append(mCancel.CancelAmtArr, common.ToUint(row3["amount"]))
			mCancel.PartialCancelCodeArr= append(mCancel.PartialCancelCodeArr, "0") // 전체취소
		}
	}

	sql = fmt.Sprintf(`
				 SELECT plan.name, date_format(FROM_UNIXTIME(instt.membership_payment_date), "%%Y%%m%%d") as join_date,
								(SELECT group_concat(user.user_seq)
									 FROM t_user user
									WHERE user.institution_seq = instt.institution_seq
										AND user.user_type LIKE '%%%v%%'
										AND user.user_status = 2) AS admin_user_arr
					 FROM t_institution instt, t_products prod, t_membership_plan plan
					WHERE instt.product_seq = prod.product_seq
					  AND prod.plan_seq = plan.plan_seq
						AND instt.institution_seq = %v`, DEF_USER_TYPE_ADMIN_SECRETARY, mCancel.Institution_seq)
	row4 := common.DB_fetch_one(sql, nil)
	mCancel.Pname = common.ToStr(row4["name"])
	mCancel.MebershipJoinDate = common.ToStr(row4["join_date"])
	mCancel.AdminUserSeqArr = common.ToStr(row4["admin_user_arr"])
	return
}

func (mCancel *MembershipCancel) MembershipCancel() (succ bool) {
	succ = false
	mCancel.GetMembershipCancelInfo()
	for idx, orderSeq := range mCancel.OrderSeqArr{
		cancOrd := CancelOrder{
			Order : Order{
				Order_seq  : orderSeq,
			},
			CancelAmt : common.ToStr(mCancel.CancelAmtArr[idx]),
			PartialCancelCode : mCancel.PartialCancelCodeArr[idx],
			CancelUserSeq : mCancel.CancelUserSeq,
		}
		cancOrd.CancelOrder()
	}

	// 기관 상태 업데이트
	instt := Institution {
		Institution_seq : mCancel.Institution_seq,
	}

	if !instt.MambershipCancel() {
		return
	}

	go mCancel.SendMembershipCancelMsg()
	succ = true
	return
}

func (mCancel *MembershipCancel) SendMembershipCancelMsg() (succ bool) {
	membershipInfo := make(map[string]string)
	membershipInfo[DEF_MEMBERSHIP_NAME]					= mCancel.Pname
	membershipJoinDate, _ := goment.New(mCancel.MebershipJoinDate)
	membershipInfo[DEF_MEMBERSHIP_JOIN_DATE]		= common.GetDateStr(membershipJoinDate)

	t1, _ := goment.New()
	membershipInfo[DEF_CANCEL_APPLICATION_DATE] = common.GetDateStr(t1)
	membershipInfo[DEF_REFUND_APPLICATION_DATE]	= common.GetDateTimeStr(t1)
	t1.Add(1, "days")
	t1.SetHour(0)
	t1.SetMinute(0)
	t1.SetSecond(0)
	membershipInfo[DEF_CANCEL_DATE]							= common.GetDateTimeStr(t1)
	membershipInfo[DEF_MEMBERSHIP_REFUND]				= common.GetCurrency(int64(mCancel.MembershipAmt))
	membershipInfo[DEF_USAGE_REFUND]						= common.GetCurrency(int64(mCancel.PlanAmt))
	membershipInfo[DEF_CANCEL_TOTAL_AMOUNT]			= common.GetCurrency(int64(mCancel.MembershipAmt + mCancel.PlanAmt))

	userArr := strings.Split(mCancel.AdminUserSeqArr, ",")
	msg := MessageMgr{}
	for _, userSeq := range userArr {
		user := User{
			User_seq : common.ToUint(userSeq),
		}
		if !user.Load(){
			continue;
		}
		msg.Msg_ID = DEF_MSG_MEMBERSHIP_CANCEL
		msg.User_info = user.Data
		msg.Mebership_info = membershipInfo
		msg.Institution_seq = mCancel.Institution_seq
		msg.SendMessage()
	}

	msg.Msg_ID = DEF_MSG_MEMBERSHIP_CANCEL_NOTI
	msg.User_info = LoadServiceAdmin()
	msg.SendMessage()

	return
}

func (resp *PaymentCancelResp) InsertResponse() (succ bool) {
	succ = false
	sql := `INSERT INTO t_resp_pg_billing_cancel(
						tid, moid, mid, result_code,
						result_msg, error_cd, error_msg,
						cancel_amt, signature, pay_method,
						cancel_date, cancel_time, cancel_num,
						remain_amt, mall_reserved
					)
					VALUES
					(?,?,?,?,
					 ?,?,?,
					 ?,?,?,
					 ?,?,?,
					 ?,?)`
	_, err := common.DBconn().Exec(sql,
																 resp.TID, resp.Moid, resp.MID, resp.ResultCode,
																 resp.ResultMsg, resp.ErrorCD, resp.ErrorMsg,
																 resp.CancelAmt, resp.Signature, resp.PayMethod,
																 resp.CancelDate, resp.CancelTime, resp.CancelNum,
																 resp.RemainAmt, resp.MallReserved)
	if nil != err {
		log.Println(err)
		return
	}

	return true
}

func (req *PaymentCancelReq) InsertRequest()(succ bool) {
	succ = false
	sql := `INSERT INTO t_req_pg_billing_cancel(
						tid, mid, moid,
						cancel_amt, cancel_msg, partial_cancel_code,
						edi_date, sign_data, mall_reserved
					)
					VALUES(
						?,?,?,
						?,?,?,
						?,?,?)`
	_, err := common.DBconn().Exec(sql,
																 req.TID, req.MID, req.Moid,
																 req.CancelAmt, req.CancelMsg, req.PartialCancelCode,
																 req.EdiDate, req.SignData, req.MallReserved)
	if nil != err {
		log.Println(err)
		return
	}

	return true
}
