function changefooterNavPre(e) {
   "0" == e && $("#app_page_navi_pre_btn").attr("onclick", "appItemSaveTemporary(appNavigationCallback, { PAGE_ID : 'PAGE_1_4'})");
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(document).ready(function () {
      function e(e) {
         var t = "smry_purpose_text";
         e.getItemData(t).applyTextValue(t);
         t = "smry_purpose_file";
         e.getItemData(t).applyMultiFileValue(1, t, "관련 자료");

         t = "smry_bckgr_evdnc_text";
         e.getItemData(t).applyTextValue(t);
         t = "smry_bckgr_evdnc_file";
         e.getItemData(t).applyMultiFileValue(1, t, "관련 자료");

         t = "smry_method_text";
         e.getItemData(t).applyTextValue(t);
         t = "smry_method_file";
         e.getItemData(t).applyMultiFileValue(1, t, "관련 자료");

         t = "smry_obsrv_check_text";
         e.getItemData(t).applyTextValue(t);
         t = "smry_obsrv_check_file";
         e.getItemData(t).applyMultiFileValue(1, t, "관련 자료");

         t = "smry_estimation_text";
         e.getItemData(t).applyTextValue(t);
         t = "smry_estimation_file";
         e.getItemData(t).applyMultiFileValue(1, t, "관련 자료");
         
         t = "general_object";
         var a = e.getItemData(t, "nosave").getStringValue("0");
         irbResetLeftNavi(a);
         t = "general_body_research";
         changefooterNavPre((a = e.getItemData(t, "nosave").getStringValue("0"))), $(".card").removeClass("hidden");
      }
      loadApplicationParams(), (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)), g_AppItemParser.load({ "filter.query_items": "smry_purpose, smry_bckgr_evdnc, smry_method, smry_obsrv_check, smry_estimation, general_division" }, e);
   });
