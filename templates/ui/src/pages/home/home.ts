import "./home.less"
import defaultAvatr from "../../assets/1649346921196.jpg"

const avatar:HTMLImageElement = document.getElementById("avatar")! as HTMLImageElement
avatar.src = defaultAvatr

const topClicks = document.querySelectorAll(".topnav .nav")

const sidebar = document.getElementById("sidebar")
const closeSidebar = document.getElementById("sidebar-close")


const showSidebar = (nodes: any) => {
    nodes.forEach((ele: Element)=>{
        ele.addEventListener("click", ()=>{
            sidebar?.classList.add("active")
        })
    })
}
showSidebar(topClicks)

closeSidebar?.addEventListener("click", ()=>{
    sidebar?.classList.remove("active")
})







