"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   "undefined" == typeof ipsap_application_common_js && document.write("<script src='/assets/js/ipsap/ipsap_application_common.js'></script>"),
   "undefined" == typeof date_utils_js && document.write("<script src='/assets/js/common/util/date_utils.js'></script>");
var g_all_userlist = [];
let g_old_member_app_seq = 0,
   g_old_end_date_app_seq = 0;
var g_min_value = 0;
const member_is_seq = 286,
   end_date_is_seq = 288;
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
   var t = "<option selected disabled>신청서 항목</option>",
      a =
         '<optgroup label="1. 심의 신청서">\n                  <option value="#form_01-1">1-1. 연구과제 기본 정보</option>\n                  <option value="#form_01-2">1-2. 연구원 및 연구시설 정보</option>\n                  <option value="#form_01-3">1-3. 실험 분류</option>\n                  <option value="#form_01-4">1-4. 실험동물 사용 유무</option>\n                  <option value="#form_01-5">1-5. 기타 관련자료의 제출</option>\n                </optgroup>',
      n =
         '<optgroup label="2. 연구 계획서">\n                  <option value="#form_02-1">2-1. 연구과제 기본 정보</option>\n                  <option value="#form_02-2">2-2. 실험 구분</option>\n                  <option value="#form_02-3">2-3. 연구목적 및 예상 성과</option>\n                  <option value="#form_02-4">2-4. 연구내용 및 범위</option>\n                  <option value="#form_02-5">2-5. 연구방법</option>\n                </optgroup>';
   (t +=
      g_min_value < 9
         ? a +
           n +
           '<optgroup label="3. 위해성 평가서">\n                  <option value="#form_03-1">3-1. 취급 생물체 및 물질 정보</option>\n                  <option value="#form_03-2">3-2. 생물안전정보</option>\n                </optgroup>'
         : g_min_value < 11 && g_min_value > 8
         ? a + n
         : a),
      $("select.move-to-selected").html(t),
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
            let n = $("#ibc_ca").children().children(":nth-child(even)");
            t.hasClass("active")
               ? (a.addClass("view_mode"), n.find("select").prop("disabled", !0), n.filter(":not(.show)").closest("ul").hide())
               : (a.removeClass("view_mode"),
                 a.find("input, textarea").prop("readonly", !1).removeClass("view_mode"),
                 n.find("select").prop("disabled", !1),
                 n.filter(":not(.show)").closest("ul").removeAttr("style")),
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
   var t = (d = e.getItemData("ibc_general_experiment")).getStringValue("0"),
      a = [];
   "" != t && (a = t.split(",")), a.length > 0 && (g_min_value = Math.min.apply(null, a));
   var n = $("#ibc_ca").children("ul").eq(0).clone(),
      i = $("#ibc_ca").children("ul").eq(1).clone();
   function r(t) {
      var a = n.clone();
      let i = `ibc_${t.is_seq}`,
         r = i + "-2",
         o = i + "-3",
         d = `ibcFile_${t.is_seq}`;
      "" != t.link_item_name && a.find("select").attr("data-selected", `form_${t.link_item_name}`),
         a.find("input").first().attr("id", i),
         a.find("input").first().attr("data-target", `#${o}`),
         a.find("label").first().attr("for", i);
      let l = a.children("li").eq(1);
      l.attr("id", o),
         l.children("label").attr("for", r),
         l.children("textarea").eq(0).attr("id", r),
         l.children("div").find("input").attr("name", d),
         l.children("div").find("input").attr("id", d),
         l.children("div").find("label").eq(3).attr("for", d);
      var s = l.children("table").eq(0).children("tbody"),
         c = l.children("table").eq(1).children("tbody"),
         p = s.children("tr").eq(0).clone(),
         _ = c.children("tr").eq(0).clone();
      s.empty(), c.empty();
      var u,
         f = [],
         m = "general_director",
         g = e.getItemData(m);
      f.push(g.dataObj.saved_data.data[0]), (u = g.dataObj.saved_data.data[0].user_seq), (g_old_member_app_seq = g.dataObj.saved_data.data[0].app_seq);
      (m = "general_expt"), (g = e.getItemData(m));
      for (var h = 0; h < Object.keys(g.dataObj.saved_data.data).length; h++) f.push(g.dataObj.saved_data.data[h]);
      return (
         $.each(f, function (e, t) {
            var a = p.clone(),
               n = _.clone(),
               i = t.info;
            let r = "";
            if (t.edu_course) {
               let a = JSON.parse(t.edu_course);
               for (e in a) r += `\n          <div class="flexMid">\n            <span class="w120 mRight5">${e}</span>\n            <div>${a[e]}</div>\n          </div>`;
            }
            u == t.user_seq
               ? (a.children("td").eq(0).find("input").prop("checked", !0), n.children("td").eq(0).find("input").prop("checked", !0), n.children("td").eq(0).find("input").prop("checked", !0))
               : (a.children("td").eq(0).find("input").prop("checked", !1), n.children("td").eq(0).find("input").prop("checked", !1), a.children("td").eq(0).find("input").prop("disabled", !0)),
               1 == t.animal_mng_flag && u != t.user_seq
                  ? (a.children("td").eq(1).find("input").prop("checked", !0), n.children("td").eq(1).find("input").prop("checked", !0))
                  : (a.children("td").eq(1).find("input").prop("checked", !1), a.children("td").eq(1).find("input").prop("disabled", !0), n.children("td").eq(1).find("input").prop("checked", !1)),
               a.children("td").eq(2).text(i.name),
               a.children("td").eq(3).text(i.dept_str),
               a.children("td").eq(4).append(r),
               s.append(a),
               n.attr("data-id", i.user_seq),
               n.children("td").eq(2).text(i.name),
               n.children("td").eq(3).text(i.dept_str),
               n.children("td").eq(4).append(makeCourseList(i.user_seq)),
               c.append(n),
               $.each(g_all_userlist, function (e, a) {
                  t.user_seq == a.user_seq && (g_all_userlist[e] = t);
               });
         }),
         a
      );
   }
   function o(t) {
      var a = i.clone();
      let n = `ibc_${t.is_seq}`,
         r = n + "-1",
         o = n + "-2",
         d = n + "-3",
         l = `ibcFile_${t.is_seq}`;
      a.find("label").eq(0).text(t.contents),
         "" != t.link_item_name ? a.find("select").attr("data-selected", `form_${t.link_item_name}`) : (a.find("select").attr("data-selected", !1), a.find(".btn_reset_select").remove()),
         a.find("input").first().attr("id", n),
         a.find("input").first().attr("data-target", `#${d}`),
         a.find("label").first().attr("for", n);
      let s = a.children("li").eq(1);
      return (
         s.attr("id", d),
         s.children("label").attr("for", r),
         s.children("textarea").eq(0).attr("id", r),
         s.children("textarea").eq(1).attr("id", o),
         s.find("input").attr("name", l),
         s.find("input").attr("id", l),
         s.find("label").eq(3).attr("for", l),
         288 == t.is_seq &&
            (s.find("label").eq(0).text("변경 전:"),
            (function (e, t) {
               var a = getTimeStrFromDateType(new Date(g_AppInfo.appObj.approved_dttm)),
                  n = e.getItemData("general_end_date", !1),
                  i = n.dataObj.saved_data.data[0],
                  r = getTimeStrFromDateType(new Date(i));
               g_old_end_date_app_seq = n.dataObj.saved_data.data.cur_app_seq;
               var o = new Date();
               o.setDate(o.getDate() + 1);
               var d = moment(o).format("YYYY-MM-DD"),
                  l = `<div class="bullet mTop5">${a} ~ ${r}</div>\n                  <label class="flex_1 mTop20">변경 후:</label>\n                  <div class="form-group flexMid">\n                  <label for="general_end_date_ca" class="bullet mRight10 mBot0">${a} ~ </label>\n                  <input class="form-control wCalendar" type="date" value="${d}" id="general_end_date_ca" min="${d}">\n                </div>`;
               t.replaceWith(l);
            })(e, s.children("textarea").eq(0))),
         a
      );
   }
   var d = e.getItemData("ibc_ca_item");
   $("#ibc_ca").empty();
   var l = d.dataObj.codes;
   if (g_min_value > 8) {
      var s = [];
      g_min_value < 11 ? (s = [289]) : g_min_value >= 11 && (s = [289, 290, 291]);
      for (let e = 0; e < s.length; e++) {
         var c = l.findIndex(function (t) {
            return t.is_seq === s[e];
         });
         c > -1 && l.splice(c, 1);
      }
   }
   for (let e = 0; e < l.length; ++e) 286 != l[e].is_seq ? $("#ibc_ca").append(o(l[e])) : $("#ibc_ca").append(r(l[e]));
   for (const e of g_all_userlist) e.edu_course && mappingCourseList(e.info.user_seq, JSON.parse(e.edu_course));
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
   var p = JSON.parse(COMM.getCookie(CONST.COOKIE.IPSAP_USER_INFO));
   $(".reviewer").text(spaceAddStr(p.name)), $(".card").removeClass("hidden");
}
function resetLeftStepNaviIcon() {
   var e = {},
      t = g_AppItemParser.getItemData("ibc_ca_item").dataObj.codes;
   for (let a = 0; a < t.length; ++a) {
      let n = `ibc_${t[a].is_seq}`;
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
      var i = n.info;
      isInArray(e, n.user_seq) ||
         ((i.tmp_phoneno = i.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
         (t += `\n      <tr data-url="" class="staff-row" onclick="onStaffSelect(${i.user_seq})">\n        <td class="text-center">${i.name}</td>\n        <td>${i.dept_str}</td>\n        <td>${i.position_str}</td>\n        <td>${i.major_field_str}</td>\n        <td>${i.tmp_phoneno}</td>\n        <td>${i.email}</td>\n        <td>${i.edu_course_num}</td>\n      </tr>`));
   }),
      (t += "</tbody>"),
      $("#other-staff").replaceWith(t);
}
function makeStaffList(e) {
   var t = e.info,
      a = `\n    <tr data-id='${
         e.user_seq
      }'>\n      <td>\n        <div class="custom-radio">\n          <input type="radio" name="task_director_new" value="1">\n        </div>\n      </td>\n      <td>\n        <div class="custom-radio">\n          <input type="radio" name="exp_manager_new" value="1">\n        </div>\n      </td>\n      <td>${
         t.name
      }</td>\n      <td>${t.dept_str}</td>\n      <td>${makeCourseList(
         t.user_seq
      )}</td>\n      <td>\n        <a href="javascript:void(0);" class="btn btn-outline-danger btn_xxs btn-staff-delete">삭제</a>\n      </td>\n    </tr>`;
   $("#changeUserTbody").append(a),
      e.exp_year_code > 0 && $(`#changeUserTbody > [data-id='${e.user_seq}']`).children("td").eq(7).children("div").children("select").val(e.exp_year_code).prop("selected", !0);
}
function mappingCourseList(e, t) {
   let a = !1;
   $.each(t, function (t, n) {
      $(`input[id='${t}_${e}']`).val(n), n && (a = !0);
   }),
      a && $(`.course_${e}`).click();
}
function makeCourseList(e) {
   let t = ["생물안전 교육", "LMO 생물안전 교육", "생물안전 3등급 교육", "동물실험 교육"],
      a = `<a href="javascript:void(0);" class="btn btn-xs btn-outline-primary course_${e} w100p collapse show" data-toggle="collapse" data-target=".course_${e}">교육이수 정보 입력</a>\n                     <div class="course_${e} collapse">`;
   for (course_no in t)
      a += `\n      <div class="flexMid text-left">\n        <span class="w110 left mRight5">${t[course_no]}</span>\n        <input type="text" id="${t[course_no]}_${e}" class="form-control form-control-sm w150">\n      </div>\n    `;
   return (a += `<a href="javascript:void(0);" class="btn btn-xs btn-outline-secondary w100p mTop5" data-toggle="collapse" data-target=".course_${e}">입력 취소</a>\n        </div>`), a;
}
function makeEduCourse(e) {
   if (!$(`.course_${e}`).eq(1).hasClass("show")) return "";
   let t = new Object();
   return (
      $(`.course_${e}`)
         .children("div")
         .each(function () {
            t[$.trim($(this).eq(0).text())] = $(this).children("input").val();
         }),
      JSON.stringify(t)
   );
}
function onClickSubmit() {
   if (0 == resetLeftStepNaviIcon()) return void alert("입력 또는 선택한 심의항목이 없습니다.");
   if (!$("#check_01").is(":checked")) return alert("제출 동의를 확인해 주세요."), void $("#check_01").focus();
   let e = {},
      t = {},
      a = [],
      n = {},
      i = [];
   var r = g_AppItemParser.getItemData("ibc_ca_item").dataObj.codes;
   if (g_min_value > 8) {
      var o = [];
      g_min_value < 11 ? (o = [289]) : g_min_value >= 11 && (o = [289, 290, 291]);
      for (let e = 0; e < o.length; e++) {
         var d = r.findIndex(function (t) {
            return t.is_seq === o[e];
         });
         d > -1 && r.splice(d, 1);
      }
   }
   for (let o = 0; o < r.length; ++o) {
      let d = `ibc_${r[o].is_seq}`,
         g = $(`#${d}-1`).val(),
         h = $(`#${d}-2`).val();
      if (1 != $(`#${d}`).is(":checked")) continue;
      var l = !0;
      if ((286 == r[o].is_seq ? (g = g_old_member_app_seq) : 288 == r[o].is_seq && (g = g_old_end_date_app_seq), ("" != g && "" != h) || (l = !1), !l))
         return alert("체크한 항목의 변경 내용 및 변경 사유를 반드시 작성해 주세요."), void $(`#${d}`).focus();
      if (286 == r[o].is_seq) {
         var s = $("#changeUserTbody").children("tr"),
            c = 0,
            p = 0;
         for (let e = 0; e < s.length; e++) {
            var _ = 0,
               u = s.eq(e).children("td"),
               f = s.eq(e).data("id"),
               m = { animal_mng_flag: 0, exp_year_code: 0, user_seq: f, edu_course: makeEduCourse(f) };
            if (
               (u.eq(0).children("div").children('input[name="task_director_new"]').is(":checked")
                  ? (c++, _++, a.push(m))
                  : u.eq(1).children("div").children('input[name="exp_manager_new"]').is(":checked")
                  ? (p++, _++, (m.animal_mng_flag = 1), i.push(m))
                  : i.push(m),
               2 == _)
            )
               return void alert("연구 책임자 와 실험담당자는 동시에 선택 될수 없습니다.");
         }
         if (1 != c) return void alert("연구 책임자는 필수 데이터 입니다.");
         if (1 != p) return void alert("실험 동물 관리 담당자는 필수 데이터 입니다.");
      } else 288 == r[o].is_seq && (n[0] = $("#general_end_date_ca").val());
      if (null == $(`#${d}`).parent().find("select").val()) return void alert("실험 계획서 항목을 선택 해주시기 바랍니다.");
      let v = $(`#${d}`).parent().find("select").val().split("_")[1];
      t[String(r[o].is_seq)] = JSON.stringify([v, g, h]);
      let b = `ibcFile_${r[o].is_seq}`;
      null != $("#" + b).prop("files")[0] && (e[`ibc_ca_file_${r[o].is_seq}`] = $("#" + b).prop("files")[0]);
   }
   var g = { parent_app_seq: g_AppItemParser.app_seq, application_type: IPSAP.APPLICATION_TYPE.CHANGE, judge_type: g_AppInfo.appObj.judge_type },
      h = new ItemParser(0);
   (h.targetSavedItemNames = []),
      (h.moreSaveFiles = e),
      h.addMoreSaveTag("application_info", { data: g }),
      h.addMoreSaveTag("ibc_ca_item", { data: t }),
      a.length > 0 && h.addMoreSaveTag("general_director", { data: a }),
      i.length > 0 && h.addMoreSaveTag("general_expt", { data: i }),
      null != n[0] && h.addMoreSaveTag("general_end_date", { data: n }),
      appItemSubmitWithParser(h, onCompleteSubmit);
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
                  var a = { animal_mng_flag: 0, exp_year_code: 0, user_seq: t.user_seq, info: t };
                  g_all_userlist.push(a);
               });
            },
         });
      })(),
      $("#all_app_r").load("/html/common/inc_application/all_app_readonly_ibc.html", function () {
         remakeApplicationInfoReadOnly(["ibc_change_application"], cbFuncAfterMapping),
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
