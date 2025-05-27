"undefined" == typeof ipsap_common_js && document.write("<script src='/assets/js/ipsap/ipsap_common.js'></script>"),
   "undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof date_utils_js && document.write("<script src='/assets/js/common/util/date_utils.js'></script>"),
   "undefined" == typeof ipsap_application_common_js && document.write("<script src='/assets/js/ipsap/ipsap_application_common.js'></script>");
var supplement_saved_list,
   checking_saved_list,
   judge_ing_saved_list,
   final_saved_list,
   iacuc_num = 0,
   ibc_num = 0,
   irb_num = 0,
   cardObjTmp = $(".myApp_body .card").eq(0).clone(),
   explistObjTmp = $(".myExps .card").eq(0).clone(),
   judge_type = JSON.parse(COMM.getCookie("institution_info")).judge_type.split(","),
   empty_apps =
      '\n<div class="card round_10 empty_apps">\n  <div class="card-header">\n    <div class="card_icon text-center"><i class="mdi mdi-file-upload-outline"></i></div>\n  </div>\n  <div class="card-body noBorder">\n    <div class="card-title text-center"><span class="empty_comment"></span> 없습니다.</div>\n  </div>\n</div>',
   empty_exps =
      '\n<div class="card round_10 empty_apps">\n  <div class="card-header">\n    <div class="card_icon text-center"><i class="mdi mdi-flask-outline"></i></div>\n  </div>\n  <div class="card-body noBorder">\n    <div class="card-title text-center">진행중인 실험이 없습니다.</div>\n  </div>\n</div>',
   empty_result =
      '\n<div class="card round_10 empty_result">\n  <div class="card-header">\n    <div class="card_icon text-center"><i class="mdi mdi-file-upload-outline"></i></div>\n  </div>\n  <div class="card-body noBorder">\n    <div class="card-title text-center">필터링된 신청서가 없습니다.</div>\n  </div>\n</div>',
   service_status = Number(JSON.parse(COMM.getCookie("institution_info")).service_status);
function onClickApplication(e, a) {
   switch (e) {
      case 0:
         switch ((g_AppInfo.initWithAppObj(supplement_saved_list[a]) || alert("Fail to ApplicationInfo Init!"), supplement_saved_list[a].judge_type)) {
            case IPSAP.JUDGE_TYPE.IACUC:
               navigateIACUC(e, a);
               break;
            case IPSAP.JUDGE_TYPE.IBC:
               navigateIBC(e, a);
               break;
            default:
               return void alert(`JUDGE_TYPE Error!! (${supplement_saved_list[a].judge_type})`);
         }
         break;
      case 1:
         switch ((g_AppInfo.initWithAppObj(checking_saved_list[a]) || alert("Fail to ApplicationInfo Init!"), checking_saved_list[a].judge_type)) {
            case IPSAP.JUDGE_TYPE.IACUC:
               navigateIACUC(e, a);
               break;
            case IPSAP.JUDGE_TYPE.IBC:
               navigateIBC(e, a);
               break;
            default:
               return void alert(`JUDGE_TYPE Error!! (${checking_saved_list[a].judge_type})`);
         }
         break;
      case 2:
         switch ((g_AppInfo.initWithAppObj(judge_ing_saved_list[a]) || alert("Fail to ApplicationInfo Init!"), judge_ing_saved_list[a].judge_type)) {
            case IPSAP.JUDGE_TYPE.IACUC:
               navigateIACUC(e, a);
               break;
            case IPSAP.JUDGE_TYPE.IBC:
               navigateIBC(e, a);
               break;
            default:
               return void alert(`JUDGE_TYPE Error!! (${judge_ing_saved_list[a].judge_type})`);
         }
         break;
      case 3:
         if (!$(".final_apps a").first().hasClass("unavailable"))
            switch ((g_AppInfo.initWithAppObj(final_saved_list[a]) || alert("Fail to ApplicationInfo Init!"), final_saved_list[a].judge_type)) {
               case IPSAP.JUDGE_TYPE.IACUC:
                  navigateIACUC(e, a);
                  break;
               case IPSAP.JUDGE_TYPE.IBC:
                  navigateIBC(e, a);
                  break;
               default:
                  return void alert(`JUDGE_TYPE Error!! (${final_saved_list[a].judge_type})`);
            }
   }
}
function navigateIACUC(e, a) {
   switch (e) {
      case 0:
         switch (supplement_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIACUC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIACUC_change(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.RENEW:
               navigateIACUC_renew(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.BRING:
               navigateIACUC_bring(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.END:
               navigateIACUC_end(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHECKLIST:
               navigateIACUC_checklist(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${supplement_saved_list[a].application_type})`);
         }
         break;
      case 1:
         switch (checking_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIACUC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIACUC_change(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.RENEW:
               navigateIACUC_renew(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.BRING:
               navigateIACUC_bring(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.END:
               navigateIACUC_end(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHECKLIST:
               navigateIACUC_checklist(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${checking_saved_list[a].application_type})`);
         }
         break;
      case 2:
         switch (judge_ing_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIACUC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIACUC_change(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.RENEW:
               navigateIACUC_renew(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.BRING:
               navigateIACUC_bring(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.END:
               navigateIACUC_end(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHECKLIST:
               navigateIACUC_checklist(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${judge_ing_saved_list[a].application_type})`);
         }
         break;
      case 3:
         switch (final_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIACUC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIACUC_change(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.RENEW:
               navigateIACUC_renew(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.BRING:
               navigateIACUC_bring(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.END:
               navigateIACUC_end(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHECKLIST:
               navigateIACUC_checklist(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${final_saved_list[a].application_type})`);
         }
   }
}
function navigateIACUC_new(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
         break;
      case 1:
         switch (parseInt(checking_saved_list[a].application_result)) {
            case IPSAP.APPLICATION_RESULT.CHECKING:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review.html");
               break;
            case IPSAP.APPLICATION_RESULT.CHECKING_2:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         }
         break;
      case 2:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review.html");
   }
}
function navigateIACUC_change(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
         break;
      case 1:
         switch (parseInt(checking_saved_list[a].application_result)) {
            case IPSAP.APPLICATION_RESULT.CHECKING:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_change.html");
               break;
            case IPSAP.APPLICATION_RESULT.CHECKING_2:
               g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         }
         break;
      case 2:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_change.html");
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_change.html");
   }
}
function navigateIACUC_renew(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         break;
      case 1:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review.html");
         break;
      case 2:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_renew.html");
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_renew.html");
   }
}
function navigateIACUC_bring(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         break;
      case 1:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_bring.html");
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_bring.html");
   }
}
function navigateIACUC_end(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-setup.html");
         break;
      case 1:
         g_AppInfo.saveParamsAndNavigate("./reviewOffice_list-review_end.html");
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./reviewConfirm_list-review_end.html");
   }
}
function navigateIACUC_checklist(e, a) {
   switch (e) {
      case 0:
      case 1:
      case 2:
      case 3:
         g_AppInfo.saveParamsAndNavigate("./inspection_list-info.html");
   }
}
function navigateIBC(e, a) {
   switch (e) {
      case 0:
         switch (supplement_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIBC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIBC_change(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${supplement_saved_list[a].application_type})`);
         }
         break;
      case 1:
         switch (checking_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIBC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIBC_change(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${supplement_saved_list[a].application_type})`);
         }
         break;
      case 2:
         switch (judge_ing_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIBC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIBC_change(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${supplement_saved_list[a].application_type})`);
         }
         break;
      case 3:
         switch (final_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIBC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIBC_change(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${supplement_saved_list[a].application_type})`);
         }
   }
}
function navigateIBC_new(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
         break;
      case 1:
         switch (parseInt(checking_saved_list[a].application_result)) {
            case IPSAP.APPLICATION_RESULT.CHECKING:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_review.html");
               break;
            case IPSAP.APPLICATION_RESULT.CHECKING_2:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
         }
         break;
      case 2:
         g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./IBC/reviewConfirm_list-IBC_review.html");
   }
}
function navigateIBC_change(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
         break;
      case 1:
         switch (parseInt(checking_saved_list[a].application_result)) {
            case IPSAP.APPLICATION_RESULT.CHECKING:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_review_change.html");
               break;
            case IPSAP.APPLICATION_RESULT.CHECKING_2:
               g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_setup.html");
         }
         break;
      case 2:
         g_AppInfo.saveParamsAndNavigate("./IBC/reviewOffice_list-IBC_review_change.html");
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./IBC/reviewConfirm_list-IBC_review_change.html");
   }
}
function originalDashboardFunction() {
   setDashboardFilters();
}
$(function () {
   if (IPSAP.DEMO_MODE) return $(".dashboard_box").removeClass("hidden"), void originalDashboardFunction();
   API.load({
      url: CONST.API.DASHBOARD.LIST,
      type: CONST.API_TYPE.GET,
      data: { "filter.dashboard_type": 1 },
      success: function (e) {
         $(".supplement_cnt").text(e.supplement_cnt), $(".checking_cnt").text(e.checking_cnt), $(".judge_ing_cnt").text(e.judge_ing_cnt), $(".final_cnt").text(e.final_cnt);
         let a = "등록된 신청서",
            t = $(".supplement_apps_body");
         if ((t.empty(), e.supplement.length))
            $.each(e.supplement, function (a, i) {
               supplement_saved_list = e.supplement;
               var s = "",
                  n = "";
               switch (i.application_type) {
                  case IPSAP.APPLICATION_TYPE.NEW:
                  case IPSAP.APPLICATION_TYPE.CHANGE:
                  case IPSAP.APPLICATION_TYPE.RENEW:
                     (s = "application"), (n = '<i class="mdi mdi-file-upload-outline"></i>'), "label gray_bg";
                     break;
                  case IPSAP.APPLICATION_TYPE.BRING:
                  case IPSAP.APPLICATION_TYPE.END:
                  case IPSAP.APPLICATION_TYPE.CHECKLIST:
                     (s = "report"), (n = '<i class="mdi mdi-finance"></i>'), "label gray_border_inner";
               }
               var r = "";
               switch (i.application_result) {
                  case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
                     r = "supplement";
                     break;
                  case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
                     r = "delayed";
               }
               switch (i.judge_type) {
                  case IPSAP.JUDGE_TYPE.IACUC:
                     iacuc_num++;
                     break;
                  case IPSAP.JUDGE_TYPE.IBC:
                     ibc_num++;
                     break;
                  case IPSAP.JUDGE_TYPE.IRB:
                     irb_num++;
               }
               var c = cardObjTmp.clone();
               c.addClass(`${i.judge_type_str}`).addClass(`${s}`).addClass(`${r}`),
                  c.find(".card_icon").html(n),
                  c.find(".committee_bg").text(`${i.application_type_str}`),
                  c.find(".committee_color").text(`${i.judge_type_str}`),
                  c.find(".card-title").text(`${i.name_ko}`),
                  c.find(".rcv_num").text(`${i.application_no}`),
                  c.find(".title_date").text(`(${i.tmp_submit_dttm})`),
                  c.find(".bullet").text(`${i.time_diff}`),
                  c.find(".link_btn").addClass("btn btn-outline-danger btn_xxs").attr("onclick", 'alert("신청자가 보완중인 문서입니다.")').text("보완중"),
                  t.append(c);
            });
         else {
            let e = t.data("comment") ? t.data("comment") : a;
            t.closest(".myApps").addClass("empty"), t.append(empty_apps).find(".empty_comment").text(e);
         }
         let i = $(".checking_apps_body");
         if ((i.empty(), e.checking.length))
            $.each(e.checking, function (a, t) {
               checking_saved_list = e.checking;
               var s = "",
                  n = "",
                  r = `onClickApplication(1, ${a})`;
               switch (t.application_type) {
                  case IPSAP.APPLICATION_TYPE.NEW:
                  case IPSAP.APPLICATION_TYPE.CHANGE:
                  case IPSAP.APPLICATION_TYPE.RENEW:
                     (s = "application"), (n = '<i class="mdi mdi-file-upload-outline"></i>'), "label gray_bg";
                     break;
                  case IPSAP.APPLICATION_TYPE.BRING:
                  case IPSAP.APPLICATION_TYPE.END:
                  case IPSAP.APPLICATION_TYPE.CHECKLIST:
                     (s = "report"), (n = '<i class="mdi mdi-finance"></i>'), "label gray_border_inner";
               }
               var c = "";
               switch (t.application_result) {
                  case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
                     c = "supplement";
                     break;
                  case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
                     c = "delayed";
               }
               switch (t.judge_type) {
                  case IPSAP.JUDGE_TYPE.IACUC:
                     iacuc_num++;
                     break;
                  case IPSAP.JUDGE_TYPE.IBC:
                     ibc_num++;
                     break;
                  case IPSAP.JUDGE_TYPE.IRB:
                     irb_num++;
               }
               var _ = cardObjTmp.clone();
               _.addClass(`${t.judge_type_str}`).addClass(`${s}`).addClass(`${c}`),
                  _.find(".card_icon").html(n),
                  _.find(".committee_bg").text(`${t.application_type_str}`),
                  _.find(".committee_color").text(`${t.judge_type_str}`),
                  _.find(".card-title").text(`${t.name_ko}`),
                  _.find(".rcv_num").text(`${t.application_no}`),
                  _.find(".title_date").text(`(${t.tmp_submit_dttm})`),
                  _.find(".bullet").text(`${t.time_diff}`),
                  _.find(".link_btn").text("행정 검토"),
                  _.find(".link_btn").addClass("btn btn-primary btn_xxs"),
                  _.find(".link_btn").attr("onclick", `${r}`),
                  i.append(_);
            });
         else {
            let e = i.data("comment") ? i.data("comment") : a;
            i.closest(".myApps").addClass("empty"), i.append(empty_apps).find(".empty_comment").text(e);
         }
         let s = $(".judge_ing_apps_body");
         if ((s.empty(), e.judge_ing.length))
            $.each(e.judge_ing, function (a, t) {
               judge_ing_saved_list = e.judge_ing;
               var i = "",
                  n = "",
                  r = `onClickApplication(2, ${a})`;
               switch (t.application_type) {
                  case IPSAP.APPLICATION_TYPE.NEW:
                  case IPSAP.APPLICATION_TYPE.CHANGE:
                  case IPSAP.APPLICATION_TYPE.RENEW:
                     (i = "application"), (n = '<i class="mdi mdi-file-upload-outline"></i>'), "label gray_bg";
                     break;
                  case IPSAP.APPLICATION_TYPE.BRING:
                  case IPSAP.APPLICATION_TYPE.END:
                  case IPSAP.APPLICATION_TYPE.CHECKLIST:
                     (i = "report"), (n = '<i class="mdi mdi-finance"></i>'), "label gray_border_inner";
               }
               var c = "";
               switch (t.application_result) {
                  case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
                     c = "supplement";
                     break;
                  case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
                     c = "delayed";
               }
               switch (t.judge_type) {
                  case IPSAP.JUDGE_TYPE.IACUC:
                     iacuc_num++;
                     break;
                  case IPSAP.JUDGE_TYPE.IBC:
                     ibc_num++;
                     break;
                  case IPSAP.JUDGE_TYPE.IRB:
                     irb_num++;
               }
               var _ = cardObjTmp.clone();
               _.addClass(`${t.judge_type_str}`).addClass(`${i}`).addClass(`${c}`),
                  _.find(".card_icon").html(n),
                  _.find(".committee_bg").text(`${t.application_type_str}`),
                  _.find(".committee_color").text(`${t.judge_type_str}`),
                  _.find(".card-title").text(`${t.name_ko}`),
                  _.find(".rcv_num").text(`${t.application_no}`),
                  _.find(".title_date").text(`(${t.tmp_submit_dttm})`),
                  "0일 지연" == t.time_diff ? _.find(".bullet").text("당일 지연") : _.find(".bullet").text(`${t.time_diff}`),
                  _.find(".link_btn").text("심사 설정 변경"),
                  6 == t.application_result ? _.find(".link_btn").addClass("btn btn-danger btn_xxs") : _.find(".link_btn").addClass("btn btn-primary btn_xxs"),
                  _.find(".link_btn").attr("onclick", `${r}`),
                  s.append(_);
            });
         else {
            let e = s.data("comment") ? s.data("comment") : a;
            s.closest(".myApps").addClass("empty"), s.append(empty_apps).find(".empty_comment").text(e);
         }
         let n = $(".final_apps_body");
         if ((n.empty(), e.final.length))
            $.each(e.final, function (a, t) {
               final_saved_list = e.final;
               var i = "",
                  s = "",
                  r = `onClickApplication(3, ${a})`;
               switch (t.application_type) {
                  case IPSAP.APPLICATION_TYPE.NEW:
                  case IPSAP.APPLICATION_TYPE.CHANGE:
                  case IPSAP.APPLICATION_TYPE.RENEW:
                     (i = "application"), (s = '<i class="mdi mdi-file-upload-outline"></i>'), "label gray_bg";
                     break;
                  case IPSAP.APPLICATION_TYPE.BRING:
                  case IPSAP.APPLICATION_TYPE.END:
                  case IPSAP.APPLICATION_TYPE.CHECKLIST:
                     (i = "report"), (s = '<i class="mdi mdi-finance"></i>'), "label gray_border_inner";
               }
               var c = "";
               switch (t.application_result) {
                  case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
                     c = "supplement";
                     break;
                  case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
                     c = "delayed";
               }
               switch (t.judge_type) {
                  case IPSAP.JUDGE_TYPE.IACUC:
                     iacuc_num++;
                     break;
                  case IPSAP.JUDGE_TYPE.IBC:
                     ibc_num++;
                     break;
                  case IPSAP.JUDGE_TYPE.IRB:
                     irb_num++;
               }
               var _ = cardObjTmp.clone();
               _.addClass(`${t.judge_type_str}`).addClass(`${i}`).addClass(`${c}`),
                  _.find(".card_icon").html(s),
                  _.find(".committee_bg").text(`${t.application_type_str}`),
                  _.find(".committee_color").text(`${t.judge_type_str}`),
                  _.find(".card-title").text(`${t.name_ko}`),
                  _.find(".rcv_num").text(`${t.application_no}`),
                  _.find(".title_date").text(`(${t.tmp_submit_dttm})`),
                  _.find(".bullet").text(`${t.time_diff}`),
                  _.find(".link_btn").text("최종 심의"),
                  _.find(".link_btn").addClass("btn btn-primary btn_xxs"),
                  _.find(".link_btn").attr("onclick", `${r}`),
                  n.append(_);
            });
         else {
            let e = n.data("comment") ? n.data("comment") : a;
            n.closest(".myApps").addClass("empty"), n.append(empty_apps).find(".empty_comment").text(e);
         }
         JSON.parse(COMM.getCookie("user_info")).user_type[1] ||
            $(".final_apps a")
               .addClass("unavailable")
               .off()
               .on("click", (e) => (alert("최종 심의는 행정 간사 권한을 가진 사용자만 수행할 수 있습니다."), e.preventDefault(), e.stopImmediatePropagation(), !1)),
            -1 == $.inArray("1", judge_type)
               ? ($(".btn.IACUC").css("display", "none"), $("#statistics").find(".IACUC").css("display", "none"))
               : -1 == $.inArray("2", judge_type)
               ? ($(".btn.IBC").css("display", "none"), $("#statistics").find(".IBC").css("display", "none"))
               : -1 == $.inArray("3", judge_type) && ($(".btn.IRB").css("display", "none"), $("#statistics").find(".IRB").css("display", "none")),
            $(".IACUC_num").text(iacuc_num),
            $(".IBC_num").text(ibc_num),
            $(".IRB_num").text(irb_num),
            $("#statistics")
               .find(".IACUC")
               .html(
                  `<td class="committee_color">IACUC</td>\n      <td data-num="${e.approved_statistics.iacuc.total_cnt}" class="border_left_thick">${e.approved_statistics.iacuc.total_cnt}</td>\n      <td data-num="${e.approved_statistics.iacuc.approved_cnt}" class="border_left_thick">${e.approved_statistics.iacuc.approved_cnt}</td>\n      <td data-num="${e.approved_statistics.iacuc.conditional_approved_cnt}" class="border_left">${e.approved_statistics.iacuc.conditional_approved_cnt}</td>\n      <td data-num="${e.approved_statistics.iacuc.require_retry_cnt}" class="border_left">${e.approved_statistics.iacuc.require_retry_cnt}</td>\n      <td data-num="${e.approved_statistics.iacuc.reject_cnt}" class="border_left">${e.approved_statistics.iacuc.reject_cnt}</td>`
               ),
            $("#statistics")
               .find(".IBC")
               .html(
                  `<td class="committee_color">IBC</td>\n      <td data-num="${e.approved_statistics.ibc.total_cnt}" class="border_left_thick">${e.approved_statistics.ibc.total_cnt}</td>\n      <td data-num="${e.approved_statistics.ibc.approved_cnt}" class="border_left_thick">${e.approved_statistics.ibc.approved_cnt}</td>\n      <td data-num="${e.approved_statistics.ibc.conditional_approved_cnt}" class="border_left">${e.approved_statistics.ibc.conditional_approved_cnt}</td>\n      <td data-num="${e.approved_statistics.ibc.require_retry_cnt}" class="border_left">${e.approved_statistics.ibc.require_retry_cnt}</td>\n      <td data-num="${e.approved_statistics.ibc.reject_cnt}" class="border_left">${e.approved_statistics.ibc.reject_cnt}</td>`
               ),
            $("#statistics")
               .find(".IRB")
               .html(
                  `<td class="committee_color">IRB</td>\n      <td data-num="${e.approved_statistics.irb.total_cnt}" class="border_left_thick">${e.approved_statistics.irb.total_cnt}</td>\n      <td data-num="${e.approved_statistics.irb.approved_cnt}" class="border_left_thick">${e.approved_statistics.irb.approved_cnt}</td>\n      <td data-num="${e.approved_statistics.irb.conditional_approved_cnt}" class="border_left">${e.approved_statistics.irb.conditional_approved_cnt}</td>\n      <td data-num="${e.approved_statistics.irb.require_retry_cnt}" class="border_left">${e.approved_statistics.irb.require_retry_cnt}</td>\n      <td data-num="${e.approved_statistics.irb.reject_cnt}" class="border_left">${e.approved_statistics.irb.reject_cnt}</td>`
               ),
            $(".myExps").empty(),
            e.performance_app_list.length
               ? $.each(e.performance_app_list, function (e, a) {
                    var t = explistObjTmp.clone();
                    t.addClass(a.judge_type_str),
                       t.find(".committee_bg").text(a.judge_type_str),
                       t.find(".card-title").text(a.name_ko),
                       t.find(".rcv_num").text(a.application_no),
                       t.find(".title_date").text(`(${getDttm(a.approved_dttm).dt} ~ ${a.general_end_date})`),
                       t.find(".end_date").html(`${a.general_end_date}<i class="mdi mdi-arrow-collapse-right mLeft5"></i>`);
                    let i = 1 == a.end_date_over ? "closed" : "";
                    t.addClass(i),
                       "closed" == i &&
                          t.addClass("tippy-btn").attr({ "data-tippy-placement": "top-start", "data-tippy-arrow": "true", title: '실험 기간이 만료되었습니다. "종료 보고서" 제출을 확인해 주세요.' });
                    let s = getYMDHMSFromUnixtime(a.approved_dttm),
                       n = getYMDHMSFromDateType(getDateByStr(a.general_end_date)),
                       r = new Date(),
                       c = getDateDiffWith10(r, s),
                       _ = getDateDiffWith10(n, s);
                    0 == _ && (_ = c);
                    let d = 335;
                    c < d
                       ? t
                            .find(".exp_period")
                            .html(
                               `<div class="progress mTop5">\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  (c / _) * 100
                               )}%">1차 년도</div>\n            </div>`
                            )
                       : c > d && c < 670
                       ? t
                            .find(".exp_period")
                            .html(
                               `<div class="progress mTop5">\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  (d / _) * 100
                               )}%">1차 년도</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" style="width: 5%">재승인</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  ((c - d) / _) * 100
                               )}%">2차 년도</div>\n            </div>`
                            )
                       : c > 670 &&
                         t
                            .find(".exp_period")
                            .html(
                               `<div class="progress mTop5">\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  (d / _) * 100
                               )}%">1차 년도</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" style="width: 5%">재승인</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  (d / _) * 100
                               )}%">2차 년도</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" style="width: 5%">재승인</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  ((c - 670) / _) * 100
                               )}%">3차 년도</div>\n            </div>`
                            ),
                       $(".myExps").append(t);
                 })
               : $(".myExps").addClass("empty").append(empty_exps),
            $(".dashboard_box").removeClass("hidden"),
            updateStackCounter($(".myApps"), ".card", ".myApp_cnt", !0),
            setTimeout(() => {
               stackCards($(".myApps, .exp_area"));
            }, 300);
      },
      complete: function () {
         1 != service_status && ($(".dashboard_box").find("*").removeAttr("href"), $(".dashboard_box").find("*").removeAttr("onclick")), originalDashboardFunction();
      },
   }),
      $("#btn_download").click(function (e) {
         var a = JSON.parse(COMM.getCookie(CONST.COOKIE.IPSAP_USER_INFO));
         location.href = CONST.API_PATH + CONST.API.DASHBOARD.DOWNLOAD + "?filter.institution_seq=" + a.institution_seq;
      });
});
