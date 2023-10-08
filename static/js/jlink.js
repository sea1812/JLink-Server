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