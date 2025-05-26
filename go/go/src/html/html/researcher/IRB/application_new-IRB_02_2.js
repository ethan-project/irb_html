"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(document).ready(function () {
      function e(e) {
         var t = "smry_stats_text";
         e.getItemData(t).applyTextValue(t);
         t = "smry_stats_file";
         e.getItemData(t).applyMultiFileValue(1, t, "관련 자료");

         t = "smry_selection_text";
         e.getItemData(t).applyTextValue(t);

         t = "smry_exception_text";
         e.getItemData(t).applyTextValue(t);

         t = "smry_produce_text";
         e.getItemData(t).applyTextValue(t);

         t = "smry_personal_info_text";
         e.getItemData(t).applyTextValue(t);

         t = "smry_specimen_text";
         e.getItemData(t).applyTextValue(t);

         t = "smry_gene_text";
         e.getItemData(t).applyTextValue(t);

         t = "smry_etc_file";
         e.getItemData(t).applyMultiFileValue(1, t, "관련 자료");

         t = "general_object";
         var a = e.getItemData(t, "nosave").getStringValue("0");
         irbResetLeftNavi(a), $(".card").removeClass("hidden");
      }
      loadApplicationParams(),
         (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
         g_AppItemParser.load({ "filter.query_items": "smry_stats_file, smry_stats_text, smry_selection_text, smry_exception_text, smry_produce_text, smry_personal_info_text, smry_specimen_text, smry_gene_text, smry_etc_file" }, e);
   });
