var g_object, g_for_people, g_research_type;
function dataMappingFunc(e) {
   0 == e.app_seq &&
      e.addMoreSaveTag("application_info", {
         data: g_AppInfo.appObj,
      }),
      getIACUCAppList(),
      modalGetIACUCAppList();
   var a = "general_title";
   (t = e.getItemData(a)).applyTextMapValue(a + "_name_ko", "name_ko"), t.applyTextMapValue(a + "_name_en", "name_en");
   a = "general_end_date";
   (t = e.getItemData(a)).applyCalanderValue(a);
   a = "ibc_general_experiment_cnt";
   (t = e.getItemData(a)).applyTextValue();
   a = "ibc_general_experiment_degree";
   (t = e.getItemData(a)).applyTextValue();
   a = "ibc_general_fund_org";
   (t = e.getItemData(a)).applyMainCheckValueInCheckBox("ibc_general_fund_org-main_check") && $("#input_01-3_group").addClass("collapse").removeClass("show");
   a = "ibc_general_fund_org_name";
   (t = e.getItemData(a)).applyTextValue();
   a = "ibc_general_fund_conflict";
   null != (t = e.getItemData(a)) && t.makeHtmlRadioType();
   a = "ibc_general_fund_start_date";
   (t = e.getItemData(a)).applyCalanderValue(a);
   a = "ibc_general_fund_end_date";
   (t = e.getItemData(a)).applyCalanderValue(a);
   a = "ibc_general_experiment";
   var t,
      n = (t = e.getItemData(a)).getStringValue("0"),
      i = [];
   "" != n && (i = n.split(",")), i.length > 0 && ibcResetLeftNavi(i), e.deleteSaveTag(["ibc_general_experiment"]), $(".card").removeClass("hidden");
}
function changePageNavigation() {
   g_research_type = $("input:radio[name ='research_type_check']:checked").val();
   if (g_research_type) {
      appNaviNextPage();
   }
}
function removeFileInput(id) {
   $("#box_04-1-" + id).remove();
}
function addFileInput() {
   const index = $("#doc_required_expt_form input").length;
   $("#doc_required_expt_form").append(`
                        <div class="flexMid mTop15 w500" id="box_04-1-${index + 1}">
                          <div class="custom-file">
                            <input type="file" name="file_04-1-${index + 1}" class="custom-file-input" id="file_04-1-${index + 1}">
                            <label class="custom-file-label" for="file_04-1-${index + 1}" data-button-text="첨부 File 선택">File을 선택해 주세요.</label>
                          </div>
                          <a href="javascript:void(0);" onclick="removeFileInput(${
                             index + 1
                          })" class="btn btn-outline-danger mLeft5 btn-file-delete" title="첨부 File 삭제"><i class="far fa-trash-alt"></i></a>
                        </div>`);
}
"undefined" == typeof api_js && document.write("<script src='/assets/js/common/api.js'></script>"),
   "undefined" == typeof const_js && document.write("<script src='/assets/js/common/const.js'></script>"),
   document.write("<script src='/html/researcher/IRB/common_application_new-IRB.js'></script>"),
   $(document).ready(function () {
      function t(t) {
         var a = "research_type_check";
         n = t.getItemData(a).getStringValue("0");

         if (a == "research_type_check") {
            saveData = t.getItemData(a).getSelectValue(0);
            $(`input:radio[name ='research_type_check']:input[value='${saveData}']`).click(), irbResetLeftNavi(n);
         }
         a = "research_danger_check";
         if (a == "research_danger_check") {
            saveData = t.getItemData(a).getSelectValue(0);
            $(`input:radio[name ='research_danger_check']:input[value='${saveData}']`).click(), irbResetLeftNavi(n);
         }
         a = "research_field_check";
         if (a == "research_field_check") {
            saveData = t.getItemData(a).getSelectValue(0);
            $(`input:radio[name ='research_field_check']:input[value='${saveData}']`).click()
            if (saveData == 4) {
               saveData = t.getItemData(a).getSelectInputValue(0, 4, 0);
               $("input[name='research_field_check_input4']").val(saveData);
            }
            irbResetLeftNavi(n);
         }
         a = "research_institution_check";
         if (a == "research_institution_check") {
            saveData = t.getItemData(a).getSelectValue(0);
            $(`input:radio[name ='research_institution_check']:input[value='${saveData}']`).click()
            if (saveData == 2) {
               saveData = t.getItemData(a).getSelectInputValue(0, 2, 0);
               $("input[name='research_institution_check_input2']").val(saveData);
            }
            irbResetLeftNavi(n);
         }

         a = "data_monitoring_check";
         if (a == "data_monitoring_check") {
            saveData = t.getItemData(a).getSelectValue(0);
            for (const id of saveData) {
               $(`input:checkbox[name='data_monitoring_check'][value='${id}']`).prop("checked", true);
               irbResetLeftNavi(n);
            }
         }

         // (a = "general_human_research"), (n = t.getItemData(a).getStringValue("0"));
         // (g_for_people = n), 1 == n ? $('input:checkbox[id="general_human_research"]').attr("checked", !0) : $('input:checkbox[id="general_human_research"]').attr("checked", !1);
         // a = "general_body_research";
         // 1 == (n = t.getItemData(a).getStringValue("0")) ? $('input:checkbox[id="general_body_research"]').attr("checked", !0) : $('input:checkbox[id="general_body_research"]').attr("checked", !1);

         //general_judgement
         (a = "general_judgement"), (n = t.getItemData(a).getStringValue("0"));
         $(`input:radio[name ='general_judgement']:input[value='${n}']`).click(), $(".card").removeClass("hidden");

         //general_object
         $('select[name="general_object"]').on("change", function () {
            const val = $(this).val();
            const $opt1 = $("#showoption1");
            const $opt2 = $("#showoption2");

            $opt1.removeClass("show").addClass("hidden").hide();
            $opt2.removeClass("show").addClass("hidden").hide();

            if (val === "1") {
               $opt1.removeClass("hidden").addClass("show").fadeIn();
            }

            if (val === "2") {
               if (!$opt2.hasClass("show")) {
                  $opt2.removeClass("hidden").addClass("show").fadeIn();
                  $("#doc_required_expt_form").append(`
                     <div class="flexMid mTop15 w500" id="box_04-1-1">
                        <div class="custom-file">
                           <input type="file" name="file_04-1-1" class="custom-file-input" id="file_04-1-1">
                           <label class="custom-file-label" for="file_04-1-1" data-button-text="첨부 File 선택">File을 선택해 주세요.</label>
                        </div>
                        <a href="javascript:void(0);" onclick="removeFileInput(1)" class="btn btn-outline-danger mLeft5 btn-file-delete" title="첨부 File 삭제">
                           <i class="far fa-trash-alt"></i>
                        </a>
                     </div>
                     `);
               }
            }
         });
         (a = "general_object"), (n = t.getItemData(a).getStringValue("0"));
         console.log("n => general_object", n);
         $('select[name="general_object"]').val(n).trigger("change"); // trigger => call change event
      }

      loadApplicationParams(),
         (g_AppItemParser = new ItemParser(g_AppInfo.appSeq)),
         g_AppItemParser.load(
            {
               "filter.query_items":
                  "research_type_check,research_danger_check,research_field_check,research_institution_check,data_monitoring_check,research_field_etc_input,general_judgement,general_object ",
            },
            t
         ),
         $(".reset_review_type").change(function () {
            1 == $(this).val()
               ? ($(".process_content").eq(1).attr("onclick", ""), $(".process_content").eq(2).attr("onclick", ""))
               : ($(".process_content").eq(1).attr("onclick", "APP_IRB_NAVIGATION.navigate('PAGE_2_1');"), $(".process_content").eq(2).attr("onclick", "APP_IRB_NAVIGATION.navigate('PAGE_3_1');")),
               e();
         });

      var a = function (e, t) {
         (this.phase = "page1_3"), (this.btn_type = t);
      };
      (a.prototype.check = function () {
         if ("nextPage" == this.btn_type) {
            /*if (
          !$("#general_human_research").is(":checked") &&
          !$("#general_body_research").is(":checked")
        )
          return void alert(
            "인간 대상 연구 혹은 인체 유래물 연구 중 반드시 하나 이상 체크해 주세요."
          );*/
            changePageNavigation();
         } else {
            /*if (
          !$("#general_human_research").is(":checked") &&
          !$("#general_body_research").is(":checked")
        )
          return void alert(
            "인간 대상 연구 혹은 인체 유래물 연구 중 반드시 하나 이상 체크해 주세요."
          );*/
            saveTemporary();
         }
      }),
         di
            .autowired(!1)
            .register("validator_nextPage")
            .as(a)
            .withConstructor()
            .withProperties()
            .prop("btn_type")
            .val("nextPage")
            .register("validator_tempSave")
            .as(a)
            .withConstructor()
            .withProperties()
            .prop("btn_type")
            .val("tempSave");

      // Show
      $(".IRB_review_target").removeClass("hidden");
   });
