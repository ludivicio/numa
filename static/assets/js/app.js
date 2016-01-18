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


/**
 * 图片上传预览，IE使用滤镜
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
        //IE下，使用滤镜
        var imgSrc = document.selection.createRange().text;
        var localImagId = document.getElementById("preview");
        //必须设置初始大小
        localImagId.style.width = "150px";
        localImagId.style.height = "180px";
        //图片异常的捕捉，防止用户修改后缀来伪造图片
        try {
            localImagId.style.filter="progid:DXImageTransform.Microsoft.AlphaImageLoader(sizingMethod=scale)";
            localImagId.filters.item("DXImageTransform.Microsoft.AlphaImageLoader").src = imgSrc;
        } catch(e) {
            alert("您上传的图片格式不正确，请重新选择!");
            return false;
        }
        
        document.selection.empty();
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
