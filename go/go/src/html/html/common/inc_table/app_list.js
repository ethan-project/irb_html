var saved_list, g_user_type, g_list_type;
function loadApplicationList(P, e, a, _) {
   null == P && (P = IPSAP.USER_TYPE.RESEARCHER), null == e && (e = IPSAP.APP_LIST_TYPE.ALL), null == a && (a = !1);
   var I = [],
      A = { "filter.app_view_type": "" };
   (g_user_type = P), (g_list_type = e);
   var t = location.search,
      s = new URLSearchParams(t),
      E = s.get("filter.start_index"),
      i = s.get("search_words"),
      S = 0;
   E && (S = E);
   var T,
      r = _;
   switch (P) {
      case IPSAP.USER_TYPE.RESEARCHER:
         setListFilter("committee", "iacuc,ibc,irb"),
            setListFilter("appType", "new"),
            setListFilter("status", "saved,approved,denied,experimenting", "last"),
            (a = !1),
            I.push("proc_exec_field", "app_view_set2"),
            $("#datatable").attr("data-buttons-text", "실험 계획서 작성"),
            $("#datatable").attr("data-buttons-icon", "fas fa-edit"),
            $("#datatable").attr("data-buttons-url", "./application_new.html"),
            (title = "전체 동물 실험 계획서 목록"),
            (titleOption = "작성중1"),
            (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.RESEARCHER_ALL),
            I.push("app_view_set3");
         break;
      case IPSAP.USER_TYPE.ADMIN_OFFICER:
         switch ((I.push("app_view_set2"), e)) {
            case IPSAP.APP_LIST_TYPE.ALL:
               I.push("app_view_set1"),
                  setListFilter("committee", "iacuc,ibc,irb"),
                  setListFilter("appType", "new"),
                  setListFilter("status", "delayed,approved,denied,experimenting", "last"),
                  (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.ADMIN_ALL);
               break;
            case IPSAP.APP_LIST_TYPE.CHECK:
               I.push("app_view_set3"),
                  setListFilter("committee", "iacuc,ibc,irb"),
                  setListFilter("status", "office_1,office_2", "last"),
                  (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.ADMIN_REVIEW_OFFICE);
               break;
            case IPSAP.APP_LIST_TYPE.JUDGE_FINISH:
               I.push("app_view_set3"),
                  setListFilter("committee", "iacuc,ibc,irb"),
                  setListFilter("progress", "expert,general"),
                  setListFilter("status", "delayed", "last"),
                  (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.ADMIN_REVIEW_CLOSE);
         }
         break;
      case IPSAP.USER_TYPE.OFFICER:
         switch ((I.push("app_view_set2"), I.push("app_view_set3"), e)) {
            case IPSAP.APP_LIST_TYPE.ALL:
               setListFilter("committee", "iacuc,ibc,irb"), setListFilter("status", "delayed,approved,denied,experimenting", "last"), (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.ADMIN_ALL);
               break;
            case IPSAP.APP_LIST_TYPE.CHECK:
               setListFilter("committee", "iacuc,ibc,irb"), setListFilter("status", "office_1,office_2", "last"), (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.ADMIN_REVIEW_OFFICE);
               break;
            case IPSAP.APP_LIST_TYPE.JUDGE_FINISH:
               setListFilter("committee", "iacuc,ibc,irb"),
                  setListFilter("progress", "expert,general"),
                  setListFilter("status", "delayed", "last"),
                  (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.ADMIN_REVIEW_CLOSE);
               break;
            case IPSAP.APP_LIST_TYPE.FINAL_FINISH:
               A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.ADMIN_REVIEW_CONFIRM;
         }
         break;
      case IPSAP.USER_TYPE.CHAIRMAN:
         switch ((I.push("app_view_set3"), e)) {
            case IPSAP.APP_LIST_TYPE.ALL:
               I.push("app_view_set2"), setListFilter("committee", "iacuc,ibc,irb", "last"), (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.CHAIRMAN_REVIEW_CONFIRM);
               break;
            case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
               I.push("proc_exec_field", "app_view_set1"), setListFilter("committee", "iacuc,ibc,irb", "last"), (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.CHAIRMAN_REVIEW_LIST);
         }
         break;
      case IPSAP.USER_TYPE.COMMITTEE:
         switch ((I.push("app_view_set3"), e)) {
            case IPSAP.APP_LIST_TYPE.ALL:
               I.push("app_view_set2"), (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.COMMITTEE_REVIEW);
               break;
            case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
               I.push("proc_exec_field", "app_view_set1"),
                  setListFilter("committee", "iacuc,ibc,irb"),
                  setListFilter("myReview", "expert,general", "last"),
                  (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.COMMITTEE_REVIEW_LIST);
         }
         break;
      case IPSAP.USER_TYPE.ALL:
         switch ((I.push("app_view_set3"), e)) {
            case IPSAP.APP_LIST_TYPE.INSPECT_RECORD:
               I.push("app_view_set2"), (A["filter.app_view_type"] = IPSAP.APP_LIST_TYPE_PARAM.INSPECT_RECORD);
               break;
            case IPSAP.APP_LIST_TYPE.ALL:
               I.push("proc_exec_field", "app_view_set2"), setListFilter("status", "delayed,approved,denied,experimenting", "last"), (A["filter.app_view_type"] = "");
         }
   }
   for (let P = 0; P < I.length; ++P) $("." + I[P]).remove();
   if (
      (a ? (T = $(".card_templete").clone()) : ($("#datatable").removeClass("card_type"), $("#datatable").addClass("list_type"), $("#type_changer").remove(), (T = $(".list_templete").clone())),
      IPSAP.DEMO_MODE)
   )
      return void n(P, e);
   let p = $("#datatable").find("tbody");
   function n(P, e, a) {
      var _ = "",
         I = "";
      switch (P) {
         case IPSAP.USER_TYPE.RESEARCHER:
            (_ = "전체 동물실험 계획서/보고서 목록"), (I = "작성중1");
            break;
         case IPSAP.USER_TYPE.ADMIN_OFFICER:
         case IPSAP.USER_TYPE.OFFICER:
            switch (e) {
               case IPSAP.APP_LIST_TYPE.ALL:
                  (_ = "전체 동물실험 계획서 목록"), (I = "작성중:1");
                  break;
               case IPSAP.APP_LIST_TYPE.CHECK:
                  (_ = "행정검토 및 심사설정 대상 목록"), (I = "작성중:1");
                  break;
               case IPSAP.APP_LIST_TYPE.JUDGE_FINISH:
                  (_ = "심사설정 대상 목록"), (I = "심사지연:1"), $(".text-muted").addClass("red");
                  break;
               case IPSAP.APP_LIST_TYPE.FINAL_FINISH:
                  (_ = "최종심의 대상 목록"), (I = "작성중:1");
            }
            break;
         case IPSAP.USER_TYPE.CHAIRMAN:
            switch (e) {
               case IPSAP.APP_LIST_TYPE.ALL:
                  _ = "최종심의 대상 목록";
                  break;
               case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                  _ = "신청서 및 보고서 심사 기록";
            }
            break;
         case IPSAP.USER_TYPE.COMMITTEE:
            switch (e) {
               case IPSAP.APP_LIST_TYPE.ALL:
                  _ = "심사 대상 목록";
                  break;
               case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                  _ = "신청서 및 보고서 심사 기록";
            }
      }
      IPSAP.APP_LIST_TYPE.INSPECT_RECORD != e &&
         ($(".card-title").text(_),
         $(".text-muted").text(I),
         $("input[name=list_type]").on({
            change: function (P) {
               $("#datatable").removeClass("card_type list_type"), $("#datatable").addClass($("input[name=list_type]:checked").val());
            },
         })),
         setTriggerTrDataUrlOnClick(),
         initDataTable("#datatable");
   }
   function c(e, a, _) {
      let I = "disabled",
         A = "btn-outline-primary";
      switch (_) {
         case IPSAP.APP_PROC_BTN_KIND.CHECK:
         case IPSAP.IBC_APP_PROC_BTN_KIND.CHECK:
         case IPSAP.IRB_APP_PROC_BTN_KIND.CHECK:
            a.application_result == IPSAP.APPLICATION_RESULT.CHECKING && ((I = ""), (A = "btn-primary"));
            break;
         case IPSAP.APP_PROC_BTN_KIND.JUDGE_SET:
         case IPSAP.IBC_APP_PROC_BTN_KIND.JUDGE_SET:
         case IPSAP.IRB_APP_PROC_BTN_KIND.JUDGE_SET:
            if ((a.application_step != IPSAP.APPLICATION_STEP.PRO_JUDGE && a.application_step != IPSAP.APPLICATION_STEP.COMM_JUDGE) || a.application_result != IPSAP.APPLICATION_RESULT.JUDGE_DELAY)
               switch (a.application_result) {
                  case IPSAP.APPLICATION_RESULT.CHECKING_2:
                  case IPSAP.APPLICATION_RESULT.JUDGE_ING:
                  case IPSAP.APPLICATION_RESULT.JUDGE_ING_2:
                     (I = ""), (A = "btn-primary");
               }
            else (I = ""), (A = "btn-primary");
            break;
         case IPSAP.APP_PROC_BTN_KIND.JUDGE_FINISH:
         case IPSAP.APP_PROC_BTN_KIND.JUDGE_FINISH2:
         case IPSAP.IBC_APP_PROC_BTN_KIND.JUDGE_FINISH:
         case IPSAP.IBC_APP_PROC_BTN_KIND.JUDGE_FINISH2:
            a.application_step == IPSAP.APPLICATION_STEP.COMM_JUDGE && a.application_result == IPSAP.APPLICATION_RESULT.JUDGE_DELAY && ((I = ""), (A = "btn-primary"));
            break;
         case IPSAP.APP_PROC_BTN_KIND.FINAL:
         case IPSAP.IBC_APP_PROC_BTN_KIND.FINAL:
         case IPSAP.IRB_APP_PROC_BTN_KIND.FINAL:
            a.application_step == IPSAP.APPLICATION_STEP.FINAL && a.application_result == IPSAP.APPLICATION_RESULT.DECISION_ING && ((I = ""), (A = "btn-primary"));
            break;
         case IPSAP.APP_PROC_BTN_KIND.PRO_JUDGE:
         case IPSAP.IBC_APP_PROC_BTN_KIND.PRO_JUDGE:
         case IPSAP.IRB_APP_PROC_BTN_KIND.PRO_JUDGE:
            a.application_step != IPSAP.APPLICATION_STEP.PRO_JUDGE ||
               (a.application_result != IPSAP.APPLICATION_RESULT.JUDGE_ING &&
                  a.application_result != IPSAP.APPLICATION_RESULT.JUDGE_ING_2 &&
                  a.application_result != IPSAP.APPLICATION_RESULT.JUDGE_DELAY) ||
               ((I = ""), (A = "btn-primary"));
            break;
         case IPSAP.APP_PROC_BTN_KIND.NOR_JUDGE:
         case IPSAP.IBC_APP_PROC_BTN_KIND.NOR_JUDGE:
         case IPSAP.IRB_APP_PROC_BTN_KIND.NOR_JUDGE:
            a.application_step != IPSAP.APPLICATION_STEP.COMM_JUDGE ||
               (a.application_result != IPSAP.APPLICATION_RESULT.JUDGE_ING &&
                  a.application_result != IPSAP.APPLICATION_RESULT.JUDGE_ING_2 &&
                  a.application_result != IPSAP.APPLICATION_RESULT.JUDGE_DELAY) ||
               ((I = ""), (A = "btn-primary"));
      }
      a.application_result == IPSAP.APPLICATION_RESULT.JUDGE_DELAY && P == IPSAP.USER_TYPE.OFFICER && (A = "btn-danger");
      var t = _.url[saved_list[e].application_type - 1];
      html = `<a href="javascript:void(0);" onclick="onAppNavigate(event, ${e}, '${t}');" class="btn ${A} btn-xs ${I}">${_.text}</a>\n    `;
      let s = !1;
      return "disabled" != I && (s = !0), { enabled: s, html: html };
   }
   p.empty(),
      (A["filter.start_index"] = S),
      (A["filter.row_cnt"] = r),
      (A.search_words = i),
      API.load({
         url: CONST.API.APPLICATION.LIST,
         type: CONST.API_TYPE.GET,
         data: A,
         success: function (_) {
            (saved_list = _.list),
               0 == _.list.length && $(".card").removeClass("hidden"),
               $.each(_.list, function (_, A) {
                  (A.tmp_reg_dttm = getDttm(A.reg_dttm)), (A.tmp_approved_dttm = getDttm(A.approved_dttm)), (A.tmp_submit_dttm = getDttm(A.submit_dttm));
                  let t = T.clone(),
                     s = 0;
                  if (
                     (IPSAP.APP_LIST_TYPE.INSPECT_RECORD != e && t.children().eq(s++).text(A.parent_app_seq),
                     $(".card").removeClass("hidden"),
                     t.children().eq(s++).text(A.application_no),
                     P == IPSAP.USER_TYPE.CHAIRMAN || P == IPSAP.USER_TYPE.COMMITTEE
                        ? P == IPSAP.USER_TYPE.COMMITTEE && e == IPSAP.APP_LIST_TYPE.JUDGE_HISTROY
                           ? ((A.tmp_normal_review_dttm = getDttm(A.normal_review_dttm)),
                             (A.tmp_expert_review_dttm = getDttm(A.expert_review_dttm)),
                             0 == A.normal_review_dttm ? t.children().eq(s++).text(A.tmp_expert_review_dttm.dt) : t.children().eq(s++).text(A.tmp_normal_review_dttm.dt))
                           : P == IPSAP.USER_TYPE.CHAIRMAN && e == IPSAP.APP_LIST_TYPE.JUDGE_HISTROY
                           ? t.children().eq(s++).text(A.tmp_approved_dttm.dt)
                           : t.children().eq(s++).text(A.tmp_reg_dttm.dt)
                        : P == IPSAP.USER_TYPE.RESEARCHER
                        ? t.children().eq(s++).text(A.tmp_reg_dttm.dt)
                        : t.children().eq(s++).text(A.tmp_submit_dttm.dt),
                     t
                        .children()
                        .eq(s++)
                        .text(A.name_ko + " (" + A.name_en + ")"),
                     t.children().eq(s++).text(A.judge_type_str),
                     P == IPSAP.USER_TYPE.CHAIRMAN && e == IPSAP.APP_LIST_TYPE.JUDGE_HISTROY)
                  ) {
                     var E = "<span class='label committee_bg'>최종 심의</span>";
                     t.children().eq(s++).append(E);
                  }
                  if (P == IPSAP.USER_TYPE.COMMITTEE && e == IPSAP.APP_LIST_TYPE.JUDGE_HISTROY) {
                     E = `<span class='label committee_bg'>${A.committee_judge_type_str}</span>`;
                     t.children().eq(s++).append(E);
                  }
                  t.children().eq(s++).text(A.application_type_str),
                     5 == A.application_step && 1 != A.application_type ? t.children().eq(s++).text("승인") : t.children().eq(s++).text(A.application_step_str),
                     (P != IPSAP.USER_TYPE.RESEARCHER && P != IPSAP.USER_TYPE.COMMITTEE && P != IPSAP.USER_TYPE.CHAIRMAN) || (5 != A.application_result && 6 != A.application_result)
                        ? P == IPSAP.USER_TYPE.OFFICER && 5 == A.application_result
                           ? t.children().eq(s++).text("심사중")
                           : P != IPSAP.USER_TYPE.RESEARCHER || (2 != A.application_result && 3 != A.application_result)
                           ? P != IPSAP.USER_TYPE.RESEARCHER || (11 != A.application_result && 12 != A.application_result)
                              ? 14 == A.application_result
                                 ? t.children().eq(s++).text("실험 종료")
                                 : t.children().eq(s++).text(A.application_result_str)
                              : t.children().eq(s++).text("실험중")
                           : t.children().eq(s++).text("검토중")
                        : t.children().eq(s++).text("심사중").removeClass("status"),
                     e == IPSAP.APP_LIST_TYPE.ALL && P == IPSAP.USER_TYPE.ADMIN_OFFICER && t.children().eq(s++).text(A.expert_member),
                     t.children().eq(s++).text(A.user_name),
                     t
                        .children()
                        .eq(s++)
                        .text(A.istt_name_ko + " (" + A.istt_name_en + ")"),
                     t.children().eq(s++).text(A.user_dept_str),
                     isInArray(I, "proc_exec_field") ||
                        (t.children().eq(s).empty(),
                        (function (_, I, A) {
                           var t = !1,
                              s = [];
                           switch (saved_list[_].judge_type) {
                              case IPSAP.JUDGE_TYPE.IACUC:
                                 switch (e) {
                                    case IPSAP.APP_LIST_TYPE.ALL:
                                       switch (P) {
                                          case IPSAP.USER_TYPE.ADMIN_OFFICER:
                                             s.push(IPSAP.APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.APP_PROC_BTN_KIND.JUDGE_SET), s.push(IPSAP.APP_PROC_BTN_KIND.FINAL);
                                             break;
                                          case IPSAP.USER_TYPE.OFFICER:
                                             s.push(IPSAP.APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.APP_PROC_BTN_KIND.JUDGE_SET);
                                             break;
                                          case IPSAP.USER_TYPE.CHAIRMAN:
                                             s.push(IPSAP.APP_PROC_BTN_KIND.FINAL);
                                             break;
                                          case IPSAP.USER_TYPE.COMMITTEE:
                                             s.push(IPSAP.APP_PROC_BTN_KIND.PRO_JUDGE), s.push(IPSAP.APP_PROC_BTN_KIND.NOR_JUDGE);
                                          case IPSAP.USER_TYPE.RESEARCHER:
                                       }
                                       break;
                                    case IPSAP.APP_LIST_TYPE.CHECK:
                                       s.push(IPSAP.APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.APP_PROC_BTN_KIND.JUDGE_SET);
                                       break;
                                    case IPSAP.APP_LIST_TYPE.JUDGE_FINISH:
                                       s.push(IPSAP.APP_PROC_BTN_KIND.JUDGE_SET);
                                       break;
                                    case IPSAP.APP_LIST_TYPE.FINAL_FINISH:
                                       (t = !0), s.push(IPSAP.APP_PROC_BTN_KIND.FINAL);
                                       break;
                                    case IPSAP.APP_LIST_TYPE.JUDGE:
                                       s.push(IPSAP.APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.APP_PROC_BTN_KIND.JUDGE_SET);
                                    case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                                 }
                                 break;
                              case IPSAP.JUDGE_TYPE.IBC:
                                 switch (e) {
                                    case IPSAP.APP_LIST_TYPE.ALL:
                                       switch (P) {
                                          case IPSAP.USER_TYPE.ADMIN_OFFICER:
                                             s.push(IPSAP.IBC_APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.IBC_APP_PROC_BTN_KIND.JUDGE_SET), s.push(IPSAP.IBC_APP_PROC_BTN_KIND.FINAL);
                                             break;
                                          case IPSAP.USER_TYPE.OFFICER:
                                             s.push(IPSAP.IBC_APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.IBC_APP_PROC_BTN_KIND.JUDGE_SET);
                                             break;
                                          case IPSAP.USER_TYPE.CHAIRMAN:
                                             s.push(IPSAP.IBC_APP_PROC_BTN_KIND.FINAL);
                                             break;
                                          case IPSAP.USER_TYPE.COMMITTEE:
                                             s.push(IPSAP.IBC_APP_PROC_BTN_KIND.PRO_JUDGE), s.push(IPSAP.IBC_APP_PROC_BTN_KIND.NOR_JUDGE);
                                          case IPSAP.USER_TYPE.RESEARCHER:
                                       }
                                       break;
                                    case IPSAP.APP_LIST_TYPE.CHECK:
                                       s.push(IPSAP.IBC_APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.IBC_APP_PROC_BTN_KIND.JUDGE_SET);
                                       break;
                                    case IPSAP.APP_LIST_TYPE.JUDGE_FINISH:
                                       s.push(IPSAP.IBC_APP_PROC_BTN_KIND.JUDGE_SET);
                                       break;
                                    case IPSAP.APP_LIST_TYPE.FINAL_FINISH:
                                       (t = !0), s.push(IPSAP.IBC_APP_PROC_BTN_KIND.FINAL);
                                       break;
                                    case IPSAP.APP_LIST_TYPE.JUDGE:
                                       s.push(IPSAP.IBC_APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.IBC_APP_PROC_BTN_KIND.JUDGE_SET);
                                    case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                                 }
                                 break;
                              case IPSAP.JUDGE_TYPE.IRB:
                                 switch (e) {
                                    case IPSAP.APP_LIST_TYPE.ALL:
                                       switch (P) {
                                          case IPSAP.USER_TYPE.ADMIN_OFFICER:
                                             s.push(IPSAP.IRB_APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.IRB_APP_PROC_BTN_KIND.JUDGE_SET), s.push(IPSAP.IRB_APP_PROC_BTN_KIND.FINAL);
                                             break;
                                          case IPSAP.USER_TYPE.OFFICER:
                                             s.push(IPSAP.IRB_APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.IRB_APP_PROC_BTN_KIND.JUDGE_SET);
                                             break;
                                          case IPSAP.USER_TYPE.CHAIRMAN:
                                             s.push(IPSAP.IRB_APP_PROC_BTN_KIND.FINAL);
                                             break;
                                          case IPSAP.USER_TYPE.COMMITTEE:
                                             s.push(IPSAP.IRB_APP_PROC_BTN_KIND.PRO_JUDGE), s.push(IPSAP.IRB_APP_PROC_BTN_KIND.NOR_JUDGE);
                                          case IPSAP.USER_TYPE.RESEARCHER:
                                       }
                                       break;
                                    case IPSAP.APP_LIST_TYPE.CHECK:
                                       s.push(IPSAP.IRB_APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.IRB_APP_PROC_BTN_KIND.JUDGE_SET);
                                       break;
                                    case IPSAP.APP_LIST_TYPE.FINAL_FINISH:
                                       (t = !0), s.push(IPSAP.IRB_APP_PROC_BTN_KIND.FINAL);
                                       break;
                                    case IPSAP.APP_LIST_TYPE.JUDGE:
                                       s.push(IPSAP.IRB_APP_PROC_BTN_KIND.CHECK), s.push(IPSAP.IRB_APP_PROC_BTN_KIND.JUDGE_SET);
                                    case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                                 }
                           }
                           a && A.addClass("wide_buttons order_end");
                           for (let P = 0; P < s.length; ++P) {
                              let e = c(_, I, s[P]);
                              if ((!t || e.enabled) && (A.append(e.html), t)) break;
                           }
                        })(_, A, t.children().eq(s++))),
                     e == IPSAP.APP_LIST_TYPE.ALL &&
                        P == IPSAP.USER_TYPE.ADMIN_OFFICER &&
                        t.children().eq(s++).children().attr("onClick", `onClickShowDeleteApplicationModal(event,${A.application_seq},'${A.application_no}')`),
                     t.attr("data-url", "javascript:void(0);"),
                     t.attr("onclick", `onClickApplication(${_})`),
                     t.attr("data-data", A.application_seq),
                     (function (e, a) {
                        if ((a.removeClass("IACUC IRB IBC"), a.addClass(e.judge_type_str), e.application_result == IPSAP.APPLICATION_RESULT.DELETED)) return void a.addClass("deleted");
                        switch (e.application_step) {
                           case IPSAP.APPLICATION_STEP.WRITE:
                              switch (e.application_result) {
                                 case IPSAP.APPLICATION_RESULT.TEMP:
                                    a.addClass("saved");
                                    break;
                                 case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
                                    a.addClass("supplement");
                              }
                              break;
                           case IPSAP.APPLICATION_STEP.FINAL:
                              switch (e.application_result) {
                                 case IPSAP.APPLICATION_RESULT.REQUIRE_RETRY:
                                 case IPSAP.APPLICATION_RESULT.REJECT:
                                    a.addClass("denied");
                              }
                              break;
                           case IPSAP.APPLICATION_STEP.PERFORMANCE:
                              switch (e.application_result) {
                                 case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
                                 case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
                                    a.addClass("inProgress");
                                    break;
                                 case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
                                    a.addClass("closed");
                                    break;
                                 case IPSAP.APPLICATION_RESULT.TASK_FINISH:
                                    a.addClass("completed");
                              }
                        }
                        e.application_result == IPSAP.APPLICATION_RESULT.JUDGE_DELAY && P == IPSAP.USER_TYPE.OFFICER && a.addClass("delayed");
                     })(A, t);
                  var i = t.wrapAll("<div/>").parent().html();
                  p.append(i);
               }),
               n(P, e, _.list);
            var A = _.pageInfo.totalCnt;
            _ = 0;
            A > (_ = 0 == S ? r * (S + 1) : r + parseInt(S)) && setBlockBtn("next", _, r), S > 0 && setBlockBtn("before", _, r, S);
            let t = $(".search_form"),
               s = $(".dataTables_filter");
            (i || A > _ || S > 0) && setServerSearch(t, s);
         },
         complete: function () {},
      });
}
function onDeleteApplication(P, e) {
   var a = CONST.API.APPLICATION.DELETE2;
   API.load({
      url: a.replace("${app_seq}", e),
      type: CONST.API_TYPE.DELETE,
      success: function (P) {},
      error: function (P) {},
      complete: function () {
         location.reload();
      },
   }),
      P.stopPropagation();
}
function onClickShowDeleteApplicationModal(P, e, a) {
   $(".app_no").text(a), $("#deleteApp").attr("onClick", `onDeleteApplication(event,'${e}')`), $("#modal_application_delete").modal("show"), P.stopPropagation();
}
function onAppNavigate(P, e, a) {
   P.stopPropagation(), g_AppInfo.initWithAppObj(saved_list[e]) || alert("Fail to ApplicationInfo Init!"), g_AppInfo.saveParamsAndNavigate(a);
}
function onClickApplication(P) {
    console.log("p", P)
   switch ((g_AppInfo.initWithAppObj(saved_list[P]) || alert("Fail to ApplicationInfo Init!"), saved_list[P].judge_type)) {
      case IPSAP.JUDGE_TYPE.IACUC:
         navigateIACUC(P);
         break;
      case IPSAP.JUDGE_TYPE.IBC:
         navigateIBC(P);
         break;
      case IPSAP.JUDGE_TYPE.IRB:
         navigateIRB(P);
         break;
      default:
         return void alert(`JUDGE_TYPE Error!! (${saved_list[P].judge_type})`);
   }
}
function navigateIACUC(P) {
   switch (saved_list[P].application_type) {
      case IPSAP.APPLICATION_TYPE.NEW:
         navigateIACUC_new(P);
         break;
      case IPSAP.APPLICATION_TYPE.CHANGE:
         navigateIACUC_change(P);
         break;
      case IPSAP.APPLICATION_TYPE.RENEW:
         navigateIACUC_renew(P);
         break;
      case IPSAP.APPLICATION_TYPE.BRING:
         navigateIACUC_bring(P);
         break;
      case IPSAP.APPLICATION_TYPE.END:
         navigateIACUC_end(P);
         break;
      case IPSAP.APPLICATION_TYPE.CHECKLIST:
         navigateIACUC_checklist(P);
         break;
      default:
         return void alert(`application_type Error!! (${saved_list[P].application_type})`);
   }
}
function navigateIACUC_new(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.TEMP:
         g_AppInfo.saveParamsAndNavigate(APP_NAVIGATION.IACUC.PAGE_INFO[0].URL);
         break;
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         g_AppInfo.saveParamsAndNavigate("./application_list-review.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_expert_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_general_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_ING_2:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_expert_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_general_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_expert_2.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_general_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.DECISION_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               break;
            case IPSAP.USER_TYPE.CHAIRMAN:
               switch (g_list_type) {
                  case IPSAP.APP_LIST_TYPE.ALL:
                     g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review.html");
                     break;
                  case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                     g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               }
         }
         break;
      case IPSAP.APPLICATION_RESULT.REJECT:
      case IPSAP.APPLICATION_RESULT.REQUIRE_RETRY:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
      case IPSAP.APPLICATION_RESULT.TASK_FINISH:
      case IPSAP.APPLICATION_RESULT.ACCEPT:
      case IPSAP.APPLICATION_RESULT.ACCEPT_AC:
         g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIACUC_change(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-review_change.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_change.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
               break;
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_ING:
      case IPSAP.APPLICATION_RESULT.JUDGE_ING_2:
      case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_expert_1_change.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_general_1_change.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.DECISION_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_change.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
               break;
            case IPSAP.USER_TYPE.CHAIRMAN:
               switch (g_list_type) {
                  case IPSAP.APP_LIST_TYPE.ALL:
                     g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_change.html");
                     break;
                  case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                     g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
               }
         }
         break;
      case IPSAP.APPLICATION_RESULT.REJECT:
      case IPSAP.APPLICATION_RESULT.REQUIRE_RETRY:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
      case IPSAP.APPLICATION_RESULT.TASK_FINISH:
      case IPSAP.APPLICATION_RESULT.ACCEPT:
      case IPSAP.APPLICATION_RESULT.ACCEPT_AC:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
            case IPSAP.USER_TYPE.COMMITTEE:
            case IPSAP.USER_TYPE.CHAIRMAN:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
         }
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIACUC_renew(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-review_renew.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_renew.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_renew.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_ING:
      case IPSAP.APPLICATION_RESULT.JUDGE_ING_2:
      case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_renew.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_expert_1_renew.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./experiment_list-info_renew.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IACUC_general_1_renew.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./experiment_list-info_renew.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_renew.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.DECISION_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_renew.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_renew.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
               break;
            case IPSAP.USER_TYPE.CHAIRMAN:
               switch (g_list_type) {
                  case IPSAP.APP_LIST_TYPE.ALL:
                     g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_renew.html");
                     break;
                  case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                     g_AppInfo.saveParamsAndNavigate("./experiment_list-info_renew.html");
               }
         }
         break;
      case IPSAP.APPLICATION_RESULT.REJECT:
      case IPSAP.APPLICATION_RESULT.REQUIRE_RETRY:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
      case IPSAP.APPLICATION_RESULT.TASK_FINISH:
      case IPSAP.APPLICATION_RESULT.ACCEPT:
      case IPSAP.APPLICATION_RESULT.ACCEPT_AC:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
            case IPSAP.USER_TYPE.COMMITTEE:
            case IPSAP.USER_TYPE.CHAIRMAN:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_renew.html");
         }
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIACUC_bring(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-review_bring.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_bring.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_bring.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.DECISION_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_bring.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_bring.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_bring.html");
               break;
            case IPSAP.USER_TYPE.CHAIRMAN:
               switch (g_list_type) {
                  case IPSAP.APP_LIST_TYPE.ALL:
                     g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_bring.html");
                     break;
                  case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                     g_AppInfo.saveParamsAndNavigate("./experiment_list-info_bring.html");
               }
         }
         break;
      case IPSAP.APPLICATION_RESULT.REJECT:
      case IPSAP.APPLICATION_RESULT.ACCEPT:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
            case IPSAP.USER_TYPE.COMMITTEE:
            case IPSAP.USER_TYPE.CHAIRMAN:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_bring.html");
         }
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIACUC_end(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-review_end.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_end.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_end.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.DECISION_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_end.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_end.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_end.html");
               break;
            case IPSAP.USER_TYPE.CHAIRMAN:
               switch (g_list_type) {
                  case IPSAP.APP_LIST_TYPE.ALL:
                     g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_end.html");
                     break;
                  case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                     g_AppInfo.saveParamsAndNavigate("./experiment_list-info_end.html");
               }
         }
         break;
      case IPSAP.APPLICATION_RESULT.REJECT:
      case IPSAP.APPLICATION_RESULT.ACCEPT:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
            case IPSAP.USER_TYPE.COMMITTEE:
            case IPSAP.USER_TYPE.CHAIRMAN:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_end.html");
         }
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIACUC_checklist(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         g_AppInfo.saveParamsAndNavigate("./inspection_list-review.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./inspection_list-info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_inspection.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.DECISION_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./inspection_list-info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_inspection.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               g_AppInfo.saveParamsAndNavigate("./inspection_list-info.html");
               break;
            case IPSAP.USER_TYPE.CHAIRMAN:
               switch (g_list_type) {
                  case IPSAP.APP_LIST_TYPE.ALL:
                     g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_inspection.html");
                     break;
                  case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                     g_AppInfo.saveParamsAndNavigate("./inspection_list-info.html");
               }
         }
         break;
      case IPSAP.APPLICATION_RESULT.REJECT:
      case IPSAP.APPLICATION_RESULT.ACCEPT:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./inspection_list-info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
            case IPSAP.USER_TYPE.COMMITTEE:
            case IPSAP.USER_TYPE.CHAIRMAN:
               g_AppInfo.saveParamsAndNavigate("./experiment_list-info_inspection.html");
         }
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIBC(P) {
   switch (saved_list[P].application_type) {
      case IPSAP.APPLICATION_TYPE.NEW:
         navigateIBC_new(P);
         break;
      case IPSAP.APPLICATION_TYPE.CHANGE:
         navigateIBC_change(P);
         break;
      default:
         return void alert(`application_type Error!! (${saved_list[P].application_type})`);
   }
}
function navigateIBC_new(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.TEMP:
         g_AppInfo.saveParamsAndNavigate(APP_IBC_NAVIGATION.IBC.PAGE_INFO[0].URL);
         break;
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_review.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_review.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IBC_expert_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IBC_general_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_ING_2:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IBC_expert_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IBC_general_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IBC_expert_2.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IBC_general_1.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.DECISION_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewConfirm_list-IBC_review.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
               break;
            case IPSAP.USER_TYPE.CHAIRMAN:
               switch (g_list_type) {
                  case IPSAP.APP_LIST_TYPE.ALL:
                     g_AppInfo.saveParamsAndNavigate("./IBC/reviewConfirm_list-IBC_review.html");
                     break;
                  case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                     g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
               }
         }
         break;
      case IPSAP.APPLICATION_RESULT.REJECT:
      case IPSAP.APPLICATION_RESULT.REQUIRE_RETRY:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
      case IPSAP.APPLICATION_RESULT.TASK_FINISH:
      case IPSAP.APPLICATION_RESULT.ACCEPT:
      case IPSAP.APPLICATION_RESULT.ACCEPT_AC:
         g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIBC_change(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_review_change.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_review_change.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
               break;
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./application_list-IBC_info.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.JUDGE_ING:
      case IPSAP.APPLICATION_RESULT.JUDGE_ING_2:
      case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               switch (saved_list[P].application_step) {
                  case IPSAP.APPLICATION_STEP.PRO_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IBC_expert_1_change.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
                     }
                     break;
                  case IPSAP.APPLICATION_STEP.COMM_JUDGE:
                     switch (g_list_type) {
                        case IPSAP.APP_LIST_TYPE.ALL:
                           g_AppInfo.saveParamsAndNavigate("./review_list-review_IBC_general_1_change.html");
                           break;
                        case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                           g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
                     }
               }
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
         }
         break;
      case IPSAP.APPLICATION_RESULT.DECISION_ING:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
               g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
               break;
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewConfirm_list-IBC_review_change.html");
               break;
            case IPSAP.USER_TYPE.COMMITTEE:
               g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
               break;
            case IPSAP.USER_TYPE.CHAIRMAN:
               switch (g_list_type) {
                  case IPSAP.APP_LIST_TYPE.ALL:
                     g_AppInfo.saveParamsAndNavigate("./IBC/reviewConfirm_list-IBC_review_change.html");
                     break;
                  case IPSAP.APP_LIST_TYPE.JUDGE_HISTROY:
                     g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
               }
         }
         break;
      case IPSAP.APPLICATION_RESULT.REJECT:
      case IPSAP.APPLICATION_RESULT.REQUIRE_RETRY:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
      case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
      case IPSAP.APPLICATION_RESULT.TASK_FINISH:
      case IPSAP.APPLICATION_RESULT.ACCEPT:
      case IPSAP.APPLICATION_RESULT.ACCEPT_AC:
         switch (g_user_type) {
            case IPSAP.USER_TYPE.RESEARCHER:
            case IPSAP.USER_TYPE.ALL:
            case IPSAP.USER_TYPE.ADMIN_OFFICER:
            case IPSAP.USER_TYPE.OFFICER:
            case IPSAP.USER_TYPE.COMMITTEE:
            case IPSAP.USER_TYPE.CHAIRMAN:
               g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
         }
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIRB(P) {
    console.log("IPSAP.APPLICATION_TYPE", IPSAP.APPLICATION_TYPE)
    retu
   switch (saved_list[P].application_type) {
      case IPSAP.APPLICATION_TYPE.NEW:
         navigateIRB_new(P);
         break;
      case IPSAP.APPLICATION_TYPE.CHANGE:
         navigateIRB_change(P);
         break;
      case IPSAP.APPLICATION_TYPE.MAINTANACE:
         navigateIRB_maintanance(P);
         break;
      case IPSAP.APPLICATION_TYPE.ADVERSE:
         navigateIRB_adverse(P);
         break;
      case IPSAP.APPLICATION_TYPE.VIOLATION:
         navigateIRB_violation(P);
         break;
      case IPSAP.APPLICATION_TYPE.PROBLEM_OCCUR:
         navigateIRB_problem(P);
         break;
      case IPSAP.APPLICATION_TYPE.END:
         navigateIRB_end(P);
         break;
      default:
         return void alert(`application_type Error!! (${saved_list[P].application_type})`);
   }
}
function navigateIRB_new(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.TEMP:
         g_AppInfo.saveParamsAndNavigate(APP_IRB_NAVIGATION.IRB.PAGE_INFO[0].URL);
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         if (g_user_type === IPSAP.USER_TYPE.RESEARCHER) g_AppInfo.saveParamsAndNavigate("./IRB/application_list-IRB_info.html");
   }
}
function navigateIRB_change(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         alert("개발중!");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         if (g_user_type === IPSAP.USER_TYPE.RESEARCHER) g_AppInfo.saveParamsAndNavigate("./IRB/experiment_list-IRB_info_change.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         alert("개발중!");
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIRB_maintanance(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         alert("개발중!");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         if (g_user_type === IPSAP.USER_TYPE.RESEARCHER) g_AppInfo.saveParamsAndNavigate("./IRB/experiment_list-IRB_info_renew.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         alert("개발중!");
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIRB_adverse(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         alert("개발중!");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         if (g_user_type === IPSAP.USER_TYPE.RESEARCHER) g_AppInfo.saveParamsAndNavigate("./IRB/experiment_list-IRB_info_adverse.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         alert("개발중!");
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIRB_violation(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         alert("개발중!");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         if (g_user_type === IPSAP.USER_TYPE.RESEARCHER) g_AppInfo.saveParamsAndNavigate("./IRB/experiment_list-IRB_info_violation.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         alert("개발중!");
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIRB_problem(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         alert("개발중!");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         if (g_user_type === IPSAP.USER_TYPE.RESEARCHER) g_AppInfo.saveParamsAndNavigate("./IRB/experiment_list-IRB_info_problem.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         alert("개발중!");
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
function navigateIRB_end(P) {
   switch (saved_list[P].application_result) {
      case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
         alert("개발중!");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING:
         if (g_user_type === IPSAP.USER_TYPE.RESEARCHER) g_AppInfo.saveParamsAndNavigate("./IRB/experiment_list-IRB_info_end.html");
         break;
      case IPSAP.APPLICATION_RESULT.CHECKING_2:
         alert("개발중!");
         break;
      default:
         return void alert(`APPLICATION_RESULT Error!! (${saved_list[P].application_result})`);
   }
}
