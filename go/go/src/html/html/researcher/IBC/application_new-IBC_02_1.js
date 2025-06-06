function dataMappingFunc(e) {
   var t = "general_title";
   (a = e.getItemData(t, !1)).applyTextMapValueReadOnly(t + "_name_ko", "name_ko"), a.applyTextMapValueReadOnly(t + "_name_en", "name_en"), mappingGeneralDirectorReadOnly(e);
   t = "ibc_plan_classification";
   (a = e.getItemData(t)).makeHtmlCheckList(t);
   t = "ibc_general_experiment";
   var a,
      n = (a = e.getItemData(t)).getStringValue("0"),
      i = [];
   if (("" != n && (i = n.split(",")), i.length > 0)) {
      ibcResetLeftNavi(i);
      let e = Math.min.apply(null, i);
      e >= 0 && e < 4
         ? $("#classification").text("국가 승인 실험")
         : e >= 4 && e < 9
         ? $("#classification").text("기관 승인 실험")
         : e >= 9 && e < 11
         ? $("#classification").text("기관 신고 실험")
         : $("#classification").text("면제실험");
   }
   e.deleteSaveTag(["ibc_general_experiment"]), $(".card").removeClass("hidden");
}
function mappingGeneralDirectorReadOnly(e) {
   var t = e.getItemData("general_director");
   if (null == t.dataObj.saved_data) return;
   t.initMembersIBC();
   let a = t.dataObj.saved_data.data[0];
   (a.info.tmp_phoneno = a.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")), (a.info.tmp_edu_date = a.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3"));
   let n = "";
   if ("" != a.edu_course) {
      let e = JSON.parse(a.edu_course);
      for (key in e) n += `\n        <div class="flexMid">\n          <span class="min_w120">${key}</span>\n          <span class="mLeft10">${e[key]}</span>\n        </div>`;
   }
   $("#general_director_name").text(a.info.name),
      $("#general_director_email").text(a.info.email),
      $("#general_director_dept").text(a.info.dept_str),
      $("#general_director_edu_date").text(a.info.tmp_edu_date),
      $("#general_director_position").text(a.info.position_str),
      $("#general_director_edu_instition").text(a.info.edu_institution_str),
      $("#general_director_major_field").text(a.info.major_field_str),
      $("#general_director_edu_course_num").text(a.info.edu_course_num),
      $("#general_director_phoneno").text(a.info.tmp_phoneno),
      $("#general_director_edu_course").append(n),
      e.deleteSaveTag(["general_director"]);
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IBC/common_application_new-IBC.js'></script>"),
   $(document).ready(function () {
      loadApplicationParams(),
         (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
         g_AppItemParser.load({ "filter.query_items": "general_title, general_director, ibc_general_experiment, ibc_plan_classification" }, dataMappingFunc);
      var e = function (e, t) {
         (this.phase = "page2_1"), (this.btn_type = t);
      };
      (e.prototype.check = function () {
         "nextPage" == this.btn_type ? appNaviNextPage() : saveTemporary();
      }),
         di
            .autowired(!1)
            .register("validator_nextPage")
            .as(e)
            .withConstructor()
            .withProperties()
            .prop("btn_type")
            .val("nextPage")
            .register("validator_tempSave")
            .as(e)
            .withConstructor()
            .withProperties()
            .prop("btn_type")
            .val("tempSave");
   });
