function dupAppAfterRederict() {
   let e = g_AppInfo.appObj.application_seq,
      c = CONST.API.APPLICATION.DUPPLICATE;
   API.load({
      url: c.replace("${app_seq}", e),
      type: CONST.API_TYPE.POST,
      success: function (e) {
         g_AppInfo.initWithAppObj(e.list), g_AppInfo.saveParamsAndNavigate(APP_IBC_NAVIGATION.IBC.PAGE_INFO[0].URL);
      },
   });
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   document.write("<script src='/html/researcher/IBC/common_application_new-IBC.js'></script>"),
   document.write("<script src='../sub_inc/1_2_ibc_general_fclty.js'></script>"),
   document.write("<script src='../sub_inc/3_2_ibc_infection_chance.js'></script>"),
   document.write("<script src='../sub_inc/3_2_ibc_risk_emergency_network.js'></script>"),
   $(function () {
      switch (
         (loadApplicationParams(),
         $("#all_app_r").load("/html/common/inc_application/all_app_readonly_ibc.html", function () {
            remakeApplicationInfoReadOnly([]);
         }),
         $(".task_title_ko").text(g_AppInfo.appObj.name_ko),
         $(".task_director_all").text("(연구 책임자 : " + g_AppInfo.appObj.user_name + ")"),
         Number(g_AppInfo.appObj.application_result))
      ) {
         case IPSAP.APPLICATION_RESULT.REJECT:
         case IPSAP.APPLICATION_RESULT.REQUIRE_RETRY:
            $(".card-footer").children().eq(0).removeClass("hidden");
            break;
         case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
         case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
         case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
         case IPSAP.APPLICATION_RESULT.TASK_FINISH:
         case IPSAP.APPLICATION_RESULT.ACCEPT:
         case IPSAP.APPLICATION_RESULT.ACCEPT_AC:
            $(".card-footer").children().eq(1).css("display", "block");
      }
      $("#btn_dup_app").on("click", function () {
         dupAppAfterRederict();
      });
   });
