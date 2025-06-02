const ipsap_item_parser_js = true;

if (typeof ipsap_item_data_js == "undefined") document.write("<script src='/assets/js/ipsap/ipsap_item_data.js'></script>");

function makeEduCourse(user_seq) {
   if (!$(`.course_${user_seq}`).eq(1).hasClass("show")) {
      return "";
   }

   let result = new Object();
   $(`.course_${user_seq}`)
      .children("div")
      .each(function () {
         result[$.trim($(this).eq(0).text())] = $(this).children("input").val();
      });
   return JSON.stringify(result);
}

class ItemParser {
   constructor(application_seq) {
      this.app_seq = application_seq;
      this.managedItems = {};
      this.targetSavedItemNames = null;
      this.moreSaveFiles = {};
   }

   load(param, callbackFunc) {
      if (IPSAP.DEMO_MODE) return;

      let class_this = this;
      let app_seq = this.app_seq;
      API.load({
         url: eval("`" + CONST.API.APPLICATION.GETINFO + "`"),
         type: CONST.API_TYPE.GET,
         data: param,
         success: function (data) {
            console.log("data api =>", data);
            if (data.rt == "ok") {
               class_this.dataObj = data.data;
               class_this.guideObj = data.guide;
               commonGuideMapping(data.guide);
               if (callbackFunc != undefined) callbackFunc(class_this);
            }
         },
      });
   }

   save(submitFlag, callbackFunc, userData) {
      var main_this = this;
      var sendingObj = this.getSendingObj();
      console.log("sendingObj", sendingObj);
      var sendingJson = JSON.stringify(sendingObj);
      var queryParamStr = "";

      var formData = new FormData();
      formData.append("param", sendingJson);
      this.setFileObj(formData);

      let urlPath = CONST.API.APPLICATION.PATCH;
      let method = CONST.API_TYPE.PATCH;
      let app_seq = this.app_seq;
      if (submitFlag) {
         urlPath = CONST.API.APPLICATION.POST;
         method = CONST.API_TYPE.POST;
      }
      if ("" != userData) {
         queryParamStr = "?" + $.param({ "filter.submit_type": userData });
      }

      API.load({
         url: urlPath.replace("${app_seq}", app_seq) + queryParamStr,
         type: method,
         enctype: "multipart/form-data",
         data: formData,
         contentType: false,
         processData: false,
         success: function (data) {
            main_this.resetFileItems(data);
            if (callbackFunc != undefined) callbackFunc(true, data, userData);
         },
         complete: function (data) {},
         error: function (data) {
            if (callbackFunc != undefined) {
               callbackFunc(false, data.responseJSON, userData);
            }
         },
      });
   }

   addMoreSaveTag(item_name, data) {
      if (this.moreTag == undefined) this.moreTag = {};

      this.moreTag[item_name] = data;
   }

   deleteSaveTag(item_arr) {
      for (var i in item_arr) {
         var targetName = item_arr[i];
         delete this.managedItems[targetName];
      }
   }

   hasParentItem() {
      return this.dataObj.parent_item != undefined;
   }

   getItemData(item_name, no_save) {
      return this.getItemFromList(this.dataObj, item_name, no_save);
   }

   getParentItemData(item_name, no_save) {
      return this.getItemFromList(this.dataObj.parent_item, item_name, no_save);
   }

   getChildItemData(item_name, no_save) {
      return this.getItemFromList(this.dataObj.child_item, item_name, no_save);
   }

   addNewKeyword(item_name, keyword) {
      if (isInArray(this.managedItems[item_name]["keyword"], keyword)) return;

      this.managedItems[item_name]["keyword"].push(keyword);
      this.managedItems[item_name]["keyword"].sort();
   }

   refreshKeyword(item_name) {
      var item_data = this.getItemData(item_name);
      item_data.refreshKeywordUI();
   }

   removeKeyword(item_name, keyword) {
      this.managedItems[item_name]["keyword"] = jQuery.grep(this.managedItems[item_name]["keyword"], function (value) {
         return value != keyword;
      });
   }

   hasMemberExists(user_seq) {
      if (this.managedItems.member_tag == undefined) return false;

      for (let i = 0; i < this.managedItems.member_tag.length; ++i) {
         let item_name = this.managedItems.member_tag[i];
         var members = this.managedItems[item_name].member;
         if (members == undefined) continue;

         for (let j = 0; j < members.length; ++j) {
            if (Number(members[j].user_seq) == Number(user_seq)) return true;
         }
      }

      return false;
   }

   insertMember(item_name, user_data) {
      if (this.managedItems.member_tag == undefined) return;

      if (this.managedItems[item_name]["member"] == undefined) this.managedItems[item_name]["member"] = [];

      var animal_mng_flag;
      if (this.managedItems.general_expt.member.length == 0) {
         animal_mng_flag = 1;
      } else {
         animal_mng_flag = 0;
      }

      var newData = {
         animal_mng_flag: animal_mng_flag,
         exp_year_code: 0,
         user_seq: user_data.user_seq,
         info: user_data,
      };
      this.managedItems[item_name].member.push(newData);
   }

   insertMemberIBC(item_name, user_data) {
      if (this.managedItems.member_tag == undefined) return;

      if (this.managedItems[item_name]["member"] == undefined) this.managedItems[item_name]["member"] = [];

      var animal_mng_flag;
      if (this.managedItems.general_expt.member.length == 0) {
         animal_mng_flag = 1;
      } else {
         animal_mng_flag = 0;
      }

      var newData = {
         animal_mng_flag: animal_mng_flag,
         exp_year_code: 0,
         edu_course: "",
         user_seq: user_data.user_seq,
         info: user_data,
      };
      this.managedItems[item_name].member.push(newData);
   }

   insertMemberIRB(item_name, user_data) {
      if (this.managedItems.member_tag == undefined) return;

      if (this.managedItems[item_name]["member"] == undefined) this.managedItems[item_name]["member"] = [];

      var newData = {
         exp_type_code: 0,
         exp_year_code: 0,
         user_seq: user_data.user_seq,
         info: user_data,
      };
      this.managedItems[item_name].member.push(newData);
   }

   removeMembers(item_name, user_seq) {
      if (this.managedItems.member_tag == undefined) return;

      var members = this.managedItems[item_name].member;
      if (members == undefined) return;

      var newList = [];
      for (let i = 0; i < members.length; ++i) {
         let user = members[i];

         if (user.user_seq != user_seq) newList.push(user);
      }
      this.managedItems[item_name].member = newList;
   }

   addMultiUIIndex(item_name, idx) {
      if (this.managedItems[item_name] == undefined) this.managedItems[item_name] = {};

      if (this.managedItems[item_name].ui_index == undefined) this.managedItems[item_name].ui_index = [];

      this.managedItems[item_name].ui_index.push(idx);
   }

   removeMultiUIIndex(item_name, idx) {
      if (this.managedItems[item_name].ui_index == undefined) return;

      this.managedItems[item_name].ui_index = jQuery.grep(this.managedItems[item_name].ui_index, function (value) {
         return value != idx;
      });
   }

   getItemFromList(listObj, item_name, no_save) {
      // console.log("listObj, item_name, no_save", listObj, item_name, no_save)
      try {
         for (const [key, value] of Object.entries(listObj)) {
            if (key == item_name) return new ItemData(item_name, value, this, no_save);
            if (value.info != undefined && value.info.item_type_str == "item_group") {
               let ret = this.getItemFromList(value.sub_items, item_name, no_save);
               if (ret != undefined) return ret;
            }
         }
      } catch (error) {}
      return undefined;
   }

   dataValidation() {
      for (const [key, value] of Object.entries(this.managedItems)) {
         if (this.targetSavedItemNames != null) {
            if (!isInArray(this.targetSavedItemNames, key)) continue;
         }

         var old_item_data = this.getItemData(key);
         if (old_item_data == undefined) continue;

         if (old_item_data.dataObj.info.item_type_str == "animal") {
            var dataArr = [];
            if (value.ui_index != undefined) {
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }

               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let aniData = {};

                  aniData.animal_code = $("#" + getIdWithIdx(`animal_select`, ui_index)).val();
                  if (aniData.animal_code == null) {
                     return { obj: $("#" + getIdWithIdx(`animal_select`, ui_index)), msg: "동물 종류를 선택하세요." };
                  } else if (aniData.animal_code == "new") {
                     var inputObj = $("#" + getIdWithIdx(`animal_select`, ui_index))
                        .parent()
                        .children("input");
                     if (inputObj.val() == "") return { obj: inputObj, msg: "동물 종류를 입력하세요." };
                  }

                  aniData.male_cnt = $("#" + getIdWithIdx("male_cnt", ui_index)).val();
                  aniData.female_cnt = $("#" + getIdWithIdx("female_cnt", ui_index)).val();
                  if (aniData.male_cnt == "") aniData.male_cnt = 0;
                  if (aniData.female_cnt == "") aniData.female_cnt = 0;
                  if (Number(aniData.male_cnt) + Number(aniData.female_cnt) == 0) return { obj: $("#" + getIdWithIdx(`male_cnt`, ui_index)), msg: "수량을 입력하세요." };

                  aniData.mb_grade = $("#" + getIdWithIdx(`mb_grade_select`, ui_index)).val();
                  if (aniData.mb_grade == null) {
                     return { obj: $("#" + getIdWithIdx(`mb_grade_select`, ui_index)), msg: "미생물학적 등급을 선택하세요." };
                  } else if (aniData.mb_grade == "new") {
                     var inputObj = $("#" + getIdWithIdx(`mb_grade_select`, ui_index))
                        .parent()
                        .children("input");
                     if (inputObj.val() == "") return { obj: inputObj, msg: "미생물학적 등급을 입력하세요." };
                  }

                  aniData.breeding_place = $("#" + getIdWithIdx(`breeding_place_select`, ui_index)).val();
                  if (aniData.breeding_place == null) {
                     return { obj: $("#" + getIdWithIdx(`breeding_place_select`, ui_index)), msg: "사육 희망장소를 선택하세요." };
                  } else if (aniData.breeding_place == "new") {
                     var inputObj = $("#" + getIdWithIdx(`breeding_place_select`, ui_index))
                        .parent()
                        .children("input");
                     if (inputObj.val() == "") return { obj: inputObj, msg: "사육 희망장소를 입력하세요." };
                  }

                  aniData.strain = $("#" + getIdWithIdx("strain", ui_index)).val();
                  if (aniData.strain == "") return { obj: $("#" + getIdWithIdx(`strain`, ui_index)), msg: "계통을 입력하세요." };

                  aniData.week_age = $("#" + getIdWithIdx("week_age", ui_index)).val();
                  if (aniData.week_age == "") return { obj: $("#" + getIdWithIdx(`week_age`, ui_index)), msg: "주령 또는 월령을 입력하세요." };

                  aniData.weight_gram = $("#" + getIdWithIdx("weight_gram", ui_index)).val();
                  if (aniData.weight_gram == "") return { obj: $("#" + getIdWithIdx(`weight_gram`, ui_index)), msg: "체중을 입력하세요." };

                  aniData.supplier_type = $("#" + getIdWithIdx(`supplier_type_select`, ui_index)).val();
                  if (aniData.supplier_type == null) {
                     return { obj: $("#" + getIdWithIdx(`supplier_type_select`, ui_index)), msg: "실험동물 공급처를 선택하세요." };
                  }

                  aniData.supplier_name = $("#" + getIdWithIdx("supplier_name", ui_index)).val();
                  if (aniData.supplier_name == "") return { obj: $("#" + getIdWithIdx(`supplier_name`, ui_index)), msg: "실험 동물 공급처를 입력하세요." };

                  var lmo_checked = $("#" + getIdWithIdx("lmo_flag", ui_index)).is(":checked");
                  if (!lmo_checked) {
                     aniData.ibc_num = $("#" + getIdWithIdx("ibc_num", ui_index)).val();

                     aniData.lmo_type = $("#" + getIdWithIdx(`lmo_type_select`, ui_index)).val();
                     if (aniData.lmo_type == null) {
                        return { obj: $("#" + getIdWithIdx(`lmo_type_select`, ui_index)), msg: "LMO 시설구분을 선택하세요." };
                     }
                  }
               }
            }
            continue;
         }

         if (value.string != undefined) {
            if (value.ui_index != undefined) {
               let values = Object.values(value.string);
               if (values.length > 0) {
                  for (let l = 0; l < value.ui_index.length; ++l) {
                     let tag_id = `${values[0]}_${value.ui_index[l]}`;
                     let strObj = $(`#${tag_id}`);
                     if (strObj.length > 0) {
                        try {
                           if (old_item_data.dataObj.items[String(l)].must > 0) {
                              let text_value = strObj.val();
                              if (old_item_data.dataObj.items[String(l)].number_only > 0 && text_value == "0") text_value = "";

                              if (text_value == "" && isObjVisible(strObj)) return { obj: strObj, msg: "필수항목을 입력해 주세요." };
                           }
                        } catch (e) {}
                     }
                  }
               }
            } else {
               for (const [s_key, s_value] of Object.entries(value.string)) {
                  let strObj = $(`#${s_value}`);
                  if (strObj.length > 0) {
                     try {
                        if (old_item_data.dataObj.items[String(s_key)].must > 0) {
                           let text_value = strObj.val();
                           if (old_item_data.dataObj.items[String(s_key)].number_only > 0 && text_value == "0") text_value = "";

                           if (text_value == "" && isObjVisible(strObj)) return { obj: strObj, msg: "필수항목을 입력해 주세요." };
                        }
                     } catch (e) {}
                  }
               }
            }
         }

         if (old_item_data.itemName == "substance_dosage_flag") {
            if (value.main_check_id != undefined) {
               if (value.ui_index != undefined) {
                  let tag_id = ``;
                  var tokens = value.main_check_id.split("_");
                  for (let l = 0; l < tokens.length - 1; ++l) tag_id += tokens[l] + "_";
                  for (let l = 0; l < value.ui_index.length; ++l) {
                     let objid = `${tag_id}${value.ui_index[l]}`;
                     let chkObj = $(`input:checkbox[id='${objid}']`);
                     if (1 == getNumberFromBool(chkObj.is(":checked"))) {
                        let checkId = `${"substance_dosage_name"}_${value.ui_index[l]}`;
                        let checkObj = $(`#${checkId}`);
                        if ("" == checkObj.val()) {
                           return { obj: checkObj, msg: "실험 물질 투여 유무가 있을 경우 투여물질명은 필수 입력 사항입니다." };
                        }
                     }
                  }
               }
            }
         }

         if (old_item_data.dataObj.info.item_type_str == "anesthetic") {
            let chkObj = $(`input:checkbox[id='${value.main_check_id}']`);
            if (1 == getNumberFromBool(chkObj.is(":checked"))) {
               if (!value.ui_index.length) {
                  return { obj: null, msg: "약물을 추가 해주시기 바랍니다." };
               }
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }
               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let tmpData = {};
                  var anesthetic_type = `anesthetic_type`;
                  var anesthetic_name = `anesthetic_name`;
                  var injection_mg = `injection_mg`;
                  var injection_route = `injection_route`;

                  if ("pain_relief_veterinary_mng" == old_item_data.itemName) {
                     tmpData.injection_time = $("#" + getIdWithIdx("injection_time", ui_index)).val();
                     tmpData.injection_cnt = $("#" + getIdWithIdx("injection_cnt", ui_index)).val();
                     if (!tmpData.injection_time) {
                        return { obj: null, msg: "투여 시점을 입력해 주세요" };
                     }

                     if (!tmpData.injection_cnt) {
                        return { obj: null, msg: "투여 횟수를 입력해 주세요" };
                     }
                     anesthetic_type += "2";
                     anesthetic_name += "2";
                     injection_mg += "2";
                     injection_route += "2";
                  }

                  tmpData.anesthetic_type = $("#" + getIdWithIdx(anesthetic_type, ui_index)).val();

                  if (!tmpData.anesthetic_type) {
                     return { obj: null, msg: "약물 종류를 선택해 주세요" };
                  }

                  if (tmpData.anesthetic_type == "new") {
                     tmpData.anesthetic_type = 0;
                     var inputObj = $("#" + getIdWithIdx(anesthetic_type, ui_index))
                        .parent()
                        .children("input");
                     tmpData.anesthetic_type_str = inputObj.val();
                     if (!tmpData.anesthetic_type_str) {
                        return { obj: null, msg: "약물 종류를 입력해 주세요" };
                     }
                  } else tmpData.anesthetic_type_str = $("#" + getIdWithIdx(anesthetic_type, ui_index) + " option:selected").text();

                  if (!tmpData.anesthetic_type_str) {
                     return { obj: null, msg: "약물 종류를 입력해 주세요" };
                  }

                  tmpData.injection_route = $("#" + getIdWithIdx(injection_route, ui_index)).val();
                  if (!tmpData.injection_route) {
                     return { obj: null, msg: "투여 경로를 선택해 주세요" };
                  }

                  if (tmpData.injection_route == "new") {
                     tmpData.injection_route = 0;
                     var inputObj = $("#" + getIdWithIdx(injection_route, ui_index))
                        .parent()
                        .children("input");
                     tmpData.injection_route_str = inputObj.val();
                     if (!tmpData.injection_route_str) {
                        return { obj: null, msg: "투여 경로를 입력해 주세요" };
                     }
                  } else tmpData.injection_route_str = $("#" + getIdWithIdx(injection_route, ui_index) + " option:selected").text();

                  if (!tmpData.injection_route_str) {
                     return { obj: null, msg: "투여 경로를 입력해 주세요" };
                  }

                  tmpData.anesthetic_name = $("#" + getIdWithIdx(anesthetic_name, ui_index)).val();
                  if (!tmpData.anesthetic_name) {
                     return { obj: null, msg: "약물명을 입력해 주세요" };
                  }
                  tmpData.injection_mg = $("#" + getIdWithIdx(injection_mg, ui_index)).val();
                  if (!tmpData.injection_mg) {
                     return { obj: null, msg: "투여량을 입력해 주세요" };
                  }
               }
            }
         }

         if (old_item_data.itemName == "ibc_general_fclty") {
            function getIdWithIdx(id, idx) {
               return `${id}_${idx}`;
            }
            for (let l = 0; l < value.ui_index.length; ++l) {
               let ui_index = value.ui_index[l];
               let tmpData = {};
               var ibc_general_fclty_biosafety = `ibc_general_fclty_biosafety`;
               var ibc_general_fclty_type = `ibc_general_fclty_type`;
               var ibc_general_fclty_grade = `ibc_general_fclty_grade`;
               var ibc_general_fclty_lmo = `ibc_general_fclty_lmo`;

               tmpData.ibc_general_fclty_biosafety = $("#" + getIdWithIdx(ibc_general_fclty_biosafety, ui_index)).val();
               if (!tmpData.ibc_general_fclty_biosafety) {
                  return { obj: null, msg: "생물안전 연구시설 번호를 입력해 주세요" };
               }

               tmpData.ibc_general_fclty_type = $("#" + getIdWithIdx(ibc_general_fclty_type, ui_index)).val();

               if (!tmpData.ibc_general_fclty_type) {
                  return { obj: null, msg: "시설 종류를 선택해 주세요" };
               }

               if (tmpData.ibc_general_fclty_type == "new") {
                  tmpData.ibc_general_fclty_type = 0;
                  var inputObj = $("#" + getIdWithIdx(ibc_general_fclty_type, ui_index))
                     .parent()
                     .children("input");
                  tmpData.ibc_general_fclty_type_str = inputObj.val();
                  if (!tmpData.ibc_general_fclty_type_str) {
                     return { obj: null, msg: "시설 종류를 입력해 주세요" };
                  }
               } else tmpData.ibc_general_fclty_type_str = $("#" + getIdWithIdx(ibc_general_fclty_type, ui_index) + " option:selected").text();

               if (!tmpData.ibc_general_fclty_type_str) {
                  return { obj: null, msg: "시설 종류를 입력해 주세요" };
               }

               tmpData.ibc_general_fclty_grade = $("#" + getIdWithIdx(ibc_general_fclty_grade, ui_index)).val();
               if (!tmpData.ibc_general_fclty_grade) {
                  return { obj: null, msg: "안전 관리 등급을 선택해 주세요" };
               }

               tmpData.ibc_general_fclty_lmo = $("#" + getIdWithIdx(ibc_general_fclty_lmo, ui_index)).val();
               if (!tmpData.ibc_general_fclty_lmo) {
                  return { obj: null, msg: "LMO 연구시설을 선택해 주세요" };
               }
            }
         }

         if (old_item_data.itemName == "ibc_risk_bios_infection_chance") {
            if (value.ui_index != undefined) {
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }
               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let tmpData = {};
                  var ibc_exposure_path = `ibc_exposure_path`;
                  var ibc_inspection_possibility = `ibc_inspection_possibility`;
                  var ibc_first_aid_method = `ibc_first_aid_method`;

                  tmpData.ibc_exposure_path = $("#" + getIdWithIdx(ibc_exposure_path, ui_index)).val();

                  if (!tmpData.ibc_exposure_path) {
                     return { obj: null, msg: "노출 경로를 선택해 주세요" };
                  }

                  if (tmpData.ibc_exposure_path == "new") {
                     tmpData.ibc_exposure_path = 0;
                     var inputObj = $("#" + getIdWithIdx(ibc_exposure_path, ui_index))
                        .parent()
                        .children("input");
                     tmpData.ibc_exposure_path_str = inputObj.val();
                     if (!tmpData.ibc_exposure_path_str) {
                        return { obj: null, msg: "노출 경로를 입력해 주세요" };
                     }
                  } else tmpData.ibc_exposure_path_str = $("#" + getIdWithIdx(ibc_exposure_path, ui_index) + " option:selected").text();

                  tmpData.ibc_inspection_possibility = $("#" + getIdWithIdx(ibc_inspection_possibility, ui_index)).val();

                  if (!tmpData.ibc_inspection_possibility) {
                     return { obj: null, msg: "감염 가능성을 선택해 주세요" };
                  }

                  tmpData.ibc_first_aid_method = $("#" + getIdWithIdx(ibc_first_aid_method, ui_index)).val();
                  if (!tmpData.ibc_first_aid_method) {
                     return { obj: null, msg: "응급 조치 방법을 입력해 주세요" };
                  }
               }
            }
         }

         if (value.member != undefined) {
            if (key == "general_director") {
               for (let i = 0; i < value.member.length; ++i) {
                  if ([value.member[i].exp_year_code] == 0) return { obj: null, msg: `연구책임자(${value.member[i].info.name})의 경력 년수를 선택하세요.` };
               }
            } else if (key == "general_expt") {
               if (value.member.length == 0) {
                  return { obj: null, msg: `실험수행자를 1명 이상 선택하세요.` };
               }
               for (let i = 0; i < value.member.length; ++i) {
                  if ([value.member[i].exp_year_code] == 0) return { obj: null, msg: `실험수행자(${value.member[i].info.name})의 경력 년수를 선택하세요.` };
               }
            }
         }

         if (value.radio != undefined) {
            for (const [s_key, s_item_name] of Object.entries(value.radio)) {
               let name = s_item_name + "_" + s_key;
               let radioObj = $(`input[name="${name}"]:checked`);
               if (radioObj.length > 0) {
                  let id = radioObj.val();
                  try {
                     for (let ii = 0; ii < old_item_data.dataObj.items.length; ++ii) {
                        if (old_item_data.dataObj.items[ii].id == String(id)) {
                           let number_only_arr = old_item_data.dataObj.items[ii].input.number_only;
                           let input_must_arr = old_item_data.dataObj.items[ii].input.must;
                           for (let col = 0; col < old_item_data.dataObj.items[ii].input.cnt; ++col) {
                              let tag_id = s_item_name + "_" + s_key + "_" + String(id) + "_input_" + String(col + 1);
                              let inputObj = $(`#${tag_id}`);
                              if (inputObj.length > 0) {
                                 if (input_must_arr[col] > 0) {
                                    let text_value = inputObj.val();
                                    if (number_only_arr[col] && text_value == "0") text_value = "";
                                    if (text_value == "" && isObjVisible(inputObj)) return { obj: inputObj, msg: "필수항목을 입력해 주세요." };
                                 }
                              }
                           }
                        }
                     }
                  } catch (e) {}
               }
            }
         }

         if (value.checkbox != undefined) {
            function checkValidateData(item_idx, tag_name) {
               let selected = [];
               $(`input[name=${tag_name}]:checked`).each(function () {
                  selected.push($(this).val());
               });

               for (let i = 0; i < selected.length; ++i) {
                  let id = selected[i];
                  try {
                     for (let ii = 0; ii < old_item_data.dataObj.items.length; ++ii) {
                        if (old_item_data.dataObj.items[ii].id == String(id)) {
                           let number_only_arr = old_item_data.dataObj.items[ii].input.number_only;
                           let input_must_arr = old_item_data.dataObj.items[ii].input.must;
                           for (let col = 0; col < old_item_data.dataObj.items[ii].input.cnt; ++col) {
                              let tag_id = tag_name + "_" + String(id) + "_input_" + String(col + 1);
                              let inputObj = $(`#${tag_id}`);
                              if (inputObj.length > 0) {
                                 if (input_must_arr[col] > 0) {
                                    let text_value = inputObj.val();
                                    if (number_only_arr[col] && text_value == "0") text_value = "";
                                    if (text_value == "" && isObjVisible(inputObj)) return { obj: inputObj, msg: "필수항목을 입력해 주세요." };
                                 }
                              }
                           }
                        }
                     }
                  } catch (e) {}
               }
               return {};
            }

            if (value.ui_index != undefined) {
               let values = Object.values(value.checkbox);
               if (values.length > 0) {
                  for (let l = 0; l < value.ui_index.length; ++l) {
                     let tag_name = `${values[0]}_${value.ui_index[l] - 1}`;
                     let ret = checkValidateData(l, tag_name);
                     if (Object.keys(ret).length > 0) return ret;
                  }
               }
            } else {
               for (const [s_key, s_item_name] of Object.entries(value.checkbox)) {
                  let tag_name = s_item_name + "_" + s_key;
                  let ret = checkValidateData(s_key, tag_name);
                  if (Object.keys(ret).length > 0) return ret;
               }
            }
         }
      }
      return {};
   }

   getSendingObj() {
      var ret = {};
      if (this.moreTag != undefined) {
         for (const [key, value] of Object.entries(this.moreTag)) {
            ret[key] = value;
         }
      }

      for (const [key, value] of Object.entries(this.managedItems)) {
         if (this.targetSavedItemNames != null) {
            if (!isInArray(this.targetSavedItemNames, key)) continue;
         }

         let dataObj = {};

         var old_item_data = this.getItemData(key);
         console.log("old_item_data", old_item_data);
         console.log("key", key);

         if (old_item_data == undefined) continue;

         if (old_item_data.dataObj.info.item_type_str == "animal") {
            var dataArr = [];
            if (value.ui_index != undefined) {
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }

               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let aniData = {};

                  aniData.animal_code = $("#" + getIdWithIdx(`animal_select`, ui_index)).val();
                  if (aniData.animal_code == "new") {
                     aniData.animal_code = 0;
                     var inputObj = $("#" + getIdWithIdx(`animal_select`, ui_index))
                        .parent()
                        .children("input");
                     aniData.animal_code_str = inputObj.val();
                  } else aniData.animal_code_str = $("#" + getIdWithIdx(`animal_select`, ui_index) + " option:selected").text();

                  aniData.mb_grade = $("#" + getIdWithIdx(`mb_grade_select`, ui_index)).val();
                  if (aniData.mb_grade == "new") {
                     aniData.mb_grade = 0;
                     var inputObj = $("#" + getIdWithIdx(`mb_grade_select`, ui_index))
                        .parent()
                        .children("input");
                     aniData.mb_grade_str = inputObj.val();
                  } else aniData.mb_grade_str = $("#" + getIdWithIdx(`mb_grade_select`, ui_index) + " option:selected").text();

                  aniData.breeding_place = $("#" + getIdWithIdx(`breeding_place_select`, ui_index)).val();
                  if (aniData.breeding_place == "new") {
                     aniData.breeding_place = 0;
                     var inputObj = $("#" + getIdWithIdx(`breeding_place_select`, ui_index))
                        .parent()
                        .children("input");
                     aniData.breeding_place_str = inputObj.val();
                  } else aniData.breeding_place_str = $("#" + getIdWithIdx(`breeding_place_select`, ui_index) + " option:selected").text();

                  aniData.supplier_type = $("#" + getIdWithIdx(`supplier_type_select`, ui_index)).val();
                  aniData.lmo_type = $("#" + getIdWithIdx(`lmo_type_select`, ui_index)).val();

                  aniData.age_unit = $("#" + getIdWithIdx(`age_unit`, ui_index)).val();
                  aniData.weight_unit = $("#" + getIdWithIdx(`weight_unit`, ui_index)).val();
                  aniData.size_unit = $("#" + getIdWithIdx(`size_unit`, ui_index)).val();

                  aniData.male_cnt = $("#" + getIdWithIdx("male_cnt", ui_index)).val();
                  aniData.female_cnt = $("#" + getIdWithIdx("female_cnt", ui_index)).val();
                  aniData.strain = $("#" + getIdWithIdx("strain", ui_index)).val();
                  aniData.week_age = $("#" + getIdWithIdx("week_age", ui_index)).val();
                  aniData.weight_gram = $("#" + getIdWithIdx("weight_gram", ui_index)).val();
                  aniData.supplier_name = $("#" + getIdWithIdx("supplier_name", ui_index)).val();
                  aniData.ibc_num = $("#" + getIdWithIdx("ibc_num", ui_index)).val();
                  aniData.genetic_type = $("#" + getIdWithIdx("genetic_type", ui_index)).val();
                  aniData.size = $("#" + getIdWithIdx("size", ui_index)).val();

                  aniData.lmo_flag = 0;
                  if ($("#" + getIdWithIdx("lmo_flag", ui_index)).is(":checked")) aniData.lmo_flag = 1;

                  if (aniData.male_cnt == "") aniData.male_cnt = 0;
                  if (aniData.female_cnt == "") aniData.female_cnt = 0;
                  if (aniData.week_age == "") aniData.week_age = 0;
                  if (aniData.weight_gram == "") aniData.weight_gram = 0;

                  if (aniData.age_unit == "") aniData.age_unit = 0;
                  if (aniData.weight_unit == "") aniData.weight_unit = 0;
                  if (aniData.size_unit == "") aniData.size_unit = 0;
                  if (aniData.size == "") aniData.size = 0;

                  if (Number(aniData.male_cnt) + Number(aniData.female_cnt) == 0) {
                     alert("동물 수량이 입력되지 않은 항목은 저장하지 않습니다.");
                     continue;
                  }

                  if (aniData.lmo_type == null) aniData.lmo_type = 0;

                  if (aniData.animal_code == null || aniData.mb_grade == null || aniData.breeding_place == null || aniData.supplier_type == null) {
                     alert("값이 지정되지 않은 항목은 저장하지 않습니다.");
                     continue;
                  }

                  dataArr.push(aniData);
               }
            }
            dataObj.data = dataArr;
            ret[key] = dataObj;
            continue;
         }

         if (old_item_data.dataObj.info.item_type_str == "anesthetic") {
            var dataArr = [];
            if (value.ui_index != undefined) {
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }
               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let tmpData = {};
                  var anesthetic_type = `anesthetic_type`;
                  var anesthetic_name = `anesthetic_name`;
                  var injection_mg = `injection_mg`;
                  var injection_route = `injection_route`;

                  if ("pain_relief_veterinary_mng" == old_item_data.itemName) {
                     tmpData.injection_time = $("#" + getIdWithIdx("injection_time", ui_index)).val();
                     tmpData.injection_cnt = $("#" + getIdWithIdx("injection_cnt", ui_index)).val();
                     anesthetic_type += "2";
                     anesthetic_name += "2";
                     injection_mg += "2";
                     injection_route += "2";
                  } else {
                     tmpData.injection_time = $("#" + getIdWithIdx("injection_time1", ui_index)).val();
                     tmpData.injection_cnt = $("#" + getIdWithIdx("injection_cnt1", ui_index)).val();
                  }

                  tmpData.anesthetic_type = $("#" + getIdWithIdx(anesthetic_type, ui_index)).val();
                  if (tmpData.anesthetic_type == "new") {
                     tmpData.anesthetic_type = 0;
                     var inputObj = $("#" + getIdWithIdx(anesthetic_type, ui_index))
                        .parent()
                        .children("input");
                     tmpData.anesthetic_type_str = inputObj.val();
                  } else tmpData.anesthetic_type_str = $("#" + getIdWithIdx(anesthetic_type, ui_index) + " option:selected").text();

                  tmpData.injection_route = $("#" + getIdWithIdx(injection_route, ui_index)).val();
                  if (tmpData.injection_route == "new") {
                     tmpData.injection_route = 0;
                     var inputObj = $("#" + getIdWithIdx(injection_route, ui_index))
                        .parent()
                        .children("input");
                     tmpData.injection_route_str = inputObj.val();
                  } else tmpData.injection_route_str = $("#" + getIdWithIdx(injection_route, ui_index) + " option:selected").text();

                  tmpData.anesthetic_name = $("#" + getIdWithIdx(anesthetic_name, ui_index)).val();
                  tmpData.injection_mg = $("#" + getIdWithIdx(injection_mg, ui_index)).val();

                  if ($(`input:checkbox[id='${value.main_check_id}']`).is(":checked") == true) {
                     if (tmpData.anesthetic_type == null) {
                        alert("값이 지정되지 않은 항목은 저장하지 않습니다.");
                        continue;
                     } else {
                        dataArr.push(tmpData);
                     }
                  }
               }
            }
            let chkData = new Object();
            dataObj.data = dataArr;
            let chkObj = $(`input:checkbox[id='${value.main_check_id}']`);
            if (chkObj != undefined) {
               chkData["0"] = getNumberFromBool(chkObj.is(":checked"));
            }
            dataObj.main_select = chkData;
            ret[key] = dataObj;
            continue;
         }

         if (old_item_data.itemName == "ibc_general_experiment") {
            var dataArr = [];
            $("input[type=checkbox][data-exp-type]").each(function () {
               if (this.checked) {
                  dataArr.push(this.value);
               }
            });

            var strData = new Object();
            strData["0"] = dataArr.toString();
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "ibc_general_fclty") {
            var dataArr = [];
            if (value.ui_index != undefined) {
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }
               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let tmpData = {};
                  var ibc_general_fclty_biosafety = `ibc_general_fclty_biosafety`;
                  var ibc_general_fclty_type = `ibc_general_fclty_type`;
                  var ibc_general_fclty_grade = `ibc_general_fclty_grade`;
                  var ibc_general_fclty_lmo = `ibc_general_fclty_lmo`;

                  tmpData.ibc_general_fclty_biosafety = $("#" + getIdWithIdx(ibc_general_fclty_biosafety, ui_index)).val();
                  tmpData.ibc_general_fclty_type = $("#" + getIdWithIdx(ibc_general_fclty_type, ui_index)).val();
                  if (tmpData.ibc_general_fclty_type == "new") {
                     tmpData.ibc_general_fclty_type = 0;
                     var inputObj = $("#" + getIdWithIdx(ibc_general_fclty_type, ui_index))
                        .parent()
                        .children("input");
                     tmpData.ibc_general_fclty_type_str = inputObj.val();
                  }

                  tmpData.ibc_general_fclty_grade = $("#" + getIdWithIdx(ibc_general_fclty_grade, ui_index)).val();
                  tmpData.ibc_general_fclty_lmo = $("#" + getIdWithIdx(ibc_general_fclty_lmo, ui_index)).val();

                  dataArr.push(JSON.stringify(tmpData));
               }
            }
            let strData = new Object();
            strData[0] = "[" + dataArr.toString() + "]";
            dataObj["data"] = strData;
            ret[key] = dataObj;
            continue;
         }

         if (old_item_data.itemName == "ibc_risk_bios_infection_chance") {
            var dataArr = [];
            if (value.ui_index != undefined) {
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }
               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let tmpData = {};
                  var ibc_exposure_path = `ibc_exposure_path`;
                  var ibc_inspection_possibility = `ibc_inspection_possibility`;
                  var ibc_first_aid_method = `ibc_first_aid_method`;

                  tmpData.ibc_exposure_path = $("#" + getIdWithIdx(ibc_exposure_path, ui_index)).val();
                  if (tmpData.ibc_exposure_path == "new") {
                     tmpData.ibc_exposure_path = 0;
                     var inputObj = $("#" + getIdWithIdx(ibc_exposure_path, ui_index))
                        .parent()
                        .children("input");
                     tmpData.ibc_exposure_path_str = inputObj.val();
                  }

                  tmpData.ibc_inspection_possibility = $("#" + getIdWithIdx(ibc_inspection_possibility, ui_index)).val();
                  tmpData.ibc_first_aid_method = $("#" + getIdWithIdx(ibc_first_aid_method, ui_index)).val();

                  dataArr.push(JSON.stringify(tmpData));
               }
            }
            let strData = new Object();
            strData[0] = "[" + dataArr.toString() + "]";
            dataObj["data"] = strData;
            ret[key] = dataObj;
            continue;
         }

         if (old_item_data.itemName == "ibc_risk_emergency_network") {
            var dataArr = [];
            if (value.ui_index != undefined) {
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }
               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let tmpData = {};

                  var ibc_risk_emergency_name = `ibc_risk_emergency_name`;
                  var ibc_risk_emergency_department = `ibc_risk_emergency_department`;
                  var ibc_risk_emergency_position = `ibc_risk_emergency_position`;
                  var ibc_risk_emergency_phone = `ibc_risk_emergency_phone`;
                  var ibc_risk_emergency_email = `ibc_risk_emergency_email`;

                  tmpData.ibc_risk_emergency_name = $("#" + getIdWithIdx(ibc_risk_emergency_name, ui_index)).val();
                  tmpData.ibc_risk_emergency_department = $("#" + getIdWithIdx(ibc_risk_emergency_department, ui_index)).val();
                  tmpData.ibc_risk_emergency_position = $("#" + getIdWithIdx(ibc_risk_emergency_position, ui_index)).val();
                  tmpData.ibc_risk_emergency_phone = $("#" + getIdWithIdx(ibc_risk_emergency_phone, ui_index)).val();
                  tmpData.ibc_risk_emergency_email = $("#" + getIdWithIdx(ibc_risk_emergency_email, ui_index)).val();

                  dataArr.push(JSON.stringify(tmpData));
               }
            }
            let strData = new Object();
            strData[0] = "[" + dataArr.toString() + "]";
            dataObj["data"] = strData;
            ret[key] = dataObj;
            continue;
         }

         if (old_item_data.itemName == "ibc_risk_bio_grade") {
            var dataArr = [];
            $("input:checkbox[name=ibc_risk_bio_grade]").each(function () {
               if (this.checked) {
                  dataArr.push(this.value);
               }
            });
            var strData = new Object();
            strData["0"] = dataArr.toString();
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "consent_date_radio") {
            var strData = new Object();
            var data = $('input[name="consent_date_radio"]:checked').val();
            strData["0"] = $('input[name="consent_date_radio"]:checked').val();
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "ibc_risk_bio_matters_check") {
            var dataArr = [];
            $("input:checkbox[name=ibc_risk_bio_matters_check]").each(function () {
               if (this.checked) {
                  dataArr.push(this.value);
               }
            });
            var strData = new Object();
            strData["0"] = dataArr.toString();
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "human_experiment_check") {
            var dataArr = [];
            $("input:checkbox[name=human_experiment_check]").each(function () {
               if (this.checked) {
                  dataArr.push(this.value);
               }
            });
            var strData = new Object();
            strData["0"] = dataArr.toString();
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "general_body_research_check") {
            var dataArr = [];
            $("input:checkbox[name=general_body_research_check]").each(function () {
               if (this.checked) {
                  dataArr.push(this.value);
               }
            });
            var strData = new Object();
            strData["0"] = dataArr.toString();
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "animal_exp_substance_method1") {
            var dataArr = [];
            if (value.ui_index != undefined) {
               function getIdWithIdx(id, idx) {
                  return `${id}_${idx}`;
               }
               for (let l = 0; l < value.ui_index.length; ++l) {
                  let ui_index = value.ui_index[l];
                  let tmpData = {};
                  var animal_exp_substance_method1 = `animal_exp_substance_method1`;
                  tmpData.animal_exp_substance_method1 = $("#" + getIdWithIdx(animal_exp_substance_method1, ui_index)).val();
                  if (tmpData.animal_exp_substance_method1 == "new") {
                     tmpData.animal_exp_substance_method1 = 0;
                     var inputObj = $("#" + getIdWithIdx(animal_exp_substance_method1, ui_index))
                        .parent()
                        .children("input");
                     tmpData.animal_exp_substance_method1_str = inputObj.val();
                  }

                  dataArr.push(JSON.stringify(tmpData));
               }
            }
            let strData = new Object();
            strData[0] = "[" + dataArr.toString() + "]";
            dataObj["data"] = strData;
            ret[key] = dataObj;
            continue;
         }
         //START: IRB 01_3
         console.log("old_item_data.itemName", old_item_data.itemName);
         if (old_item_data.itemName == "research_type_check") {
            var strData = new Object();
            var data = $('input[name="research_type_check"]:checked').val();
            strData["0"] = { select_ids: [$('input[name="research_type_check"]:checked').val()] };
            dataObj["data"] = strData;
         }
         if (old_item_data.itemName == "research_danger_check") {
            var strData = new Object();
            var data = $('input[name="research_danger_check"]:checked').val();
            strData["0"] = { select_ids: [$('input[name="research_danger_check"]:checked').val()] };
            dataObj["data"] = strData;
         }
         if (old_item_data.itemName == "research_field_check") {
            var strData = new Object();
            var data = $('input[name="research_field_check"]:checked').val();
            if ($('input[name="research_field_check"]:checked').val() == 4) {
               strData["0"] = {
                  select_ids: [$('input[name="research_field_check"]:checked').val()],
                  inputs: {
                     4: {
                        0: $('input[name="research_field_check_input4"]').val(),
                     },
                  },
               };
            } else {
               strData["0"] = {
                  select_ids: [$('input[name="research_field_check"]:checked').val()],
               };
            }

            dataObj["data"] = strData;
         }
         if (old_item_data.itemName == "research_institution_check") {
            var strData = new Object();
            var data = $('input[name="research_institution_check"]:checked').val();
               if ($('input[name="research_institution_check"]:checked').val() == 2) {
               strData["0"] = {
                  select_ids: [$('input[name="research_institution_check"]:checked').val()],
                  inputs: {
                     2: {
                        0: $('input[name="research_institution_check_input2"]').val(),
                     },
                  },
               };
            } else {
               strData["0"] = {
                  select_ids: [$('input[name="research_institution_check"]:checked').val()],
               };
            }
            dataObj["data"] = strData;
         }
         if (old_item_data.itemName == "data_monitoring_check") {
            var strData = new Object();
            var data = $('input[name="data_monitoring_check"]:checked').val();
            strData["0"] = {
               select_ids: $('input[name="data_monitoring_check"]:checked')
                  .map(function () {
                     return $(this).val();
                  })
                  .get(),
            };
            dataObj["data"] = strData;
         }
         //END: IRB 01_3

         if (old_item_data.itemName == "general_object") {
            var strData = new Object();
            strData["0"] = $('select[name="general_object"]').val();
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "general_judgement") {
            var strData = new Object();
            var data = $('input[name="general_judgement"]:checked').val();
            strData["0"] = $('input[name="general_judgement"]:checked').val();
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "general_human_research") {
            var strData = new Object();
            strData["0"] = $('input:checkbox[id="general_human_research"]').is(":checked");
            dataObj["data"] = strData;
         }

         if (old_item_data.itemName == "general_body_research") {
            var strData = new Object();
            strData["0"] = $('input:checkbox[id="general_body_research"]').is(":checked");
            dataObj["data"] = strData;
         }

         if (key == "general_director" || key == "general_expt") {
            for (const memberObj of value.member) {
               // if ('edu_course' in memberObj) {
               memberObj.edu_course = makeEduCourse(memberObj.user_seq);
               // }
            }
         }

         if (value.main_check_id != undefined) {
            let chkData = new Object();
            if (value.ui_index != undefined) {
               let tag_id = ``;
               var tokens = value.main_check_id.split("_");
               for (let l = 0; l < tokens.length - 1; ++l) tag_id += tokens[l] + "_";

               for (let l = 0; l < value.ui_index.length; ++l) {
                  let objid = `${tag_id}${value.ui_index[l]}`;
                  let chkObj = $(`input:checkbox[id='${objid}']`);
                  if (chkObj.length > 0) chkData[String(l)] = getNumberFromBool(chkObj.is(":checked"));
               }
            } else {
               let chkObj = $(`input:checkbox[id='${value.main_check_id}']`);
               if (chkObj != undefined) {
                  chkData["0"] = getNumberFromBool(chkObj.is(":checked"));
               } else {
                  for (let i = 0; ; ++i) {
                     let objId = value.main_check_id + "_" + String(i);
                     let chkObj = $(`input:checkbox[id='${objId}']`);
                     if (chkObj.length == 0) break;
                     chkData[String(i)] = getNumberFromBool(chkObj.is(":checked"));
                  }
               }
            }
            dataObj["main_select"] = chkData;
         }

         if (value.string != undefined) {
            let strData = new Object();
            if (value.ui_index != undefined) {
               let values = Object.values(value.string);
               if (values.length > 0) {
                  for (let l = 0; l < value.ui_index.length; ++l) {
                     let tag_id = `${values[0]}_${value.ui_index[l]}`;
                     let strObj = $(`#${tag_id}`);
                     if (strObj.length > 0) strData[String(l)] = strObj.val();
                  }
               }
            } else {
               for (const [s_key, s_value] of Object.entries(value.string)) {
                  let strObj = $(`#${s_value}`);
                  if (strObj.length > 0) {
                     strData[s_key] = strObj.val();
                  }
               }
            }
            dataObj["data"] = strData;
         }

         if (value.radio != undefined) {
            let strData = new Object();
            for (const [s_key, s_item_name] of Object.entries(value.radio)) {
               let name = s_item_name + "_" + s_key;
               var data = {};

               let radioObj = $(`input[name="${name}"]:checked`);
               if (radioObj.length > 0) {
                  let id = radioObj.val();
                  data.select_ids = [id];

                  var input_data = {};
                  try {
                     for (let ii = 0; ii < old_item_data.dataObj.items.length; ++ii) {
                        if (old_item_data.dataObj.items[ii].id == String(id)) {
                           var row_input_data = {};
                           for (let col = 0; col < old_item_data.dataObj.items[ii].input.cnt; ++col) {
                              let tag_id = s_item_name + "_" + s_key + "_" + String(id) + "_input_" + String(col + 1);
                              let inputObj = $(`#${tag_id}`);
                              if (inputObj.length > 0) {
                                 let text_value = inputObj.val();
                                 row_input_data[String(col + 1)] = text_value;
                              }
                           }
                           if (Object.keys(row_input_data).length > 0) input_data[String(id)] = row_input_data;
                        }
                     }
                  } catch (e) {}
                  if (Object.keys(input_data).length > 0) data.inputs = input_data;
               }
               strData[s_key] = data;
            }
            dataObj["data"] = strData;
         }

         if (value.checkbox != undefined) {
            let strData = new Object();

            function makeChkData(strData, item_idx, tag_name) {
               let selected = [];
               $(`input[name=${tag_name}]:checked`).each(function () {
                  selected.push($(this).val());
               });
               selected.sort();

               var input_data = {};
               for (let i = 0; i < selected.length; ++i) {
                  let id = selected[i];
                  try {
                     for (let ii = 0; ii < old_item_data.dataObj.items.length; ++ii) {
                        if (old_item_data.dataObj.items[ii].id == String(id)) {
                           var row_input_data = {};
                           for (let col = 0; col < old_item_data.dataObj.items[ii].input.cnt; ++col) {
                              let tag_id = tag_name + "_" + String(id) + "_input_" + String(col + 1);
                              let inputObj = $(`#${tag_id}`);
                              if (inputObj.length > 0) {
                                 let text_value = inputObj.val();
                                 row_input_data[String(col + 1)] = text_value;
                              }
                           }
                           if (Object.keys(row_input_data).length > 0) input_data[String(id)] = row_input_data;
                        }
                     }
                  } catch (e) {}
               }
               var data = { select_ids: selected };
               if (Object.keys(input_data).length > 0) data.inputs = input_data;
               strData[String(item_idx)] = data;
            }

            if (value.ui_index != undefined) {
               let values = Object.values(value.checkbox);
               if (values.length > 0) {
                  for (let l = 0; l < value.ui_index.length; ++l) {
                     let tag_name = `${values[0]}_${value.ui_index[l] - 1}`;
                     makeChkData(strData, l, tag_name);
                  }
               }
            } else {
               for (const [s_key, s_item_name] of Object.entries(value.checkbox)) {
                  let tag_name = s_item_name + "_" + s_key;
                  makeChkData(strData, s_key, tag_name);
               }
            }
            dataObj["data"] = strData;
         }

         if (value.keyword != undefined) {
            dataObj["data"] = value.keyword;
         }

         if (value.member != undefined) {
            dataObj["data"] = value.member;
         }

         if (value.file != undefined) {
            let delFiles = {};
            if (value.file.del != undefined) {
               for (let i = 0; i < value.file.del.length; ++i) {
                  delFiles[value.file.del[i]] = false;
               }
            }
            dataObj["data"] = delFiles;
         }

         if (key == "member_tag") {
         } else {
            if (ret[key] == undefined) ret[key] = dataObj;
         }
      }

      return ret;
   }

   setFileObj(formData) {
      for (const [key, value] of Object.entries(this.managedItems)) {
         if (value.file != undefined) {
            var tag_id = value.file["0"];
            var new_tag_name = tag_id + "_new";

            var fileObjs = $(`input[name=${new_tag_name}]`);
            for (let i = 0; i < fileObjs.length; ++i) {
               if (fileObjs[i].files[0] != undefined) {
                  formData.append(tag_id, fileObjs[i].files[0]);
               }
            }
         }
      }

      for (const [key, value] of Object.entries(this.moreSaveFiles)) {
         formData.append(key, value);
      }
   }

   resetFileItems(data) {
      var rcvItemParser = new ItemParser(this.app_seq);
      rcvItemParser.dataObj = data.files;

      for (const [key, value] of Object.entries(this.managedItems)) {
         if (value.file != undefined) {
            var oldItemData = this.getItemData(key);
            var newItemData = rcvItemParser.getItemData(key);
            if (newItemData) oldItemData.dataObj = newItemData.dataObj;

            var max_cnt = value.file.info.max_cnt;
            var prefix_text = value.file.info.prefix_text;
            var label_class = value.file.info.labelclass;
            oldItemData.applyMultiFileValue(max_cnt, key, prefix_text, label_class);
            value["del"] = [];
         }
      }
   }
}
