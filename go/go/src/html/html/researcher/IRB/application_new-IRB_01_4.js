function changefooterNavPre(e) {
   "1" == e && $("#app_page_navi_pre_btn").attr("onclick", "appItemSaveTemporary(appNavigationCallback, { PAGE_ID : 'PAGE_1_3'})");
}
function changefooterNavNext(e, a) {
   "0" == a &&
      ("1" == e
         ? $("#app_page_navi_next_btn").attr("onclick", "appItemSaveTemporary(appNavigationCallback, { PAGE_ID : 'PAGE_2_1'})")
         // : $("#app_page_navi_next_btn").attr("onclick", "appItemSaveTemporary(appNavigationCallback, { PAGE_ID : 'PAGE_4_1'})"));
                  : $("#app_page_navi_next_btn").attr("onclick", "appItemSaveTemporary(appNavigationCallback, { PAGE_ID : 'PAGE_2_1'})"));

}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(document).ready(function () {
      //Test show/hide component by data
      $("#form_01_4_A").hide();
      $("#form_01_4_B").hide();
      $("#form_01_4_C").hide();
      $("#form_01_4_D").hide();
      $("#form_01_4_E").hide();
      //5 . 기증자 정보
      $("#form_01_4_E_B5_1_3").on("change", function () {
         if ($(this).is(":checked")) {
            $("#form_01_4_E_B5_2_row").show();
         } else {
            $("#form_01_4_E_B5_2_row").hide();
         }
      });
      // Optional: handle initial state on page load
      if ($("#form_01_4_E_B5_1_3").is(":checked")) {
         $("#form_01_4_E_B5_2_row").show();
      } else {
         $("#form_01_4_E_B5_2_row").hide();
      }

      var e = "#general_experiment_method_etc_check-main_check",
         a = "#general_experiment_purpose_etc_check-main_check",
         t = "#general_investgate_etc_check-main_check";
      function n() {
         $("#general_experiment").is(":checked")
            ? ($("input:checkbox[name=human_experiment_check]").each(function (e) {
                 this.value >= 17 && this.value <= 23 && (this.disabled = !1);
              }),
              $(e).attr("disabled", !1),
              $(a).attr("disabled", !1))
            : ($("input:checkbox[name=human_experiment_check]").each(function (e) {
                 this.value >= 17 && this.value <= 23 && ((this.checked = !1), (this.disabled = !0));
              }),
              $(e).is(":checked") && ($(e).prop("checked", !1), $("#input_01_4-4_4_group").removeClass("show"), $("#general_experiment_method_etc_input").val("")),
              $(e).attr("disabled", !0),
              $(a).is(":checked") && ($(a).prop("checked", !1), $("#input_01_4-6_5_group").removeClass("show"), $("#general_experiment_purpose_etc_input").val("")),
              $(a).attr("disabled", !0));
      }
      function c() {
         $("#investgate").is(":checked")
            ? ($("input:checkbox[name=human_experiment_check]").each(function (e) {
                 this.value >= 25 && this.value <= 30 && (this.disabled = !1);
              }),
              $(t).attr("disabled", !1))
            : ($("input:checkbox[name=human_experiment_check]").each(function (e) {
                 this.value >= 25 && this.value <= 30 && ((this.checked = !1), (this.disabled = !0));
              }),
              $(t).is(":checked") && ($(t).prop("checked", !1), $("#input_01_4-6_7_group").removeClass("show"), $("#general_investgate_etc_input").val("")),
              $(t).attr("disabled", !0));
      }
      function i(e) {
         {
            let e = JSON.parse(COMM.getCookie("institution_info")).name_ko;
            $(".inst_name").text(e);
         }
         var a = "general_human_all";
         e.getItemData(a).applyTextValue();
         a = "general_human_institution";
         e.getItemData(a).applyTextValue();
         var t = e.getItemData("human_experiment_check").getStringValue("0"),
            i = new Object();
         "" != t && (i = t.split(",")),
            $("input:checkbox[name=human_experiment_check]").each(function (e) {
               isInArray(i, this.value) && ((this.checked = !0), 3 == this.value && $("#form_01-4-2_3_group").addClass("show"));
            });
         a = "general_recruit_etc_check";
         var _ = e.getItemData(a).applyMainCheckValueInCheckBox("general_recruit_etc_check-main_check"),
            p = "general_recruit_etc_input";
         e.getItemData(p).applyTextValue(), _ ? ($("#input_01_4-3_4_group").addClass("show"), $(`#${p}`).css("display", "block")) : $(`#${p}`).val("");
         (a = "general_experiment_method_etc_check"),
            (_ = e.getItemData(a).applyMainCheckValueInCheckBox("general_experiment_method_etc_check-main_check")),
            (p = "general_experiment_method_etc_input");
         e.getItemData(p).applyTextValue(), _ ? ($("#input_01_4-4_4_group").addClass("show"), $(`#${p}`).css("display", "block")) : $(`#${p}`).val("");
         (a = "general_experiment_purpose_etc_check"),
            (_ = e.getItemData(a).applyMainCheckValueInCheckBox("general_experiment_purpose_etc_check-main_check")),
            (p = "general_experiment_purpose_etc_input");
         e.getItemData(p).applyTextValue(), _ ? ($("#input_01_4-6_5_group").addClass("show"), $(`#${p}`).css("display", "block")) : $(`#${p}`).val("");
         (a = "general_investgate_etc_check"), (_ = e.getItemData(a).applyMainCheckValueInCheckBox("general_investgate_etc_check-main_check")), (p = "general_investgate_etc_input");
         e.getItemData(p).applyTextValue(), _ ? ($("#input_01_4-6_7_group").addClass("show"), $(`#${p}`).css("display", "block")) : $(`#${p}`).val("");
         (a = "general_object"), (t = e.getItemData(a, "nosave").getStringValue("0"));
         irbResetLeftNavi(t), changefooterNavPre(t);
         //research_type_check
         // a = "research_type_check";
         // e.getItemData(a).applyTextValue();
         // t = e.getItemData("research_type_check").getSelectValue(0);
         research_type_check = e?.dataObj?.research_type_check?.saved_data?.data?.[0]?.select_ids || "0";
         console.log("research_type_check", typeof research_type_check[0]);
         //Test show/hide component by data
         switch (research_type_check[0]) {
            case "1":
               $("#form_01_4_A").show();
               break;
            case "2":
               $("#form_01_4_B").show();
               break;
            case "3":
               $("#form_01_4_C").show();
               break;
            case "4":
               $("#form_01_4_D").show();
               break;
            case "5":
               $("#form_01_4_E").show();
               break;
            default:
               break;
         }

         console.log(t);
         changefooterNavNext(t, e.getItemData("general_body_research", "nosave").getStringValue("0")), n(), c(), $(".card").removeClass("hidden");
      }

      $("#general_experiment").on("click", function () {
         n();
      }),
         $("#investgate").on("click", function () {
            c();
         }),
         loadApplicationParams(),
         (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
         g_AppItemParser.load({ "filter.query_items": "human_experiment, general_division, research_type_check" }, i);
   });
