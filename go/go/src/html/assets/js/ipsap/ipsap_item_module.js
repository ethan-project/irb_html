const ipsap_item_module_js = true;

if (typeof ipsap_item_parser_js == 'undefined') document.write("<script src='/assets/js/ipsap/ipsap_item_parser.js'></script>");

var g_AppItemParser = undefined;

function appItemSaveTemporary(callbackFunc, userData) {
  if (g_AppItemParser == undefined) return;

  g_AppItemParser.save(false, callbackFunc, userData);
}

function appItemSubmitWithParser(itemParser, callbackFunc, userData) {
  if (itemParser == undefined) return;

  itemParser.save(true, callbackFunc, userData);
}

function appItemSubmit(callbackFunc, userData) {
  if (g_AppItemParser == undefined) return;

  g_AppItemParser.save(true, callbackFunc, userData);
}

function commonGuideMapping(guideObject) {
  try {
    for (const [key, value] of Object.entries(guideObject)) {
      let guide_id = key + "_guide";
      let html = `<ul id="${guide_id}">` + value + `</ul>`;
      $('#' + guide_id).replaceWith(html);
    }
  } catch (error) {}
}

function makeComboList(combo_name, first_option, comboDataObj, key_id, option) {
  const data = comboDataObj[key_id];

  html = `<select class="form-control btn-outline-primary small_select" id="${combo_name}">`;
  html += first_option;
  if (option != undefined) html += `<optgroup label="—">`;
  $.each(data, function (id, value) {
    html += `<option value="${id}">${value}</option>`;
  });
  if (option != undefined) {
    html += `</optgroup>
            <optgroup label="—">
              <option value="new" data-custom-toggle='visible'>${option}</option>
            </optgroup>`;
  }
  html += `</select>`;
  return html;
}

function makeComboList2(combo_name, comboDataObj, key_id, addClass) {
  const data = comboDataObj[key_id];

  html = `<select class="form-control btn-outline-primary small_select ${addClass}" id="${combo_name}">`;
  $.each(data, function (id, value) {
    html += `<option value="${id}">${value}</option>`;
  });
  html += `</select>`;
  return html;
}

function getValueFromCode(codeList, code) {
  for (let i = 0; i < codeList.length; ++i) {
    if (codeList[i].id == Number(code)) {
      return codeList[i].value;
    }
  }
  return "";
}