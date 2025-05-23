
!function ($) {
    "use strict";

    var Components = function () {};

    Components.prototype.initPopoverPlugin = function () {
        $.fn.popover && $('[data-toggle="popover"]').popover();
    }, Components.prototype.initSlimScrollPlugin = function () {
        $.fn.slimScroll && $(".slimscroll-alt").slimScroll({ position: 'right', size: "5px", color: '#98a6ad', wheelStep: 10 });
    }, Components.prototype.initRangeSlider = function () {
        $.fn.slider && $('[data-plugin="range-slider"]').slider({});
    }, Components.prototype.initCounterUp = function () {
        var delay = $(this).attr('data-delay') ? $(this).attr('data-delay') : 100;
        var time = $(this).attr('data-time') ? $(this).attr('data-time') : 1200;
        $('[data-plugin="counterup"]').each(function (idx, obj) {
            $(this).counterUp({
                delay: 100,
                time: 1200
            });
        });
    }, Components.prototype.initToast = function () {
        $.fn.toast && $('[data-toggle="toast"]').toast();
    }, Components.prototype.initAccordionBg = function () {
        $(".collapse.show").each(function () {
            $(this).prev(".card-header").addClass("custom-accordion");
        });

        $(".collapse").on('show.bs.collapse', function () {
            $(this).prev(".card-header").addClass("custom-accordion");
        }).on('hide.bs.collapse', function () {
            $(this).prev(".card-header").removeClass("custom-accordion");
        });
    }, Components.prototype.initValidation = function () {
        window.addEventListener('load', function () {
            var forms = document.getElementsByClassName('needs-validation');

            var validation = Array.prototype.filter.call(forms, function (form) {
                form.addEventListener('submit', function (event) {
                    if (form.checkValidity() === false) {
                        event.preventDefault();
                        event.stopPropagation();
                    }
                    form.classList.add('was-validated');
                }, false);
            });
        }, false);
    }, Components.prototype.initPrettify = function () {
        var entityMap = {
            '&': '&amp;',
            '<': '&lt;',
            '>': '&gt;',
            '"': '&quot;',
            "'": '&#39;',
            '/': '&#x2F;',
            '`': '&#x60;',
            '=': '&#x3D;'
        };
        function escapeHtml(string) {
            return String(string).replace(/[&<>"'`=\/]/g, function (s) {
                return entityMap[s];
            });
        }
        $(function () {
            $(".escape").each(function (i, e) {
                $(e).html(escapeHtml($(e).html()).trim());
            });
        });
    }, Components.prototype.init = function () {
        var $this = this;
        this.initPopoverPlugin(), this.initSlimScrollPlugin(), this.initRangeSlider(), this.initCounterUp(), this.initToast(), this.initAccordionBg(), this.initValidation(), this.initPrettify();
    }, $.Components = new Components(), $.Components.Constructor = Components;
}(window.jQuery), function ($) {
    "use strict";

    $.Components.init();
}(window.jQuery);