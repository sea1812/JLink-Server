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