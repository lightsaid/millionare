
class Container{
	private static isExist: boolean = false
    private toastBox: HTMLDivElement = document.createElement('div')
    private ulBox: HTMLUListElement = document.createElement("ul")
	constructor(text: string, delay: number) {
		if (!Container.isExist) {
            this.toastBox.classList.add('app-toast')
            this.ulBox.classList.add("ul")
            this.toastBox.appendChild(this.ulBox)
		    document.body.append(this.toastBox)
			Container.isExist = true
		}
    }
    
}

export default class Toast extends Container {
    private static id: number = 0
    private container: HTMLDivElement | null;
	private li: HTMLLIElement;
	private text: string;
    private delay: number;
    private timer: any;
    private id: number;
	constructor(text: string, delay: number) {
        super(text, delay)
		if (text.length > 48) {
			text.slice(0, 45) + '...'
        }
        this.container = document.querySelector(".app-toast .ul")
        this.li = document.createElement('li')
        this.li.setAttribute("id", Toast.id.toString())
        let br = document.createElement('br')
        this.li.classList.add("msg")
		this.text = text
        this.li.innerText = text
        this.li.appendChild(br)
		this.delay = delay
        this.timer = null
        Toast.id += Toast.id
        this.id = Toast.id
        
        // 显示
        this.show()
    }

    // 显示msg
    show(){
        this.container?.appendChild(this.li)
        // 此时的 this.container? 和页面的 .app-toast 是不一样的
        document.querySelector(".app-toast")?.classList.remove("app-toast-hidden");
        setTimeout(()=>{
            this.close()
        },this.delay)
    }

    // 关闭当前msg 
    close() {
        this.container?.childNodes.forEach((node)=>{
            let item = node as (HTMLLIElement)
            if (item.id == this.id.toString()){
                this.container?.removeChild(node)
            }
        })
        clearTimeout(this.timer)
        if (this.container?.childNodes && this.container?.childNodes.length == 0) {
           document.querySelector(".app-toast")?.classList.add("app-toast-hidden")
        }
    }
    
    // 清除所有msg
    clear(){
        this.container?.childNodes.forEach((node)=>{
            this.container?.removeChild(node)
        })
        document.querySelector(".app-toast")?.classList.add("app-toast-hidden")
    }
}
