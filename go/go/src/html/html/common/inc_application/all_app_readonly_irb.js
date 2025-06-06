"undefined" == typeof ipsap_common_func_js && document.write("<script src='/assets/js/ipsap/ipsap_common_func.js'></script>");
var g_general_obj_value = 0,
   checkBoxId1 = "#general_experiment_method_etc_check-main_check",
   checkBoxId2 = "#general_experiment_purpose_etc_check-main_check",
   checkBoxId3 = "#general_investgate_etc_check-main_check";
function remakeApplicationInfoReadOnly(e, a, t) {
   console.log("e, a, t", e, a, t);
   var n = "irb_all";
   for (let a = 0; a < e.length; ++a) n = n + "," + e[a];
   let l = { "filter.query_items": n };
   if (null != t && g_AppInfo.childAppSeq > 0 && t.length > 0) {
      var _ = t[0];
      for (let e = 1; e < t.length; ++e) _ = _ + "," + t[e];
      (l["filter.child_app_seq"] = g_AppInfo.childAppSeq), (l["filter.child_items"] = _);
   }
   console.log("e", e);

   if (isInArray(e, "supplement")) {
      for (const [e, a] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
         console.log("e, a", e, a);

         let t = "app_supplement_" + e,
            n = "app_supplement_text_" + e;
         (html = `\n      <div class="col-sm-9 mTop20">\n        <div class="checkbox checkbox-danger">\n          <input type="checkbox" id="${t}" value="1" data-toggle="collapse" data-target='#${n}' checked>\n          <label for="${t}">${a}</label>\n        </div>\n      </div>\n      <div class="col-sm-9 collapse show" id="${n}">\n        <textarea class="form-control red_border_shallow input_review" rows="2" placeholder="${"구체적인 보완 요청 사항을 입력해 주세요"}">${""}</textarea>\n      </div>`),
            console.log("t", "#" + t);
         $("#" + t).replaceWith(html);
      }
   } else{
    for (const [e, a] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
        //  $("#app_supplement_" + e).remove();
         $("#app_supplement_" + e).hide();
      }
   }
      
   function c() {
      var e = [];
      1 == $("input:radio[name ='general_object']:checked").val()
         ? ($(".general_judgement_area").removeClass("hidden"), e.push("심의 대상"))
         : ($(".general_judgement_area").addClass("hidden"), e.push("심의 면제")),
         $('input:checkbox[id="general_human_research"]').is(":checked") && e.push("인간대상 연구"),
         $('input:checkbox[id="general_body_research"]').is(":checked") && e.push("인체 유래물 연구"),
         (function (e) {
            $("#select_content").empty();
            var a = '<span class="fw_bold mRight10">선택 항목:</span>';
            for (let t = 0; t < e.length; t++) a += `<button type="button" class="btn btn-outline-primary btn-round btn_tag mRight5"> \n                  ${e[t]}\n                </button>`;
            $("#select_content").append(a);
         })(e);
   }
   function r() {
      $("#general_experiment").is(":checked")
         ? ($("input:checkbox[name=human_experiment_check]").each(function (e) {
              this.value >= 17 && this.value <= 23 && (this.disabled = !1);
           }),
           $(checkBoxId1).attr("disabled", !1),
           $(checkBoxId2).attr("disabled", !1))
         : ($("input:checkbox[name=human_experiment_check]").each(function (e) {
              this.value >= 17 && this.value <= 23 && ((this.checked = !1), (this.disabled = !0));
           }),
           $(checkBoxId1).is(":checked") && ($(checkBoxId1).prop("checked", !1), $("#input_01_4-4_4_group").removeClass("show"), $("#general_experiment_method_etc_input").val("")),
           $(checkBoxId1).attr("disabled", !0),
           $(checkBoxId2).is(":checked") && ($(checkBoxId2).prop("checked", !1), $("#input_01_4-6_5_group").removeClass("show"), $("#general_experiment_purpose_etc_input").val("")),
           $(checkBoxId2).attr("disabled", !0));
   }
   function o() {
      $("#investgate").is(":checked")
         ? ($("input:checkbox[name=human_experiment_check]").each(function (e) {
              this.value >= 25 && this.value <= 30 && (this.disabled = !1);
           }),
           $(checkBoxId3).attr("disabled", !1))
         : ($("input:checkbox[name=human_experiment_check]").each(function (e) {
              this.value >= 25 && this.value <= 30 && ((this.checked = !1), (this.disabled = !0));
           }),
           $(checkBoxId3).is(":checked") && ($(checkBoxId3).prop("checked", !1), $("#input_01_4-6_7_group").removeClass("show"), $("#general_investgate_etc_input").val("")),
           $(checkBoxId3).attr("disabled", !0));
   }
   function i() {
      3 == $('input[name="consent_method_0"]:checked').val() ? $("#consent_exmpt").attr("disabled", !1) : ($("#consent_exmpt").val(""), $("#consent_exmpt").attr("disabled", !0));
   }
   (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
      g_AppItemParser.load(l, function (t) {
         !(function (a) {
            if (!isInArray(e, "supplement")) return;
            var t = a.getItemData("supplement");
            for (const [e, a] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
               let a = t.dataObj.saved_data.data[e];
               if (null == a || "" == a) continue;
               let n = "app_supplement_" + e,
                  l = "app_supplement_text_" + e;
               0 == a.substring(0, 1)
                  ? ($("#" + n).prop("checked", !1),
                    $("#" + n)
                       .parent()
                       .removeClass("checkbox-danger"),
                    $("#" + n)
                       .siblings()
                       .addClass("blue_deep"),
                    $("#" + l).addClass("collapse"),
                    $("#" + l).removeClass("show"))
                  : $("#" + l)
                       .children("textarea")
                       .val(a.substring(2));
            }
         })(t);
         {
            let e = (l = t.getItemData("application_info")).getStringValue("application_no");
            "" != e && $(".register_num").text(e);
         }
         var n = "general_title";
         (l = t.getItemData(n)).applyTextMapValueReadOnly(n + "_name_ko", "name_ko"), l.applyTextMapValueReadOnly(n + "_name_en", "name_en");
         n = "general_end_date";
         var l = t.getItemData(n);
         if (($("#general_end_date").text(l.applyTextValueReadOnly()), $(".task_period").text(l.applyTextValueReadOnly()), 0 != g_AppInfo.appObj.approved_dttm)) {
            let e = getDateDiffWith10($("#general_end_date").text(), g_AppInfo.appObj.approved_dttm);
            $("#total_date").text("총 " + e + " 일"), $("#approved_date").text(g_AppInfo.appObj.approved_dttm), $(".task_approved_date").text(g_AppInfo.appObj.approved_dttm);
         } else $("#approved_date").text("IRB 위원회 승인일"), $("#total_date").text("위원회 승인 이후부터 실험 일수를 계산합니다.");
         (n = "general_fund_org"), (l = t.getItemData(n));
         (d = l.applyMainCheckValueInCheckBoxReadOnly("general_fund_org-main_check")) &&
            ($("#general_fund_org-main_check").attr("onclick", "return false;"), $("#input_01-3_group").addClass("collapse").removeClass("show"));
         n = "general_fund_org_name";
         (l = t.getItemData(n)).applyTextValueReadOnly();
         n = "general_fund_conflict";
         null != (l = t.getItemData(n)) && l.makeHtmlRadioTypeReadOnly();
         (function (e) {
            var a = "general_director",
               t = e.getItemData(a);
            if (null == t.dataObj.saved_data) return;
            t.initMembers();
            let n = t.dataObj.saved_data.data[0];
            (n.info.tmp_phoneno = n.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
               (n.info.tmp_edu_date = n.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3")),
               (n.exp_year_str = ""),
               1 == n.exp_year_code ? (n.exp_year_str = "1년 미만") : 2 == n.exp_year_code ? (n.exp_year_str = "1년 ~ 3년") : (n.exp_year_str = "3년 이상");
            $("#general_director_name").text(n.info.name),
               $("#general_director_email").text(n.info.email),
               $("#general_director_dept").text(n.info.dept_str),
               $("#general_director_edu_date").text(n.info.tmp_edu_date),
               $("#general_director_position").text(n.info.position_str),
               $("#general_director_edu_instition").text(n.info.edu_institution_str),
               $("#general_director_major_field").text(n.info.major_field_str),
               $("#general_director_edu_course_num").text(n.info.edu_course_num),
               $("#general_director_phoneno").text(n.info.tmp_phoneno),
               $("#general_director_select").text(n.exp_year_str);
         })(t),
            (function (e) {
               var a = "general_expt",
                  t = e.getItemData(a);
               if (null == t.dataObj.saved_data) return;
               t.initMembers(),
                  (function () {
                     var e = "general_expt",
                        a = g_AppItemParser.getItemData(e),
                        t = `<tbody class="data" id="${e}">`;
                     let n = a.dataObj.saved_data.data;
                     $.each(n, function (a, n) {
                        (n.info.tmp_phoneno = n.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
                           (n.info.tmp_edu_date = n.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3")),
                           (n.exp_year_str = ""),
                           1 == n.exp_year_code ? (n.exp_year_str = "1년 미만") : 2 == n.exp_year_code ? (n.exp_year_str = "1년 ~ 3년") : (n.exp_year_str = "3년 이상"),
                           (t += `\n          <tr data-id="${n.user_seq}" item_name="${e}">\n              <td>${n.exp_type_code_str}</td>\n              <td>${n.info.name}</td>\n              <td>${n.info.dept_str}</td>\n              <td>${n.info.position_str}</td>\n              <td>${n.info.tmp_phoneno}</td>\n              <td class="ellipsis max_w150" title="${n.info.email}">${n.info.email}</td>\n              <td>${n.info.edu_course_num}</td>\n              <td>${n.info.edu_institution_str}</td>\n              <td>${n.exp_year_code_str}</td>\n          </tr>\n          `);
                     }),
                        $("#general_expt").replaceWith(t),
                        (t += "</tbody>");
                  })();
            })(t);
         n = "general_object";
         var _ = (l = t.getItemData(n)).getStringValue("0");
         $(`input:radio[name ='general_object']:input[value='${_}']`).click(), $("input:radio[name ='general_object']").prop("readonly", !0), (g_general_obj_value = _);
         (n = "general_human_research"), (l = t.getItemData(n));
         1 == (_ = l.getStringValue("0"))
            ? ($('input:checkbox[id="general_human_research"]').attr("checked", !0), $(".general_human_research").removeClass("hidden"))
            : $('input:checkbox[id="general_human_research"]').attr("checked", !1);
         $('input:checkbox[id="general_human_research"]').prop("readonly", !0);
         (n = "general_body_research"), (l = t.getItemData(n));
         1 == (_ = l.getStringValue("0"))
            ? ($('input:checkbox[id="general_body_research"]').attr("checked", !0), $(".general_body_research").removeClass("hidden"))
            : $('input:checkbox[id="general_body_research"]').attr("checked", !1);
         $('input:checkbox[id="general_body_research"]').prop("readonly", !0);
         (n = "general_judgement"), (_ = (l = t.getItemData(n)).getStringValue("0"));
         if (($(`input:radio[name ='general_judgement']:input[value='${_}']`).click(), $("input:radio[name ='general_judgement']").prop("readonly", !0), c(), 2 == g_general_obj_value)) {
            n = "general_exempt_reason_text";
            (l = t.getItemData(n)).applyTextValueReadOnly(n);
            n = "general_exempt_reason_file";
            (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "관련 자료"), $(".general_exempt_reason").removeClass("hidden");
         } else $("#form_02").removeClass("hidden"), $("#form_03").removeClass("hidden");
         {
            let e = JSON.parse(COMM.getCookie("institution_info")).name_ko;
            $(".inst_name").text(e);
         }
         n = "general_human_all";
         (l = t.getItemData(n)).applyTextValueReadOnly();
         n = "general_human_institution";
         (l = t.getItemData(n)).applyTextValueReadOnly();
         var i = "human_experiment_check",
            p = ((_ = (l = t.getItemData(i)).getStringValue("0")), new Object());
         "" != _ && (p = _.split(","));
         $("input:checkbox[name=human_experiment_check]").each(function (e) {
            isInArray(p, this.value) && ((this.checked = !0), 3 == this.value && $("#form_01-4-2_3_group").addClass("show"));
         });
         n = "general_recruit_etc_check";
         var d = (l = t.getItemData(n)).applyMainCheckValueInCheckBox("general_recruit_etc_check-main_check"),
            s = "general_recruit_etc_input";
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#input_01_4-3_4_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "general_experiment_method_etc_check"),
            (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBox("general_experiment_method_etc_check-main_check")),
            (s = "general_experiment_method_etc_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#input_01_4-4_4_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "general_experiment_purpose_etc_check"),
            (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBox("general_experiment_purpose_etc_check-main_check")),
            (s = "general_experiment_purpose_etc_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#input_01_4-6_5_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "general_investgate_etc_check"), (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBox("general_investgate_etc_check-main_check")), (s = "general_investgate_etc_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#input_01_4-6_7_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         r(), o();
         n = "general_body_all";
         (l = t.getItemData(n)).applyTextValueReadOnly();
         n = "general_body_institution";
         (l = t.getItemData(n)).applyTextValueReadOnly();
         (i = "general_body_research_check"), (_ = (l = t.getItemData(i)).getStringValue("0")), (p = new Object());
         "" != _ && (p = _.split(","));
         $("input:checkbox[name=general_body_research_check]").each(function (e) {
            isInArray(p, this.value) && (this.checked = !0);
         });
         (n = "general_sample_type_etc_check"), (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("general_sample_type_etc_check-main_check")), (s = "general_sample_type_etc_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#input_01_4-9_5_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "general_collect_check"), (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("general_collect_check-main_check")), (s = "general_collect_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#input_01_4-10_4_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "general_sample_type_check"), (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("general_sample_type_check-main_check")), (s = "general_sample_type_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#input_01_4-10_7_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "general_anonymity_check"), (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("general_anonymity_check-main_check")), (s = "general_anonymity_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#input_01_4-11_4_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         n = "general_bdoy_store_flag";
         null != (l = t.getItemData(n)) && l.makeHtmlRadioTypeReadOnly();
         n = "general_genetic_store_flag";
         null != (l = t.getItemData(n)) && l.makeHtmlRadioTypeReadOnly();
         n = "smry_purpose_text";
         (l = t.getItemData(n)).applyTextValueReadOnly(n);
         n = "smry_purpose_file";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "관련 자료", "mRight10 mBot0 w110");
         n = "smry_bckgr_evdnc_text";
         (l = t.getItemData(n)).applyTextValueReadOnly(n);
         n = "smry_bckgr_evdnc_file";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "관련 자료", "mRight10 mBot0 w110");
         n = "smry_method_text";
         (l = t.getItemData(n)).applyTextValueReadOnly(n);
         n = "smry_method_file";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "관련 자료", "mRight10 mBot0 w110");
         n = "smry_obsrv_check_text";
         (l = t.getItemData(n)).applyTextValueReadOnly(n);
         n = "smry_obsrv_check_file";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "관련 자료", "mRight10 mBot0 w110");
         n = "smry_estimation_text";
         (l = t.getItemData(n)).applyTextValueReadOnly(n);
         n = "smry_estimation_file";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "관련 자료", "mRight10 mBot0 w110");
         n = "consent_method";
         (l = t.getItemData(n)).makeHtmlRadioTypeReadOnly(IPSAP.COL.HORI);
         n = "consent_exmpt";
         (l = t.getItemData(n)).applyTextValueReadOnly();
         n = "consent_contact_flag";
         (l = t.getItemData(n)).makeHtmlRadioTypeReadOnly(IPSAP.COL.HORI);
         n = "consent_pin_flag";
         (l = t.getItemData(n)).makeHtmlRadioTypeReadOnly(IPSAP.COL.HORI);
         n = "consent_authority_check";
         (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("consent_authority_check-main_check");
         (n = "consent_authority_attorney_check"),
            (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("consent_authority_attorney_check-main_check")),
            (s = "consent_authority_attorney_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#form_03_1-5_3_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "consent_authority_etc_check"), (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("consent_authority_etc_check-main_check")), (s = "consent_authority_etc_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#form_03_1-5_5_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         n = "consent_manager_check";
         (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("consent_manager_check-main_check");
         (n = "consent_manager_collaborator_check"),
            (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("consent_manager_collaborator_check-main_check")),
            (s = "consent_manager_collaborator_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#form_03_1-6_3_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "consent_manager_etc_check"), (d = (l = t.getItemData(n)).applyMainCheckValueInCheckBoxReadOnly("consent_manager_etc_check-main_check")), (s = "consent_manager_etc_input");
         t.getItemData(s).applyTextValueReadOnly(), d ? ($("#form_03_1-6_5_group").addClass("show"), $(`#${s}`).css("display", "block")) : $(`#${s}`).val("");
         (n = "consent_date_radio"), (_ = (l = t.getItemData(n)).getStringValue("0"));
         $(`input:radio[name ='consent_date_radio']:input[value='${_}']`).click();
         s = "consent_date_input";
         t.getItemData(s).applyTextValueReadOnly();
         n = "consent_place";
         (l = t.getItemData(n)).applyTextValueReadOnly();
         n = "doc_required_expt_form";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_desc_agree";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_provide";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_obtained_info";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_recruitment";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_conflict_pledge";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_edu_cert";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_director_resume";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_substance_transfer";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_license";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_additional_permission";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group");
         n = "doc_etc";
         (l = t.getItemData(n)).applyMultiFileValueReadOnly(n, "", "mBot0"),
            $("#" + n)
               .children()
               .children()
               .removeClass("form-group"),
            $(".task_director").text(g_AppInfo.appObj.user_name);
         var m = getDttm(g_AppInfo.appObj.submit_dttm).dt.split("-"),
            u = m[0] + "년 " + m[1] + "월 " + m[2] + "일";
         0 == g_AppInfo.appObj.submit_dttm && (u = getTimeStrFromDateType(new Date()));
         $(".agree_date").text(u), $(".card").removeClass("hidden"), null != a && a(t);
         setTimeout(function () {
            $("body").hasClass("noPrint") || setPrintableForm();
         }, 500);
      }),
      $(".reset_review_type").change(function () {
         c();
      }),
      $("#general_experiment").on("click", function () {
         r();
      }),
      $("#investgate").on("click", function () {
         o();
      }),
      i(),
      $('input[name="consent_method_0"]').change(function () {
         i();
      });
}
function makeSupplementData() {
   var e = {};
   for (const [t, a] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
      let a = "app_supplement_text_" + t,
         i = $("#" + ("app_supplement_" + t)).is(":checked"),
         _ = $("#" + a)
            .children("textarea")
            .val();
      if ((i || (_ = ""), !i || "" != _)) {
         let a = "0|";
         i && (a = "1|"), (a += _), (e[t] = a);
      }
   }
   return { data: e };
}
function cbFuncAfterMapping(e) {
   resetLeftStepNaviIcon(),
      $("input[type=checkbox]").on({
         change: function (e) {
            resetLeftStepNaviIcon(),
               $(this).is(":checked")
                  ? ($(this).parent().addClass("checkbox-danger"), $(this).siblings().removeClass("blue_deep"))
                  : ($(this).parent().removeClass("checkbox-danger"), $(this).siblings().addClass("blue_deep"), $(this).parent().parent().next().find(".input_review").val(""));
         },
      }),
      $(".input_review").on({
         keyup: function (e) {
            resetLeftStepNaviIcon();
         },
      });
   var t = e.getItemData("ibc_general_experiment").getStringValue("0"),
      s = [];
   "" != t && (s = t.split(",")), Math.min.apply(null, s) <= 8 ? $("#btn_finish_review_fast").hide() : $("#btn_finish_review").hide();
}
function resetLeftStepNaviIcon() {
   let e = 0,
      t = !0,
      s = !0,
      i = !1,
      n = !1;
   for (const [a, p] of Object.entries(IPSAP.IBC_REVIEW_TAG)) {
      let r = "app_supplement_" + a,
         o = "app_supplement_text_" + a,
         l = p.split("-")[0];
      if (e != l) {
         if (e > 0) {
            let s = $(`.process_content[data-process-num=${e}]`);
            s.removeClass("completed pending supplement"), t ? (i ? s.addClass("supplement") : s.addClass("completed")) : s.addClass("pending");
         }
         (e = l), (t = !0), (i = !1);
      }
      t &&
         $("#" + r).is(":checked") &&
         "" ==
            $("#" + o)
               .children("textarea")
               .val() &&
         ((t = !1), (s = !1)),
         $("#" + r).is(":checked") &&
            "" !=
               $("#" + o)
                  .children("textarea")
                  .val() &&
            ((i = !0), (n = !0));
   }
   if (e > 0) {
      let s = $(`.process_content[data-process-num=${e}]`);
      s.removeClass("completed pending supplement"), t ? (i ? s.addClass("supplement") : s.addClass("completed")) : s.addClass("pending");
   }
   $("#btn_request").removeClass("disabled"),
      $("#btn_finish_review").removeClass("disabled"),
      $("#btn_finish_review_fast").removeClass("disabled"),
      (s && n) || $("#btn_request").addClass("disabled"),
      (s && !n) || ($("#btn_finish_review").addClass("disabled"), $("#btn_finish_review_fast").addClass("disabled"));
}
$(document).ready(function () {
   remakeApplicationInfoReadOnly(["supplement"], cbFuncAfterMapping);
});
