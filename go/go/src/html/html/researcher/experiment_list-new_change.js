"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   "undefined" == typeof ipsap_application_common_js && document.write("<script src='/assets/js/ipsap/ipsap_application_common.js'></script>"),
   "undefined" == typeof date_utils_js && document.write("<script src='/assets/js/common/util/date_utils.js'></script>");
var g_all_userlist = [];
let g_old_member_app_seq = 0,
   g_old_end_date_app_seq = 0;
function originalJSFunction() {
   var e = !0;
   $(".modal.draggable>.modal-dialog").draggable({
      drag: function (e, t) {
         t.position.top <= 0 && (t.position.top = 0), t.position.top >= window.innerHeight - 80 && (t.position.top = window.innerHeight - 80);
      },
      cursor: "move",
      handle: ".modal-header",
   }),
      $(".modal.draggable>.modal-dialog>.modal-content>.modal-header").css("cursor", "move"),
      $(".modal.draggable>.modal-dialog").on({
         "click drag": function (t) {
            // height
            var screenHeight = $(window).height();
            console.log("screenHeight", screenHeight)
            var modalHeight = $('#modal_changeApp .modal-content').outerHeight();
            console.log("modalHeight", modalHeight)
            //margin
            var marginTop = parseInt($('.modal-dialog').css('margin-top'), 10);
            var marginBottom = parseInt($('.modal-dialog').css('margin-bottom'), 10);
            console.log('Top Margin:', marginTop, 'px');
            console.log('Bottom Margin:', marginBottom, 'px');
            //calcu 
            var topPosition = screenHeight - modalHeight - marginTop - marginBottom;
            console.log("topPosition", topPosition)
            $(".modal_hint").hide(),
               e &&
                  ($(".modal[data-backdrop=false]").css({ "background-color": "rgba(0,0,0, 0.0)", "pointer-events": "none" }),
                  $(".modal-dialog-centered.ui-draggable").css({ "justify-content": "flex-start" }),
                  "click" == t.type && $(this).css({ "margin-right": "0", left: "0", top: topPosition +"px" })),
               (e = !1);
         },
      }),
      $(".modal.fade").each(function (e, t) {
         $(t).hasClass("in") && $(t).modal("show"), "modal_changeApp" == $(t).attr("id") && $("body").css({ overflow: "auto" });
      }),
      $("input[id^=regular_], input[id^=simple_]").on({
         change: function (e) {
            $(this).prop("checked")
               ? ($(this).closest("li").addClass("active"), $($(this).data("target")).addClass("active"))
               : ($(this).closest("li").removeClass("active"), $($(this).data("target")).removeClass("active"));
         },
      }),
      $("select.move-to-selected").html(
         '\n    <option selected disabled>실험 계획서 항목</option>\n    <optgroup label="1. 일반정보">\n      <option value="#form_01-1">1-1. 동물실험 과제명</option>\n      <option value="#form_01-2">1-2. 연구기간 및 동물실험 횟수</option>\n      <option value="#form_01-3">1-3. 연구비 지원 기관</option>\n      <option value="#form_01-4">1-4. 연구 책임자</option>\n      <option value="#form_01-5">1-5. 실험 수행자</option>\n      <option value="#form_01-6">1-6. 관련 자료 첨부</option>\n    </optgroup>\n    <optgroup label="2. 동물실험 목적">\n      <option value="#form_02-1">2-1. 동물실험 목적에 따른 분류 Ⅰ</option>\n      <option value="#form_02-2">2-2. 동물실험 목적에 따른 분류 Ⅱ</option>\n      <option value="#form_02-3">2-3. 동물실험원칙(3Rs)에 다른 대안방법 모색</option>\n      <option value="#form_02-4">2-4. 검색어 (Key Words)</option>\n      <option value="#form_02-5">2-5. 정보 검색일</option>\n      <option value="#form_02-6">2-6. 대안 방법 검토결과</option>\n    </optgroup>\n    <optgroup label="3. 실험동물 정보">\n      <option value="#form_03-1">3-1. 실험 동물의 종류</option>\n      <option value="#form_03-2">3-2. 해당 동물 종(Species)과 계통(Strain)을 선택한 합리적인 이유</option>\n      <option value="#form_03-3">3-3. 사용 동물 수에 대한 합리적인 근거</option>\n    </optgroup>\n    <optgroup label="4. 실험동물 정보">\n      <option value="#form_04-1">4-1. 실험물질 투여 유무</option>\n    </optgroup>\n    <optgroup label="5. 동물실험 방법">\n      <option value="#form_05-1">5-1. 동물 실험의 개요 및 일정</option>\n      <option value="#form_05-2">5-2. 동물 실험의 형태</option>\n      <option value="#form_05-3">5-3. 보정법</option>\n      <option value="#form_05-4">5-4. 식별법</option>\n      <option value="#form_05-5">5-5. 외과적 처치</option>\n      <option value="#form_05-6">5-6. 동물 반출</option>\n      <option value="#form_05-7">5-7. 시료 채취</option>\n      <option value="#form_05-8">5-8. 실험 물질</option>\n    </optgroup>\n    <optgroup label="6. 실험동물 정보">\n      <option value="#form_06-1">6-1. 고통 등급 분류</option>\n      <option value="#form_06-2">6-2. 고통등급 D (실험 동물의 고통 경감 방안)</option>\n      <option value="#form_06-3">6-3. 고통등급 E (동물실험을 수행하는 사유 및 관리 방안)</option>\n    </optgroup>\n    <optgroup label="7. 고통경감 방안">\n      <option value="#form_07-1">7-1. 고통 및 스트레스에 대한 평가 방법</option>\n      <option value="#form_07-2">7-2. 인도적인 종료시점을 적용할 수 없는 사유</option>\n      <option value="#form_07-3">7-3. 마약, 향정신성 의약품 사용 유무</option>\n      <option value="#form_07-4">7-4. 동물 의약품 사용 유무</option>\n      <option value="#form_07-5">7-5. 고통경감을 위한 수의학적 관리</option>\n    </optgroup>\n    <optgroup label="8. 사육관리">\n      <option value="#form_08-1">8-1. 특별한 주거 및 사육 조건</option>\n      <option value="#form_08-2">8-2. 사육 환경</option>\n      <option value="#form_08-3">8-3. 풍부화 도구</option>\n      <option value="#form_08-4">8-4. 풍부화 불가능 사유</option>\n    </optgroup>\n    <optgroup label="9. 안락사">\n      <option value="#form_09-1">9-1. 안락사 방법</option>\n      <option value="#form_09-2">9-2. 사체 처리 방법</option>\n      <option value="#form_09-3">9-3. 실험동물 유래자원 공유</option>\n    </optgroup>\n    <optgroup label="10. 연구책임자 준수사항">\n      <option value="#form_10-1">10-1. 작업 환경 및 실험도구의 안전성 관리</option>\n      <option value="#form_10-2">10-2. 준수 사항</option>\n    </optgroup>\n    <optgroup label="기타">\n      <option value="#">해당 사항 없음</option>\n    </optgroup>\n  '
      ),
      $("select.move-to-selected").each(function (e, t) {
         $(t)
            .find('option[value="#' + $(t).data("selected") + '"]')
            .attr("selected", !0);
      }),
      $("select.move-to-selected")
         .closest(".list-group-item")
         .on({
            click: function (e) {
               let t = $(this).find("select").val();
               t && "#" != t && $("html, body").animate({ scrollTop: $(t).offset().top }, 300);
            },
         }),
      $(".btn_reset_select").on({
         click: function (e) {
            $(this).prev("select").find("option[selected=selected]").prop("selected", !0);
         },
      }),
      $("#review_summary").on({
         click: function (e) {
            var t = $(this),
               a = t.closest(".modal");
            t.toggleClass("active"), t.hasClass("active") ? t.text("변경항목 수정") : t.text("변경항목 모아보기"), a.find("input, textarea").prop("readonly", !0);
            let n = $(".list-group-item");
            t.hasClass("active")
               ? (a.addClass("view_mode"), n.find("select").prop("disabled", !0), n.filter(":not(.active)").closest("ul").hide())
               : (a.removeClass("view_mode"),
                 a.find("input, textarea").prop("readonly", !1).removeClass("view_mode"),
                 n.find("select").prop("disabled", !1),
                 n.filter(":not(.active)").closest("ul").removeAttr("style"));
         },
      }),
      scrollSpy(),
      $(window).on({
         load: function (e) {
            $(".supplement_comment").length &&
               ($(".modal-body").animate({ scrollTop: $(document).height() }, 500),
               setTimeout(function () {
                  $("#review_summary").trigger("click");
               }, 300));
         },
         scroll: function () {
            scrollSpy();
         },
      });
}
function cbFuncAfterMapping(e) {
   mappingChangeAppInfo(e), originalJSFunction();
}
$(function () {
   g_AppInfo.loadParams(),
      $("#all_app_r").load("/html/common/inc_application/all_app_readonly.html", function () {
         if (IPSAP.DEMO_MODE) originalJSFunction();
         else {
            switch (
               ((function () {
                  let e = CONST.API.INSTITUTION.USER;
                  API.load({
                     url: e.replace("${inst_seq}", g_AppInfo.appObj.institution_seq),
                     data: { "filter.user_type": IPSAP.USER_TYPE.RESEARCHER },
                     type: CONST.API_TYPE.GET,
                     success: function (e) {
                        $.each(e, function (e, t) {
                           var a = { animal_mng_flag: 0, exp_year_code: 0, user_seq: t.user_seq, info: t };
                           g_all_userlist.push(a);
                        });
                     },
                  });
               })(),
               remakeApplicationInfoReadOnly(["change_application"], cbFuncAfterMapping),
               $(".apptype_title").text(g_AppInfo.appObj.application_type_str),
               parseInt(g_AppInfo.appObj.application_result))
            ) {
               case IPSAP.APPLICATION_RESULT.SUPPLEMENT:
                  $(".apptype_date").text("보완중");
                  break;
               case IPSAP.APPLICATION_RESULT.CHECKING:
               case IPSAP.APPLICATION_RESULT.CHECKING_2:
                  $(".apptype_date").text("검토중");
                  break;
               case IPSAP.APPLICATION_RESULT.JUDGE_ING:
               case IPSAP.APPLICATION_RESULT.JUDGE_ING_2:
                  $(".apptype_date").text("심사중");
                  break;
               case IPSAP.APPLICATION_RESULT.DECISION_ING:
                  $(".apptype_date").text("심의중");
                  break;
               case IPSAP.APPLICATION_RESULT.REJECT:
                  $(".apptype_area").addClass("denied"), $(".apptype_date").text("반려");
                  break;
               case IPSAP.APPLICATION_RESULT.REQUIRE_RETRY:
                  $(".apptype_area").addClass("denied"), $(".apptype_date").text("보완후 재심");
                  break;
               case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
                  $(".apptype_date").text("실험 중 (승인)");
                  break;
               case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
                  $(".apptype_date").text("실험 중 (조건부 승인)");
                  break;
               case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
                  $(".apptype_date").text("실험 만료");
                  break;
               case IPSAP.APPLICATION_RESULT.TASK_FINISH:
                  $(".apptype_date").text("실험 종료");
                  break;
               case IPSAP.APPLICATION_RESULT.ACCEPT:
                  $(".apptype_date").text("승인 : " + getDttm(g_AppInfo.appObj.approved_dttm).dt);
                  break;
               case IPSAP.APPLICATION_RESULT.ACCEPT_AC:
                  $(".apptype_date").text("조건부 승인 : " + getDttm(g_AppInfo.appObj.approved_dttm).dt);
            }
            switch (parseInt(g_AppInfo.appObj.application_step)) {
               case 0:
               default:
                  $(".progress_area").children().removeClass("active"), $(".progress_area").children().eq(0).addClass("active");
                  break;
               case 1:
                  $(".progress_area").children().removeClass("active"), $(".progress_area").children().eq(1).addClass("active");
                  break;
               case 2:
                  $(".progress_area").children().removeClass("active"), $(".progress_area").children().eq(2).addClass("active");
                  break;
               case 3:
                  $(".progress_area").children().removeClass("active"), $(".progress_area").children().eq(3).addClass("active");
                  break;
               case 4:
                  $(".progress_area").children().removeClass("active"), $(".progress_area").children().eq(4).addClass("active");
                  break;
               case 5:
                  $(".progress_area").children().removeClass("active"), $(".progress_area").children().eq(5).addClass("active");
            }
            $(".task_title_ko").text(g_AppInfo.appObj.name_ko),
               $(".task_title_en").text(g_AppInfo.appObj.name_en),
               $(".task_director_all").text(`(연구책임자 : ${g_AppInfo.appObj.user_name})`),
               $(".task_director").text(g_AppInfo.appObj.user_name),
               $(".task_ditrect_dept").text(g_AppInfo.appObj.user_dept_str),
               $(".task_director_contact").text(g_AppInfo.appObj.user_phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
               $(".judge_type_str").text(g_AppInfo.appObj.judge_type_str),
               $(".today_str").text(getTimeStrFromDateType(new Date()));
            let e = getYMDHMSFromDateType(getDateByStr(g_AppInfo.appObj.approved_dttm)),
               t = getYMDHMSFromDateType(getDateByStr(g_AppInfo.appObj.general_end_date));
            0 == g_AppInfo.appObj.approved_dttm ? $(".task_dayz").text("00") : $(".task_dayz").text(getDateDiffWith10(t, e));
         }
      }),
      $(document).on("click", ".btn-staff-delete", function () {
         $(this).closest("tr").remove();
      });
});
const animal_is_seq = 104,
   member_is_seq = 111,
   end_date_is_seq = 112;
function mappingChangeAppInfo(e) {
   function t() {
      return '\n    <tr>\n      <td class="bullet pLeft10">미등록</td>\n      <td class="pLeft10"><span class="unit mRight5">(M/F)</span><span class="data"><span>-</span><span class="unit mHside5">/</span><span>-</span></span></td>\n      <td class="pHside20"><i class="fas fa-arrow-right"></i></td>\n      <td class="flexMid">\n        <input type="text" size="20" maxlength="20" class="form-control form-control-sm w100" value="">\n        <input type="number" min=0 size="4" maxlength="4" class="form-control form-control-sm text-center w60 mLeft5" value="0"><span class="unit mHside5">/</span>\n        <input type="number" min=0 size="4" maxlength="4" class="form-control form-control-sm text-center w60" value="0">\n        <a href="javascript:void(0);" class="btn btn_add_species btn-outline-primary mLeft10"><i class="fas fa-plus"></i></a>\n        <a href="javascript:void(0);" class="btn btn_remove_row btn-outline-danger mLeft3"><i class="far fa-trash-alt"></i></a>\n      </td>\n    </tr>';
   }
   var a = $("#ca_regular").children("ul").first().clone(),
      n = $("#ca_fast").children("ul").first().clone();
   function o(t) {
      var a = n.clone();
      let o = `regular_${t.is_seq}`,
         i = o + "-2",
         r = o + "-3",
         s = `regFile_${t.is_seq}`;
      a.find("span").eq(1).text(t.contents),
         "" == t.link_item_name ? a.find("select").attr("data-selected", "") : a.find("select").attr("data-selected", `form_${t.link_item_name}`),
         a.find("input").first().attr("id", o),
         a.find("input").first().attr("data-target", `#${r}`),
         a.find("label").first().attr("for", o);
      let l = a.children("li").eq(1);
      l.removeClass("regular_01"),
         l.attr("id", r),
         l.children("label").attr("for", i),
         l.children("textarea").eq(0).attr("id", i),
         l.children("div").find("input").attr("name", s),
         l.children("div").find("input").attr("id", s),
         l.children("div").find("label").eq(3).attr("for", s);
      var p = l.children("table").eq(0).children("tbody"),
         d = l.children("table").eq(1).children("tbody"),
         c = p.children("tr").eq(0).clone(),
         _ = d.children("tr").eq(0).clone();
      p.empty(), d.empty();
      var f,
         m = [],
         u = "general_director",
         g = e.getItemData(u);
      m.push(g.dataObj.saved_data.data[0]), (f = g.dataObj.saved_data.data[0].user_seq), (g_old_member_app_seq = g.dataObj.saved_data.data[0].app_seq);
      (u = "general_expt"), (g = e.getItemData(u));
      for (var h = 0; h < Object.keys(g.dataObj.saved_data.data).length; h++) m.push(g.dataObj.saved_data.data[h]);
      return (
         $.each(m, function (e, t) {
            var a = c.clone(),
               n = _.clone(),
               o = t.info;
            f == t.user_seq
               ? (a.children("td").eq(0).find("input").prop("checked", !0), n.children("td").eq(0).find("input").prop("checked", !0), n.children("td").eq(0).find("input").prop("checked", !0))
               : (a.children("td").eq(0).find("input").prop("checked", !1), n.children("td").eq(0).find("input").prop("checked", !1), a.children("td").eq(0).find("input").prop("disabled", !0)),
               1 == t.animal_mng_flag && f != t.user_seq
                  ? (a.children("td").eq(1).find("input").prop("checked", !0), n.children("td").eq(1).find("input").prop("checked", !0))
                  : (a.children("td").eq(1).find("input").prop("checked", !1), a.children("td").eq(1).find("input").prop("disabled", !0), n.children("td").eq(1).find("input").prop("checked", !1)),
               a.children("td").eq(2).text(o.name),
               a.children("td").eq(3).text(o.dept_str),
               a.children("td").eq(4).text(o.position_str),
               a.children("td").eq(5).text(o.major_field_str),
               a.children("td").eq(6).text(o.tmp_phoneno),
               a.children("td").eq(7).text(t.exp_year_code_str),
               p.append(a),
               n.attr("data-id", o.user_seq),
               n.children("td").eq(2).text(o.name),
               n.children("td").eq(3).text(o.dept_str),
               n.children("td").eq(4).text(o.position_str),
               n.children("td").eq(5).text(o.major_field_str),
               n.children("td").eq(6).text(o.tmp_phoneno),
               n.children("td").eq(7).children("div").children("select").val(t.exp_year_code).prop("selected", !0),
               d.append(n),
               $.each(g_all_userlist, function (e, a) {
                  t.user_seq == a.user_seq && (g_all_userlist[e] = t);
               });
         }),
         a
      );
   }
   function i(t, n, o, i) {
      if (116 == o.is_seq) return;
      var r = a.clone();
      r.removeClass("mTop30"),
         r.addClass(i),
         r.find("span").eq(0).removeClass("blue_deep"),
         r.find("span").eq(0).addClass(t),
         r
            .find("span")
            .eq(0)
            .text("[" + n + "]");
      let s = `regular_${o.is_seq}`,
         l = s + "-1",
         p = s + "-2",
         d = s + "-3",
         c = `regFile_${o.is_seq}`;
      r.find("span").eq(1).text(o.contents),
         "" == o.link_item_name ? r.find("select").attr("data-selected", "") : r.find("select").attr("data-selected", `form_${o.link_item_name}`),
         r.find("input").first().attr("id", s),
         r.find("input").first().attr("data-target", `#${d}`),
         r.find("label").first().attr("for", s);
      let _ = r.children("li").eq(1);
      return (
         _.removeClass("regular_01"),
         _.attr("id", d),
         _.children("label").attr("for", l),
         _.children("textarea").eq(0).attr("id", l),
         _.children("textarea").eq(1).attr("id", p),
         _.find("input").attr("name", c),
         _.find("input").attr("id", c),
         _.find("label").eq(3).attr("for", c),
         104 == o.is_seq &&
            (function (e, t) {
               var a = e.getItemData("animal_type_final");
               0 == a.dataObj.saved_data.data.length && (a = e.getItemData("animal_type"));
               var n = "<table><tbody>",
                  o = a.dataObj.saved_data.data;
               for (let e = 0; e < o.length; ++e)
                  n += `\n        <tr>\n          <td class="bullet pLeft10"><span class="data">${o[e].animal_code_str}</span></td>\n          <td class="pLeft10"><span class="unit mRight5">(M/F)</span><span class="data"><span>${o[e].male_cnt}</span><span class="unit mHside5">/</span><span>${o[e].female_cnt}</span></span></td>\n          <td class="pHside20"><i class="fas fa-arrow-right"></i></td>\n          <td class="flexMid">\n            <input type="text" size="20" maxlength="20" class="form-control form-control-sm w100" value="${o[e].animal_code_str}">\n            <input type="number" min="0" size="4" maxlength="4" class="form-control form-control-sm text-center w60 mLeft5" value="${o[e].male_cnt}"><span class="unit mHside5">/</span>\n            <input type="number" min="0" size="4" maxlength="4" class="form-control form-control-sm text-center w60" value="${o[e].female_cnt}">\n            <a href="javascript:void(0);" class="btn btn_add_species btn-outline-primary mLeft10"><i class="fas fa-plus"></i></a>\n            <a href="javascript:void(0);" class="btn btn_remove_row btn-outline-danger mLeft3"><i class="far fa-trash-alt"></i></a>\n          </td>\n        </tr>`;
               (n += "</tbody></table>"), t.replaceWith(n);
            })(e, _.children("textarea").eq(0)),
         112 == o.is_seq &&
            (_.find("label").eq(0).text("변경 전:"),
            (function (e, t) {
               var a = getTimeStrFromDateType(new Date(g_AppInfo.appObj.approved_dttm)),
                  n = e.getItemData("general_end_date", !1),
                  o = n.dataObj.saved_data.data[0],
                  i = getTimeStrFromDateType(new Date(o));
               g_old_end_date_app_seq = n.dataObj.saved_data.data.cur_app_seq;
               var r = new Date();
               r.setDate(r.getDate() + 1);
               var s = moment(r).format("YYYY-MM-DD"),
                  l = `<div class="bullet mTop5">${a} ~ ${i}</div>\n                  <label class="flex_1 mTop20">변경 후:</label>\n                  <div class="form-group flexMid">\n                  <label for="general_end_date_ca" class="bullet mRight10 mBot0">${a} ~ </label>\n                  <input class="form-control wCalendar" type="date" value="${s}" id="general_end_date_ca" min="${s}">\n                </div>`;
               t.replaceWith(l);
            })(e, _.children("textarea").eq(0))),
         r
      );
   }
   var r = e.getItemData("ca_regular_item");
   $("#ca_regular").empty();
   var s = r.dataObj.codes;
   for (let e = 0; e < s.length; ++e) {
      let t = "mTop2";
      0 == e && (t = "mTop30"), $("#ca_regular").append(i("blue_deep", "정규", s[e], t));
   }
   r = e.getItemData("ca_fast_item");
   $("#ca_fast").empty();
   s = r.dataObj.codes;
   for (let e = 0; e < s.length; ++e) {
      let t = "mTop2";
      0 == e && (t = "mTop30"), 111 != s[e].is_seq ? $("#ca_fast").append(i("green", "신속", s[e], t)) : $("#ca_fast").append(o(s[e]));
   }
   setTriggerCheckBoxOnChange(),
      setTriggerCheckBoxOnClickKeydownChange(),
      resetLeftStepNaviIcon(),
      $("input[type=checkbox]").on({
         change: function (e) {
            resetLeftStepNaviIcon();
         },
      }),
      $("select.move-to-selected").on({
         change: function (e) {
            resetLeftStepNaviIcon();
            let t = $(this).val();
            t && "#" != t && $("html, body").animate({ scrollTop: $(t).offset().top }, 300);
         },
      }),
      $(".btn-file-delete").on({
         click: function (e) {
            var t = $(this).parent().find("input").attr("name"),
               a = `<input type="file" name="${t}" class="custom-file-input" id="${t}">`;
            $(this).parent().find("input").replaceWith(a), $(this).parent().find("label").eq(1).text("");
         },
      });
   var l = JSON.parse(COMM.getCookie(CONST.COOKIE.IPSAP_USER_INFO));
   $(".reviewer").text(spaceAddStr(l.name)),
      (function e() {
         $(".btn_add_species").off(),
            $(".btn_add_species").on({
               click: function (a) {
                  "" != $(this).closest("tbody").children("tr").last().find("input").eq(0).val() && ($(this).closest("tbody").append(t()), e());
               },
            }),
            $(".btn_remove_row").off(),
            $(".btn_remove_row").on({
               click: function (a) {
                  var n = $(this).closest("tbody");
                  $(this).closest("tr").remove(), 0 == n.children("tr").length && (n.append(t()), e());
               },
            });
      })();
}
function onStaffSelect(e) {
   $.each(g_all_userlist, function (t, a) {
      if (Number(a.user_seq) == e) return makeStaffList(a), void $("#modal_staff").modal("hide");
   });
}
function getSeletecUserSeqArr() {
   var e = $("#changeUserTbody").children("tr"),
      t = [];
   for (let a = 0; a < e.length; a++) t.push(e.eq(a).data("id"));
   return t;
}
function openModalStaffChange() {
   makeOtherStaffList(), $("#modal_staff").modal("show");
}
function makeOtherStaffList() {
   var e = getSeletecUserSeqArr();
   let t = '<tbody id="other-staff">';
   $.each(g_all_userlist, function (a, n) {
      var o = n.info;
      isInArray(e, n.user_seq) ||
         ((o.tmp_phoneno = o.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
         (t += `\n      <tr data-url="" class="staff-row" onclick="onStaffSelect(${o.user_seq})">\n        <td class="text-center">${o.name}</td>\n        <td>${o.dept_str}</td>\n        <td>${o.position_str}</td>\n        <td>${o.major_field_str}</td>\n        <td>${o.tmp_phoneno}</td>\n        <td>${o.email}</td>\n        <td>${o.edu_course_num}</td>\n      </tr>`));
   }),
      (t += "</tbody>"),
      $("#other-staff").replaceWith(t);
}
function makeStaffList(e) {
   var t = e.info,
      a = `<tr data-id='${e.user_seq}'>\n                <td>\n                  <div class="custom-radio">\n                    <input type="radio" name="task_director_new" value="1">\n                  </div>\n                </td>\n                <td>\n                  <div class="custom-radio">\n                    <input type="radio" name="exp_manager_new" value="1">\n                  </div>\n                </td>\n                <td>${t.name}</td>\n                <td>${t.dept_str}</td>\n                <td>${t.position_str}</td>\n                <td>${t.major_field_str}</td>\n                <td class="hidden">${t.tmp_phoneno}</td>\n                <td>\n                <div class="btn-group">\n                  <select class="form-control btn-outline-primary small_select narrow_padding">\n                    <option selected disabled>경력 선택</option>\n                    <option value="1">1년 미만</option>\n                    <option value="2">1년 ~ 3년</option>\n                    <option value="3">3년 이상</option>\n                  </select>\n                </div>\n                </td>\n                <td class="hidden">${t.email}</td>\n                <td>\n                  <a href="javascript:void(0);" class="btn btn-outline-danger btn_xxs btn-staff-delete">삭제</a>\n                </td>\n              </tr>`;
   $("#changeUserTbody").append(a),
      e.exp_year_code > 0 && $(`#changeUserTbody > [data-id='${e.user_seq}']`).children("td").eq(7).children("div").children("select").val(e.exp_year_code).prop("selected", !0);
}
function resetLeftStepNaviIcon() {
   var e = {},
      t = 0,
      a = 0,
      n = g_AppItemParser.getItemData("ca_regular_item").dataObj.codes;
   for (let a = 0; a < n.length; ++a) {
      let o = `regular_${n[a].is_seq}`;
      if (1 == $(`#${o}`).is(":checked")) {
         try {
            let t = $(`#${o}`).parent().find("select").val().split("_")[1].split("-");
            e[String(Number(t[0]))] = !0;
         } catch (e) {}
         ++t;
      }
   }
   n = g_AppItemParser.getItemData("ca_fast_item").dataObj.codes;
   for (let t = 0; t < n.length; ++t) {
      let o = `regular_${n[t].is_seq}`;
      if (1 == $(`#${o}`).is(":checked")) {
         try {
            let t = $(`#${o}`).parent().find("select").val().split("_")[1].split("-");
            e[String(Number(t[0]))] = !0;
         } catch (e) {}
         ++a;
      }
   }
   return (
      $("#ca_regular_item_cnt").attr("data-cnt", t),
      $("#ca_regular_item_cnt").text("(" + t + ")"),
      $("#ca_fast_item_cnt").attr("data-cnt", a),
      $("#ca_fast_item_cnt").text("(" + a + ")"),
      $(".process_content").each(function (t, a) {
         $(a).removeClass("supplement"), 1 == e[$(a).attr("data-process-num")] && $(a).addClass("supplement");
      }),
      Object.keys(e).length
   );
}
function onClickSubmit() {
   if (IPSAP.DEMO_MODE) return void (window.location.href = "./experiment_list.html");
   if (0 == resetLeftStepNaviIcon()) return void alert("입력한 심의항목이 없습니다.\n변경할 항목의 내용을 입력해 주세요.");
   if (!$("#check_01").is(":checked")) return alert("제출 동의를 확인해 주세요."), void $("#check_01").focus();
   let e = {},
      t = {},
      a = {},
      n = [],
      o = {},
      i = [];
   var r = g_AppItemParser.getItemData("ca_regular_item").dataObj.codes;
   for (let a = 0; a < r.length; ++a) {
      let n = `regular_${r[a].is_seq}`,
         o = $(`#${n}-1`).val(),
         i = $(`#${n}-2`).val();
      if (1 != $(`#${n}`).is(":checked")) continue;
      if ("" == o || "" == i) return alert("체크한 항목의 변경 내용 및 변경 사유를 반드시 작성해 주세요."), void $(`#${n}`).focus();
      let s = $(`#${n}`).parent().find("select").val().split("_")[1];
      if (104 == r[a].is_seq) {
         let e = [];
         if (
            ($("#" + n)
               .closest("ul")
               .find("tbody")
               .children("tr")
               .each(function (t, a) {
                  let n = $(a).find("input").eq(0).val(),
                     o = Number($(a).find("input").eq(1).val()),
                     i = Number($(a).find("input").eq(2).val());
                  if ("" != n && o + i > 0) {
                     let t = [],
                        r = $(a).children("td").eq(0).children("span");
                     r.length > 0 ? t.push(r.text()) : t.push($(a).children("td").eq(0).text()),
                        t.push($(a).children("td").eq(1).children("span").eq(1).children("span").eq(0).text()),
                        t.push($(a).children("td").eq(1).children("span").eq(1).children("span").eq(2).text()),
                        t.push(n),
                        t.push(o),
                        t.push(i),
                        e.push(t);
                  }
               }),
            0 == e.length)
         )
            return void alert("변경하고는 하는 동물의 종 또는 수량이 없습니다.");
         o = JSON.stringify(e);
      }
      t[String(r[a].is_seq)] = JSON.stringify([s, o, i]);
      let l = `regFile_${r[a].is_seq}`;
      null != $("#" + l).prop("files")[0] && (e[`ca_regular_file_${r[a].is_seq}`] = $("#" + l).prop("files")[0]);
   }
   r = g_AppItemParser.getItemData("ca_fast_item").dataObj.codes;
   for (let t = 0; t < r.length; ++t) {
      let h = `regular_${r[t].is_seq}`,
         v = $(`#${h}-1`).val(),
         b = $(`#${h}-2`).val();
      if (1 != $(`#${h}`).is(":checked")) continue;
      var s = !0;
      if ((111 == r[t].is_seq ? (v = g_old_member_app_seq) : 112 == r[t].is_seq && (v = g_old_end_date_app_seq), ("" != v && "" != b) || (s = !1), !s))
         return alert("체크한 항목의 변경 내용 및 변경 사유를 반드시 작성해 주세요."), void $(`#${h}`).focus();
      if (111 == r[t].is_seq) {
         var l = $("#changeUserTbody").children("tr"),
            p = 0,
            d = 0;
         for (let e = 0; e < l.length; e++) {
            var c = 0,
               _ = l.eq(e).children("td"),
               f = l.eq(e).data("id"),
               m = _.eq(7).children("div").children("select").val();
            if (null == m) {
               var u = _.eq(2).text();
               return void alert(u + "님의 동물 실험 경력 년수를 선택 해주시기 바랍니다.");
            }
            var g = { animal_mng_flag: 0, exp_year_code: m, user_seq: f };
            if (
               (_.eq(0).children("div").children('input[name="task_director_new"]').is(":checked")
                  ? (p++, c++, n.push(g))
                  : _.eq(1).children("div").children('input[name="exp_manager_new"]').is(":checked")
                  ? (d++, c++, (g.animal_mng_flag = 1), i.push(g))
                  : i.push(g),
               2 == c)
            )
               return void alert("연구 책임자 와 실험담당자는 동시에 선택 될수 없습니다.");
         }
         if (1 != p) return void alert("연구 책임자는 필수 데이터 입니다.");
         if (1 != d) return void alert("실험 동물 관리 담당자는 필수 데이터 입니다.");
      } else 112 == r[t].is_seq && (o[0] = $("#general_end_date_ca").val());
      let I = $(`#${h}`).parent().find("select").val().split("_")[1];
      a[String(r[t].is_seq)] = JSON.stringify([I, v, b]);
      let q = `regFile_${r[t].is_seq}`;
      null != $("#" + q).prop("files")[0] && (e[`ca_fast_file_${r[t].is_seq}`] = $("#" + q).prop("files")[0]);
   }
   var h = { parent_app_seq: g_AppItemParser.app_seq, application_type: IPSAP.APPLICATION_TYPE.CHANGE, judge_type: g_AppInfo.appObj.judge_type },
      v = new ItemParser(0);
   (v.targetSavedItemNames = []),
      (v.moreSaveFiles = e),
      v.addMoreSaveTag("application_info", { data: h }),
      v.addMoreSaveTag("ca_regular_item", { data: t }),
      v.addMoreSaveTag("ca_fast_item", { data: a }),
      n.length > 0 && v.addMoreSaveTag("general_director", { data: n }),
      i.length > 0 && v.addMoreSaveTag("general_expt", { data: i }),
      null != o[0] && v.addMoreSaveTag("general_end_date", { data: o }),
      appItemSubmitWithParser(v, onCompleteSubmit, IPSAP.APP_SUBMIT_TYPE_PARAM.CHILD_SUBMIT);
}
function onCompleteSubmit(e, t, a) {
   e ? (window.location.href = "./experiment_list.html") : alert(`변경 승인 신청을 실패했습니다.\n(${t.em})`);
}
$("#modal_changeApp").on("hidden.bs.modal", (e) => {
   $(e.currentTarget).modal("show");
});
