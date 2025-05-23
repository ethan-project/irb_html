const ipsap_common_func_js = true;

function isObjVisible(obj) {
  if (obj.length == 0) return false;

  try {
    if (obj.css("display") == "none") return false;
  } catch (e) {
    return true;
  }

  var parentObj = obj.parent();
  if (parentObj.length == 0) return true;

  return isObjVisible(parentObj);
}

function getCommaStrFromNumArr(arr) {
  var ret = ``;
  for (let i = 0; i < arr.length; ++i) {
    if (ret != "") ret += `,`;
    ret += `"` + arr[i] + `"`;
  }
  return ret;
}

function isInArray(arr, data) {
  try {
    for (let i = 0; i < arr.length; i++) {
      if (arr[i] == data) return true;
    }
  } catch (e) {}
  return false;
}

function removeInArray(arr, data) {
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] == data) {
      arr.splice(i, 1);
    }
    return true;
  }
  return false;
}

function getNumberFromBool(flag) {
  if (flag) return 1;

  return 0;
}

function getValidItemIdx(item_idx) {
  if (item_idx == undefined) item_idx = 0;

  return String(item_idx);
}

function getHtmlReplaceId(item_name, item_id, col_idx) {
  let tag_full_name = item_name;

  if (item_id != undefined) tag_full_name += "_" + getValidItemIdx(item_id);

  if (col_idx != undefined) tag_full_name += "_" + String(col_idx);

  return tag_full_name;
}

function getHtmlTagId(item_name, item_idx, code_id, col_idx) {
  let tag_full_name = item_name + "_" + getValidItemIdx(item_idx);

  if (code_id != undefined) tag_full_name += "_" + String(code_id);

  if (col_idx != undefined) tag_full_name += "_" + String(col_idx);

  return tag_full_name;
}

String.prototype.replaceAll = function (org, dest) {
  return this.split(org).join(dest);
};

function spaceAddStr(char) {
  if (char == "") return char;

  char = char.replace(/ +/gi, "");
  var len = char.length;
  var resultStr = "";
  for (var i = 0; i < len; i++) {
    var str = char.substr(i, 1);
    resultStr += str + " ";
  }
  return resultStr;
}

function stringShortDot(str, len) {
  if (str.length <= len) return str;

  return str.substring(0, len - 3) + "...";
}

$.urlParam = function (name) {
  var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
  if (results == null) {
    return null;
  } else {
    return results[1] || 0;
  }
};