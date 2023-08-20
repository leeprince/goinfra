/* eslint-disable */
function qidianMonitorCallBack() {
  setTimeout(function () {
      // 为了给嵌入业务使用，会挂载到全局上
      window.QdMonitor = window.qidianMonitor && new window.qidianMonitor({
        projectName: "qidian-webim-v3",
        // 增加log日志上报不需要白名单的限制
        rumIgnoreWhiteList: true,
        // 上报之前
        beforeReport: function(log) {       
            if ((/trtc.js/.test(log.msg)) || (/trtc_4.14.4.js/.test(log.msg))) {
                return false;
            }
            // 对低版本浏览器进行屏蔽。chrome,FF 为18年开始的版本
            var LIMIT = {
                firefox: 58,
                chrome: 64,
                IE: 10
            };
            var limit = false;
            var ua = navigator.userAgent;
            var fM = ua.match(/Firefox\/(\d+)/i);
            var cM = ua.match(/Chrome\/(\d+)/i);
            var eM = ua.match(/MSIE (\d+)/i);
            if (fM) {
                if (Number(fM[1]) < LIMIT.firefox) {
                    limit = true;
                }
            }else if (cM) {
                if (Number(cM[1]) < LIMIT.chrome) {
                    limit = true;
                }
            }else if (eM) {
                if (Number(eM[1]) < LIMIT.IE) {
                    limit = true;
                }
            }
            if (limit) {
                return false;
            }

            // 增加script加载的文件进行收集
            try {
                if (log.msg && log.msg.trim().indexOf('Script error. @ (:0:0)')>-1){
                    return false;
                 }
                 // 移除chrome插件的报错
                 if(log.msg && log.msg.indexOf('chrome-extension://')>-1){
                    return false;
                 }
                 // 屏蔽 https://oazhiwen.qidian.qq.com/olapi/robot/complete-words
                 // AJAX_ERROR: request abort 
                 if(log.msg && log.msg.indexOf('AJAX_ERROR: request abort')>-1 && log.ctx && log.ctx.aegisUrl && log.ctx.aegisUrl.indexOf('/robot/complete-words') > -1){
                    return false;
                 }
                // 移除发生在*.html上的报错。这些报错都是注入代码引起的。
                 if(log.msg && /@\s*\(.*?\/chatv3([\w-]+)?\/[\w\-]+?\.html/i.test(log.msg)){
                    return false;
                 }
            } catch (error) {
                console.log('Script error',error);
            }
            return true // 返回true则进行上报，返回false则不进行上报
        },
        // 上报之后
        afterReport: function(log) {
            // console.log('afterReport', log) 
        },
        // 上报测速之前
        beforeReportSpeed: function(log) {
            // console.log('beforeReportSpeed', log)
            return log
        },
        // 发送请求之前，可以改变上报内容，必须返回需要上报的内容，否则会阻止上报
        beforeRequest: function(log) {
            // console.log('beforeRequest', log)
            return log
        },
        // 针对某监控平台特殊的钩子, 该钩子就只会给RUM执行
        rumBeforeRequest: function(log) {
            // console.log('beforeRequest', log)
            return log
        }
    });
  });
}