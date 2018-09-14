/**
 * 首页
 * Created by haibao on 17/6/23.
 */
$(function () {
    //轮播图
    var timer;
    var cArr = [];
    var index = 0;
    var $banner = $(".bannerList li");
    var $box = $(".bannerBox");

    $banner.each(function(i,e){
        //动态获取 图片列表li的class类名数组
        cArr.unshift($(e).attr('class'));
    });

    var cArrLen = cArr.length; //图片数量

    //下一张点击
    $(".next").stop().on('click', function(){
        nextimg();
    });

    //上一张
    //function previmg(){
    //    cArr.unshift(cArr[cArrLen-1]);
    //    cArr.pop();
    //    //i是元素的索引，从0开始
    //    //e为当前处理的元素
    //    //each循环，当前处理的元素移除所有的class，然后添加数组索引i的class
    //    $banner.each(function(i,e){
    //        $(e).removeClass().addClass(cArr[i]);
    //    });
    //    index--;
    //    if (index < 0) {
    //        index = cArrLen - 1;
    //    }
    //}

    //下一张
    function nextimg() {
        cArr.push(cArr[0]);
        cArr.shift();
        $banner.each(function(i, e){
            $(e).removeClass().addClass(cArr[i]);
        });
        index++;
        if (index > cArrLen - 1) {
            index = 0;
        }
    }

    //大于1张图片时启动定时器
    function start() {
        if(cArrLen > 1){
            timer = setInterval(nextimg, 4000);
        }
    }

    //鼠标移入/移出box时清除定时器
    $box.on('mouseover', function(){
        clearInterval(timer);
    }).on('mouseleave', function(){
        start();
    });

    //进入页面自动开始定时器
    start();

});