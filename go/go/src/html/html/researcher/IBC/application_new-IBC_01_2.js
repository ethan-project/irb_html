function makeEduCourse(e) {
   if (!$(`.course_${e}`).eq(1).hasClass("show")) return "";
   let t = new Object();
   return (
      $(`.course_${e}`)
         .children("div")
         .each(function () {
            t[$.trim($(this).eq(0).text())] = $(this).children("input").val();
         }),
      JSON.stringify(t)
   );
}
function makeExptUserList() {
   var e = "general_expt",
      t = g_AppItemParser.getItemData(e),
      n = '<tbody id="general_expt">';
   let r = g_AppItemParser.managedItems[e].member;
   $.each(r, function (e, t) {
      (t.info.tmp_phoneno = t.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")), (t.info.tmp_edu_date = t.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3"));
      var r = `general_expt_${t.user_seq}_radio`,
         a = makeCourseList(t.user_seq);
      n += `\n      <tr data-id="${t.user_seq}" item_name="general_expt">\n        <td>\n          <div class="custom-radio">\n            <input type="radio" id="${r}" name="exp_manager" value="${t.user_seq}">\n          </div>\n        </td>\n        <td>${t.info.name}</td>\n        <td>${t.info.dept_str}</td>\n        <td>${t.info.position_str}</td>\n        <td>${t.info.major_field_str}</td>\n        <td>${t.info.tmp_phoneno}</td>\n        <td>${t.info.email}</td>\n        <td>${a}</td>\n        <td>\n          <a href="javascript:void(0);" class="btn btn-xs btn-outline-danger btn-staff-delete"><i class="far fa-trash-alt mRight5"></i>삭제</a>\n        </td>\n      </tr>`;
   }),
      (n +=
         '\n  <tr class="add_row">\n    <td colspan="11">\n      <a href="javascript:void(0);" onclick="openModelOtherStaff();" class="btn btn-xs btn-outline-primary" ><i class="fas fa-user-plus mRight5"></i>실험 수행자 선택 추가</a>\n    </td>\n  </tr>\n  </tbody>'),
      $("#general_expt").replaceWith(n),
      $.each(r, function (e, t) {
         var n = `general_expt_${t.user_seq}_radio`;
         t.animal_mng_flag > 0 && $("#" + n).prop("checked", !0), t.edu_course && mappingCourseList(t.user_seq, JSON.parse(t.edu_course));
      }),
      $("input[name=exp_manager]").on("change", function () {
         var e = Number($(this).val());
         t.changeAttrForMember(e, "animal_mng_flag", 1);
      });
}
function mappingCourseList(e, t) {
   let n = !1;
   $.each(t, function (t, r) {
      $(`input[id='${t}_${e}']`).val(r), r && (n = !0);
   }),
      n && $(`.course_${e}`).click();
}
function makeCourseList(e) {
   let t = $("#general_director").data("course-title").split(","),
      n = `\n    <button class="btn btn-xs btn-outline-primary course_${e} show w100p" data-toggle="collapse" data-target=".course_${e}">교육이수 정보 입력</button>\n    <div class="course_${e} collapse">`;
   for (course_no in t)
      n += `\n      <div class="flexMid text-left">\n        <span class="w110 left mRight5">${t[course_no]}</span>\n        <input type="text" id="${t[course_no]}_${e}" class="form-control form-control-sm w150">\n      </div>\n    `;
   return (n += `\n      <button class="btn btn-xs btn-outline-secondary w100p mTop5" data-toggle="collapse" data-target=".course_${e}">입력 취소</button>\n    </div>`), n;
}
function openModelOtherStaff() {
   makeOtherStaffList(), $("#modal_staff").modal("show");
   var e = "general_expt";
   let t = g_AppItemParser.managedItems[e].member;
   $.each(t, function (t, n) {
      g_AppItemParser.managedItems[e].member[t].edu_course = makeEduCourse(n.user_seq);
   });
}
function onStaffSelect(e) {
   $.each(g_other_userlist, function (t, n) {
      if (Number(n.user_seq) == e && !g_AppItemParser.hasMemberExists(n.user_seq)) return g_AppItemParser.insertMemberIBC("general_expt", n), makeExptUserList(), void $("#modal_staff").modal("hide");
   });
}
function makeOtherStaffList() {
   let e = '<tbody id="other-staff">';
   $.each(g_other_userlist, function (t, n) {
      var r = n.user_type.split(",");
      g_AppItemParser.hasMemberExists(n.user_seq) ||
         -1 == $.inArray("3", r) ||
         ((n.tmp_phoneno = n.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
         (e += `\n      <tr data-url="" class="staff-row" onclick="onStaffSelect(${n.user_seq})">\n        <td class="text-center">${n.name}</td>\n        <td>${n.dept_str}</td>\n        <td>${n.position_str}</td>\n        <td>${n.major_field_str}</td>\n        <td>${n.tmp_phoneno}</td>\n        <td>${n.email}</td>\n        <td>${n.edu_course_num}</td>\n        <td>${n.edu_institution_str}</td>\n      </tr>`));
   }),
      (e += "</tbody>"),
      $("#other-staff").replaceWith(e);
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IBC/common_application_new-IBC.js'></script>"),
   document.write("<script src='../sub_inc/1_2_ibc_general_fclty.js'></script>"),
   $(document).ready(function () {
      function e(e) {
         !(function (e) {
            var t = e.getItemData("general_director");
            if (null == t.dataObj.saved_data) return;
            (t.dataObj.saved_data.data[0].animal_mng_flag = 0), t.initMembersIBC();
            let n = t.dataObj.saved_data.data[0];
            (n.info.tmp_phoneno = n.info.phoneno.replace(/(^02.{0}|^01.{1}|[0-9]{3})([0-9]+)([0-9]{4})/, "$1-$2-$3")),
               (n.info.tmp_edu_date = n.info.edu_date.replace(/(\d{4})(\d{2})(\d{2})/, "$1-$2-$3")),
               $("#general_director_name").text(n.info.name),
               $("#general_director_email").text(n.info.email),
               $("#general_director_dept").text(n.info.dept_str),
               $("#general_director_edu_date").text(n.info.tmp_edu_date),
               $("#general_director_position").text(n.info.position_str),
               $("#general_director_edu_instition").text(n.info.edu_institution_str),
               $("#general_director_major_field").text(n.info.major_field_str),
               $("#general_director_edu_course_num").text(n.info.edu_course_num),
               $("#general_director_phoneno").text(n.info.tmp_phoneno),
               $("#director_course_info").append(makeCourseList(n.user_seq)),
               n.edu_course && mappingCourseList(n.user_seq, JSON.parse(n.edu_course));
         })(e),
            (function (e) {
               var t = e.getItemData("general_expt");
               null != t.dataObj.saved_data && (t.initMembersIBC(), makeExptUserList());
            })(e),
            mappingIbcFacility(e);
         var t = e.getItemData("ibc_general_experiment").getStringValue("0"),
            n = [];
         "" != t && (n = t.split(",")), n.length > 0 && ibcResetLeftNavi(n), e.deleteSaveTag(["ibc_general_experiment"]), $(".card").removeClass("hidden");
      }
      loadApplicationParams(),
         (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
         g_AppItemParser.load({ "filter.query_items": "ibc_general_researcher_fclty, ibc_general_experiment" }, e),
         (function () {
            let e = CONST.API.INSTITUTION.USER;
            API.load({
               url: e.replace("${inst_seq}", g_AppInfo.appObj.institution_seq),
               data: { "filter.user_type": IPSAP.USER_TYPE.RESEARCHER },
               type: CONST.API_TYPE.GET,
               success: function (e) {
                  g_other_userlist = e;
               },
            });
         })(),
         $(document).on("click", ".btn-staff-delete", function () {
            if ($(this).closest("tr").find("input:radio").is(":checked")) return !1;
            var e = $(this).closest("tr").attr("data-id"),
               t = $(this).closest("tr").attr("item_name");
            g_AppItemParser.removeMembers(t, e), $(this).closest("tr").remove();
         });
      var t = function (e, t) {
         (this.phase = "page1_2"), (this.btn_type = t);
      };
      (t.prototype.check = function () {
         "nextPage" == this.btn_type ? appNaviNextPage() : saveTemporary();
      }),
         di
            .autowired(!1)
            .register("validator_nextPage")
            .as(t)
            .withConstructor()
            .withProperties()
            .prop("btn_type")
            .val("nextPage")
            .register("validator_tempSave")
            .as(t)
            .withConstructor()
            .withProperties()
            .prop("btn_type")
            .val("tempSave");
   });
