"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   "undefined" == typeof date_utils_js && document.write("<script src='/assets/js/common/util/date_utils.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(function () {
      loadApplicationParams(),
         $("#all_app_r").load("/html/common/inc_application/all_app_readonly_irb.html", function () {
            remakeApplicationInfoReadOnly([]);
         }),
         $(".task_title_ko").text(g_AppInfo.appObj.name_ko),
         $(".task_director_all").text("(연구 책임자 : " + g_AppInfo.appObj.user_name + ")");
   });
