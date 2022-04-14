import * as doc from "../../common/helper"

interface IDate {
    year: number,
    month: number,
    day: number,
    isCuurrentMonth: boolean
}

// str -> yyyy-mm-dd
const get_week_by_month_first_day = (str: string): number => {
    let date = new Date(str)
    return date.getDay()
}

const get_days_by_month = (year: string | number, month: string | number) => {
    var date=new Date(Number(year), Number(month), 0);
    var days=date.getDate();
    return days;
}

const weeks = ["日","一","二","三","四","五","六"]
const months = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]

// 构造上面的模版
const GenerateCalendarPopup = (): DocumentFragment => {
    var startDate = new Date("2022-01-01")
    var endDate = new Date()

    // 使用文档碎片提高效率
    var  fragment = document.createDocumentFragment()

    // 容器
    var appPopup = doc.docc("div")
    appPopup.classList.add("app-popup")
    appPopup.classList.add("app-popup-calendar")
    // appPopup.classList.add("calen-popup-show")

    var popupContent = doc.docc("div")
    popupContent.classList.add("calendar-popup")
    appPopup.appendChild(popupContent)

    // 标题
    var title = doc.docc("div")
    title.classList.add("title")
    title.innerHTML = `<span id="close-calendar-popup">&times;</span><span>请选择日期</span>`
    popupContent.appendChild(title)

    // 日期头部
    var header = doc.docc("ul")
    header.classList.add("top-header")
    let hli = ""
    weeks.forEach(w=> {
        hli += `<li>${w}</li>`
    })
    header.innerHTML = hli
    popupContent.appendChild(header)
    
    
    //遍历生成日期
    let start = startDate.getFullYear()
    let end = endDate.getFullYear()
    for(let i=start; i<=end; i++){
        for(let m of months){
            var calendarDom = doc.docc("div")
            calendarDom.classList.add("calendars")
            // 此处构造每一个月份的日期
            
            // 1. 获取当前月份总共多少天
            let currMonthAllDays = get_days_by_month(i, m)

            // 2. 这个月份从周几开始
            let firstWeekDay = `${i}-${m}-01`
            let week = get_week_by_month_first_day(firstWeekDay)

            // 3. 创建 IDate[] 数据，本月还没到的 isCuurrentMonth 设置为 false
            let monthDays: IDate[] = []
            for(let j =1; j<=currMonthAllDays; j++) {
                let flag = true
                // 未来的日子
                if (i==end && m ==endDate.getMonth()+1 &&  j>endDate.getDate()){
                    flag = false
                }
                let dd: IDate = {year: i, month: m, day: j, isCuurrentMonth: flag}
                monthDays.push(dd)
            }

            // 4. 填充每月开始那几天，可能是上个月份的日期
            // 如果 dayValue 超出了月份的合理范围，setDate 将会相应地更新 Date 对象。
            // 例如，如果为 dayValue 指定0，那么日期就会被设置为上个月的最后一天。
            // dateObj.setDate(dayValue)
            if(week > 0){ // 此月份不是从周日开始，需要填充上月天数
                const date = new Date(firstWeekDay)
                date.setDate(0)
                let tailDay = date.getDate()
                for(let k = week; k > 0; k--){
                    let dd: IDate = {year: date.getFullYear(), month: date.getMonth(), day: tailDay, isCuurrentMonth: false }
                    monthDays.unshift(dd)
                    tailDay--;
                }
            }

            let html = `
            <p class="year">${i}年${m}月</p>
            <ul class="ul">
                ${createCalendarHtml(monthDays)}
            </ul>`
    
            calendarDom.innerHTML = html
            popupContent.appendChild(calendarDom)

            // 5.2 如果超出当前时间就退出
            if(i === end && m > endDate.getMonth()) {
                break
            }
        }
    }

    // 遍历 days 生成 li html
    function createCalendarHtml (days: IDate[]): string{
        let listr = ""
        let date = new Date()
        days.forEach(d=>{
            let classic = d.isCuurrentMonth === false?'disble ':''
            if (d.year === date.getFullYear() && d.month === date.getMonth()+1 && date.getDate() === d.day) {
                classic += "selected "
            }
            let id = d.year.toString() + "-" + d.month.toString() + "-" + d.day.toString()

            listr += `<li id=${id} class=${classic}>${d.day}</li>`
        })
        return listr
    }

    // 挂在到文档碎片
    fragment.appendChild(appPopup)
    return fragment
}

const CloseCalendarPopupListen = () => {
    doc.gid("close-calendar-popup")?.addEventListener("click", (e) => {
        e.preventDefault()
        e.stopPropagation()
        popupTogger()
    })
}

const CloseCalendarPopup = () => {
    popupTogger()
    doc.qas(".calendar-popup .ul li.active").forEach(ele=>{
        ele.classList.remove("active")
    })
}

const OpenCalendarPopup = () => {
    popupTogger()
}
enum OptClass {
    add = "add",
    remove = "remove"
}
const setCalendarPopupMarginBotton = (opt: OptClass, classSelector: string) => {
    const calendarPop = doc.qss(".app-popup .calendar-popup") as HTMLDivElement
    calendarPop.classList[opt](classSelector)
}

const popupTogger = () => {
    let timer: any;
    let appPopup = doc.qss(".app-popup")
    // 隐藏，先等calendarContent做完动画,appPopup再display：none
    if (appPopup?.classList.contains("calen-popup-show")){
        setCalendarPopupMarginBotton(OptClass.add,"hidden")
        setCalendarPopupMarginBotton(OptClass.remove, "show")
        setTimeout(()=>{
            appPopup?.classList.toggle("calen-popup-show")
        }, 320)
    }else{
        // 显示，先等appPop 显示出来，calendarpop再做动画
        appPopup?.classList.toggle("calen-popup-show")
        timer = setTimeout(()=>{
            setCalendarPopupMarginBotton(OptClass.add,"show")
            setCalendarPopupMarginBotton(OptClass.remove,"hidden")
            clearTimeout(timer)
        }, 20)
    }
}

const SelectorCall = (cb: Function) =>  {
    doc.qss(".app-popup-calendar")?.addEventListener("click", (e)=>{
        e.stopPropagation()
        e.preventDefault()
        let tagget = e.target as HTMLLIElement
        tagget.classList.add("active")
        cb && cb(tagget.id)
    })
}

export  {
    GenerateCalendarPopup,
    CloseCalendarPopupListen,
    SelectorCall,
    CloseCalendarPopup,
    OpenCalendarPopup
}