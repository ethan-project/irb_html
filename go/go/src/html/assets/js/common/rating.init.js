

function ratingEnable($objects, rate, readonly) {
  rate = !rate ? null : rate;
  readonly = !readonly ? false : true;

  $objects.barrating('show', {
    initialRating: rate,
    readonly: readonly,
    theme: 'bars-square',
    showValues: true,
    showSelectedRating: false
  });
}

function ratingDisable() {
  $('select').barrating('destroy');
}