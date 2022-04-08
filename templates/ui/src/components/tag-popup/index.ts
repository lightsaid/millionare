import * as doc from "../../common/helper"

import { ITags, ITagsParams } from "./types";

// 构造此html dom

// <div class="app-popup tag-popup-show">
// <div class="tag-popup">
//     <div class="title">
//         <span>&times;</span>
//         <span>请选择类型</span>
//     </div>
//     <div class="tags">
//         <p class="tag-title">支出</p>
//         <ul class="ul">
//             <li class="tag active">人间烟火</li>
//             <li class="tag">人间烟火</li>
//             <li class="tag">人间烟火</li>
//             <li class="tag">人间烟火</li>
//             <li class="tag">人间烟火</li>
//             <li class="tag">人间烟火</li>
//         </ul>
//     </div>
// </div>
// </div>


// 构造上面的模版
export const GenerateTagPopup = (oTags: ITagsParams[]): DocumentFragment => {
    // 使用文档碎片提高效率
    var  fragment = document.createDocumentFragment()

    // 容器
    var appPopup = doc.docc("div")
    appPopup.classList.add("app-popup")
    appPopup.classList.add("app-popup-tag")
    // appPopup.classList.add("tag-popup-show")

    var popupContent = doc.docc("div")
    popupContent.classList.add("tag-popup")
    appPopup.appendChild(popupContent)

    // 标题
    var title = doc.docc("div")
    title.classList.add("title")
    title.innerHTML = `<span id="close-tag-popup">&times;</span><span>请选择类型</span>`
    popupContent.appendChild(title)
    
    //遍历生成tag
    for(let tag of oTags){
        var monthDom = doc.docc("div")
        monthDom.classList.add("tags")
        let html = `
        <p class="tag-title">${tag.title}</p>
        <ul class="ul">
            ${createTagsHtml(tag.tags)}
        </ul>`

        // 添加 支出/收入 的tag
        monthDom.innerHTML = html
        popupContent.appendChild(monthDom)
    }

    // 遍历 tags 生成 li html
    function createTagsHtml (tags: ITags[]): string{
        let listr = ""
        tags.forEach(tag=>{
            // 月份
            listr += `<li id=${tag.id}>${tag.name}</li>`
        })
        return listr
    }

    // 挂在到文档碎片
    fragment.appendChild(appPopup)
    return fragment
}


export const CloseTagPopupListen = () => {
    doc.gid("close-tag-popup")?.addEventListener("click", (e) => {
        e.preventDefault()
        e.stopPropagation()
        popupTogger()
    })
}

export const CloseTagPopup = () => {
    popupTogger()
    doc.qas(".ul li.active").forEach(ele=>{
        ele.classList.remove("active")
    })
}

export const OpenTagPopup = () => {
    popupTogger()
}

export const SelectorCall = (cb: Function) =>  {
    doc.qss(".tag-popup")?.addEventListener("click", (e)=>{
        let tagget = e.target as HTMLLIElement
        tagget.classList.add("active")
        cb && cb(tagget.id)
    })
}

const popupTogger = () => {
    doc.qss(".app-popup-tag")?.classList.toggle("tag-popup-show")
}

