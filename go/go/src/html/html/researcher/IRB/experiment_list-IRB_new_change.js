"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   "undefined" == typeof ipsap_application_common_js && document.write("<script src='/assets/js/ipsap/ipsap_application_common.js'></script>"),
   "undefined" == typeof date_utils_js && document.write("<script src='/assets/js/common/util/date_utils.js'></script>");
var g_all_userlist = [];
let g_old_member_app_seq = 0,
   g_old_end_date_app_seq = 0;
const member_is_seq = 299,
   end_date_is_seq = 300;
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
            $(".modal_hint").hide(),
               e &&
                  ($(".modal[data-backdrop=false]").css({ "background-color": "rgba(0,0,0, 0.0)", "pointer-events": "none" }),
                  $(".modal-dialog-centered.ui-draggable").css({ "justify-content": "flex-start" }),
                  "click" == t.type && $(this).css({ "margin-right": "0", left: "0", top: "0" })),
               (e = !1),
               "drag" == t.type && $(".btn_shrink i").switchClass("fa-arrow-up", "fa-arrow-down");
         },
      }),
      $(".modal.fade").each(function (e, t) {
         $(t).hasClass("in") && $(t).modal("show"), "modal_changeApp" == $(t).attr("id") && $("body").css({ overflow: "auto" });
      }),
      $("input[id^=change_]").on({
         change: function (e) {
            $(this).prop("checked")
               ? ($(this).closest("li").addClass("active"), $($(this).data("target")).addClass("active"))
               : ($(this).closest("li").removeClass("active"), $($(this).data("target")).removeClass("active"));
         },
      });
   $("select.move-to-selected").html(
      '<option selected disabled>신청서 항목</option><optgroup label="1. 심의 신청서">\n                  <option value="#form_01-1">1-1. 연구과제 기본 정보</option>\n                  <option value="#form_01-2">1-2. 연구원 정보</option>\n                  <option value="#form_01-3">1-3. 연구유형 및 심의분류</option>\n                  <option value="#form_01-4" disabled>1-4. 연구대상</option>\n                </optgroup><optgroup label="2. 연구 상세 요약서">\n                  <option value="#form_02-1">2-1. 연구 목적</option>\n                  <option value="#form_02-2">2-2. 배경 및 이론적 증거</option>\n                  <option value="#form_02-3">2-3. 연구 방법</option>\n                  <option value="#form_02-4">2-4. 관찰 및 검사 항목</option>\n                  <option value="#form_02-5">2-5. 평가 기준 및 평가 방법</option>\n                </optgroup><optgroup label="3. 동의취득">\n                  <option value="#form_03-1">3-1. 동의취득 정보</option>\n                </optgroup><optgroup label="4. 제출서류">\n                  <option value="#form_04-1">4-1. 필수 제출 서류</option>\n                  <option value="#form_04-2">4-2. 추가 제출 서류</option>\n                  <option value="#form_04-3">4-3. 기타 제출 서류</option>\n                </optgroup>'
   ),
      $("select.move-to-selected").each((e, t) => {
         $(t)
            .find('option[value="#' + $(t).data("selected") + '"]')
            .attr("selected", !0);
      }),
      $("select.move-to-selected")
         .closest(".list-group-item")
         .on({
            click: (e) => {
               let t = $(e.currentTarget).find("select").val();
               t && "#" != t && $("html, body").animate({ scrollTop: $(t).offset().top }, 300);
            },
         }),
      $("a.move-to-selected").on({
         click: (e) => {
            let t = $(e.currentTarget).data("move-to-target");
            t && "#" != t && $("html, body").animate({ scrollTop: $(t).offset().top }, 300);
         },
      }),
      $(".btn_reset_select").on({
         click: (e) => {
            $(e.currentTarget).prev("select").find("option[selected=selected]").prop("selected", !0);
         },
      }),
      $("#review_summary").on({
         click: function (e) {
            var t = $(this),
               a = t.closest(".modal");
            t.toggleClass("active"), t.hasClass("active") ? t.text("변경항목 수정") : t.text("변경항목 모아보기"), a.find("input, textarea").prop("readonly", !0);
            let n = $("#irb_ca").children().children(":nth-child(even)");
            t.hasClass("active")
               ? (a.addClass("view_mode"),
                 n.find("select").prop("disabled", !0),
                 n.filter(":not(.show)").closest("ul").hide(),
                 n.filter(".show").length ||
                    $("#irb_ca").append(
                       '\n            <ul class="list-group empty_list">\n              <li class="list-group-item text-center">\n                <div class="underline">등록한 변경 사항이 없습니다.</div>\n              </li>\n            </ul>'
                    ))
               : (a.removeClass("view_mode"),
                 a.find("input, textarea").prop("readonly", !1).removeClass("view_mode"),
                 n.find("select").prop("disabled", !1),
                 n.filter(":not(.show)").closest("ul").removeAttr("style"),
                 $("#irb_ca").find(".empty_list").remove()),
               $("#check_01").prop("readonly", !1);
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
function mappingChangeAppInfo(e) {
   setExpCategory(e);
   var t = $("#irb_ca").children("ul").eq(1).clone(),
      a = $("#irb_ca").children("ul").eq(4).clone();
   function n(a) {
      var n = t.clone();
      let o = `irb_${a.is_seq}`,
         i = o + "-2",
         r = o + "-3",
         l = `irbFile_${a.is_seq}`;
      "" != a.link_item_name && n.find("select").attr("data-selected", `form_${a.link_item_name}`),
         n.find("input").first().attr("id", o),
         n.find("input").first().attr("data-target", `#${r}`),
         n.find("label").first().attr("for", o);
      let s = n.children("li").eq(1);
      s.attr("id", r),
         s.children("label").attr("for", i),
         s.children("textarea").eq(0).attr("id", i),
         s.children("div").find("input").attr("name", l),
         s.children("div").find("input").attr("id", l),
         s.children("div").find("label").eq(3).attr("for", l);
      var d = s.children("table").eq(0).children("tbody"),
         p = s.children("table").eq(1).children("tbody"),
         c = d.children("tr").eq(0).clone(),
         _ = p.children("tr").eq(0).clone();
      d.empty(), p.empty();
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
            f == t.user_seq ? a.children("td").eq(0).text("연구 책임자") : a.children("td").eq(0).text(t.exp_type_code_str),
               a.children("td").eq(1).text(o.name),
               a.children("td").eq(2).text(o.dept_str),
               a.children("td").eq(3).text(o.position_str),
               a.children("td").eq(4).text(o.major_field_str),
               a.children("td").eq(6).text(t.exp_year_code_str),
               d.append(a),
               n.attr("data-id", o.user_seq),
               n.children("td").eq(0).children("select").val(t.exp_type_code).prop("selected", !0),
               n.children("td").eq(1).text(o.name),
               n.children("td").eq(2).text(o.dept_str),
               n.children("td").eq(3).text(o.position_str),
               n.children("td").eq(4).text(o.major_field_str),
               n.children("td").eq(6).children("div").children("select").val(t.exp_year_code).prop("selected", !0),
               p.append(n),
               $.each(g_all_userlist, function (e, a) {
                  t.user_seq == a.user_seq && (g_all_userlist[e] = t);
               });
         }),
         n
      );
   }
   function o(t) {
      var n = a.clone();
      let o = `irb_${t.is_seq}`,
         i = o + "-1",
         r = o + "-2",
         l = o + "-3",
         s = `irbFile_${t.is_seq}`;
      n.find("label").eq(0).text(t.contents),
         "" != t.link_item_name ? n.find("select").attr("data-selected", `form_${t.link_item_name}`) : (n.find("select").attr("data-selected", !1), n.find(".btn_reset_select").remove()),
         n.find("input").first().attr("id", o),
         n.find("input").first().attr("data-target", `#${l}`),
         n.find("label").first().attr("for", o);
      let d = n.children("li").eq(1);
      return (
         d.attr("id", l),
         d.children("label").attr("for", i),
         d.children("textarea").eq(0).attr("id", i),
         d.children("textarea").eq(1).attr("id", r),
         d.find("input").attr("name", s),
         d.find("input").attr("id", s),
         d.find("label").eq(3).attr("for", s),
         300 == t.is_seq &&
            (d.find("label").eq(0).text("변경 전:"),
            (function (e, t) {
               var a = getTimeStrFromDateType(new Date(g_AppInfo.appObj.approved_dttm)),
                  n = e.getItemData("general_end_date", !1),
                  o = n.dataObj.saved_data.data[0],
                  i = getTimeStrFromDateType(new Date(o));
               g_old_end_date_app_seq = n.dataObj.saved_data.data.cur_app_seq;
               var r = new Date();
               r.setDate(r.getDate() + 1);
               var l = moment(r).format("YYYY-MM-DD"),
                  s = `<div class="bullet mTop5">${a} ~ ${i}</div>\n                  <label class="flex_1 mTop20">변경 후:</label>\n                  <div class="form-group flexMid">\n                  <label for="general_end_date_ca" class="bullet mRight10 mBot0">${a} ~ </label>\n                  <input class="form-control wCalendar" type="date" value="${l}" id="general_end_date_ca" min="${l}">\n                </div>`;
               t.replaceWith(s);
            })(e, d.children("textarea").eq(0))),
         n
      );
   }
   var i = e.getItemData("irb_ca_item");
   $("#irb_ca").empty();
   var r = i.dataObj.codes;
   for (let e = 0; e < r.length; ++e) 299 == r[e].is_seq ? $("#irb_ca").append(n(r[e])) : $("#irb_ca").append(o(r[e]));
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
   $(".reviewer").text(spaceAddStr(l.name));
}
function resetLeftStepNaviIcon() {
   var e = {},
      t = g_AppItemParser.getItemData("irb_ca_item").dataObj.codes;
   for (let a = 0; a < t.length; ++a) {
      let n = `irb_${t[a].is_seq}`;
      if (1 == $(`#${n}`).is(":checked"))
         try {
            let t = $(`#${n}`).parent().find("select").val().split("_")[1].split("-");
            e[String(Number(t[0]))] = !0;
         } catch (e) {}
   }
   return (
      $(".process_content").each(function (t, a) {
         $(a).removeClass("supplement"), 1 == e[$(a).attr("data-process-num")] && $(a).addClass("supplement");
      }),
      Object.keys(e).length
   );
}
function appendExpCategory(e) {
   $("#exp_category").empty();
   var t = '<span class="sum_label">실험 분류:</span>';
   for (let a = 0; a < e.length; a++) t += `<button type="button" class="btn btn-outline-primary btn-round btn_tag btn_tag_sm mRight5">\n                  ${e[a]}\n                </button>`;
   $("#exp_category").append(t);
}
function setExpCategory(e) {
   var t = [],
      a = "general_object";
   1 == e.getItemData(a, "nosave").getStringValue("0") ? t.push("심의 대상") : t.push("면제 대상");
   a = "general_human_research";
   1 == e.getItemData(a, "nosave").getStringValue("0") && t.push("인간대상 연구");
   a = "general_body_research";
   1 == e.getItemData(a, "nosave").getStringValue("0") && t.push("인체 유래물 연구"), appendExpCategory(t);
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
      a = `<tr data-id='${e.user_seq}'>\n              <td>\n              <select class="form-control btn-outline-primary small_select">\n                <option selected disabled>구분 선택</option>\n                <option value="0">연구 책임자</option>\n                <option value="1">연구 담당자</option>\n                <option value="2">공동 연구자</option>\n                <option value="3">연구원</option>\n                <option value="4">지도 교수</option>\n              </select>\n              </td>\n              <td>${t.name}</td>\n              <td>${t.dept_str}</td>\n              <td>${t.position_str}</td>\n              <td>${t.major_field_str}</td>\n              <td class="hidden">${t.tmp_phoneno}</td>\n              <td>\n              <div class="btn-group">\n                <select class="form-control btn-outline-primary small_select narrow_padding">\n                  <option selected disabled>경력 선택</option>\n                  <option value="1">1년 미만</option>\n                  <option value="2">1년 ~ 3년</option>\n                  <option value="3">3년 이상</option>\n                </select>\n              </div>\n              </td>\n              <td class="hidden">${t.email}</td>\n              <td>\n                <a href="javascript:void(0);" class="btn btn-outline-danger btn_xxs btn-staff-delete">삭제</a>\n              </td>\n              </tr>`;
   $("#changeUserTbody").append(a),
      e.exp_year_code > 0 && $(`#changeUserTbody > [data-id='${e.user_seq}']`).children("td").eq(6).children("div").children("select").val(e.exp_year_code).prop("selected", !0),
      e.exp_type_code > 0 && $(`#changeUserTbody > [data-id='${e.user_seq}']`).children("td").eq(0).children("select").val(e.exp_type_code).prop("selected", !0);
}
function onClickSubmit() {
   if (0 == resetLeftStepNaviIcon()) return void alert("입력 또는 선택한 심의항목이 없습니다.");
   if (!$("#check_01").is(":checked")) return alert("제출 동의를 확인해 주세요."), void $("#check_01").focus();
   let e = {},
      t = {},
      a = [],
      n = {},
      o = [];
   var i = g_AppItemParser.getItemData("irb_ca_item").dataObj.codes;
   for (let m = 0; m < i.length; ++m) {
      let u = `irb_${i[m].is_seq}`,
         g = $(`#${u}-1`).val(),
         h = $(`#${u}-2`).val();
      if (1 != $(`#${u}`).is(":checked")) continue;
      if ((299 == i[m].is_seq ? (g = g_old_member_app_seq) : 300 == i[m].is_seq && (g = g_old_end_date_app_seq), "" == g || "" == h))
         return alert("체크한 항목의 변경 내용 및 변경 사유를 반드시 작성해 주세요."), void $(`#${u}`).focus();
      if (299 == i[m].is_seq) {
         var r = $("#changeUserTbody").children("tr"),
            l = 0;
         for (let e = 0; e < r.length; e++) {
            var s = r.eq(e).children("td"),
               d = r.eq(e).data("id"),
               p = s.eq(0).children("select").val(),
               c = s.eq(6).children("div").children("select").val(),
               _ = s.eq(1).text();
            if (null == c) return void alert(_ + "님의 동물 실험 경력 년수를 선택 해주시기 바랍니다.");
            if (null == p) return void alert(_ + "님의 연구원 구분을 선택 해주시기 바랍니다.");
            var f = { animal_mng_flag: 0, exp_type_code: p, exp_year_code: c, user_seq: d };
            0 == p ? (l++, a.push(f)) : o.push(f);
         }
         if (1 != l) return void alert("연구 책임자는 한명 이어야 됩니다.");
      } else 300 == i[m].is_seq && (n[0] = $("#general_end_date_ca").val());
      if (null == $(`#${u}`).parent().find("select").val()) return void alert("실험 계획서 항목을 선택 해주시기 바랍니다.");
      let v = $(`#${u}`).parent().find("select").val().split("_")[1];
      t[String(i[m].is_seq)] = JSON.stringify([v, g, h]);
      let b = `irbFile_${i[m].is_seq}`;
      null != $("#" + b).prop("files")[0] && (e[`irb_ca_file_${i[m].is_seq}`] = $("#" + b).prop("files")[0]);
   }
   var m = { parent_app_seq: g_AppItemParser.app_seq, application_type: IPSAP.APPLICATION_TYPE.CHANGE, judge_type: g_AppInfo.appObj.judge_type },
      u = new ItemParser(0);
   (u.targetSavedItemNames = []),
      (u.moreSaveFiles = e),
      u.addMoreSaveTag("application_info", { data: m }),
      u.addMoreSaveTag("irb_ca_item", { data: t }),
      a.length > 0 && u.addMoreSaveTag("general_director", { data: a }),
      o.length > 0 && u.addMoreSaveTag("general_expt", { data: o }),
      null != n[0] && u.addMoreSaveTag("general_end_date", { data: n }),
      appItemSubmitWithParser(u, onCompleteSubmit);
}
function onCompleteSubmit(e, t, a) {
   e ? (window.location.href = "../experiment_list.html") : alert(`변경 승인 신청을 실패했습니다.\n(${t.em})`);
}
$(function () {
   g_AppInfo.loadParams(),
      (function () {
         let e = CONST.API.INSTITUTION.USER;
         API.load({
            url: e.replace("${inst_seq}", g_AppInfo.appObj.institution_seq),
            data: { "filter.user_type": IPSAP.USER_TYPE.RESEARCHER },
            type: CONST.API_TYPE.GET,
            success: function (e) {
               $.each(e, function (e, t) {
                  var a = { exp_type_code: 0, exp_year_code: 0, user_seq: t.user_seq, info: t };
                  g_all_userlist.push(a);
               });
            },
         });
      })(),
      $("#all_app_r").load("/html/common/inc_application/all_app_readonly_irb.html", function () {
         remakeApplicationInfoReadOnly(["irb_change_application"], cbFuncAfterMapping),
            $(".apptype_title").text(g_AppInfo.appObj.application_type_str),
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
      }),
      $(document).on("click", ".btn-staff-delete", function () {
         $(this).closest("tr").remove();
      });
}),
   $("#modal_changeApp").on("hidden.bs.modal", (e) => {
      $(e.currentTarget).modal("show");
   });
