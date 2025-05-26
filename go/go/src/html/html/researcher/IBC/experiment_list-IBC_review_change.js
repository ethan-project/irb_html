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
   var t = "<option selected disabled>실험 계획서 항목</option>",
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
         load: (e) => {},
         scroll: (e) => {
            scrollSpy();
         },
      }),
      $("#modal_changeApp").on({
         "shown.bs.modal": (e) => {
            $(".supplement_comment").length &&
               ($(".modal-body").animate({ scrollTop: $(document).height() }, 500),
               setTimeout(() => {
                  $("#review_summary").trigger("click");
               }, 300)),
               $("a.process_content").each((e, t) => {
                  let a = $(t).data("offset") ? $(t).data("offset") : "0, 0",
                     n = $(t).data("tippy-target") ? $(t).data("tippy-target").split(",") : t;
                  $(t).attr("title", "변경 신청 내용이 있습니다."),
                     tippy(t, { offset: a, triggerTarget: n, arrow: !0, placement: "top-end" }),
                     $(t).removeAttr("title"),
                     t._tippy.disable(),
                     $(t).hasClass("supplement") && t._tippy.enable();
               });
         },
      });
}
function cbFuncAfterMapping(e) {
   mappingChangeAppInfo(e), originalJSFunction(), resetSelectAndFileData(e), resetLeftStepNaviIcon();
}
function resetSelectAndFileData(e) {
   function t(e, t) {
      for (let a = 0; a < e.length; ++a) if (e[a].file_idx == t) return e[a];
   }
   var a = e.getChildItemData("ibc_ca_file"),
      n = e.getChildItemData("ibc_ca_item"),
      i = n.dataObj.codes;
   for (let e = 0; e < i.length; ++e) {
      let s = `ibcFile_${i[e].is_seq}`,
         o = {},
         c = {};
      var r = 0,
         l = 1;
      288 == i[e].is_seq && ((r = 2), (l = 3)),
         (o = $("#" + s)
            .parent()
            .parent()
            .parent()
            .children("div")
            .eq(r)),
         (c = $("#" + s)
            .parent()
            .parent()
            .parent()
            .children("div")
            .eq(l)),
         c.attr("id", `ibcOldFile_${i[e].is_seq}`);
      let p = n.dataObj.saved_data.data[String(i[e].is_seq)];
      if (null != p) {
         var d = JSON.parse(p);
         $(`#ibc_${i[e].is_seq}`).parent().find("select").val(`#form_${d[0]}`).prop("selected", !0);
         let n = t(a.dataObj.saved_data.data, i[e].is_seq);
         if (null != n) {
            let e = c.children("div").find("a").first();
            e.attr("href", n.src), e.text(n.org_file_name), o.hide();
         } else c.remove();
      } else c.remove();
   }
}
function mappingChangeAppInfo(e) {
   var t = (n = e.getItemData("ibc_general_experiment")).getStringValue("0"),
      a = [];
   "" != t && (a = t.split(",")), a.length > 0 && (g_min_value = Math.min.apply(null, a));
   {
      var n = e.getItemData("application_info");
      let t = Number(n.getStringValue("approved_dttm"));
      let a = e.getItemData("general_end_date", !1).getStringValue("0");
      $(".task_days").text(getDateDiffWith10(new Date(a), new Date(getDttm(t).dt)));
   }
   n = e.getChildItemData("ca_supplement");
   try {
      $(".supplement_comment").children("div").eq(1).text(n.dataObj.saved_data.data[0]);
   } catch (e) {}
   var i = $("#ibc_ca").children("ul").eq(0).clone(),
      r = $("#ibc_ca").children("ul").eq(1).clone();
   function l(t, a) {
      var n = i.clone();
      let r = `ibc_${t.is_seq}`,
         l = r + "-2",
         d = r + "-3",
         s = `ibcFile_${t.is_seq}`;
      "" != t.link_item_name && n.find("select").attr("data-selected", `form_${t.link_item_name}`),
         n.find("input").first().attr("id", r),
         n.find("input").first().attr("data-target", `#${d}`),
         n.find("label").first().attr("for", r);
      let o = n.children("li").eq(1);
      o.attr("id", d),
         o.children("label").attr("for", l),
         o.children("textarea").eq(0).attr("id", l),
         o.children("div").find("input").attr("name", s),
         o.children("div").find("input").attr("id", s),
         o.children("div").find("label").eq(3).attr("for", s);
      var c = o.children("table").eq(0).children("tbody"),
         p = o.children("table").eq(1).children("tbody"),
         _ = c.children("tr").eq(0).clone(),
         u = p.children("tr").eq(0).clone();
      if ((c.empty(), p.empty(), null != a[String(t.is_seq)])) {
         var f = JSON.parse(a[String(t.is_seq)]);
         o.children("textarea").eq(0).text(f[2]);
      }
      var m,
         h = [],
         g = [],
         v = "general_director",
         b = {},
         q = e.getChildItemData(v, !1),
         S = e.getItemData(v);
      (b = q.dataObj.saved_data.data[0] ? q : S),
         h.push(S.dataObj.saved_data.data[0]),
         g.push(b.dataObj.saved_data.data[0]),
         (m = b.dataObj.saved_data.data[0].user_seq),
         (g_old_member_app_seq = S.dataObj.saved_data.data[0].app_seq);
      (v = "general_expt"), (b = {}), (q = e.getChildItemData(v, !1)), (S = e.getItemData(v));
      b = q.dataObj.saved_data.data[0] ? q : S;
      for (var y = 0; y < Object.keys(S.dataObj.saved_data.data).length; y++) h.push(S.dataObj.saved_data.data[y]);
      for (y = 0; y < Object.keys(b.dataObj.saved_data.data).length; y++) g.push(b.dataObj.saved_data.data[y]);
      return (
         $.each(h, function (e, t) {
            var a = _.clone(),
               n = t.info;
            let i = "";
            if (t.edu_course) {
               let a = JSON.parse(t.edu_course);
               for (e in a) i += `\n          <div class="flexMid">\n            <span class="w120 mRight5">${e}</span>\n            <div>${a[e]}</div>\n          </div>`;
            }
            m == t.user_seq
               ? a.children("td").eq(0).find("input").prop("checked", !0)
               : (a.children("td").eq(0).find("input").prop("checked", !1), a.children("td").eq(0).find("input").prop("disabled", !0)),
               1 == t.animal_mng_flag && m != t.user_seq
                  ? a.children("td").eq(1).find("input").prop("checked", !0)
                  : (a.children("td").eq(1).find("input").prop("checked", !1), a.children("td").eq(1).find("input").prop("disabled", !0)),
               a.children("td").eq(2).text(n.name),
               a.children("td").eq(3).text(n.dept_str),
               a.children("td").eq(4).append(i),
               c.append(a);
         }),
         $.each(g, function (e, t) {
            var a = u.clone(),
               n = t.info;
            m == t.user_seq
               ? (a.children("td").eq(0).find("input").prop("checked", !0), a.children("td").eq(0).find("input").prop("checked", !0))
               : a.children("td").eq(0).find("input").prop("checked", !1),
               1 == t.animal_mng_flag && m != t.user_seq ? a.children("td").eq(1).find("input").prop("checked", !0) : a.children("td").eq(1).find("input").prop("checked", !1),
               a.attr("data-id", n.user_seq),
               a.children("td").eq(2).text(n.name),
               a.children("td").eq(3).text(n.dept_str),
               a.children("td").eq(4).append(makeCourseList(n.user_seq)),
               p.append(a),
               $.each(g_all_userlist, function (e, a) {
                  t.user_seq == a.user_seq && (g_all_userlist[e] = t);
               });
         }),
         n
      );
   }
   function d(t, a) {
      var n = r.clone();
      let i = `ibc_${t.is_seq}`,
         l = i + "-1",
         d = i + "-2",
         s = i + "-3",
         o = `ibcFile_${t.is_seq}`;
      n.find("label").eq(0).text(t.contents),
         "" != t.link_item_name ? n.find("select").attr("data-selected", `form_${t.link_item_name}`) : (n.find("select").attr("data-selected", !1), n.find(".btn_reset_select").remove()),
         n.find("input").first().attr("id", i),
         n.find("input").first().attr("data-target", `#${s}`),
         n.find("label").first().attr("for", i);
      let c = n.children("li").eq(1);
      if (
         (c.attr("id", s),
         c.children("label").attr("for", l),
         c.children("textarea").eq(0).attr("id", l),
         c.children("textarea").eq(1).attr("id", d),
         c.find("input").attr("name", o),
         c.find("input").attr("id", o),
         c.find("label").eq(3).attr("for", o),
         null != a[String(t.is_seq)])
      ) {
         var p = JSON.parse(a[String(t.is_seq)]);
         c.children("textarea").eq(0).text(p[1]), c.children("textarea").eq(1).text(p[2]);
      }
      return (
         288 == t.is_seq &&
            (c.find("label").eq(0).text("변경 전:"),
            (function (e, t) {
               var a = e.getItemData("application_info");
               let n = 1e3 * Number(a.getStringValue("approved_dttm"));
               var i = getTimeStrFromDateType(new Date(n)),
                  r = e.getChildItemData("general_end_date", !1).getStringValue("0"),
                  l = e.getItemData("general_end_date"),
                  d = l.getStringValue("0"),
                  s = new Date();
               s.setDate(s.getDate() + 1);
               var o = moment(s).format("YYYY-MM-DD");
               "" == r && (r = o), (g_old_end_date_app_seq = l.dataObj.saved_data.data.cur_app_seq);
               var c = `<div class="bullet mTop5">${i} ~ ${d}</div>\n                  <label class="flex_1 mTop20">변경 후:</label>\n                  <div class="form-group flexMid">\n                  <label for="general_end_date_ca" class="bullet mRight10 mBot0">${i} ~ </label>\n                  <input class="form-control wCalendar" type="date" value="${r}" id="general_end_date_ca" min="${o}">\n                </div>`;
               t.replaceWith(c);
            })(e, c.children("textarea").eq(0))),
         n
      );
   }
   n = e.getChildItemData("ibc_ca_item");
   $("#ibc_ca").empty();
   var s = n.dataObj.codes;
   if (g_min_value > 8) {
      var o = [];
      g_min_value < 11 ? (o = [289]) : g_min_value >= 11 && (o = [289, 290, 291]);
      for (let e = 0; e < o.length; e++) {
         var c = s.findIndex(function (t) {
            return t.is_seq === o[e];
         });
         c > -1 && s.splice(c, 1);
      }
   }
   var p = n.dataObj.saved_data.data;
   for (let e = 0; e < s.length; ++e) 286 != s[e].is_seq ? $("#ibc_ca").append(d(s[e], p)) : $("#ibc_ca").append(l(s[e], p));
   setTriggerCheckBoxOnChange(), setTriggerCheckBoxOnClickKeydownChange();
   for (const e of g_all_userlist) e.edu_course && mappingCourseList(e.info.user_seq, JSON.parse(e.edu_course));
   var _ = (n = e.getChildItemData("ibc_ca_item")).dataObj.saved_data.data;
   s = n.dataObj.codes;
   for (let e = 0; e < s.length; ++e)
      if (null != _[String(s[e].is_seq)]) {
         let t = `ibc_${s[e].is_seq}`;
         $("#" + t).click();
      }
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
            var t = $(this).parent().find("input").attr("name");
            if (null == t) {
               let e = $(this).parent();
               e.parent().children("div").eq(0).show(), e.parent().children("div").eq(2).show(), e.remove();
            } else {
               var a = `<input type="file" name="${t}" class="custom-file-input" id="${t}">`;
               $(this).parent().find("input").replaceWith(a), $(this).parent().find("label").eq(1).text("");
            }
         },
      });
   var u = JSON.parse(COMM.getCookie(CONST.COOKIE.IPSAP_USER_INFO));
   $(".reviewer").text(spaceAddStr(u.name)), $(".card").removeClass("hidden");
}
function resetLeftStepNaviIcon() {
   var e = {},
      t = g_AppItemParser.getChildItemData("ibc_ca_item").dataObj.codes;
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
      a = `<tr data-id='${
         e.user_seq
      }'>\n                <td>\n                  <div class="custom-radio">\n                    <input type="radio" name="task_director_new" value="1">\n                  </div>\n                </td>\n                <td>\n                  <div class="custom-radio">\n                    <input type="radio" name="exp_manager_new" value="1">\n                  </div>\n                </td>\n                <td>${
         t.name
      }</td>\n                <td>${t.dept_str}</td>\n                <td>${makeCourseList(
         t.user_seq
      )}</td>\n                <td>\n                  <a href="javascript:void(0);" class="btn btn-outline-danger btn_xxs btn-staff-delete">삭제</a>\n                </td>\n              </tr>`;
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
   var r = g_AppItemParser.getChildItemData("ibc_ca_item").dataObj.codes;
   for (let f = 0; f < r.length; ++f) {
      let m = `ibc_${r[f].is_seq}`,
         h = $(`#${m}-1`).val(),
         g = $(`#${m}-2`).val();
      if (1 != $(`#${m}`).is(":checked")) continue;
      var l = !0;
      if ((286 == r[f].is_seq ? (h = g_old_member_app_seq) : 288 == r[f].is_seq && (h = g_old_end_date_app_seq), ("" != h && "" != g) || (l = !1), !l))
         return alert("체크한 항목의 변경 내용 및 변경 사유를 반드시 작성해 주세요."), void $(`#${m}`).focus();
      if (286 == r[f].is_seq) {
         var d = $("#changeUserTbody").children("tr"),
            s = 0,
            o = 0;
         for (let e = 0; e < d.length; e++) {
            var c = 0,
               p = d.eq(e).children("td"),
               _ = d.eq(e).data("id"),
               u = { animal_mng_flag: 0, exp_year_code: 0, user_seq: _, edu_course: makeEduCourse(_) };
            if (
               (p.eq(0).children("div").children('input[name="task_director_new"]').is(":checked")
                  ? (s++, c++, a.push(u))
                  : p.eq(1).children("div").children('input[name="exp_manager_new"]').is(":checked")
                  ? (o++, c++, (u.animal_mng_flag = 1), i.push(u))
                  : i.push(u),
               2 == c)
            )
               return void alert("연구 책임자 와 실험담당자는 동시에 선택 될수 없습니다.");
         }
         if (1 != s) return void alert("연구 책임자는 필수 데이터 입니다.");
         if (1 != o) return void alert("실험 동물 관리 담당자는 필수 데이터 입니다.");
      } else 288 == r[f].is_seq && (n[0] = $("#general_end_date_ca").val());
      if (null == $(`#${m}`).parent().find("select").val()) return void alert("실험 계획서 항목을 선택 해주시기 바랍니다.");
      let v = $(`#${m}`).parent().find("select").val().split("_")[1];
      if (((t[String(r[f].is_seq)] = JSON.stringify([v, h, g])), $(`#ibcOldFile_${r[f].is_seq}`).length > 0)) e[`ibc_ca_file_${r[f].is_seq}`] = String(r[f].is_seq);
      else {
         let t = `ibcFile_${r[f].is_seq}`;
         null != $("#" + t).prop("files")[0] && (e[`ibc_ca_file_${r[f].is_seq}`] = $("#" + t).prop("files")[0]);
      }
   }
   var f = new ItemParser(g_AppInfo.childAppSeq);
   (f.targetSavedItemNames = []),
      (f.moreSaveFiles = e),
      f.addMoreSaveTag("ibc_ca_item", { data: t }),
      a.length > 0 && f.addMoreSaveTag("general_director", { data: a }),
      i.length > 0 && f.addMoreSaveTag("general_expt", { data: i }),
      null != n[0] && f.addMoreSaveTag("general_end_date", { data: n }),
      appItemSubmitWithParser(f, onCompleteSubmit, IPSAP.APP_SUBMIT_TYPE_PARAM.CHILD_SUBMIT);
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
         remakeApplicationInfoReadOnly([], cbFuncAfterMapping, ["ibc_change_application", "ca_supplement", "general_end_date"]),
            $(".apptype_title").text(g_AppInfo.appObj.application_type_str),
            $(".task_title_ko").text(g_AppInfo.appObj.name_ko),
            $(".task_title_en").text(g_AppInfo.appObj.name_en),
            $(".task_director_all").text(`(연구책임자 : ${g_AppInfo.appObj.user_name})`),
            $(".task_director").text(g_AppInfo.appObj.user_name),
            $(".task_ditrect_dept").text(g_AppInfo.appObj.user_dept_str),
            $(".task_director_contact").text(g_AppInfo.appObj.user_phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
            $(".judge_type_str").text(g_AppInfo.appObj.judge_type_str),
            $(".today_str").text(getTimeStrFromDateType(new Date()));
      }),
      $(document).on("click", ".btn-staff-delete", function () {
         $(this).closest("tr").remove();
      });
}),
   $("#modal_changeApp").on("hidden.bs.modal", (e) => {
      $(e.currentTarget).modal("show");
   });
