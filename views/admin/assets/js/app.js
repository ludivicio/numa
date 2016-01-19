(function ($) {
  'use strict';

  $(function () {
    var $fullText = $('.admin-fullText');
    $('#admin-fullscreen').on('click', function () {
      $.AMUI.fullscreen.toggle();
    });

    $(document).on($.AMUI.fullscreen.raw.fullscreenchange, function () {
      $fullText.text($.AMUI.fullscreen.isFullscreen ? '退出全屏' : '开启全屏');
    });
  });
})(jQuery);


/*
 * 图片上传预览， IE是用了滤镜
 */
function previewImage(file) {
  var MAXWIDTH = 260;
  var MAXHEIGHT = 180;
  if (file.files && file.files[0]) {
    var img = document.getElementById('preview');
    img.onload = function () {
      var rect = clacImgZoomParam(MAXWIDTH, MAXHEIGHT, img.offsetWidth, img.offsetHeight);
      img.width = rect.width;
      img.height = rect.height;
      // img.style.marginLeft = rect.left + 'px';
      // img.style.marginTop = rect.top + 'px';
    }
    var reader = new FileReader();
    reader.onload = function (evt) { img.src = evt.target.result; }
    reader.readAsDataURL(file.files[0]);
  } else {
    //兼容IE
    file.select();
    var src = document.selection.createRange().text;
    var img = document.getElementById('preview');
    img.filters.item('DXImageTransform.Microsoft.AlphaImageLoader').src = src;
  }
}

function clacImgZoomParam(maxWidth, maxHeight, width, height) {
  var param = { top: 0, left: 0, width: width, height: height };
  if (width > maxWidth || height > maxHeight) {
    rateWidth = width / maxWidth;
    rateHeight = height / maxHeight;

    if (rateWidth > rateHeight) {
      param.width = maxWidth;
      param.height = Math.round(height / rateWidth);
    } else {
      param.width = Math.round(width / rateHeight);
      param.height = maxHeight;
    }
  }

  param.left = Math.round((maxWidth - param.width) / 2);
  param.top = Math.round((maxHeight - param.height) / 2);
  return param;
}

