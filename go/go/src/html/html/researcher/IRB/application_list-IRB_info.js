function onReviewRequest() {
   (g_AppItemParser.targetSavedItemNames = []), g_AppItemParser.addMoreSaveTag("supplement", makeSupplementData()), appItemSubmit(onCompleteReviewRequest, IPSAP.APP_SUBMIT_TYPE_PARAM.SUPPLEMENT);
}
function onCompleteReviewRequest(e, t, s) {
    console.log("e, t, s", e, t, s)
   e ? (window.location.href = "./application_list-IRB_info.html") : appSubmitHandleError("행정 보완 요청을 실패했습니다.", t);
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   "undefined" == typeof ipsap_item_module_js && document.write("<script src='/assets/js/ipsap/ipsap_item_module.js'></script>"),
   "undefined" == typeof date_utils_js && document.write("<script src='/assets/js/common/util/date_utils.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(function () {
      loadApplicationParams(),
         $("#all_app_r").load("/html/common/inc_application/all_app_readonly_irb.html", function () {
            switch ((remakeApplicationInfoReadOnly([]), g_AppInfo.appObj.application_result)) {
               case IPSAP.APPLICATION_RESULT.CHECKING:
               case IPSAP.APPLICATION_RESULT.CHECKING_2:
               case IPSAP.APPLICATION_RESULT.JUDGE_ING:
               case IPSAP.APPLICATION_RESULT.JUDGE_ING_2:
               case IPSAP.APPLICATION_RESULT.DECISION_ING:
               case IPSAP.APPLICATION_RESULT.JUDGE_DELAY:
               case IPSAP.APPLICATION_RESULT.REJECT:
               case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_A:
               case IPSAP.APPLICATION_RESULT.EXPERIMENT_ING_AC:
                  break;
               case IPSAP.APPLICATION_RESULT.EXPERIMENT_FINISH:
               case IPSAP.APPLICATION_RESULT.TASK_FINISH:
               case IPSAP.APPLICATION_RESULT.ACCEPT:
               case IPSAP.APPLICATION_RESULT.ACCEPT_AC:
                  $(".card-footer").children().eq(2).removeClass("hidden");
            }
         }),
         $(".task_title_ko").text(g_AppInfo.appObj.name_ko),
         $(".task_director_all").text("(연구 책임자 : " + g_AppInfo.appObj.user_name + ")");
         $("#btn_dup_app").on("click", function () {
            onReviewRequest();
         });
   });
