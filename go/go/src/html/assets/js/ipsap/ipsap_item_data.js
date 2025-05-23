const ipsap_item_data_js = true;

if (typeof ipsap_common_js == 'undefined') document.write("<script src='/assets/js/ipsap/ipsap_common.js'></script>");

class ItemData {
  constructor(item_name, dataObj, itemParser, no_save) {
    this.itemName = item_name;
    this.dataObj = dataObj;
    this.itemParser = itemParser;

    if (no_save == undefined) {
      if (itemParser.managedItems[this.itemName] == undefined) itemParser.managedItems[this.itemName] = {};
    }
  }

  getDataCount() {
    if (this.dataObj == undefined) return 0;

    var data = this.dataObj.saved_data.data;
    if (data == undefined) return 0;

    return Object.keys(data).length;
  }

  getMainChecked(item_idx) {
    if (this.dataObj == undefined) return false;

    if (item_idx == undefined) item_idx = 0;

    try {
      const map = new Map(Object.entries(this.dataObj.saved_data.main_select));
      const value = map.get(String(item_idx));
      if (value != undefined) return value > 0;
    } catch (error) {}

    return false;
  }

  getStringValue(tag_name) {
    try {
      if (this.dataObj.info.item_type_str != "basic" && this.dataObj.info.item_type_str != "string") return "";
      const map = new Map(Object.entries(this.dataObj.saved_data.data));
      var value = map.get(tag_name);
      if (value != undefined) return value;
    } catch (error) {}
    return "";
  }

  getSelectValue(item_idx) {
    try {
      if (item_idx == undefined) item_idx = 0;
      if (this.dataObj.info.item_type_str != "select") return;
      const map = new Map(Object.entries(this.dataObj.saved_data.data));
      const idx_value = map.get(String(item_idx));
      if (idx_value == undefined) return new Array();

      return idx_value.select_ids;
    } catch (error) {}
    return new Array();
  }

  getSelectInputValue(item_idx, row_id, col_no) {
    try {
      if (this.dataObj.info.item_type_str != "select") return "";
      const map = new Map(Object.entries(this.dataObj.saved_data.data));
      const idx_value = map.get(String(item_idx));
      if (idx_value == undefined) return "";

      const map_input = new Map(Object.entries(idx_value.inputs));
      const map_row = new Map(Object.entries(map_input.get(String(row_id))));
      return map_row.get(String(col_no));
    } catch (error) {}
    return "";
  }

  applyMainCheckValueInCheckBox(checkbox_id, item_idx) {
    this.itemParser.managedItems[this.itemName]["main_check_id"] = checkbox_id;
    var checked = this.getMainChecked(item_idx);
    var real_checkbox_id = checkbox_id;

    var obj = $(`input:checkbox[id='${checkbox_id}']`);
    if (obj.length == 0) {
      if (item_idx != undefined) real_checkbox_id += "_" + String(item_idx);
      obj = $(`input:checkbox[id='${real_checkbox_id}']`);
    }

    if (obj.length > 0) {
      obj.prop("checked", checked);
    }


    return checked;
  }

  applyMainCheckValueInCheckBoxReadOnly(checkbox_id, item_idx) {

    var checked = this.getMainChecked(item_idx);
    var real_checkbox_id = checkbox_id;

    var obj = $(`input:checkbox[id='${checkbox_id}']`);
    if (obj.length == 0) {
      if (item_idx != undefined) real_checkbox_id += "_" + String(item_idx);
      obj = $(`input:checkbox[id='${real_checkbox_id}']`);
    }

    if (obj.length > 0) {
      obj.prop("checked", checked);

      if (checked) {
        obj.prop("readonly", true);
        obj.siblings('label')[0].htmlFor = '';
      } else if (!checked) {
        obj.prop("disabled", true);
      }
    }

    return checked;
  }

  applyTextValue(item_idx) {
    var text_id = this.itemName;

    if (item_idx == undefined) item_idx = 0;

    return this.applyTextMapValue(text_id, item_idx);
  }

  applyTextValueReadOnly(item_idx) {
    var text_id = this.itemName;

    if (item_idx == undefined) item_idx = 0;

    return this.applyTextMapValueReadOnly(text_id, item_idx);
  }

  applyTextMapValue(text_id, item_idx) {
    if (text_id == undefined || item_idx == undefined) {
      alert("text_id 또는 item_idx 값이 필요합니다.");
      return;
    }

    if (this.itemParser.managedItems[this.itemName]["string"] == undefined) this.itemParser.managedItems[this.itemName]["string"] = {};
    this.itemParser.managedItems[this.itemName]["string"][String(item_idx)] = text_id;

    var value = this.getStringValue(String(item_idx));

    try {
      const input_format = this.dataObj.items[String(item_idx)];
      const max_len = input_format.max_len;
      const placeholder = input_format.placeholder;
      const number_only = input_format.number_only;

      if (number_only > 0 && value == "") value = "0";

      if (max_len > 0) $(`#${text_id}`).prop("maxlength", max_len);
      if (placeholder != "") $(`#${text_id}`).prop("placeholder", placeholder);
    } catch (e) {}

    var targetObj = $(`#${text_id}`);
    if (targetObj.length == 0) targetObj = $(`#${text_id}_${String(item_idx + 1)}`);

    if (targetObj.length > 0) targetObj.val(value);

    return value;
  }

  applyTextMapValueReadOnly(text_id, item_idx) {
    if (text_id == undefined || item_idx == undefined) {
      alert("text_id 또는 item_idx 값이 필요합니다.");
      return;
    }

    var value = this.getStringValue(String(item_idx));

    var targetObj = $(`#${text_id}`);
    if (targetObj.length == 0) targetObj = $(`#${text_id}_${String(item_idx + 1)}`);

    if (targetObj.length > 0) targetObj.text(value);

    return value;
  }

  applyCalanderValue(text_id, item_idx) {
    if (item_idx == undefined) item_idx = 0;

    if (this.itemParser.managedItems[this.itemName]["string"] == undefined) this.itemParser.managedItems[this.itemName]["string"] = {};
    this.itemParser.managedItems[this.itemName]["string"][String(item_idx)] = text_id;

    let value = this.getStringValue(String(item_idx));

    $(`#${text_id}`).val(value);
    return value;
  }

  applyKeywordValue() {
    let item_name = this.itemName;
    let replace_id = getHtmlReplaceId(item_name);

    if (this.itemParser.managedItems[this.itemName]["keyword"] == undefined) this.itemParser.managedItems[this.itemName]["keyword"] = [];

    var keywordArr = this.dataObj.saved_data.data;
    for (var i = 0; i < keywordArr.length; ++i) {
      let keyword = keywordArr[i];
      this.itemParser.managedItems[this.itemName]["keyword"].push(keyword);
    }
    this.itemParser.managedItems[this.itemName]["keyword"].sort();
    this.refreshKeywordUI();
  }

  applyKeywordValueReadOnly() {
    let item_name = this.itemName;
    let replace_id = getHtmlReplaceId(item_name);

    var keywordArr = this.dataObj.saved_data.data;
    var html = ``;
    for (var i = 0; i < keywordArr.length; ++i) {
      let keyword = keywordArr[i];
      html += `<button type="button" class="btn btn-outline-primary btn-round btn_tag" item_name="${this.itemName}" keyword="${keyword}">
        ${keyword}
      </button>`;
    }
    $('#' + replace_id).html(html);
  }

  refreshKeywordUI() {
    if (this.itemParser.managedItems[this.itemName]["keyword"] == undefined) return;

    let item_name = this.itemName;
    let replace_id = getHtmlReplaceId(item_name);

    var keywordArr = this.itemParser.managedItems[this.itemName]["keyword"];
    var html = ``;
    for (var i = 0; i < keywordArr.length; ++i) {
      let keyword = keywordArr[i];
      html += `<button type="button" class="btn btn-outline-primary btn-round btn_tag" item_name="${this.itemName}" keyword="${keyword}">
        ${keyword}
        <i class="fas fa-times mLeft10 keyword-delete" title="검색어 삭제"></i>
      </button>`;
    }
    $('#' + replace_id).html(html);
  }

  initMembers() {
    var item_type_str = this.dataObj.info.item_type_str;
    if (item_type_str != "member") return;

    item_type_str += "_tag";

    if (this.itemParser.managedItems[item_type_str] == undefined) this.itemParser.managedItems[item_type_str] = [];
    this.itemParser.managedItems[item_type_str].push(this.itemName);

    if (this.itemParser.managedItems[this.itemName]["member"] == undefined) this.itemParser.managedItems[this.itemName]["member"] = [];

    let users = this.dataObj.saved_data.data;
    for (let i = 0; i < users.length; ++i) {
      var newData = {
        "animal_mng_flag": users[i].animal_mng_flag,
        "exp_year_code": users[i].exp_year_code,
        "user_seq": users[i].user_seq,
        "info": users[i].info
      };
      this.itemParser.managedItems[this.itemName]["member"].push(newData);
    }
  }

  initMembersIBC() {
    var item_type_str = this.dataObj.info.item_type_str;
    if (item_type_str != "member") return;

    item_type_str += "_tag";

    if (this.itemParser.managedItems[item_type_str] == undefined) this.itemParser.managedItems[item_type_str] = [];
    this.itemParser.managedItems[item_type_str].push(this.itemName);

    if (this.itemParser.managedItems[this.itemName]["member"] == undefined) this.itemParser.managedItems[this.itemName]["member"] = [];

    let users = this.dataObj.saved_data.data;
    for (let i = 0; i < users.length; ++i) {
      var newData = {
        "animal_mng_flag": users[i].animal_mng_flag,
        "edu_course": users[i].edu_course,
        "exp_year_code": 0,
        "user_seq": users[i].user_seq,
        "info": users[i].info
      };
      this.itemParser.managedItems[this.itemName]["member"].push(newData);
    }
  }

  initMembersIRB() {
    var item_type_str = this.dataObj.info.item_type_str;
    if (item_type_str != "member") return;

    item_type_str += "_tag";

    if (this.itemParser.managedItems[item_type_str] == undefined) this.itemParser.managedItems[item_type_str] = [];
    this.itemParser.managedItems[item_type_str].push(this.itemName);

    if (this.itemParser.managedItems[this.itemName]["member"] == undefined) this.itemParser.managedItems[this.itemName]["member"] = [];

    let users = this.dataObj.saved_data.data;
    for (let i = 0; i < users.length; ++i) {
      var newData = {
        "exp_type_code": users[i].exp_type_code,
        "exp_year_code": users[i].exp_year_code,
        "user_seq": users[i].user_seq,
        "info": users[i].info
      };
      this.itemParser.managedItems[this.itemName]["member"].push(newData);
    }
  }

  changeAttrForMember(user_seq, tag_name, val) {
    var item_type_str = this.dataObj.info.item_type_str;
    if (item_type_str != "member") return;
    item_type_str += "_tag";

    if (this.itemParser.managedItems[item_type_str] == undefined) return;

    if (this.itemParser.managedItems[this.itemName]["member"] == undefined) return;

    let users = this.itemParser.managedItems[this.itemName]["member"];
    if ('animal_mng_flag' == tag_name) {
      for (let i = 0; i < users.length; ++i) {
        users[i][tag_name] = 0;
      }
    }

    for (let i = 0; i < users.length; ++i) {
      if (users[i].user_seq == Number(user_seq)) {
        users[i][tag_name] = Number(val);
        break;
      }
    }
  }

  applyMultiFileValue(max_cnt, tag_id, prefix_text, labelclass) {
    var item_this = this;

    if (labelclass == undefined) {
      labelclass = "mRight10 mBot0 w80";
    }

    if (this.itemParser.managedItems[this.itemName]["file"] == undefined) this.itemParser.managedItems[this.itemName]["file"] = {};
    this.itemParser.managedItems[this.itemName]["file"]["0"] = tag_id;
    this.itemParser.managedItems[this.itemName]["file"]["info"] = { max_cnt: max_cnt, prefix_text: prefix_text, labelclass: labelclass };

    var staticFileNo = 0;
    var fileCnt = 0;

    var addBtnId = this.itemName + "_add_btn";
    var removeBtnId = this.itemName + "_remove_btn";

    var addBtnHtml = ``;
    if (max_cnt > 1) addBtnHtml = `<a class="${addBtnId} btn btn-outline-primary mLeft10 btn-file-add" title="첨부 File 추가"><i class="fas fa-plus"></i></a>`;

    makeFileListUI(true);

    function makeFileListUI(makeEmpty) {
      fileCnt = 0;
      staticFileNo = 0;
      var old_tag_id = tag_id + "_old";
      var fileArr = item_this.dataObj.saved_data.data;
      var html = ``;
      var label_text = ``;
      for (var i = 0; i < fileArr.length; ++i) {
        var fileInfo = fileArr[i];

        if (isInArray(item_this.itemParser.managedItems[item_this.itemName]["file"]["del"], fileInfo.filepath)) continue;

        ++fileCnt;
        ++staticFileNo;
        var file_id = old_tag_id + "_" + staticFileNo;

        if (prefix_text != "") {
          label_text = prefix_text + " :";
        }

        html += `
        <div class="flexMid flex_1 mTop5" item_name="${item_this.itemName}" file_path="${fileInfo.filepath}">
          <label for="input_01-6_1" class="${labelclass}">${label_text}</label>
          <div class="custom-file flex_1 uploaded_file_box">
            <a href="${fileInfo.src}" class="data">${fileInfo.org_file_name}</a>
          </div>
          ${addBtnHtml}
          <a class="${removeBtnId} btn btn-outline-danger mLeft5 btn-file-delete" title="첨부 File 삭제"><i class="far fa-trash-alt"></i></a>
        </div>`;
      }
      $('#' + tag_id).html(html);

      if (fileCnt < max_cnt && (makeEmpty || fileCnt == 0)) addEmptyFileInput();
    }

    function addEmptyFileInput() {
      ++fileCnt;
      ++staticFileNo;
      var new_tag_name = tag_id + "_new";
      var file_id = new_tag_name + "_" + staticFileNo;
      var label_text = ``;

      if (prefix_text != "") {
        label_text = prefix_text + " :";
      }

      var html = `
        <div class="flexMid flex_1 mTop5">
          <label for="${file_id}" class="${labelclass}">${label_text}</label>
          <div class="custom-file flex_1">
            <input type="file" name="${new_tag_name}" class="custom-file-input" id="${file_id}">
            <label class="custom-file-label" for="${file_id}" data-button-text="첨부 File 선택">File을 선택해 주세요.</label>
          </div>
          ${addBtnHtml}
          <a class="${removeBtnId} btn btn-outline-danger mLeft5 btn-file-delete" title="첨부 File 삭제"><i class="far fa-trash-alt"></i></a>
        </div>`;

      $('#' + tag_id).append(html);
    }

    $(document).off('click', '.' + addBtnId);
    $(document).on('click', '.' + addBtnId, function () {
      if (fileCnt >= max_cnt || fileCnt < 0) return false;
      addEmptyFileInput();
    });

    $(document).off('click', '.' + removeBtnId);
    $(document).on('click', '.' + removeBtnId, function () {
      let target_div = $(this).parent();
      let item_name = target_div.attr("item_name");
      let file_path = target_div.attr("file_path");

      if (item_name != undefined) {
        if (item_this.itemParser.managedItems[item_name]["file"]["del"] == undefined) item_this.itemParser.managedItems[item_name]["file"]["del"] = [];
        item_this.itemParser.managedItems[item_name]["file"]["del"].push(file_path);
      }

      target_div.remove();
      --fileCnt;
      if (fileCnt == 0) addEmptyFileInput();
    });
  }

  applyMultiFileValueReadOnly(tag_id, prefix_text, labelclass) {
    var item_this = this;

    if (labelclass == undefined) {
      labelclass = "mRight10 mBot0";
    }

    var label_text = ``;

    if (prefix_text != "") {
      label_text = prefix_text + " :";
    }

    var staticFileNo = 0;
    var fileCnt = 0;

    makeFileListUI(true);

    function makeFileListUI(makeEmpty) {
      fileCnt = 0;
      staticFileNo = 0;
      var old_tag_id = tag_id + "_old";
      var fileArr = item_this.dataObj.saved_data.data;
      var html = `<div class="col-sm-9">`;

      if (fileArr.length == 0) {
        html += `
        <div class="form-group row flexMid flexNoWrap">
          <label class="${labelclass}">${label_text}</label>
            <span class="data">첨부파일 없음</span>
        </div>`;
      }

      for (var i = 0; i < fileArr.length; ++i) {
        var fileInfo = fileArr[i];

        ++fileCnt;
        ++staticFileNo;
        var file_id = old_tag_id + "_" + staticFileNo;

        html += `
        <div class="form-group row flexMid flexNoWrap" item_name="${item_this.itemName}" file_path="${fileInfo.filepath}">
          <label class="${labelclass}">${label_text}</label>
            <a href="${fileInfo.src}"  class="data">${fileInfo.org_file_name}</a>
        </div>`;
      }
      html += `</div>`;
      $('#' + tag_id).html(html);
    }
  }

  makeHtmlRadioType(target, item_idx) {
    if (target == undefined) target = IPSAP.COL.ALL;

    let item_name = this.itemName;
    let replace_id = getHtmlReplaceId(item_name);
    if (this.itemParser.managedItems[this.itemName]["radio"] == undefined) this.itemParser.managedItems[this.itemName]["radio"] = {};
    this.itemParser.managedItems[this.itemName]["radio"][getValidItemIdx(item_idx)] = replace_id;

    var i = 0;
    let add = 1;
    switch (target) {
      case IPSAP.COL.ODD:
        add = 2;
        replace_id += "_odd";
        break;
      case IPSAP.COL.EVEN:
        i = 1;
        add = 2;
        replace_id += "_even";
        break;
    }

    let select_ids = this.getSelectValue(item_idx);
    let html = ``;

    let showMap = {};
    for (; i < this.dataObj.items.length; i = i + add) {
      const item = this.dataObj.items[i];
      let id = getHtmlTagId(item_name, item_idx, item.id);

      if (select_ids.length == 0) select_ids.push(item.id);

      let name = getHtmlTagId(item_name, item_idx);
      let selected = false;
      let moreAttr = ` data-custom-toggle='visible'`;
      if (item.id == select_ids[0]) {
        selected = true;
        moreAttr += ` checked`;
      }

      let divLeftAttr = ``;
      let labelClass = `class="custom-control-label"`;
      let inputHtml = ``;
      let inputbox_id = ``;

      if (item.input != undefined) {
        const input_cnt = item.input.cnt;
        const input_w_class = item.input.w_class;
        const input_iw_class = item.input.iw_class;
        const input_max_len = item.input.max_len;
        const input_prefix = item.input.prefix;
        const input_suffix = item.input.suffix;
        const input_number = item.input.number_only;
        const input_placeholder = item.input.placeholder;

        if (input_w_class == "") labelClass = ` class="custom-control-label w50"`;else labelClass = ` class="custom-control-label ${input_w_class}"`;

        let input_value = [];
        try {
          const inputs = this.dataObj.saved_data.data[getValidItemIdx(item_idx)].inputs[String(item.id)];
          for (let col = 0; col < input_cnt; ++col) {
            input_value.push(inputs[String(col + 1)]);
          }
        } catch (e) {}
        if (input_value.length < input_cnt) {
          for (let col = input_value.length; col < input_cnt; ++col) input_value.push(``);
        }

        for (let ino = 0; ino < input_cnt; ++ino) {
          let col_no = ino + 1;
          divLeftAttr = ` flexMid pLeft3`;
          inputbox_id = `${id}_input_${col_no}`;

          let input2Addclass = ``;
          if (input_cnt > 1) input2Addclass = `w50`;

          if (input_prefix[ino] != ``) inputHtml += `<span class="mLeft20 mRight5 ${input2Addclass} hidden" id="${inputbox_id}2" for="${inputbox_id}">${input_prefix[ino]}</span>`;

          let moreclass = ``;
          if (input_cnt == 1) moreclass = `flex_1`;

          let input_type = "text";
          if (input_number[ino]) input_type = "number";
          inputHtml += `<input type="${input_type}" min="0" class="form-control form-control-sm ${moreclass} ${input_iw_class[ino]}  hidden"
                          id="${inputbox_id}" value="${input_value[ino]}"
                          maxlength="${input_max_len[ino]}" placeholder="${input_placeholder[ino]}">`;

          if (input_suffix[ino] != ``) inputHtml += `<span class="mLeft5 hidden" id="${inputbox_id}3">${input_suffix[ino]}</span>`;

          if (selected && inputbox_id != ``) {
            if (showMap[id] == undefined) showMap[id] = [];
            showMap[id].push(inputbox_id);
          }
        }
      }

      let single_line = ``;
      if (target == IPSAP.COL.HORI) {
        single_line = `
          <div class="form-check-inline">
            <div class="custom-control custom-radio">
              <input type="radio" id="${id}" name="${name}" value="${item.id}" class="custom-control-input" ${moreAttr}>
              <label ${labelClass} for="${id}">${item.value}</label>
              ${inputHtml}
            </div>
          </div>`;
      } else {
        single_line = `
          <div class="custom-control custom-radio flexMid mTop10">
            <input type="radio" id="${id}" name="${name}" value="${item.id}"
            class="custom-control-input" ${moreAttr}>
            <label ${labelClass} for="${id}">${item.value}</label>
            ${inputHtml}
          </div>`;
      }

      html += single_line;
    };

    var targetObj = $('#' + replace_id);
    if (targetObj.length == 0) {
      targetObj = $(`#${replace_id}_${item_idx + 1}`);
    }

    if (targetObj.length > 0) {
      targetObj.html(html);
      targetObj.children('div').first().removeClass('mTop10');
      $('input[type=radio]').on({
        click: function (e) {
          let $this_radio = $(this),
              radio_name = $this_radio.attr('name'),
              $radios = $('input[name=' + radio_name + ']');

          if ($this_radio.data('custom-toggle') == 'visible') {
            $this_radio.parent().find('.hidden').show().filter('input:first').focus();
            $radios.not($this_radio).each(function () {
              $(this).parent().find('.hidden').hide();
            });
          }
        }
      });

      for (const [key, value] of Object.entries(showMap)) {
        for (let i = 0; i < value.length; ++i) {
          $('#' + value[i]).css('display', 'block');
          $('#' + value[i] + "2").css('display', 'block');
          $('#' + value[i] + "3").css('display', 'block');
        }
      }
    }
  }

  makeHtmlRadioTypeReadOnly(target, item_idx) {
    if (target == undefined) target = IPSAP.COL.ALL;

    let item_name = this.itemName;
    let replace_id = getHtmlReplaceId(item_name);

    var i = 0;
    let add = 1;
    switch (target) {
      case IPSAP.COL.ODD:
        add = 2;
        replace_id += "_odd";
        break;
      case IPSAP.COL.EVEN:
        i = 1;
        add = 2;
        replace_id += "_even";
        break;
    }

    let select_ids = this.getSelectValue(item_idx);
    let html = ``;

    let showMap = {};
    for (; i < this.dataObj.items.length; i = i + add) {
      const item = this.dataObj.items[i];
      let id = getHtmlTagId(item_name, item_idx, item.id);

      if (select_ids.length == 0) select_ids.push(item.id);

      let name = getHtmlTagId(item_name, item_idx);
      let selected = false;
      let moreAttr = ` data-custom-toggle='visible'`;
      if (item.id == select_ids[0]) {
        selected = true;
        moreAttr += ` checked readonly`;
      } else moreAttr += ' disabled';

      let divLeftAttr = ``;
      let labelClass = `class="custom-control-label"`;
      let inputHtml = ``;
      let inputbox_id = ``;

      if (item.input != undefined) {
        const input_cnt = item.input.cnt;
        const input_w_class = item.input.w_class;
        const input_iw_class = item.input.iw_class;
        const input_prefix = item.input.prefix;
        const input_suffix = item.input.suffix;
        const input_number = item.input.number_only;

        if (input_w_class == "") labelClass = ` class="custom-control-label w50"`;else labelClass = ` class="custom-control-label ${input_w_class}"`;

        let input_value = [];
        try {
          const inputs = this.dataObj.saved_data.data[getValidItemIdx(item_idx)].inputs[String(item.id)];
          for (let col = 0; col < input_cnt; ++col) {
            input_value.push(inputs[String(col + 1)]);
          }
        } catch (e) {}
        if (input_value.length < input_cnt) {
          for (let col = input_value.length; col < input_cnt; ++col) input_value.push(``);
        }

        for (let ino = 0; ino < input_cnt; ++ino) {
          let col_no = ino + 1;
          divLeftAttr = ` flexMid pLeft3`;
          inputbox_id = `${id}_input_${col_no}`;

          let input2Addclass = ``;
          if (input_cnt > 1) input2Addclass = `w50`;

          if (input_prefix[ino] != ``) inputHtml += `<span class="mLeft20 mRight5 ${input2Addclass}" id="${inputbox_id}2" for="${inputbox_id}">${input_prefix[ino]}</span>`;

          let moreclass = ``;
          if (input_cnt == 1) moreclass = `flex_1`;

          let input_type = "text";
          if (input_number[ino]) input_type = "number";

          inputHtml += `<span class="data ${input_iw_class[ino]}">${input_value[ino]}</span>`;

          if (input_suffix[ino] != ``) inputHtml += `<span class="mLeft5" id="${inputbox_id}3">${input_suffix[ino]}</span>`;

          if (selected && inputbox_id != ``) {
            if (showMap[id] == undefined) showMap[id] = [];
            showMap[id].push(inputbox_id);
          }
        }
      }

      if (!selected) {
        inputHtml = ``;
      }

      let single_line = ``;
      if (target == IPSAP.COL.HORI) {
        single_line = `
          <div class="form-check-inline">
            <div class="custom-control custom-radio">
              <input type="radio" id="${id}" name="${name}" value="${item.id}" class="custom-control-input" ${moreAttr}>
              <label ${labelClass}>${item.value}</label>
              ${inputHtml}
            </div>
          </div>`;
      } else {
        single_line = `
          <div class="custom-control custom-radio flexMid mTop10">
            <input type="radio" id="${id}" name="${name}" value="${item.id}"
            class="custom-control-input" ${moreAttr}>
            <label ${labelClass}>${item.value}</label>
            ${inputHtml}
          </div>`;
      }

      html += single_line;
    };

    var targetObj = $('#' + replace_id);
    if (targetObj.length == 0) {
      targetObj = $(`#${replace_id}_${item_idx + 1}`);
    }

    if (targetObj.length > 0) {
      targetObj.html(html);
      targetObj.children('div').first().removeClass('mTop10');
    }
  }

  makeHtmlCheckList(target, item_idx) {
    if (target == undefined) target = IPSAP.COL.ALL;

    let item_name = this.itemName;
    let replace_id = getHtmlReplaceId(this.itemName);
    if (this.itemParser.managedItems[this.itemName]["checkbox"] == undefined) this.itemParser.managedItems[this.itemName]["checkbox"] = {};
    this.itemParser.managedItems[this.itemName]["checkbox"][getValidItemIdx(item_idx)] = replace_id;

    var i = 0;
    let add = 1;
    switch (target) {
      case IPSAP.COL.ODD:
        add = 2;
        replace_id += "_odd";
        break;
      case IPSAP.COL.EVEN:
        i = 1;
        add = 2;
        replace_id += "_even";
        break;
    }

    let select_ids = this.getSelectValue(item_idx);
    let html = ``;

    let showMap = {};
    for (; i < this.dataObj.items.length; i = i + add) {
      const item = this.dataObj.items[i];
      let id = `${item_name}_${getValidItemIdx(item_idx)}_${item.id}`;

      let name = getHtmlTagId(item_name, item_idx);
      let selected = false;
      let moreAttr = ``;
      if (isInArray(select_ids, item.id)) {
        selected = true;
        moreAttr += ` checked`;
      }

      let labelClass = ``;
      let divLeftAttr = ``;
      let inputHtml = ``;
      let inputbox_id = ``;

      if (item.input != undefined) {
        const input_cnt = item.input.cnt;
        const input_w_class = item.input.w_class;
        const input_iw_class = item.input.iw_class;
        const input_max_len = item.input.max_len;
        const input_prefix = item.input.prefix;
        const input_suffix = item.input.suffix;
        const input_number = item.input.number_only;
        const input_placeholder = item.input.placeholder;

        if (input_w_class == "") labelClass = ` class="w50"`;else labelClass = ` class="${input_w_class}"`;

        let input_value = [];
        try {
          const inputs = this.dataObj.saved_data.data[getValidItemIdx(item_idx)].inputs[String(item.id)];
          for (let col = 0; col < input_cnt; ++col) {
            input_value.push(inputs[String(col + 1)]);
          }
        } catch (e) {}
        if (input_value.length < input_cnt) {
          for (let col = input_value.length; col < input_cnt; ++col) input_value.push(``);
        }

        for (let ino = 0; ino < input_cnt; ++ino) {
          let col_no = ino + 1;
          divLeftAttr = ` flexMid pLeft3`;
          moreAttr += ` data-custom-toggle='visible'`;
          inputbox_id = `${id}_input_${col_no}`;

          if (input_prefix[ino] != ``) inputHtml += `<span class="mLeft20 mRight5 hidden" id="${inputbox_id}2" for="${inputbox_id}">${input_prefix[ino]}</span>`;

          let input_type = "text";
          if (input_number[ino]) input_type = "number";
          inputHtml += `<input type="${input_type}" min="0" class="form-control form-control-sm ${input_iw_class[ino]} hidden"
                          id="${inputbox_id}" value="${input_value[ino]}"
                          maxlength="${input_max_len[ino]}" placeholder="${input_placeholder[ino]}">`;

          if (input_suffix[ino] != ``) inputHtml += `<span class="mLeft5 hidden" id="${inputbox_id}3">${input_suffix[ino]}</span>`;

          if (selected && inputbox_id != ``) {
            if (showMap[id] == undefined) showMap[id] = [];
            showMap[id].push(inputbox_id);
          }
        }
      }

      html += `
          <div class="checkbox checkbox-primary${divLeftAttr}">
            <input type="checkbox" name="${name}" id="${id}" value="${item.id}"${moreAttr}>
            <label ${labelClass} for="${id}">${item.value}</label>
            ${inputHtml}
          </div>`;
    }

    var targetObj = $('#' + replace_id);
    if (targetObj.length == 0) {
      targetObj = $(`#${replace_id}_${item_idx + 1}`);
    }
    if (targetObj.length > 0) {
      if (targetObj.children().hasClass('content_title')) {
        let h6 = targetObj.find(".content_title");
        h6.siblings().empty();
        h6.after(html);
      } else {
        targetObj.html(html);
      }

      setTriggerCheckBoxOnChange();

      for (const [key, value] of Object.entries(showMap)) {
        for (let i = 0; i < value.length; ++i) {
          $('#' + value[i]).css('display', 'block');
          $('#' + value[i] + "2").css('display', 'block');
          $('#' + value[i] + "3").css('display', 'block');
        }
      }
    }
  }

  makeHtmlCheckListReadOnly(target, item_idx) {
    if (target == undefined) target = IPSAP.COL.ALL;

    let item_name = this.itemName;
    let replace_id = getHtmlReplaceId(this.itemName);


    var i = 0;
    let add = 1;
    switch (target) {
      case IPSAP.COL.ODD:
        add = 2;
        replace_id += "_odd";
        break;
      case IPSAP.COL.EVEN:
        i = 1;
        add = 2;
        replace_id += "_even";
        break;
    }

    let select_ids = this.getSelectValue(item_idx);
    let html = ``;

    let showMap = {};
    for (; i < this.dataObj.items.length; i = i + add) {
      const item = this.dataObj.items[i];
      let id = `${item_name}_${getValidItemIdx(item_idx)}_${item.id}`;

      let name = getHtmlTagId(item_name, item_idx);
      let selected = false;
      let moreAttr = ``;
      if (isInArray(select_ids, item.id)) {
        selected = true;
        moreAttr += ` checked readonly`;
      } else moreAttr += ' disabled';

      let labelClass = ``;
      let divLeftAttr = ``;
      let inputHtml = ``;
      let inputbox_id = ``;

      if (item.input != undefined) {
        const input_cnt = item.input.cnt;
        const input_w_class = item.input.w_class;
        const input_iw_class = item.input.iw_class;
        const input_prefix = item.input.prefix;
        const input_suffix = item.input.suffix;
        const input_number = item.input.number_only;

        if (input_w_class == "") labelClass = ` class="w50"`;else labelClass = ` class="${input_w_class}"`;

        let input_value = [];
        try {
          const inputs = this.dataObj.saved_data.data[getValidItemIdx(item_idx)].inputs[String(item.id)];
          for (let col = 0; col < input_cnt; ++col) {
            if ("" != inputs[String(col + 1)]) {
              input_value.push(inputs[String(col + 1)]);
            }
          }
        } catch (e) {}
        if (input_value.length < input_cnt) {
          for (let col = input_value.length; col < input_cnt; ++col) input_value.push(``);
        }

        for (let ino = 0; ino < input_cnt; ++ino) {
          let col_no = ino + 1;
          divLeftAttr = ` flexMid pLeft3`;

          inputbox_id = `${id}_input_${col_no}`;

          if (input_prefix[ino] != ``) inputHtml += `<span class="mLeft20 mRight5" id="${inputbox_id}2" for="${inputbox_id}">${input_prefix[ino]}</span>`;

          let input_type = "text";
          if (input_number[ino]) input_type = "number";

          if ("" != input_value[ino]) {
            inputHtml += `<span class="data ${input_iw_class[ino]}">${input_value[ino]}</span>`;
          }

          if (input_suffix[ino] != ``) inputHtml += `<span class="mLeft5" id="${inputbox_id}3">${input_suffix[ino]}</span>`;

          if (selected && inputbox_id != ``) {
            if (showMap[id] == undefined) showMap[id] = [];
            showMap[id].push(inputbox_id);
          }
        }
      }

      if (!selected) {
        inputHtml = ``;
      }

      html += `
          <div class="checkbox checkbox-primary${divLeftAttr}">
            <input type="checkbox" name="${name}" id="${id}" value="${item.id}"${moreAttr}>
            <label ${labelClass}>${item.value}</label>
            ${inputHtml}
          </div>`;
    }

    var targetObj = $('#' + replace_id);
    if (targetObj.length == 0) {
      targetObj = $(`#${replace_id}_${item_idx + 1}`);
    }
    if (targetObj.length > 0) {
      if (targetObj.children().hasClass('content_title')) {
        let h6 = targetObj.find(".content_title");
        h6.siblings().empty();
        h6.after(html);
      } else {
        targetObj.html(html);
      }
    }
  }
}