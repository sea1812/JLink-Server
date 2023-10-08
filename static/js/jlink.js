/*
* jlink js
* */

//显示新增节点模式窗
function ShowInserNodeModal(){
    $("#myNewServerModal").modal();
}
//显示编辑节点模式窗
function ShowEditNodeModal(){

}
//跳转到页面
function JumpPage(url){
    window.location.href = url
}
//返回上页
function goBack(){
    history.back(-1);
}
//启动指定服务器上的指定进程
function startProcess(AProcessName, ANodeUniid){
    $.ajax({
        method:"post",
        url:"/api/startprocess",
        data:{"processname":AProcessName, "nodeuniid":ANodeUniid},
        success:function (result) {
            if(result=="ok"){
                window.location.reload();
            }else{
                alert(result);
            }
        }}
    );
}
//停止指定服务器上的指定进程
function stopProcess(AProcessName, ANodeUniid){
    $.ajax({
        method:"post",
        url:"/api/stopprocess",
        data:{"processname":AProcessName, "nodeuniid":ANodeUniid},
        success:function (result) {
            if(result=="ok"){
                window.location.reload();
            }else{
                alert(result);
            }
        }}
    );
}
//查看服务器上的标准错误日志输出
function stdError(AProcessName, ANodeUniid){
    $.ajax({
        method:"post",
        url:"/api/stderror",
        data:{"processname":AProcessName, "nodeuniid":ANodeUniid},
        success:function (result) {
            if(result.code==200){
                //alert(result.data)
                $("#title").html(AProcessName+"标准错误输出");
                $("#log").html(result.data);
                $("#myLogModal").modal();
            }else{
                alert(result);
            }
        }}
    );
}
//查看服务器上的标准输出日志
function stdOut(AProcessName, ANodeUniid){
    $.ajax({
        method:"post",
        url:"/api/stdout",
        data:{"processname":AProcessName, "nodeuniid":ANodeUniid},
        success:function (result) {
            if(result.code==200){
                //alert(result.data)
                $("#title").html(AProcessName+"标准日志输出");
                $("#log").html(result.data);
                $("#myLogModal").modal();
            }else{
                alert(result);
            }
        }}
    );
}
//查看进程信息
function processInfo(AProcessName, ANodeUniid){
    $.ajax({
        method:"post",
        url:"/api/processinfo",
        data:{"processname":AProcessName, "nodeuniid":ANodeUniid},
        success:function (result) {
            if(result.code===200){
                //alert(JSON.stringify(result, null, 1));
                $("#title").html(AProcessName+"进程信息");
                $("#log").html("<pre>"+result.data.content+"</pre>");
                $("#myLogModal").modal();
            }else{
                alert(result);
            }
        }}
    );
}