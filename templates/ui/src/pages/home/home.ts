import './home.less'
import "../../components/tag-popup/index.less"
import "../../components/month-popup/index.less"
import * as popup from '../../components/month-popup/index'
import * as tagsPopup from "../../components/tag-popup/index";
import { ITags, ITagsParams } from "../../components/tag-popup/types";

import defaultAvatr from '../../assets/1649346921196.jpg'
import { gid } from '../../common/helper'

const avatar: HTMLImageElement = document.getElementById('avatar')! as HTMLImageElement
avatar.src = defaultAvatr

const topClicks = document.querySelectorAll('.topnav .nav')

const sidebar = document.getElementById('sidebar')
const closeSidebar = document.getElementById('sidebar-close')

const addRecord = (ele: Element) => {
	ele.addEventListener('click', () => {
		console.log('add')
	})
}

// ------------------- sidebar

const showSidebar = (nodes: NodeListOf<Element>) => {
	nodes.forEach((ele: Element) => {
		if (ele.id == 'add') {
			addRecord(ele)
		} else {
			ele.addEventListener('click', () => {
				sidebar?.classList.add('active')
			})
		}
	})
}
showSidebar(topClicks)

closeSidebar?.addEventListener('click', (e) => {
	sidebar?.classList.remove('active')
})

// --------------------------------- monthPopup

gid('select-date')?.addEventListener('click', (e) => {
    e.stopPropagation()
	popup.OpenMonthPopup()
})

// 生成 monthPopup
const apppopup = popup.GenerateMonthPopup()
document.body.appendChild(apppopup)

// 渲染dom完成执行监听close事件
popup.CloseMonthPopupListen()

// 监听点击事件
popup.SelectorCall((date: string) => {
	console.log(date)
	popup.CloseMonthPopup()
})

// --------------------------------------------- tags-popup
// 生成 tags-popup
const allTags: ITagsParams[] = [
    {
        title: "收入",
        tags:[
            	{ id: 1, name: '生意' },
            	{ id: 1, name: '工资' },
            	{ id: 1, name: '领奖金' },
            	{ id: 1, name: '赚外快' },
            	{ id: 1, name: '收红包' },
            	{ id: 1, name: '收转款' },
            	{ id: 1, name: '退款' },
            	{ id: 1, name: '其他收入' },
            ]
    },
    {
        title: "支出",
        tags: [
            { id: 1, name: '人间烟火' },
            { id: 2, name: '服饰美容' },
            { id: 3, name: '环球出行' },
            { id: 4, name: '人情世故' },
            { id: 5, name: '健康运动' },
            { id: 6, name: '知识投资' },
            { id: 7, name: '看病就医' },
            { id: 8, name: '生活娱乐' },
            { id: 9, name: '子女教育' },
            { id: 10, name: '慷慨解囊' },
            { id: 11, name: '生活缴费' },
            { id: 11, name: '其他支出' },
        ]
    }
]
var tagsDom  = tagsPopup.GenerateTagPopup(allTags)
document.body.appendChild(tagsDom)

gid("select-category")?.addEventListener("click", (e) => {
    e.stopPropagation()
    tagsPopup.OpenTagPopup()
})
// 启动close监听
tagsPopup.CloseTagPopupListen()
tagsPopup.SelectorCall((data: string)=>{
    console.log(data)
    tagsPopup.CloseTagPopup()
})