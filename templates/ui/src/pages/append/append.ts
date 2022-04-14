import "./append.less"
import "../../components/calendar-popup/index.less"
import * as popup from "../../components/calendar-popup/index";

// 添加所有点击事件
const handlers = document.querySelectorAll(".handler")
handlers.forEach(ele=>{
    // 获取操作符号, 操作符号在 html 里定义
    const opt = ele.getAttribute("opt")
    let func = (e: Event) => {}
    switch (opt){
        case "SELECT_DATE":
            func = openCalendar
            break;
        case "":
            break;
    }
    ele.addEventListener("click", func)
})

// 打开 日期选择器
function openCalendar(e:Event) {
    console.log(111)
    popup.OpenCalendarPopup()
}

var calendar = popup.GenerateCalendarPopup()
document.body.appendChild(calendar)
popup.CloseCalendarPopupListen()