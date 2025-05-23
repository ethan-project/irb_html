$(document).ready(function () {
  rowTippy = [];

  var rowGroup_enable = $('#datatable').hasClass('rowGroup'),
      Head = $('#datatable').find('thead'),
      head_TH = Head.find('th'),
      Body = $('#datatable').find('tbody'),
      rowGroup_column_Arry = (Head.data('rowgroup') + '').split(','),
      groupStart_column = Head.data('rowgroup-start'),
      isFolded = !Head.data('rowgroup-folded') ? false : Head.data('rowgroup-folded'),
      node_cnt = Body.find('tr:first-of-type').children().length,
      orderFixed = [$("input[name='app_name_order']:checked").data('order'), $("input[name='app_name_order']:checked").val()],
      rowGroup_order_Arry = [],
      initPage_length = !Head.data('page-length') ? 10 : Head.data('page-length'),
      childCol_index = head_TH.index($('#datatable').find('thead th.child_column')),
      save_state = !Head.data('save-state') ? true : Head.data('save-state'),
      zeroRecords = '검색된 결과가 없습니다.',
      block_unit = !$('#datatable thead').data('block-unit') ? '건' : $('#datatable thead').data('block-unit');


  $.each(rowGroup_column_Arry, function (index, order_col) {
    rowGroup_order_Arry.push([order_col, 'asc']);
  });


  let listOrder = [];
  if (typeof Head.data('init-order') !== 'undefined') {
    let initOrderArr = Head.data('init-order').toString().split(','),
        orderDirectionArr = Head.data('order-direction').split(','),
        correction_num = 0;
    for (let i = 0; i < initOrderArr.length; i++) {
      let tmp = [];
      tmp.push(initOrderArr[i] - correction_num);
      tmp.push(orderDirectionArr[i]);
      listOrder.push(tmp);
    }
  }

  stateSaveParams_data = '';
  dataTableOption = {
    dom: 'lfrtipB',


    paging: true,
    stateSave: save_state,
    stateSaveParams: (settings, data) => {
      data.length = initPage_length;
      stateSaveParams_data = data;
    },
    info: true,
    destroy: true,

    searching: true,
    ordering: true,
    lengthChange: false,
    pageLength: initPage_length,
    iDisplayLength: 10,
    language: {
      zeroRecords: zeroRecords,
      search: '<span class="filter_label"><i class="fas fa-search mRight5"></i>결과 내 검색:</span>',
      paginate: {
        'previous': '이전',
        'next': '다음'
      }
    },
    order: listOrder,
    columnDefs: [{ targets: 'no-order', orderable: false }, { targets: 'hidden', visible: false }, { targets: 'child_column', visible: false }],
    buttons: [{
      extend: 'excelHtml5',
      text: '',
      className: 'btn btn-sm btn-outline-primary noMinWidth',
      header: true,
      exportOptions: {
        columns: 'th:not(.noExport)'
      },
      filename: () => {
        return $('.page-title').text() || '신청서 및 보고서';
      },
      title: () => {
        return $('.card-title').text() || '신청서 및 보고서 목록';
      },

      sheetName: '실험 계획서 목록'
    }],

    orderFixed: {
      pre: orderFixed
    },
    rowGroup: {
      enable: rowGroup_enable,
      emptyDataGroup: null,
      className: 'rowGroup',
      responsive: {
        details: false
      },
      startRender: function (rows, group, level) {
        var rowGroups = $(rows.nodes()),
            result_title = `${rowGroups.eq(0).find('[data-title]').data('title')} <span class="subInfo mLeft5">[ ${rowGroups.eq(0).data('appno')} ]</span>`,
            result_Row = `<td colspan="${node_cnt}"></td>`,
            rows_Counter = rows.count(),
            isFolded = !Head.data('rowgroup-folded') ? false : Head.data('rowgroup-folded'),
            fold_type = isFolded ? 'folded' : 'unfolded',
            group_ID,
            group_name,
            group_query,
            group_Starter,
            new_group_icon,
            group_Starter_cols,
            group_sum_col,
            action_buttons,
            btn_change_disabled,
            btn_bring_disabled,
            btn_closed_disabled,
            exp_period;

        if (group == '-' || !group) {
          group = false;
        }

        if (group) {
          group_name = result_title;
          group_Starter = $('<tr/>').addClass(fold_type).append(result_Row);
          group_Starter_cols = group_Starter.children();
          group_Starter_cols.eq(groupStart_column).html(`<span class="group-folder mRight5 ${fold_type}"></span><i class="mRight5"></i>`);

          rowGroups.each(function (idx, ele) {
            var loginUserSeq = parseInt(JSON.parse(COMM.getCookie("user_info"))["user_seq"]);
            var trUserSeq = $(ele).data('truser');
            if (idx < rowGroups.length - 1) {
              $(ele).removeClass('group-sub-end').addClass('group-sub');
            } else {
              $(ele).removeClass('group-sub').addClass('group-sub-end');
            }

            if (idx == 0) {
              var key = $(ele).data('key');
              if ($(ele).data('status').split(',')[0] == 'inProgress') {
                group_Starter.addClass('inProgress').attr({
                  'data-tippy-placement': 'top-start',
                  'data-tippy-animation': 'shift-away',
                  'data-tippy-arrow': 'true'
                });

                if ($('body').hasClass('researcher')) {
                  group_Starter.attr({ 'title': '실험이 진행중입니다.' });
                  if ($(ele).hasClass('IACUC')) {
                    action_buttons = `
                    <span class="mLeft30">
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 2)" class="btn btn-outline-primary btn-xs">변경 신청</a>
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 4)" class="btn btn-outline-primary btn-xs">반입 보고</a>
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 6)" class="btn btn-outline-primary btn-xs">종료 보고</a>
                    </span>
                    `;
                  } else if ($(ele).hasClass('IBC')) {
                    action_buttons = `<span class="mLeft30">
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIBC(${key})" class="btn btn-outline-primary btn-xs">변경 신청</a>
                                      </span>`;
                  } else if ($(ele).hasClass('IRB')) {
                    action_buttons = ` <span class="mLeft30">
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 2)" class="btn btn-outline-primary btn-xs">변경 신청</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 7)" class="btn btn-outline-primary btn-xs">지속 심의(중간보고)</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 8)" class="btn btn-outline-primary btn-xs">중대한 이상 반응</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 9)" class="btn btn-outline-primary btn-xs">연구계획 위반/이탈</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 10)" class="btn btn-outline-primary btn-xs">예상치 못한 문제 발생</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 6)" class="btn btn-outline-primary btn-xs">종료 보고</a>
                                       </span>`;
                  }

                  if (loginUserSeq != trUserSeq) {
                    action_buttons = ``;
                  }
                } else if ($('body').hasClass('officer')) {
                  if ($(ele).data('status').indexOf('assigned') < 0) {
                    if ($(ele).hasClass('IACUC')) {
                      group_Starter.attr({ 'title': '실험이 진행중입니다. 승인 후 점검위원을 지정해 주세요.' });

                      action_buttons = `
                      <span class="mLeft30">
                        <a href="javascript:void(0);" onClick="modalAssign(${key})" class="btn btn-outline-danger btn-xs" data-toggle="modal" data-target="#modal_assign">점검위원 지정</a>
                      </span>`;
                    }
                  } else {
                    group_Starter.attr({ 'title': '실험이 진행중입니다.' });

                    action_buttons = `
                    <span class="mLeft30">
                      <a href="javascript:void(0);" onClick="modalAssign(${key})" class="btn btn-outline-primary btn-xs" data-toggle="modal" data-target="#modal_assign">점검위원 변경</a>
                    </span>`;
                  }
                } else {
                  group_Starter.attr({ 'title': '실험이 진행중입니다.' });
                  action_buttons = ``;
                }
                new_group_icon = 'fas fa-flask';
                exp_period = '<span class="blue_deep">실험중</span> : ' + $(ele).data('period');
              };
              if ($(ele).data('status').split(',')[0] == 'to_expire') {
                group_Starter.addClass('closed').attr({
                  'data-tippy-placement': 'top-start',
                  'data-tippy-animation': 'shift-away',
                  'data-tippy-arrow': 'true'
                });

                if ($('body').hasClass('researcher')) {
                  group_Starter.attr({ 'title': '실험 유효기간 만료가 도래하여, 재승인이 필요합니다.' });
                  action_buttons = `
                  <span class="mLeft30">
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 3)" class="btn btn-outline-danger btn-xs">재승인 신청</a>
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 2)" class="btn btn-outline-primary btn-xs">변경 신청</a>
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 4)" class="btn btn-outline-primary btn-xs">반입 보고</a>
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 6)" class="btn btn-outline-primary btn-xs">종료 보고</a>
                  </span>
                  `;

                  if (loginUserSeq != trUserSeq) {
                    action_buttons = ``;
                  }
                } else if ($('body').hasClass('officer') && (!$(ele).data('status').split(',')[1] || $.trim($(ele).data('status').split(',')[1]) != 'assigned')) {
                  group_Starter.attr({ 'title': '실험 유효기간 만료가 도래하여, 재승인이 필요합니다.' });
                  action_buttons = `
                  <span class="mLeft30">
                    <a href="./experimenting_list-reviewer.html" class="btn btn-outline-danger btn-xs">점검위원 지정</a>
                  </span>
                  `;
                } else {
                  group_Starter.attr({ 'title': '실험 유효기간 만료가 도래하여, 재승인이 필요합니다.' });
                  action_buttons = ``;
                }
                new_group_icon = 'fas fa-flask';
                exp_period = "<span class='red'>재승인 만료 " + $(ele).data('expire_date') + "</span> : " + $(ele).data('period');
              };
              if ($(ele).data('status').split(',')[0] == 'closed') {
                group_Starter.addClass('closed').attr({
                  'data-tippy-placement': 'top-start',
                  'data-tippy-animation': 'shift-away',
                  'data-tippy-arrow': 'true'
                });

                if ($('body').hasClass('researcher')) {
                  group_Starter.attr({ 'title': '실험 기간이 만료되었습니다. "종료 보고서"를 제출해 주세요.' });
                  action_buttons = `
                  <span class="mLeft30">
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 2)" class="btn btn-outline-primary btn-xs">변경 신청</a>
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 4)" class="btn btn-outline-primary btn-xs">반입 보고</a>
                    <a href="javascript:void(0);" onClick="onClickActionButton(${key}, 6)" class="btn btn-outline-primary btn-xs">종료 보고</a>
                  </span>
                  `;

                  if ($(ele).hasClass('IBC')) {
                    group_Starter.attr({ 'title': '실험 기간이 만료되었습니다.' });
                    action_buttons = `<span class="mLeft30">
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIBC(${key})" class="btn btn-outline-primary btn-xs">변경 신청</a>
                                      </span>`;
                  } else if ($(ele).hasClass('IRB')) {
                    action_buttons = ` <span class="mLeft30">
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 2)" class="btn btn-outline-primary btn-xs">변경 신청</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 7)" class="btn btn-outline-primary btn-xs">지속 심의(중간보고)</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 8)" class="btn btn-outline-primary btn-xs">중대한 이상 반응</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 9)" class="btn btn-outline-primary btn-xs">연구계획 위반/이탈</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 10)" class="btn btn-outline-primary btn-xs">예상치 못한 문제 발생</a>
                                        <a href="javascript:void(0);" onClick="onClickActionButtonIRB(${key}, 6)" class="btn btn-outline-primary btn-xs">종료 보고</a>
                                       </span>`;
                  }

                  if (loginUserSeq != trUserSeq) {
                    action_buttons = ``;
                  }
                } else if ($('body').hasClass('officer')) {
                  group_Starter.attr({ 'title': '실험 기간이 만료되었습니다. "종료 보고서" 제출을 확인해 주세요.' });
                  if ($(ele).hasClass('IBC')) {
                    group_Starter.attr({ 'title': '실험 기간이 만료되었습니다.' });
                  }
                  action_buttons = ``;
                } else {
                  group_Starter.attr({ 'title': '실험 기간이 만료되었습니다. "종료 보고서" 제출을 확인해 주세요.' });
                  if ($(ele).hasClass('IBC')) {
                    group_Starter.attr({ 'title': '실험 기간이 만료되었습니다.' });
                  }
                  action_buttons = ``;
                }
                new_group_icon = 'fas fa-flask';
                let expired_period = $(ele).data('period').split(' ~ ')[0] + ' ~ <span class="red">' + $(ele).data('period').split(' ~ ')[1] + '</span>';
                exp_period = '<span class="red">실험 기간 만료</span> : ' + expired_period;
              };
              if ($(ele).data('status').split(',')[0] == 'completed') {
                group_Starter.addClass('completed').attr({
                  'data-tippy-placement': 'top-start',
                  'data-tippy-animation': 'shift-away',
                  'data-tippy-arrow': 'true',
                  'title': '모든 과정이 종료된 실험입니다.'
                });

                action_buttons = ``;
                new_group_icon = 'fas fa-check';
                exp_period = '<span class="black">실험 종료</span> : ' + $(ele).data('period');

                group_Starter.switchClass('unfolded', 'folded');
                group_Starter.find('.unfolded').switchClass('unfolded', 'folded');

                isFolded = true;
              };

              if (group_Starter.hasClass('inProgress') || group_Starter.hasClass('closed') || group_Starter.hasClass('completed')) {
                let tippy_offset = !group_Starter.data('offset') ? '0, 0' : group_Starter.data('offset'),
                    tippy_target = !group_Starter.data('tippy-target') ? ele : group_Starter.data('tippy-target').split(',');

                tippy(group_Starter[0], {
                  offset: tippy_offset,
                  triggerTarget: tippy_target
                });
              }
            }

            if (isFolded) {
              $(ele).css({ 'display': 'none' });
            }

            $(ele).off('mouseenter mouseleave').on({
              mouseenter: e => {
                let $group_tr = $(e.currentTarget).prevAll('.rowGroup').first(),
                    $sub_trs = $group_tr.add($group_tr.nextUntil('.rowGroup')),
                    group_idx = $('tr.rowGroup').index($group_tr);

                if (typeof rowTippy !== 'undefined') {
                  clearTimeout(rowTippy[group_idx]);
                }
                try {
                  $group_tr[0]._tippy.show();
                } catch (error) {}
                $sub_trs.addClass('active');
              },
              mouseleave: e => {
                $(e.currentTarget).siblings().add($(e.currentTarget)).removeClass('active');

                let $group_tr = $(e.currentTarget).prevAll('.rowGroup').first(),
                    group_idx = $('tr.rowGroup').index($group_tr);

                rowTippy[group_idx] = setTimeout(function () {
                  try {
                    $group_tr[0]._tippy.hide();
                  } catch (error) {}
                }, 300);
              }
            });
          });

          group_name += `<span class="subInfo">( ${exp_period} )</span>`;
          if (!!action_buttons) {
            group_name += action_buttons;
          }

          group_Starter.find('.group-folder+i').addClass(new_group_icon).after(group_name);
        } else {
          group_Starter = $('<tr class="nonGroup"/>').addClass(fold_type).append(result_Row);
          group_Starter_cols = group_Starter.children();
          group_Starter_cols.eq(groupStart_column).html(`<span class="group-folder ${fold_type}"></span>`);
          rowGroups.each(function (idx, ele) {
            if (idx < rowGroups.length - 1) {
              $(ele).removeClass('group-sub-end').addClass('nonGroup-sub group-sub');
            } else {
              $(ele).removeClass('group-sub').addClass('nonGroup-sub group-sub-end');
            }
          });

          group_Starter.find('.group-folder').after(group_name);
        }

        group_Starter.find('a').on('click', function (e) {
          e.stopPropagation();
        });

        let filtered_rows = $('#datatable').DataTable().rows().data().filter(function (data, index) {
          return data[0] == group ? true : false;
        }),
            group_comment = '※ 이 실험의 관련 문서가 더 있습니다. <span class="gray">검색 또는 페이지 나뉨으로 목록에서 보이지 않습니다.</span>';

        if (rows.count() < $(filtered_rows).length) {
          group_Starter.find('td').append(`<span class="red_deep thin onlyUnfolded mLeft10">${group_comment}</span>`);
        };

        group_Starter.off().on({
          mouseenter: function (e) {
            let $group_tr = $(this),
                $sub_trs = $group_tr.add($group_tr.nextUntil('.rowGroup')),
                group_idx = $('tr.rowGroup').index(this);

            $sub_trs.addClass('active');
            if (typeof rowTippy !== 'undefined') {
              clearTimeout(rowTippy[group_idx]);
            }
          },
          mouseleave: function (e) {
            $(this).siblings().add($(this)).removeClass('active');
          },
          click: function (e) {
            e.stopPropagation();
            var folded_icon = $(this).find('.group-folder'),
                same_level = group_Starter.nextUntil('.dtrg-level-' + level),
                upper_level = group_Starter.nextUntil('.dtrg-level-' + (level - 1)),
                inter_group = same_level.filter(upper_level);

            if ($(this).nextUntil('.rowGroup').filter('.selected').length) {
              $(this).attr('title', $(this).data('content'));
            } else {
              if (folded_icon.hasClass('unfolded')) {
                $(this).switchClass('unfolded', 'folded');
                folded_icon.switchClass('unfolded', 'folded');
                level > 0 ? inter_group.hide() : same_level.hide();
              } else {
                $(this).switchClass('folded', 'unfolded');
                folded_icon.switchClass('folded', 'unfolded');
                level > 0 ? inter_group.show() : same_level.show();
              }
            }
          }
        });

        return group_Starter;
      },
      endRender: function (rows, group, level) {},
      dataSrc: rowGroup_column_Arry
    },

    infoCallback: function (settings, start, end, max, total, pre) {

      var $tHead = $(this).find('thead'),
          $tBody = $(this).find('tbody'),
          $search_info = $('.search_info');

      if ($search_info.length) {
        if (total < max) {
          var total_counts = '<span class="mHside5">' + total + ' / ' + max + '</span>';
        } else {
          var total_counts = '<span class="mHside5">전체: ' + max + '</span>';
        }
      } else {
        if (total < max) {
          var total_counts = '<span class="mHside5">전체 결과: ' + total + ' / ' + max + '</span>';
        } else {
          var total_counts = '<span class="mHside5">전체 결과: </span>' + max;
        }
      }

      if (end > 0) {
        var result_range = start + ' ~ ' + end;
      } else {
        var result_range = 0;
      }

      if (!$tHead.hasClass('noSearch_info')) {

        if ($tHead.data('info-type') == 'mini') {
          $search_info.html(result_range + `<span class="mHside5">( 전체:</span>${max} )`);
        } else if (typeof $tHead.data('info-type') === 'undefined') {
          $search_info.html(`검색 결과: ${result_range}<span class="mLeft5">( </span>${total_counts} )`);
        }
      } else {
        $search_info.empty();
      }
    },


    initComplete: function () {
      $('.export_buttons').append($('.dt-buttons button').append('<i class="fas fa-download align-self-center icon-xs text-white"></i>'));
      $('.dt-buttons').remove();
      $('.export_buttons button').attr({ 'title': '목록 Download' });
    },


    drawCallback: function (settings, json) {

      var $table = $(this),
          $tHead = $(this).find('thead'),
          $tBody = $(this).find('tbody');

      if (settings.aiDisplay.length == 0) {
        $tHead.find('th').removeAttr('style');
        $tBody.find('tr').addClass('zerorecords');
      }

      if (!$table.hasClass('rowGroup')) {
        $tBody.find('tr').each((idx, ele) => {
          if ($(ele).hasClass('inProgress')) {
            $(ele).attr({
              'title': "실험이 진행중입니다."
            });
          }
          if ($(ele).hasClass('closed')) {
            $(ele).attr({
              'title': "실험 기간이 만료되었습니다."
            });
          }
          if ($(ele).hasClass('completed')) {
            $(ele).attr({
              'title': "실험이 종료되었습니다."
            });
          }
          if ($(ele).hasClass('supplement')) {
            $(ele).attr({
              'title': "신청서를 보완해주세요."
            });
          }
          if ($(ele).hasClass('delayed')) {
            $(ele).attr({
              'title': "심사가 지연됐습니다."
            });
          }
          if ($(ele).hasClass('to_expire')) {
            $(ele).attr({
              'title': "재승인 기간입니다. 재승인을 신청해주세요."
            });
          }
          $(ele).attr({
            'data-tippy-placement': "top-start",
            'data-tippy-animation': "shift-away",
            'data-tippy-arrow': "true",
            'data-tippy-target': ".title, .status"
          });

          tippy(ele);
        });
      }

      if (typeof $table.data('buttons-url') !== 'undefined') {
        let button_text_arry = $table.data('buttons-text').split(','),
            button_icon_arry = $table.data('buttons-icon').split(','),
            button_url_arry = $table.data('buttons-url').split(','),
            $pagination = $table.siblings('.dataTables_paginate').find('ul.pagination'),
            button_icon = '';

        $pagination.append('<div class="buttons"></div>');
        $.each(button_url_arry, (idx, val) => {
          if (!button_icon_arry) {
            button_icon = '';
          } else if (!!button_icon_arry[idx].trim()) {
            button_icon = `<i class="${button_icon_arry[idx].trim()} mRight5"></i>`;
          }
          $pagination.find('.buttons').append(`
            <a href="${button_url_arry[idx]}" class="btn btn-sm btn-primary button_new">${button_icon}${button_text_arry[idx]}</a>`);
        });
      }

      if ($table.data('next-index') > 0) {
        $('.pagination').append(`<a href="?filter.start_index=${$table.data('next-index')}" class="btn btn-outline-secondary btn-xs flexMid mLeft20" id="go_next_block">다음 ${$table.data('rowCnt')} ${block_unit}<i class="fas fa-angle-double-right mLeft5"></i></a>`);
      }
      if ($table.data('before-index') > -1) {
        $('.pagination').prepend(`<a href="?filter.start_index=${$table.data('before-index')}" class="btn btn-outline-secondary btn-xs flexMid mRight20" id="go_prev_block"><i class="fas fa-angle-double-left mRight5"></i>이전 ${$table.data('rowCnt')} ${block_unit}</a>`);
      }

      let $search_form = $('.search_form');
      if ($tHead.hasClass('noSearch')) {
        $('.dataTables_filter').remove();
      } else if (!$search_form.find('.dataTables_filter').length) {
        let $tables_filter = $('.dataTables_filter');
        $tables_filter.find('input[type=search]').attr('placeholder', '입력 즉시 검색합니다.');
        $search_form.prepend($tables_filter);
      }

      var filter_arr = ['.myReview_filters', '.progress_filters', '.committee_filters', '.appType_filters', '.status_filters', '.user_filters', '.payment_filters'];
      var $check_filter_group = $(filter_arr.join(',')),
          regExp = true;

      $check_filter_group.each((ids, ele) => {
        var $check_filters = $check_filter_group.find('.filter');

        $check_filters.each((idx2, ele2) => {
          var $filter = $(ele2);

          $filter.off().on({
            change: function (e) {
              e.stopPropagation();
              var $this_filters = $filter.closest('.check_filter'),
                  $filter_inputs = $this_filters.find('input[type=checkbox]'),
                  filter_col = $this_filters.data('column'),
                  $other_check_filters = $check_filter_group.not($this_filters),
                  $other_filter_inputs = $other_check_filters.find('input[type=checkbox]'),
                  other_filter_col = $other_check_filters.data('column'),
                  filter_vals,
                  other_filter_vals;

              filter_vals = '';
              $filter_inputs.each(function (idx, ele) {
                if ($(ele).prop('checked')) {
                  filter_vals += $(ele).filter(function (idx, ele) {
                    return $(ele).prop('checked');
                  }).val() + '|';
                }
              });
              filter_vals = filter_vals.slice(0, -1);
              datatable.column(filter_col).search(filter_vals, regExp, false).draw();

              other_filter_vals = '';
              $other_filter_inputs.each(function (idx, ele) {
                if ($(ele).prop('checked')) {
                  other_filter_vals += $(ele).filter(function (idx, ele) {
                    return $(ele).prop('checked');
                  }).val() + '|';
                }
              });
              other_filter_vals = other_filter_vals.slice(0, -1);
              if (other_filter_vals) {
                datatable.column(other_filter_col).search(other_filter_vals, regExp, false).draw();
              }
            }
          });
        });
      });

      $(this).find('tbody .child_row').each(function (idx, ele) {
        var $childRow = $(ele);

        $childRow.each(function (idx2, ele2) {
          var triggerWrap = $(ele2),
              trigger = triggerWrap.find('.child_trigger');

          trigger.off().on({
            click: function (e) {
              e.stopPropagation();
              var row = datatable.row(triggerWrap);
              if ($(this).hasClass('active')) {
                $(this).removeClass('active');
                row.child.hide();
              } else {
                $(this).addClass('active');

                row.child(row.data()[childCol_index]).show();
                $(row.child()).addClass('child');
                $(row.child()).on({
                  mouseenter: function () {
                    $(this).prev('tr').addClass('hovered');
                  },
                  mouseleave: function () {
                    $(this).prev('tr').removeClass('hovered');
                  }
                });
              }
            }
          });
        });
      });
    }
  };


  $("input[name='app_name_order']:radio").on({
    change: function (e) {
      let dataTable = $('#datatable').DataTable();
      let orderArr = [this.dataset.order, this.value];
      dataTable.order.fixed({
        pre: orderArr
      }).draw();
    }
  });
});

function initDataTable(table_id) {
  if ($(table_id).closest('.board_01').length) {
    dataTableOption.language.search = `<span class="filter_label"><i class="fas fa-search mRight5"></i>공지사항 검색:</span>`;
  }
  if ($(table_id).closest('.board_02').length) {
    dataTableOption.language.search = `<span class="filter_label"><i class="fas fa-search mRight5"></i>자료실 검색:</span>`;
  }
  if ($(table_id).closest('.board_03').length) {
    dataTableOption.language.search = `<span class="filter_label"><i class="fas fa-search mRight5"></i>FAQ 검색:</span>`;
  }
  if ($(table_id).closest('.experiment').length) {
    dataTableOption.language.search = `<span class="filter_label"><i class="fas fa-search mRight5"></i>실험 및 문서 검색:</span>`;
  }
  if ($(table_id).hasClass('noOrder')) {
    dataTableOption.ordering = false;
  }
  if ($(table_id).hasClass('noPaging')) {
    dataTableOption.paging = false;
  }
  datatable = $(table_id).DataTable(dataTableOption);

  let urlParams = new URLSearchParams(window.location.search),
      resultVal = urlParams.get('state');

  if (resultVal == 'clear') {
    datatable.state.clear();
    $(table_id).DataTable().search('').columns().search('').draw();

    $('.check_filters input:checked').eq(0).trigger('input');

    history.replaceState(undefined, undefined, window.location.href.split('?')[0]);
  } else {
    $('.check_filters input:checked').eq(0).trigger('input');

    if (stateSaveParams_data) {
      $.each(stateSaveParams_data.columns, (idx, val) => {
        if (val.search.search) {
          $('.check_filters [type=checkbox]').filter((idx2, ele2) => {
            return val.search.search.split('|').indexOf($(ele2).val()) > -1;
          }).prop('checked', true);
        }
      });
    }
  }
}

var filter_data = {
  iacuc: {
    str: '<span class="IACUC_color">IACUC</span>',
    val: 'IACUC'
  },
  ibc: {
    str: '<span class="IBC_color">IBC</span>',
    val: 'IBC'
  },
  irb: {
    str: '<span class="IRB_color">IRB</span>',
    val: 'IRB'
  },
  new: {
    str: '신규 승인',
    val: '신규 승인'
  },
  change: {
    str: '변경 승인',
    val: '변경 승인'
  },
  renew: {
    str: '재승인',
    val: '재승인'
  },
  bring: {
    str: '반입 보고',
    val: '반입 보고'
  },
  end: {
    str: '종료 보고',
    val: '종료 보고'
  },
  inspection: {
    str: '승인후 점검',
    val: '승인후 점검'
  },
  general: {
    str: '일반 심사',
    val: '일반심사'
  },
  expert: {
    str: '전문 심사',
    val: '전문심사'
  },
  saved: {
    str: '임시 저장, <span class="red">보완중</span>',
    val: '임시 저장|보완중'
  },
  delayed: {
    str: '<span class="red">심사 지연</span>',
    val: '심사 지연'
  },
  approved: {
    str: '승인',
    val: '승인'
  },
  denied: {
    str: '<span class="red">반려</span>',
    val: '반려'
  },
  office_1: {
    str: '행정 검토',
    val: '^((?!심사).)*$'
  },
  office_2: {
    str: '심사 설정',
    val: '심사 설정'
  },
  experimenting: {
    str: '실험중',
    val: '실험중'
  },
  registered: {
    str: '정상',
    val: '정상'
  },
  unregistered: {
    str: '등록대기',
    val: '등록대기'
  },
  withdrawn: {
    str: '탈퇴',
    val: '탈퇴'
  },
  expelled: {
    str: '강제탈퇴',
    val: '강제탈퇴'
  },
  canceled: {
    str: '결제취소',
    val: '전체 취소|부분 취소'
  }
};

function setListFilter(type, filter, flag) {
  let $wrapper = $(`.${type}_filters`),
      filter_arr = filter.split(',');

  $.each(filter_arr, (idx, val) => {
    if (typeof val !== 'undefined') {
      let filter = val;
      if (['IACUC', 'IBC', 'IRB'].indexOf(filter_data[filter].val) > -1) {
        var committee_color = ' ' + filter_data[filter].val;
      }
      $wrapper.append(`
      <div class="checkbox checkbox-primary filter filter_${filter}${committee_color}">
        <input type="checkbox" id="filter_${filter}" value="${filter_data[filter].val}">
        <label for="filter_${filter}" class="mBot0">${filter_data[filter].str}</label>
      </div>`);
    }

    if (flag == 'last' && idx == filter_arr.length - 1) {
      $('.check_filter').each((idx, ele) => {
        if (!$(ele).children().length) {
          $(ele).remove();
        }
      });
    }
  });
  $('.check_filters').removeClass('hidden');
}

function setServerSearch($search_form, $tables_filter) {
  let urlParams = new URLSearchParams(window.location.search),
      param_key = 'search_words',
      result = urlParams.get(param_key) || '';

  $search_form.append(`
  <span class="search_divider alt_search collapse show">|</span>
  <form method="get" action="" class="alt_search collapse">
    <div class="flexMid">
      <span class="filter_label"><i class="fas fa-search-plus mRight5"></i>전체 목록 검색:</span>
      <input type="search" name="${param_key}" placeholder="전체 목록에서 검색합니다." class="form-control form-control-sm" value="${result}">
      <button type="submit" class="btn btn-sm btn-primary mLeft5">검색 실행</button>
      <button type="reset" class="btn btn-sm btn-secondary mLeft5">초기화</button>
      <span class="search_divider">|</span>
      <button type="button" class="btn btn-sm btn-outline-primary" data-toggle="collapse" data-target=".alt_search"><i class="fas fa-exchange-alt mRight5"></i>결과 내에서 검색</button>
    </div>
  </form>
  <button type="button" class="btn btn-sm btn-outline-primary alt_search collapse show" data-toggle="collapse" data-target=".alt_search"><i class="fas fa-exchange-alt mRight5"></i>전체 목록에서 검색</button>`);
  $tables_filter.addClass('collapse show alt_search');
  $search_form.find('button[type=reset]').on({
    click: e => {
      e.preventDefault;
      $search_form.find('input[type=search]').val('');
      $search_form.find('form').submit();
    }
  });
  if (result) {
    $('button.alt_search').trigger('click');
  }
}

function setBlockBtn(btn_type, data, rowCnt, startIndex) {
  let urlParams = new URLSearchParams(window.location.search),
      param_key = 'search_words',
      result = urlParams.get(param_key);
  result = result ? `&${param_key}=${urlParams.get(param_key)}` : '';

  if (btn_type == 'next') {
    $('.pagination').append(`<a href="?filter.start_index=${data}&filter.row_cnt=${rowCnt}${result}" class="btn btn-outline-secondary btn-xs flexMid mLeft20" id="go_next_block">다음 ${rowCnt} 건<i class="fas fa-angle-double-right mLeft5"></i></a>`);
    $('#datatable').data({ 'next-index': data, 'rowCnt': rowCnt });
  } else {
    $('.pagination').prepend(`<a href="?filter.start_index=${startIndex - rowCnt}&filter.row_cnt=${rowCnt}${result}" class="btn btn-outline-secondary btn-xs flexMid mRight20" id="go_prev_block"><i class="fas fa-angle-double-left mRight5"></i>이전 ${rowCnt} 건</a>`);
    $('#datatable').data({ 'before-index': startIndex - rowCnt, 'rowCnt': rowCnt });
  }
}