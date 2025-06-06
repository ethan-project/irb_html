var writing_saved_list, checking_saved_list, judge_ing_saved_list, final_saved_list;
"undefined" == typeof ipsap_common_js && document.write("<script src='/assets/js/ipsap/ipsap_common.js'></script>"),
   "undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof date_utils_js && document.write("<script src='/assets/js/common/util/date_utils.js'></script>"),
   "undefined" == typeof ipsap_application_common_js && document.write("<script src='/assets/js/ipsap/ipsap_application_common.js'></script>");
var cardObjTmp = $(".myApp_body .card").eq(0).clone(),
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
         switch ((g_AppInfo.initWithAppObj(writing_saved_list[a]) || alert("Fail to ApplicationInfo Init!"), writing_saved_list[a].judge_type)) {
            case IPSAP.JUDGE_TYPE.IACUC:
               navigateIACUC(e, a);
               break;
            case IPSAP.JUDGE_TYPE.IBC:
               navigateIBC(e, a);
               break;
            default:
               return void alert(`JUDGE_TYPE Error!! (${writing_saved_list[a].judge_type})`);
         }
         break;
      case 1:
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
         switch (writing_saved_list[a].application_type) {
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
               return void alert(`application_type Error!! (${writing_saved_list[a].application_type})`);
         }
         break;
      case 1:
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
         switch (writing_saved_list[a].application_result) {
            case 0:
               g_AppInfo.saveParamsAndNavigate(APP_NAVIGATION.IACUC.PAGE_INFO[0].URL);
               break;
            case 1:
               g_AppInfo.saveParamsAndNavigate("./application_list-review.html");
         }
         break;
      case 1:
         break;
      case 2:
      case 3:
         g_AppInfo.saveParamsAndNavigate("./application_list-info.html");
   }
}
function navigateIACUC_change(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-review_change.html");
         break;
      case 1:
         break;
      case 2:
      case 3:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-info_change.html");
   }
}
function navigateIACUC_renew(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-review_renew.html");
         break;
      case 1:
         break;
      case 2:
      case 3:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-info_renew.html");
   }
}
function navigateIACUC_bring(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-review_bring.html");
         break;
      case 1:
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-info_bring.html");
   }
}
function navigateIACUC_end(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-review_end.html");
         break;
      case 1:
         break;
      case 3:
         g_AppInfo.saveParamsAndNavigate("./experiment_list-info_end.html");
   }
}
function navigateIACUC_checklist(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./inspection_list-review.html");
         break;
      case 1:
         break;
      case 2:
      case 3:
         g_AppInfo.saveParamsAndNavigate("./inspection_list-info.html");
   }
}
function navigateIBC(e, a) {
   switch (e) {
      case 0:
         switch (writing_saved_list[a].application_type) {
            case IPSAP.APPLICATION_TYPE.NEW:
               navigateIBC_new(e, a);
               break;
            case IPSAP.APPLICATION_TYPE.CHANGE:
               navigateIBC_change(e, a);
               break;
            default:
               return void alert(`application_type Error!! (${writing_saved_list[a].application_type})`);
         }
         break;
      case 1:
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
               return void alert(`application_type Error!! (${judge_ing_saved_list[a].application_type})`);
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
               return void alert(`application_type Error!! (${final_saved_list[a].application_type})`);
         }
   }
}
function navigateIBC_new(e, a) {
   switch (e) {
      case 0:
         switch (writing_saved_list[a].application_result) {
            case 0:
               g_AppInfo.saveParamsAndNavigate(APP_IBC_NAVIGATION.IBC.PAGE_INFO[0].URL);
               break;
            case 1:
               g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_review.html");
         }
         break;
      case 1:
         break;
      case 2:
      case 3:
         g_AppInfo.saveParamsAndNavigate("./IBC/application_list-IBC_info.html");
   }
}
function navigateIBC_change(e, a) {
   switch (e) {
      case 0:
         g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_review_change.html");
         break;
      case 1:
         break;
      case 2:
      case 3:
         g_AppInfo.saveParamsAndNavigate("./IBC/experiment_list-IBC_info_change.html");
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
      data: { "filter.dashboard_type": 2 },
      success: function (e) {
         $(".writing_cnt").text(e.writing_cnt), $(".checking_cnt").text(e.checking_cnt), $(".judge_ing_cnt").text(e.judge_ing_cnt), $(".final_cnt").text(e.final_cnt);
         let a = "등록된 신청서",
            t = $(".writing_apps_body");
         if ((t.empty(), e.writing.length))
            $.each(e.writing, function (a, i) {
               writing_saved_list = e.writing;
               var n = "",
                  s = "",
                  r = `onClickApplication(0 , ${a})`;
               switch (i.application_type) {
                  case IPSAP.APPLICATION_TYPE.NEW:
                  case IPSAP.APPLICATION_TYPE.CHANGE:
                  case IPSAP.APPLICATION_TYPE.RENEW:
                     (n = "application"), (s = '<i class="mdi mdi-file-upload-outline"></i>'), "label gray_bg";
                     break;
                  case IPSAP.APPLICATION_TYPE.BRING:
                  case IPSAP.APPLICATION_TYPE.END:
                  case IPSAP.APPLICATION_TYPE.CHECKLIST:
                     (n = "report"), (s = '<i class="mdi mdi-finance"></i>'), "label gray_border_inner";
               }
               var d = "";
               if (i.application_result === IPSAP.APPLICATION_RESULT.SUPPLEMENT) d = "supplement";
               var _ = cardObjTmp.clone();
               _.addClass(`${i.judge_type_str}`).addClass(`${n}`).addClass(`${d}`),
                  _.find(".card_icon").html(s),
                  _.find(".committee_bg").text(`${i.application_type_str}`),
                  _.find(".committee_color").text(`${i.judge_type_str}`),
                  _.find(".card-title").text(`${i.name_ko}`),
                  _.find(".rcv_num").text(`${i.application_no}`),
                  _.find(".title_date").text(`(${i.tmp_submit_dttm})`),
                  _.find(".bullet").text(`${i.time_diff}`),
                  0 == i.application_result
                     ? (_.find(".link_btn").text("계속 작성"), _.find(".link_btn").addClass("btn btn-primary btn_xxs"))
                     : (_.find(".link_btn").text("보완 작성"), _.find(".link_btn").addClass("btn btn-danger btn_xxs")),
                  _.find(".link_btn").attr("onclick", `${r}`),
                  t.append(_);
            });
         else {
            let e = t.data("comment") ? t.data("comment") : a;
            t.closest(".myApps").addClass("empty"), t.append(empty_apps).find(".empty_comment").text(e);
         }
         let i = $(".checking_apps_body");
         if ((i.empty(), e.checking.length))
            $.each(e.checking, function (a, t) {
               checking_saved_list = e.checking;
               var n = "",
                  s = "",
                  r = `onClickApplication(1, ${a})`;
               switch (t.application_type) {
                  case IPSAP.APPLICATION_TYPE.NEW:
                  case IPSAP.APPLICATION_TYPE.CHANGE:
                  case IPSAP.APPLICATION_TYPE.RENEW:
                     (n = "application"), (s = '<i class="mdi mdi-file-upload-outline"></i>'), "label gray_bg";
                     break;
                  case IPSAP.APPLICATION_TYPE.BRING:
                  case IPSAP.APPLICATION_TYPE.END:
                  case IPSAP.APPLICATION_TYPE.CHECKLIST:
                     (n = "report"), (s = '<i class="mdi mdi-finance"></i>'), "label gray_border_inner";
               }
               var d = "";
               if (t.application_result === IPSAP.APPLICATION_RESULT.SUPPLEMENT) d = "supplement";
               var _ = cardObjTmp.clone();
               _.addClass(`${t.judge_type_str}`).addClass(`${n}`).addClass(`${d}`),
                  _.find(".card_icon").html(s),
                  _.find(".committee_bg").text(`${t.application_type_str}`),
                  _.find(".committee_color").text(`${t.judge_type_str}`),
                  _.find(".card-title").text(`${t.name_ko}`),
                  _.find(".rcv_num").text(`${t.application_no}`),
                  _.find(".title_date").text(`(${t.tmp_submit_dttm})`),
                  _.find(".bullet").text(`${t.time_diff}`),
                  _.find(".link_btn").text("검토중"),
                  _.find(".link_btn").addClass("btn btn-outline-secondary btn_xxs disabled"),
                  _.find(".link_btn").attr("onclick", `${r}`),
                  i.append(_);
            });
         else {
            let e = i.data("comment") ? i.data("comment") : a;
            i.closest(".myApps").addClass("empty"), i.append(empty_apps).find(".empty_comment").text(e);
         }
         let n = $(".judge_ing_apps_body");
         if ((n.empty(), e.judge_ing.length))
            $.each(e.judge_ing, function (a, t) {
               judge_ing_saved_list = e.judge_ing;
               var i = "",
                  s = "",
                  r = `onClickApplication(2, ${a})`;
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
               var d = "";
               if (t.application_result === IPSAP.APPLICATION_RESULT.SUPPLEMENT) d = "supplement";
               var _ = cardObjTmp.clone();
               _.addClass(`${t.judge_type_str}`).addClass(`${i}`).addClass(`${d}`),
                  _.find(".card_icon").html(s),
                  _.find(".committee_bg").text(`${t.application_type_str}`),
                  _.find(".committee_color").text(`${t.judge_type_str}`),
                  _.find(".card-title").text(`${t.name_ko}`),
                  _.find(".rcv_num").text(`${t.application_no}`),
                  _.find(".title_date").text(`(${t.tmp_submit_dttm})`),
                  _.find(".bullet").text(`${t.time_diff}`),
                  _.find(".link_btn").text("심사중"),
                  _.find(".link_btn").addClass("btn btn-primary btn_xxs"),
                  _.find(".link_btn").attr("onclick", `${r}`),
                  n.append(_);
            });
         else {
            let e = n.data("comment") ? n.data("comment") : a;
            n.closest(".myApps").addClass("empty"), n.append(empty_apps).find(".empty_comment").text(e);
         }
         let s = $(".final_apps_body");
         if ((s.empty(), e.final.length))
            $.each(e.final, function (a, t) {
               final_saved_list = e.final;
               var i = "",
                  n = "",
                  r = `onClickApplication(3, ${a})`;
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
               var d = "";
               if (t.application_result === IPSAP.APPLICATION_RESULT.SUPPLEMENT) d = "supplement";
               var _ = cardObjTmp.clone();
               _.addClass(`${t.judge_type_str}`).addClass(`${i}`).addClass(`${d}`),
                  _.find(".card_icon").html(n),
                  _.find(".committee_bg").text(`${t.application_type_str}`),
                  _.find(".committee_color").text(`${t.judge_type_str}`),
                  _.find(".card-title").text(`${t.name_ko}`),
                  _.find(".rcv_num").text(`${t.application_no}`),
                  _.find(".title_date").text(`(${t.tmp_submit_dttm})`),
                  _.find(".bullet").text(`${t.time_diff}`),
                  _.find(".link_btn").text("심의중"),
                  _.find(".link_btn").addClass("btn btn-primary btn_xxs"),
                  _.find(".link_btn").attr("onclick", `${r}`),
                  s.append(_);
            });
         else {
            let e = s.data("comment") ? s.data("comment") : a;
            s.closest(".myApps").addClass("empty"), s.append(empty_apps).find(".empty_comment").text(e);
         }
         -1 == $.inArray("1", judge_type)
            ? $(".btn.IACUC").css("display", "none")
            : -1 == $.inArray("2", judge_type)
            ? $(".btn.IBC").css("display", "none")
            : -1 == $.inArray("3", judge_type) && $(".btn.IRB").css("display", "none"),
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
                          t.addClass("tippy-btn").attr({ "data-tippy-placement": "top-start", "data-tippy-arrow": "true", title: '실험 기간이 만료되었습니다. "종료 보고서"를 제출해 주세요.' });
                    let n = getYMDHMSFromUnixtime(a.approved_dttm),
                       s = getYMDHMSFromDateType(getDateByStr(a.general_end_date)),
                       r = new Date(),
                       d = getDateDiffWith10(r, n),
                       _ = getDateDiffWith10(s, n);
                    0 == _ && (_ = d);
                    let c = 335;
                    d < c
                       ? t
                            .find(".exp_period")
                            .html(
                               `<div class="progress mTop5">\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  (d / _) * 100
                               )}%">1차 년도</div>\n            </div>`
                            )
                       : d > c && d < 670
                       ? t
                            .find(".exp_period")
                            .html(
                               `<div class="progress mTop5">\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  (c / _) * 100
                               )}%">1차 년도</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" style="width: 5%">재승인</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  ((d - c) / _) * 100
                               )}%">2차 년도</div>\n            </div>`
                            )
                       : d > 670 &&
                         t
                            .find(".exp_period")
                            .html(
                               `<div class="progress mTop5">\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  (c / _) * 100
                               )}%">1차 년도</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" style="width: 5%">재승인</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  (c / _) * 100
                               )}%">2차 년도</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" style="width: 5%">재승인</div>\n              <div class="progress-bar progress-bar-striped progress-bar-animated" style="width: ${Math.floor(
                                  ((d - 670) / _) * 100
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
         location.href = CONST.API_PATH + CONST.API.DASHBOARD.DOWNLOAD + "?filter.reg_user_seq=" + a.user_seq + "&filter.institution_seq=" + a.institution_seq;
      });
});
