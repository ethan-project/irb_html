function changefooterNavPre(e, a) {
   "2" == e &&
      ("0" == a
         ? $("#app_page_navi_pre_btn").attr("onclick", "appItemSaveTemporary(appNavigationCallback, { PAGE_ID : 'PAGE_1_4'})")
         : $("#app_page_navi_pre_btn").attr("onclick", "appItemSaveTemporary(appNavigationCallback, { PAGE_ID : 'PAGE_1_5'})"));
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(document).ready(function () {
      function e(e) {
         var a = "doc_required_expt_form";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_desc_agree";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_provide";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_obtained_info";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_recruitment";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_conflict_pledge";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_edu_cert";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_director_resume";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_substance_transfer";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_license";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_permission";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_embryonic";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_compensation";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "doc_additional_experimental";
         e.getItemData(a).applyMultiFileValue(1, a, "", "");
         a = "general_object";
         var t = e.getItemData(a, "nosave").getStringValue("0");
         irbResetLeftNavi(t);
         changefooterNavPre(t, e.getItemData("general_body_research", "nosave").getStringValue("0")), $(".card").removeClass("hidden");
      }
      loadApplicationParams(), (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)), g_AppItemParser.load({ "filter.query_items": "doc_required, doc_additional, general_division" }, e);
   });
