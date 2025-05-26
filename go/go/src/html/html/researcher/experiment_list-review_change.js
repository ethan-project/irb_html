"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   document.write("<script src='./application_new-0.js'></script>");
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
            $(".modal_hint").hide(),
               e &&
                  ($(".modal[data-backdrop=false]").css({ "background-color": "rgba(0,0,0, 0.0)", "pointer-events": "none" }),
                  $(".modal-dialog-centered.ui-draggable").css({ "justify-content": "flex-start" }),
                  "click" == t.type && $(this).css({ "margin-right": "0", left: "0", top: "0" })),
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
   {
      var t = e.getItemData("application_info");
      let a = Number(t.getStringValue("approved_dttm"));
      $(".apptype_date").text("승인:" + getDttm(a).dt);
   }
   mappingChangeAppInfo(e), originalJSFunction(), resetSelectAndFileData(e), resetLeftStepNaviIcon();
}
$(function () {
   loadApplicationParams(),
      $("#all_app_r").load("/html/common/inc_application/all_app_readonly.html", function () {
         $(".apptype_title").text("신규 승인"),
            IPSAP.DEMO_MODE
               ? originalJSFunction()
               : (remakeApplicationInfoReadOnly([], cbFuncAfterMapping, ["change_application", "ca_supplement", "general_end_date", "animal_type_final", "animal_type"]),
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
                 })());
      }),
      $(document).on("click", ".btn-staff-delete", function () {
         $(this).closest("tr").remove();
      });
});
const animal_is_seq = 104,
   member_is_seq = 111,
   end_date_is_seq = 112;
function resetSelectAndFileData(e) {
   function t(e, t) {
      for (let a = 0; a < e.length; ++a) if (e[a].file_idx == t) return e[a];
   }
   var a = [
      ["ca_regular_file", "ca_regular_item"],
      ["ca_fast_file", "ca_fast_item"],
   ];
   for (let d = 0; d < a.length; ++d) {
      var n = e.getChildItemData(a[d][0]),
         i = e.getChildItemData(a[d][1]),
         l = i.dataObj.codes;
      for (let e = 0; e < l.length; ++e) {
         let a = `regFile_${l[e].is_seq}`,
            d = {},
            p = {};
         var r = 0,
            o = 1;
         112 == l[e].is_seq && ((r = 2), (o = 3)),
            (d = $("#" + a)
               .parent()
               .parent()
               .parent()
               .children("div")
               .eq(r)),
            (p = $("#" + a)
               .parent()
               .parent()
               .parent()
               .children("div")
               .eq(o)),
            p.attr("id", `regOldFile_${l[e].is_seq}`);
         let c = i.dataObj.saved_data.data[String(l[e].is_seq)];
         if (null != c) {
            var s = JSON.parse(c);
            $(`#regular_${l[e].is_seq}`).parent().find("select").val(`#form_${s[0]}`).prop("selected", !0);
            let a = t(n.dataObj.saved_data.data, l[e].is_seq);
            if (null != a) {
               let e = p.children("div").find("a").first();
               e.attr("href", a.src), e.text(a.org_file_name), d.hide();
            } else p.remove();
         } else p.remove();
      }
   }
}
function mappingChangeAppInfo(e) {
   var t = e.getChildItemData("ca_supplement");
   try {
      $(".supplement_comment").children("div").eq(1).text(t.dataObj.saved_data.data[0]);
   } catch (e) {}
   function a() {
      return '\n    <tr>\n      <td class="bullet pLeft10">미등록</td>\n      <td class="pLeft10"><span class="unit mRight5">(M/F)</span><span class="data"><span>-</span><span class="unit mHside5">/</span><span>-</span></span></td>\n      <td class="pHside20"><i class="fas fa-arrow-right"></i></td>\n      <td class="flexMid">\n        <input type="text" size="20" maxlength="20" class="form-control form-control-sm w100" value="">\n        <input type="number" min=0 size="4" maxlength="4" class="form-control form-control-sm text-center w60 mLeft5" value="0"><span class="unit mHside5">/</span>\n        <input type="number" min=0 size="4" maxlength="4" class="form-control form-control-sm text-center w60" value="0">\n        <a href="javascript:void(0);" class="btn btn_add_species btn-outline-primary mLeft10"><i class="fas fa-plus"></i></a>\n        <a href="javascript:void(0);" class="btn btn_remove_row btn-outline-danger mLeft3"><i class="far fa-trash-alt"></i></a>\n      </td>\n    </tr>';
   }
   var n = $("#ca_regular").children("ul").first().clone(),
      i = $("#ca_fast").children("ul").first().clone();
   function l(t, a) {
      var n = i.clone();
      let l = `regular_${t.is_seq}`,
         r = l + "-2",
         o = l + "-3",
         s = `regFile_${t.is_seq}`;
      n.find("span").eq(1).text(t.contents),
         "" == t.link_item_name ? n.find("select").attr("data-selected", "") : n.find("select").attr("data-selected", `form_${t.link_item_name}`),
         n.find("input").first().attr("id", l),
         n.find("input").first().attr("data-target", `#${o}`),
         n.find("label").first().attr("for", l);
      let d = n.children("li").eq(1);
      if ((d.removeClass("regular_01"), d.attr("id", o), null != a[String(t.is_seq)])) {
         var p = JSON.parse(a[String(t.is_seq)]);
         d.children("textarea").eq(0).text(p[2]);
      }
      d.children("label").attr("for", r),
         d.children("textarea").eq(0).attr("id", r),
         d.children("div").find("input").attr("name", s),
         d.children("div").find("input").attr("id", s),
         d.children("div").find("label").eq(3).attr("for", s);
      var c = d.children("table").eq(0).children("tbody"),
         _ = d.children("table").eq(1).children("tbody"),
         f = c.children("tr").eq(0).clone(),
         m = _.children("tr").eq(0).clone();
      c.empty(), _.empty();
      var u,
         h = [],
         g = [],
         v = "general_director",
         b = {},
         q = e.getChildItemData(v, !1),
         y = e.getItemData(v);
      (b = q.dataObj.saved_data.data[0] ? q : y),
         h.push(y.dataObj.saved_data.data[0]),
         g.push(b.dataObj.saved_data.data[0]),
         (u = b.dataObj.saved_data.data[0].user_seq),
         (g_old_member_app_seq = y.dataObj.saved_data.data[0].app_seq);
      (v = "general_expt"), (b = {}), (q = e.getChildItemData(v, !1)), (y = e.getItemData(v));
      b = q.dataObj.saved_data.data[0] ? q : y;
      for (var x = 0; x < Object.keys(y.dataObj.saved_data.data).length; x++) h.push(y.dataObj.saved_data.data[x]);
      for (x = 0; x < Object.keys(b.dataObj.saved_data.data).length; x++) g.push(b.dataObj.saved_data.data[x]);
      return (
         $.each(h, function (e, t) {
            var a = f.clone(),
               n = t.info;
            u == t.user_seq
               ? a.children("td").eq(0).find("input").prop("checked", !0)
               : (a.children("td").eq(0).find("input").prop("checked", !1), a.children("td").eq(0).find("input").prop("disabled", !0)),
               1 == t.animal_mng_flag && u != t.user_seq
                  ? a.children("td").eq(1).find("input").prop("checked", !0)
                  : (a.children("td").eq(1).find("input").prop("checked", !1), a.children("td").eq(1).find("input").prop("disabled", !0)),
               a.children("td").eq(2).text(n.name),
               a.children("td").eq(3).text(n.dept_str),
               a.children("td").eq(4).text(n.position_str),
               a.children("td").eq(5).text(n.major_field_str),
               a.children("td").eq(6).text(n.tmp_phoneno),
               a.children("td").eq(7).text(t.exp_year_code_str),
               c.append(a);
         }),
         $.each(g, function (e, t) {
            var a = m.clone(),
               n = t.info;
            u == t.user_seq
               ? (a.children("td").eq(0).find("input").prop("checked", !0), a.children("td").eq(0).find("input").prop("checked", !0))
               : a.children("td").eq(0).find("input").prop("checked", !1),
               1 == t.animal_mng_flag && u != t.user_seq ? a.children("td").eq(1).find("input").prop("checked", !0) : a.children("td").eq(1).find("input").prop("checked", !1),
               a.attr("data-id", n.user_seq),
               a.children("td").eq(2).text(n.name),
               a.children("td").eq(3).text(n.dept_str),
               a.children("td").eq(4).text(n.position_str),
               a.children("td").eq(5).text(n.major_field_str),
               a.children("td").eq(6).text(n.tmp_phoneno),
               a.children("td").eq(7).children("div").children("select").val(t.exp_year_code).prop("selected", !0),
               _.append(a),
               $.each(g_all_userlist, function (e, a) {
                  t.user_seq == a.user_seq && (g_all_userlist[e] = t);
               });
         }),
         n
      );
   }
   function r(t, a, i, l, r) {
      var o = n.clone();
      o.removeClass("mTop30"),
         o.addClass(l),
         o.find("span").eq(0).removeClass("blue_deep"),
         o.find("span").eq(0).addClass(t),
         o
            .find("span")
            .eq(0)
            .text("[" + a + "]");
      let s = `regular_${i.is_seq}`,
         d = s + "-1",
         p = s + "-2",
         c = `regFile_${i.is_seq}`;
      o.find("span").eq(1).text(i.contents),
         "" == i.link_item_name ? o.find("select").attr("data-selected", "") : o.find("select").attr("data-selected", `form_${i.link_item_name}`),
         o.find("input").first().attr("id", s),
         o.find("input").first().attr("data-target", `.${s}`),
         o.find("label").first().attr("for", s);
      let _ = o.children("li").eq(1);
      if (
         (_.removeClass("regular_01"),
         _.addClass(s),
         _.children("label").attr("for", d),
         _.children("textarea").eq(0).attr("id", d),
         _.children("textarea").eq(1).attr("id", p),
         _.find("input").attr("name", c),
         _.find("input").attr("id", c),
         _.find("label").eq(3).attr("for", c),
         null != r[String(i.is_seq)])
      ) {
         var f = JSON.parse(r[String(i.is_seq)]);
         _.children("textarea").eq(0).text(f[1]), _.children("textarea").eq(1).text(f[2]);
      }
      if (104 == i.is_seq) {
         let t;
         try {
            t = JSON.parse(JSON.parse(r[String(i.is_seq)])[1]);
         } catch (e) {}
         !(function (e, t, a) {
            var n = "<table><tbody>";
            if (null == t) {
               var i = a.getItemData("animal_type_final");
               0 == i.dataObj.saved_data.data.length && (i = a.getItemData("animal_type"));
               var l = i.dataObj.saved_data.data;
               for (let e = 0; e < l.length; ++e)
                  n += `\n          <tr>\n            <td class="bullet pLeft10"><span class="data">${l[e].animal_code_str}</span></td>\n            <td class="pLeft10"><span class="unit mRight5">(M/F)</span><span class="data"><span>${l[e].male_cnt}</span><span class="unit mHside5">/</span><span>${l[e].female_cnt}</span></span></td>\n            <td class="pHside20"><i class="fas fa-arrow-right"></i></td>\n            <td class="flexMid">\n              <input type="text" size="20" maxlength="20" class="form-control form-control-sm w100" value="${l[e].animal_code_str}">\n              <input type="number" min="0" size="4" maxlength="4" class="form-control form-control-sm text-center w60 mLeft5" value="${l[e].male_cnt}"><span class="unit mHside5">/</span>\n              <input type="number" min="0" size="4" maxlength="4" class="form-control form-control-sm text-center w60" value="${l[e].female_cnt}">\n              <a href="javascript:void(0);" class="btn btn_add_species btn-outline-primary mLeft10"><i class="fas fa-plus"></i></a>\n              <a href="javascript:void(0);" class="btn btn_remove_row btn-outline-danger mLeft3"><i class="far fa-trash-alt"></i></a>\n            </td>\n          </tr>`;
            } else
               for (let e = 0; e < t.length; ++e) {
                  let a = t[e];
                  n += `\n          <tr>\n            <td class="bullet pLeft10"><span class="data">${a[0]}</span></td>\n            <td class="pLeft10"><span class="unit mRight5">(M/F)</span><span class="data"><span>${a[1]}</span><span class="unit mHside5">/</span><span>${a[2]}</span></span></td>\n            <td class="pHside20"><i class="fas fa-arrow-right"></i></td>\n            <td class="flexMid">\n              <input class="form-control form-control-sm hidden mTop5 w100" type="text" value="${a[3]}" style="display: block;">\n              <input type="number" min="0" size="4" maxlength="4" class="form-control form-control-sm text-center w60 mLeft5" value="${a[4]}"><span class="unit mHside5">/</span>\n              <input type="number" min="0" size="4" maxlength="4" class="form-control form-control-sm text-center w60" value="${a[5]}">\n              <a href="javascript:void(0);" class="btn btn_add_species btn-outline-primary mLeft10"><i class="fas fa-plus"></i></a>\n              <a href="javascript:void(0);" class="btn btn_remove_row btn-outline-danger mLeft3"><i class="far fa-trash-alt"></i></a>\n            </td>\n          </tr>`;
               }
            (n += "</tbody></table>"), e.replaceWith(n);
         })(_.children("textarea").eq(0), t, e);
      }
      if (112 == i.is_seq) {
         _.find("label").eq(0).text("변경 전:");
         try {
            !(function (e, t) {
               var a = e.getItemData("application_info");
               let n = 1e3 * Number(a.getStringValue("approved_dttm"));
               var i = getTimeStrFromDateType(new Date(n)),
                  l = e.getChildItemData("general_end_date", !1).getStringValue("0"),
                  r = e.getItemData("general_end_date"),
                  o = r.getStringValue("0"),
                  s = new Date();
               s.setDate(s.getDate() + 1);
               var d = moment(s).format("YYYY-MM-DD");
               "" == l && (l = d), (g_old_end_date_app_seq = r.dataObj.saved_data.data.cur_app_seq);
               var p = `<div class="bullet mTop5">${i} ~ ${o}</div>\n                  <label class="flex_1 mTop20">변경 후:</label>\n                  <div class="form-group flexMid">\n                  <label for="general_end_date_ca" class="bullet mRight10 mBot0">${i} ~ </label>\n                  <input class="form-control wCalendar" type="date" value="${l}" id="general_end_date_ca" min="${d}">\n                </div>`;
               t.replaceWith(p);
            })(e, _.children("textarea").eq(0));
         } catch (e) {}
      }
      return o;
   }
   t = e.getChildItemData("ca_regular_item");
   $("#ca_regular").empty();
   var o = t.dataObj.codes;
   for (let e = 0; e < o.length; ++e) {
      let a = "mTop2";
      0 == e && (a = "mTop30"), $("#ca_regular").append(r("blue_deep", "정규", o[e], a, t.dataObj.saved_data.data));
   }
   t = e.getChildItemData("ca_fast_item");
   $("#ca_fast").empty();
   o = t.dataObj.codes;
   for (let e = 0; e < o.length; ++e) {
      let a = "mTop2";
      0 == e && (a = "mTop30"), 111 != o[e].is_seq ? $("#ca_fast").append(r("green", "신속", o[e], a, t.dataObj.saved_data.data)) : $("#ca_fast").append(l(o[e], t.dataObj.saved_data.data));
   }
   setTriggerCheckBoxOnChange(), setTriggerCheckBoxOnClickKeydownChange();
   var s = (t = e.getChildItemData("ca_regular_item")).dataObj.saved_data.data;
   o = t.dataObj.codes;
   for (let e = 0; e < o.length; ++e)
      if (null == s[String(o[e].is_seq)]) {
         let t = `regular_${o[e].is_seq}`;
         $("#" + t).click(),
            $("#" + t)
               .closest("li")
               .removeClass("active");
      }
   (s = (t = e.getChildItemData("ca_fast_item")).dataObj.saved_data.data), (o = t.dataObj.codes);
   for (let e = 0; e < o.length; ++e)
      if (null == s[String(o[e].is_seq)]) {
         let t = `regular_${o[e].is_seq}`;
         $("#" + t).click(),
            $("#" + t)
               .closest("li")
               .removeClass("active");
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
   var d = JSON.parse(COMM.getCookie(CONST.COOKIE.IPSAP_USER_INFO));
   $(".reviewer").text(spaceAddStr(d.name)),
      (function e() {
         $(".btn_add_species").off(),
            $(".btn_add_species").on({
               click: function (t) {
                  "" != $(this).closest("tbody").children("tr").last().find("input").eq(0).val() && ($(this).closest("tbody").append(a()), e());
               },
            }),
            $(".btn_remove_row").off(),
            $(".btn_remove_row").on({
               click: function (t) {
                  var n = $(this).closest("tbody");
                  $(this).closest("tr").remove(), 0 == n.children("tr").length && (n.append(a()), e());
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
      a = `<tr data-id='${e.user_seq}'>\n                <td>\n                  <div class="custom-radio">\n                    <input type="radio" name="task_director_new" value="1">\n                  </div>\n                </td>\n                <td>\n                  <div class="custom-radio">\n                    <input type="radio" name="exp_manager_new" value="1">\n                  </div>\n                </td>\n                <td>${t.name}</td>\n                <td>${t.dept_str}</td>\n                <td>${t.position_str}</td>\n                <td>${t.major_field_str}</td>\n                <td class="hidden">${t.tmp_phoneno}</td>\n                <td>\n                <div class="btn-group">\n                  <select class="form-control btn-outline-primary small_select narrow_padding">\n                    <option selected disabled>경력 선택</option>\n                    <option value="1">1년 미만</option>\n                    <option value="2">1년 ~ 3년</option>\n                    <option value="3">3년 이상</option>\n                  </select>\n                </div>\n                </td>\n                <td class="hidden">${t.email}</td>\n                <td>\n                  <a href="javascript:void(0);" class="btn btn-outline-danger btn_xxs btn-staff-delete">삭제</a>\n                </td>\n              </tr>`;
   $("#changeUserTbody").append(a),
      e.exp_year_code > 0 && $(`#changeUserTbody > [data-id='${e.user_seq}']`).children("td").eq(7).children("div").children("select").val(e.exp_year_code).prop("selected", !0);
}
function resetLeftStepNaviIcon() {
   var e = {},
      t = 0,
      a = 0,
      n = g_AppItemParser.getChildItemData("ca_regular_item").dataObj.codes;
   for (let a = 0; a < n.length; ++a) {
      let i = `regular_${n[a].is_seq}`;
      if (1 == $(`#${i}`).is(":checked")) {
         try {
            let t = $(`#${i}`).parent().find("select").val().split("_")[1].split("-");
            e[String(Number(t[0]))] = !0;
         } catch (e) {}
         ++t;
      }
   }
   n = g_AppItemParser.getChildItemData("ca_fast_item").dataObj.codes;
   for (let t = 0; t < n.length; ++t) {
      let i = `regular_${n[t].is_seq}`;
      if (1 == $(`#${i}`).is(":checked")) {
         try {
            let t = $(`#${i}`).parent().find("select").val().split("_")[1].split("-");
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
      $(".process_content").each((t, a) => {
         $(a).removeClass("supplement"), a._tippy && a._tippy.disable(), 1 == e[$(a).attr("data-process-num")] && ($(a).addClass("supplement"), a._tippy && a._tippy.enable());
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
      i = {},
      l = [];
   var r = g_AppItemParser.getChildItemData("ca_regular_item").dataObj.codes;
   for (let a = 0; a < r.length; ++a) {
      let n = `regular_${r[a].is_seq}`,
         i = $(`#${n}-1`).val(),
         l = $(`#${n}-2`).val();
      if (1 != $(`#${n}`).is(":checked")) continue;
      if ("" == i || "" == l) return alert("체크한 항목의 변경 내용 및 변경 사유를 반드시 작성해 주세요."), void $(`#${n}`).focus();
      let o = $(`#${n}`).parent().find("select").val().split("_")[1];
      if (104 == r[a].is_seq) {
         let e = [];
         if (
            ($("#" + n)
               .closest("ul")
               .find("tbody")
               .children("tr")
               .each(function (t, a) {
                  let n = $(a).find("input").eq(0).val(),
                     i = Number($(a).find("input").eq(1).val()),
                     l = Number($(a).find("input").eq(2).val());
                  if ("" != n && i + l > 0) {
                     let t = [],
                        r = $(a).children("td").eq(0).children("span");
                     r.length > 0 ? t.push(r.text()) : t.push($(a).children("td").eq(0).text()),
                        t.push($(a).children("td").eq(1).children("span").eq(1).children("span").eq(0).text()),
                        t.push($(a).children("td").eq(1).children("span").eq(1).children("span").eq(2).text()),
                        t.push(n),
                        t.push(i),
                        t.push(l),
                        e.push(t);
                  }
               }),
            0 == e.length)
         )
            return void alert("변경하고는 하는 동물의 종 또는 수량이 없습니다.");
         i = JSON.stringify(e);
      }
      if (((t[String(r[a].is_seq)] = JSON.stringify([o, i, l])), $(`#regOldFile_${r[a].is_seq}`).length > 0)) e[`ca_regular_file_${r[a].is_seq}`] = String(r[a].is_seq);
      else {
         let t = `regFile_${r[a].is_seq}`;
         null != $("#" + t).prop("files")[0] && (e[`ca_regular_file_${r[a].is_seq}`] = $("#" + t).prop("files")[0]);
      }
   }
   r = g_AppItemParser.getChildItemData("ca_fast_item").dataObj.codes;
   for (let t = 0; t < r.length; ++t) {
      let g = `regular_${r[t].is_seq}`,
         v = $(`#${g}-1`).val(),
         b = $(`#${g}-2`).val();
      if (1 != $(`#${g}`).is(":checked")) continue;
      var o = !0;
      if ((111 == r[t].is_seq ? (v = g_old_member_app_seq) : 112 == r[t].is_seq && (v = g_old_end_date_app_seq), ("" != v && "" != b) || (o = !1), !o))
         return alert("체크한 항목의 변경 내용 및 변경 사유를 반드시 작성해 주세요."), void $(`#${g}`).focus();
      if (111 == r[t].is_seq) {
         var s = $("#changeUserTbody").children("tr"),
            d = 0,
            p = 0;
         for (let e = 0; e < s.length; e++) {
            var c = 0,
               _ = s.eq(e).children("td"),
               f = s.eq(e).data("id"),
               m = _.eq(7).children("div").children("select").val();
            if (null == m) {
               var u = _.eq(2).text();
               return void alert(u + "님의 동물 실험 경력 년수를 선택 해주시기 바랍니다.");
            }
            var h = { animal_mng_flag: 0, exp_year_code: m, user_seq: f };
            if (
               (_.eq(0).children("div").children('input[name="task_director_new"]').is(":checked")
                  ? (d++, c++, n.push(h))
                  : _.eq(1).children("div").children('input[name="exp_manager_new"]').is(":checked")
                  ? (p++, c++, (h.animal_mng_flag = 1), l.push(h))
                  : l.push(h),
               2 == c)
            )
               return void alert("연구 책임자 와 실험담당자는 동시에 선택 될수 없습니다.");
         }
         if (1 != d) return void alert("연구 책임자는 필수 데이터 입니다.");
         if (1 != p) return void alert("실험 동물 관리 담당자는 필수 데이터 입니다.");
      } else 112 == r[t].is_seq && (i[0] = $("#general_end_date_ca").val());
      let q = $(`#${g}`).parent().find("select").val().split("_")[1];
      if (((a[String(r[t].is_seq)] = JSON.stringify([q, v, b])), $(`#regOldFile_${r[t].is_seq}`).length > 0)) e[`ca_fast_file_${r[t].is_seq}`] = String(r[t].is_seq);
      else {
         let a = `regFile_${r[t].is_seq}`;
         null != $("#" + a).prop("files")[0] && (e[`ca_fast_file_${r[t].is_seq}`] = $("#" + a).prop("files")[0]);
      }
   }
   var g = new ItemParser(g_AppInfo.childAppSeq);
   (g.targetSavedItemNames = []),
      (g.moreSaveFiles = e),
      g.addMoreSaveTag("ca_regular_item", { data: t }),
      g.addMoreSaveTag("ca_fast_item", { data: a }),
      n.length > 0 && g.addMoreSaveTag("general_director", { data: n }),
      l.length > 0 && g.addMoreSaveTag("general_expt", { data: l }),
      null != i[0] && g.addMoreSaveTag("general_end_date", { data: i }),
      appItemSubmitWithParser(g, onCompleteSubmit, IPSAP.APP_SUBMIT_TYPE_PARAM.CHILD_SUBMIT);
}
function onCompleteSubmit(e, t, a) {
   e ? (window.location.href = "./experiment_list.html") : alert(`변경 승인 신청을 실패했습니다.\n(${t.em})`);
}
$("#modal_changeApp").on("hidden.bs.modal", (e) => {
   $(e.currentTarget).modal("show");
});
