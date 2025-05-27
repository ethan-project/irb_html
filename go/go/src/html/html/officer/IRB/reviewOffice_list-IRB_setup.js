"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>");
var inst_info = JSON.parse(COMM.getCookie("institution_info")),
   inst_seq = null == inst_info ? 0 : inst_info.institution_seq,
   user_info = JSON.parse(COMM.getCookie("user_info")),
   user_seq = null == user_info ? 0 : user_info.user_seq,
   contents = $(".page-content"),
   checked = [];
function loadAppItemData() {
   getCommittee(), $("#btn_setting_finish").addClass("disabled");
   let e = { "filter.query_items": "main_general,general_director,judge_setting,normal_review_result,expert_review_result" };
   g_AppInfo.appObj.application_result != IPSAP.APPLICATION_RESULT.CHECKING_2 && $("#btn_review_retry").remove(),
      g_AppInfo.appObj.application_step > IPSAP.APPLICATION_STEP.PRO_JUDGE && $("#btnPro").remove(),
      g_AppInfo.appObj.application_step != IPSAP.APPLICATION_STEP.CHECK ? $("#btn_setting_finish").remove() : $("#btn_setting_change").remove();
   var t = g_AppInfo.appSeq;
   switch (g_AppInfo.appObj.application_type) {
      case IPSAP.APPLICATION_TYPE.CHANGE:
      case IPSAP.APPLICATION_TYPE.RENEW:
      case IPSAP.APPLICATION_TYPE.BRING:
      case IPSAP.APPLICATION_TYPE.CHECKLIST:
      case IPSAP.APPLICATION_TYPE.END:
         (t = g_AppInfo.childAppSeq), (e["filter.parent_items"] = "general_title,general_end_date,application_info,general_director");
   }
   (g_AppItemParser = new ItemParser(t)), g_AppItemParser.load(e, dataMappingFunc);
}
$(function () {
   loadApplicationParams(), loadAppItemData();
});
var g_already_review_member = [],
   g_already_nomarl_review_member_obj = {},
   g_already_expert_review_member_obj = {};
function dataMappingFunc(e) {
   mappingAppInfo(e), mappingGeneralDirector(e);
   var t = "expert_review_result",
      a = e.getItemData(t, !0);
   try {
      g_already_expert_review_member_obj = a.dataObj.saved_data.data;
   } catch (e) {}
   (t = "normal_review_result"), (a = e.getItemData(t, !0));
   try {
      (g_already_review_member = Object.keys(a.dataObj.saved_data.data)), (g_already_nomarl_review_member_obj = a.dataObj.saved_data.data);
   } catch (e) {}
   {
      t = "expert_member";
      let n = (a = e.getItemData(t, !0)).dataObj.saved_data.data;
      try {
         for (let e = 0; e < n.length; ++e) addMember("pro", { user_seq: n[e].user_seq, name: n[e].info.name, email: n[e].info.email });
      } catch (e) {}
   }
   {
      t = "committee_in_member";
      let n = (a = e.getItemData(t, !0)).dataObj.saved_data.data;
      try {
         for (let e = 0; e < n.length; ++e) addMember("norm_in", { user_seq: n[e].user_seq, name: n[e].info.name, email: n[e].info.email });
      } catch (e) {}
   }
   {
      t = "committee_ex_member";
      let n = (a = e.getItemData(t, !0)).dataObj.saved_data.data;
      try {
         for (let e = 0; e < n.length; ++e) addMember("norm_out", { user_seq: n[e].user_seq, name: n[e].info.name, email: n[e].info.email });
      } catch (e) {}
   }
   {
      t = "judge_expert_deadline";
      var n = (a = e.getItemData(t, !0)).getStringValue("0");
      let r = new Date();
      r.setDate(r.getDate() + 1);
      let i = r.toISOString().slice(0, 11) + "00:00:00";
      "" == n ? $("#date_1").val(i) : $("#date_1").val(n),
         $("#date_1").on("change", function () {
            let e = $("#date_1").val(),
               t = new Date(e),
               a = new Date();
            a.getTime() > t.getTime() && $("#date_1").val(moment(a).format("YYYY-MM-DD HH:mm"));
         }),
         Object.keys(g_already_expert_review_member_obj).length > 0 && $("#date_1").attr("disabled", !0);
   }
   t = "judge_normal_deadline";
   if ("" == (n = (a = e.getItemData(t, !0)).getStringValue("0"))) $("#date_2").val("7");
   else {
      var r = n;
      if (n.length > 1) {
         "H" == n.slice(-1) && $(".small_select option:eq(1)").prop("selected", !0), (r = n.substring(0, n.length - 1));
      }
      $("#date_2").val(Number(r));
   }
   t = "judge_alarm";
   (a = e.getItemData(t)).makeHtmlCheckList(), $(".card").removeClass("hidden");
}
function mappingAppInfo(e) {
   contents.find(".alt_title").text(`${g_AppInfo.appObj.sub_title}`),
      contents.find("span.title_desc").text(`(연구 책임자 : ${g_AppInfo.appObj.user_name})`),
      contents.find("span.data").eq(0).text(g_AppInfo.appObj.name_ko),
      contents.find("span.data").eq(1).text(g_AppInfo.appObj.name_en),
      $("#date_1").val(new Date().toISOString().slice(0, 11) + "00:00:00"),
      $("#date_2").val("7");
   let t = contents.find("span.data").eq(2);
   if (e.hasParentItem()) {
      var a;
      {
         var n = e.getParentItemData("application_info");
         let t = Number(n.getStringValue("approved_dttm"));
         a = getDttm(t).dt + " ~ ";
      }
      n = e.getParentItemData("general_end_date");
      t.text(a + n.getStringValue("0"));
   } else {
      n = e.getItemData("general_end_date");
      t.text("위원회 승인일 ~ " + n.getStringValue("0"));
   }
}
function mappingGeneralDirector(e) {
   var t = null;
   if (null == (t = g_AppInfo.appObj.application_type == IPSAP.APPLICATION_TYPE.NEW ? e.getItemData("general_director") : e.getParentItemData("general_director")).dataObj.saved_data) return;
   t.initMembersIRB();
   let a = contents.find(".info_table").eq(0).find("tbody"),
      n = t.dataObj.saved_data.data[0];
   (n.info.tmp_phoneno = n.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")), (n.info.tmp_edu_date = n.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3"));
   let r = "";
   if ("" != n.edu_course) {
      let e = JSON.parse(n.edu_course);
      for (key in e) r += `\n      <div class="flexMid">\n        <span class="min_w120">${key}</span>\n        <span class="mLeft10">${e[key]}</span>\n      </div>`;
   }
   a.find("tr").eq(0).find("td.data").eq(0).text(n.info.name),
      a.find("tr").eq(0).find("td.data").eq(1).text(n.info.dept_str),
      a.find("tr").eq(0).find("td.data").eq(2).text(n.info.tmp_phoneno),
      a.find("tr").eq(1).find("td.data").eq(0).text(n.info.edu_institution_str),
      a.find("tr").eq(1).find("td.data").eq(1).text(n.info.position_str),
      a.find("tr").eq(1).find("td.data").eq(2).text(n.info.email),
      a.find("tr").eq(2).find("td.data").eq(0).text(n.info.tmp_edu_date),
      a.find("tr").eq(2).find("td.data").eq(1).text(n.info.major_field_str),
      a.find("tr").eq(2).find("td.data").eq(2).append(r);
}
function getCommittee() {
   var tbody_name = "committee",
      html = `<tbody id="${tbody_name}">`;
   API.load({
      url: eval("`" + CONST.API.INSTITUTION.USER + "`"),
      type: CONST.API_TYPE.GET,
      data: { "filter.user_type": 4 },
      success: function (e) {
         $.each(e, function (e, t) {
            if (!g_AppItemParser.hasMemberExists(t.user_seq)) {
               var a = `<div class="checkbox checkbox-primary check-list">\n                <input type="checkbox" value="1" id="check_${t.user_seq}">\n                <label for="check_${t.user_seq}"></label>\n             </div>`;
               (t.tmp_phoneno = t.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
                  (t.tmp_edu_date = t.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3")),
                  (html += `\n              <tr data-user_seq="${t.user_seq}">\n                <td>${t.name}</td>\n                <td>${t.dept_str}</td>\n                <td>${t.position_str}</td>\n                <td>${t.major_field_str}</td>\n                <td>${t.tmp_phoneno}</td>\n                <td>${t.email}</td>\n                <td>${t.tmp_edu_date}</td>\n                <td>${t.edu_institution_str}</td>\n                <td>${t.edu_course_num}</td>\n                <td class="text-center">${a}</td>\n              </tr>`);
            }
         }),
            (html += "</tbody>"),
            $("#committee").replaceWith(html);
      },
      complete: function () {
         setProCommittee(), setNormCommittee();
      },
      error: function (e) {},
   });
}
function setProCommittee(e) {
   $("#btnPro").on("click", function () {
      if (0 == $(".check-list input:checkbox:checked").length) return void alert("배정을 원하는 심사위원을 선택하세요.");
      if ($(".check-list input:checkbox:checked").length > 1 && $(".check-list input:checkbox:disabled").length < 1) return void $("#modal_pro_check").modal("show");
      if ($(".member_pro").length > 0) return void $("#modal_pro_check").modal("show");
      try {
         var t = e.getItemData("expert_review_dttm", "nosave").getStringValue("0");
      } catch (e) {}
      let a = t;
      $(".check-list input:checkbox:checked").each(function (e, t) {
         let n = $(t),
            r = { user_seq: n.closest("tr").data("user_seq"), name: n.closest("tr").children("td").eq(0).text(), email: n.closest("tr").children("td").eq(5).text() };
         n.prop("disabled") || addMember("pro", r, a), $("#btnPass").attr("disabled", !0), $("#pass_comment").removeClass("box_shadow_red_on");
      });
   });
}
function setNormCommittee(e) {
   $("#btnNorm_in").on("click", function () {
      if (0 == $(".check-list input:checkbox:checked").length) return void alert("배정을 원하는 심사위원을 선택하세요.");
      try {
         var t = e.getItemData("normal_review_dttm", "nosave").getStringValue("0");
      } catch (e) {}
      $(".check-list input:checkbox:checked").each(function (e, a) {
         let n = $(a),
            r = { user_seq: n.closest("tr").data("user_seq"), name: n.closest("tr").children("td").eq(0).text(), email: n.closest("tr").children("td").eq(5).text() },
            i = t;
         n.prop("disabled") || addMember("norm_in", r, i);
      });
   }),
      $("#btnNorm_out").on("click", function () {
         if (0 == $(".check-list input:checkbox:checked").length) return void alert("배정을 원하는 심사위원을 선택하세요.");
         try {
            var t = e.getItemData("normal_review_dttm").getStringValue("0");
         } catch (e) {}
         $(".check-list input:checkbox:checked").each(function (e, a) {
            let n = $(a),
               r = { user_seq: n.closest("tr").data("user_seq"), name: n.closest("tr").children("td").eq(0).text(), email: n.closest("tr").children("td").eq(5).text() },
               i = t;
            n.prop("disabled") || addMember("norm_out", r, i);
         });
      });
}
function addMember(e, t) {
   var a = `<tr class="member_${e}">\n    <td>${t.name}</td>\n    <td class="text-left">${t.email}</td>\n    <td data-user_seq="${t.user_seq}">\n      <button type="button" class="btn btn-danger btn_xxs noMinWidth coding_btn_delete">취소</button>\n    </td>\n    </tr>`;
   if (
      ($(`#${e}`).append(a),
      memberListSetDisable(t.user_seq),
      "pro" == e && g_already_expert_review_member_obj[0] && ($(`#${e}`).find("button").last().attr("disabled", !0), Object.keys(g_already_expert_review_member_obj).length > 0))
   ) {
      var n = $(`#${e}`).find("button").last(),
         r = "";
      1 == g_already_expert_review_member_obj[0].select_ids[0] || 2 == g_already_expert_review_member_obj[0].select_ids[0]
         ? ((r = "찬성"), n.removeClass("btn-danger").addClass("btn-outline-primary"))
         : ((r = "반대"), n.removeClass("btn-danger").addClass("btn-outline-danger")),
         n.text(r);
   }
   if (isInArray(g_already_review_member, String(t.user_seq)) && ($(`#${e}`).find("button").last().attr("disabled", !0), Object.keys(g_already_nomarl_review_member_obj).length > 0)) {
      (n = $(`#${e}`).find("button").last()), (r = "");
      1 == g_already_nomarl_review_member_obj[t.user_seq].select_ids[0]
         ? ((r = "찬성"), n.removeClass("btn-danger").addClass("btn-outline-primary"))
         : ((r = "반대"), n.removeClass("btn-danger").addClass("btn-outline-danger")),
         n.text(r);
   }
   $(".coding_btn_delete").off("click"),
      $(".coding_btn_delete").on("click", function () {
         $(this).parent().parent().parent().attr("id");
         var e = $(this).parent().attr("data-user_seq");
         $(this).parent().parent().remove(),
            memberListSetEnable(e),
            $(this).closest("tr.member_pro").length && ($("#btnPro, #btnPass").attr("disabled", !1), $("#pass_comment").removeClass("box_shadow_red_on"));
      }),
      resetFinishBtnStatus();
}
function memberListSetDisable(e) {
   var t = `check_${e}`;
   $("#" + t).prop("checked", !0), $("#" + t).prop("disabled", !0);
}
function memberListSetEnable(e) {
   var t = `check_${e}`;
   $("#" + t).prop("checked", !1), $("#" + t).prop("disabled", !1), resetFinishBtnStatus();
}
function resetFinishBtnStatus() {
   setMemberCount(),
      0 != $("#pro").children("tr").length && $("#norm_in").children("tr").length + $("#norm_out").children("tr").length != 0
         ? $("#btn_setting_finish").removeClass("disabled")
         : $("#btn_setting_finish").addClass("disabled");
}
function onSettingFinish(e) {
   if (0 != $(".member_pro").length)
      if ($(".member_norm_in").length + $(".member_norm_out").length != 0) {
         if (IPSAP.DEMO_MODE) return (window.location.href = "./application_list.html"), void $("#committee").removeClass("hidden");
         g_AppItemParser.targetSavedItemNames = ["judge_alarm"];
         {
            let e = [];
            $(".member_pro").each(function (t, a) {
               e.push({ user_seq: $(a).children("td").eq(2).data("user_seq") });
            }),
               g_AppItemParser.addMoreSaveTag("expert_member", { data: e });
         }
         {
            let e = [];
            $(".member_norm_in").each(function (t, a) {
               e.push({ user_seq: $(a).children("td").eq(2).data("user_seq") });
            }),
               g_AppItemParser.addMoreSaveTag("committee_in_member", { data: e });
         }
         {
            let e = [];
            $(".member_norm_out").each(function (t, a) {
               e.push({ user_seq: $(a).children("td").eq(2).data("user_seq") });
            }),
               g_AppItemParser.addMoreSaveTag("committee_ex_member", { data: e });
         }
         g_AppItemParser.addMoreSaveTag("judge_expert_deadline", { data: { 0: $("#date_1").val() } });
         var t = $("#date_2").val() + $(".small_select").val();
         g_AppItemParser.addMoreSaveTag("judge_normal_deadline", { data: { 0: t } }),
            e ? appItemSubmit(onCompleteSettingFinish, IPSAP.APP_SUBMIT_TYPE_PARAM.CHECKING_2) : appItemSaveTemporary(onCompleteChangeSettingFinish, IPSAP.APP_SUBMIT_TYPE_PARAM.CHECKING_2);
      } else alert("일반 위원이 배정되지 않았습니다.");
   else alert("전문 위원이 배정되지 않았습니다.");
}
function onCompleteChangeSettingFinish(e, t, a) {
   e ? (window.location.href = "/html/officer/reviewClose_list.html") : alert(`심사 설정 변경 적용을 실패했습니다.\n(${t.em})`);
}
function onCompleteSettingFinish(e, t, a) {
   e ? (window.location.href = "/html/officer/application_list.html") : alert(`심사 설정 완료를 실패했습니다.\n(${t.em})`);
}
function onClickReviewRetry() {
   $("#modal_officeReview").modal("hide"),
      (g_AppItemParser.targetSavedItemNames = ["judge_alarm"]),
      g_AppItemParser.addMoreSaveTag("judge_expert_deadline", { data: { 0: "" } }),
      g_AppItemParser.addMoreSaveTag("judge_normal_deadline", { data: { 0: "" } }),
      g_AppItemParser.addMoreSaveTag("committee_in_member", { data: [] }),
      g_AppItemParser.addMoreSaveTag("committee_ex_member", { data: [] }),
      appItemSubmit(onCompleteReviewRetry, IPSAP.APP_SUBMIT_TYPE_PARAM.RETRY_CHECKING);
}
function onCompleteReviewRetry(e, t, a) {
   e ? (window.location.href = "/html/officer/IRB/reviewOffice_list-IRB_review.html") : alert(`행정 검토 재실행을 실패했습니다.\n(${t.em})`);
}
function setMemberCount() {
   $(".pro_cnt").text($("#pro > tr").length), $(".norm_in_cnt").text($("#norm_in > tr").length), $(".norm_out_cnt").text($("#norm_out > tr").length);
}
function onClickGoStafInfo() {
   IPSAP.setStor("user_seq", parseInt(user_seq)), (location.href = "/html/officer/staff_list-info.html");
}
function onClickJumpFinal() {
   $("#modal_btn_jump_fianl").modal("hide"), appItemSubmit(onCompleteJumpFinal, IPSAP.APP_SUBMIT_TYPE_PARAM.JUMP_FINAL);
}
function onCompleteJumpFinal(e, t, a) {
   e ? (window.location.href = "./application_list.html") : alert(`최종 심의 진행을 실패했습니다.\n(${t.em})`);
}
$("#btnPass").on("click", function () {
   if ($(".check-list input:checkbox:checked").length > 1 && $(".check-list input:checkbox:disabled").length < 1) return void $("#modal_pro_check").modal("show");
   $(".member_pro").length > 0
      ? $("#modal_pro_check").modal("show")
      : user_info.user_type[4]
      ? (addMember("pro", { user_seq: user_seq, name: user_info.name, email: user_info.email }, 0), $("#btnPro").attr("disabled", !0), $("#pass_comment").addClass("box_shadow_red_on"))
      : $("#modal_pro_pass").modal("show");
});
