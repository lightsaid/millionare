import * as doc from "../../common/helper"

// let html = `    
// <div class="app-popup">
//     <div class="month-popup">
//         <div class="title">
//             <span>&times;</span>
//             <span>请选择月份</span>
//         </div>
//         <div class="months">
//             <p class="year">2020年</p>
//             <ul class="ul">
//                 <li>1月</li>
//                 <li class="active">2月</li>
//                 <li>3月</li>
//                 <li>4月</li>
//                 <li>5月</li>
//                 <li>6月</li>
//                 <li>7月</li>
//                 <li>8月</li>
//                 <li>9月</li>
//                 <li>10月</li>
//                 <li>11月</li>
//                 <li>12月</li>
//             </ul>
//         </div>
//     </div>
// </div>`

// 构造上面的模版
const GenerateMonthPopup = (): DocumentFragment => {
    var startDate = new Date("2020-01-01")
    var endDate = new Date()
    var mongths = ["1月", "2月", "3月", "4月","5月","6月","7月","8月","9月","10月","11月","12月"]

    // 使用文档碎片提高效率
    var  fragment = document.createDocumentFragment()

    // 容器
    var appPopup = doc.docc("div")
    appPopup.classList.add("app-popup")
    appPopup.classList.add("app-popup-month")
    // appPopup.classList.add("app-popup-show")

    var popupContent = doc.docc("div")
    popupContent.classList.add("month-popup")
    appPopup.appendChild(popupContent)

    // 标题
    var title = doc.docc("div")
    title.classList.add("title")
    title.innerHTML = `<span id="close-month-popup">&times;</span><span>请选择月份</span>`
    popupContent.appendChild(title)
    
    //遍历生成月份
    let start = startDate.getFullYear()
    let end = endDate.getFullYear()
    for(let i=start; i<=end; i++){
        var monthDom = doc.docc("div")
        monthDom.classList.add("months")
        let html = `
        <p class="year">${i}年</p>
        <ul class="ul">
            ${createMonthsHtml(i)}
        </ul>`

        // 每一年的月份，追加到appPopup
        monthDom.innerHTML = html
        popupContent.appendChild(monthDom)
    }

    // 遍历 mongths 生成 li html
    function createMonthsHtml (year: number): string{
        let listr = ""
        mongths.forEach(month=>{
            // 月份
            listr += `<li id=${year.toString() + month}>${month}</li>`
        })
        return listr
    }

    // 挂在到文档碎片
    fragment.appendChild(appPopup)
    return fragment
}

const CloseMonthPopupListen = () => {
    doc.gid("close-month-popup")?.addEventListener("click", (e) => {
        e.preventDefault()
        e.stopPropagation()
        popupTogger()
    })
}

const CloseMonthPopup = () => {
    popupTogger()
    doc.qas(".month-popup .ul li.active").forEach(ele=>{
        ele.classList.remove("active")
    })
}

const OpenMonthPopup = () => {
    popupTogger()
}
enum OptClass {
    add = "add",
    remove = "remove"
}
const setMonthPopupMarginBotton = (opt: OptClass, classSelector: string) => {
    const monthPop = doc.qss(".app-popup .month-popup") as HTMLDivElement
    monthPop.classList[opt](classSelector)
}

const popupTogger = () => {
    let timer: any;
    let appPopup = doc.qss(".app-popup")
    // 隐藏，先等monthContent做完动画,appPopup再display：none
    if (appPopup?.classList.contains("app-popup-show")){
        setMonthPopupMarginBotton(OptClass.add,"hidden")
        setMonthPopupMarginBotton(OptClass.remove, "show")
        setTimeout(()=>{
            appPopup?.classList.toggle("app-popup-show")
        }, 320)
    }else{
        // 显示，先等appPop 显示出来，monthpop再做动画
        appPopup?.classList.toggle("app-popup-show")
        timer = setTimeout(()=>{
                setMonthPopupMarginBotton(OptClass.add,"show")
                setMonthPopupMarginBotton(OptClass.remove,"hidden")
            clearTimeout(timer)
        }, 20)
    }
}

const SelectorCall = (cb: Function) =>  {
    doc.qss(".app-popup-month")?.addEventListener("click", (e)=>{
        e.stopPropagation()
        e.preventDefault()
        let tagget = e.target as HTMLLIElement
        tagget.classList.add("active")
        cb && cb(tagget.id)
    })
}

export  {
    GenerateMonthPopup,
    CloseMonthPopupListen,
    SelectorCall,
    CloseMonthPopup,
    OpenMonthPopup
}